package consensus

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/database"

	"github.com/desmos-labs/juno/client"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	constypes "github.com/forbole/bdjuno/x/consensus/types"
)

// ListenOperation allows to start listening to new consensus events properly
func ListenOperation(cp *client.Proxy, db *database.BigDipperDb) {
	events := []string{
		tmtypes.EventNewRound,
		tmtypes.EventNewRoundStep,
		tmtypes.EventCompleteProposal,
		tmtypes.EventVote,
		tmtypes.EventPolka,
		tmtypes.EventValidBlock,
	}

	// This channel will be used to gather all the events
	var eventChan = make(chan tmctypes.ResultEvent, 10)

	for _, event := range events {
		go subscribeConsensusEvent(event, cp, eventChan)
	}

	for event := range eventChan {
		handleEvent(event, cp, db)
	}
}

// subscribeConsensusEvent allows to subscribe to the consensus event having the given name,
// and returns a read-only channel emitting all the events
func subscribeConsensusEvent(event string, cp *client.Proxy, eventChan chan<- tmctypes.ResultEvent) {
	query := fmt.Sprintf("tm.event = '%s'", event)

	eventCh, cancel, err := cp.SubscribeEvents("juno", query)
	if err != nil {
		log.Error().Str("module", "consensus").Err(err).Msg("error while subscribing to event")
		return
	}
	defer cancel()

	for event := range eventCh {
		eventChan <- event
	}
}

// handleEvent handles the given event storing its data inside the database properly
func handleEvent(event tmctypes.ResultEvent, cp *client.Proxy, db *database.BigDipperDb) {
	consEvent := mapEvent(event, cp)
	if consEvent == nil {
		return
	}

	// Save the event
	log.Debug().Str("module", "consensus").
		Int64("height", consEvent.Height).
		Int32("round", consEvent.Round).
		Str("step", consEvent.Step).
		Msg("saving consensus event")

	err := db.SaveConsensus(consEvent)
	if err != nil {
		log.Error().Str("module", "consensus").Err(err).Msg("error while saving consensus event")
	}
}

// mapEvent converts the given ResultEvent to a ConsensusEvent instance
func mapEvent(event tmctypes.ResultEvent, cp *client.Proxy) *constypes.ConsensusEvent {
	switch data := event.Data.(type) {
	case tmtypes.EventDataNewRound:
		return constypes.NewConsensusEvent(data.Height, data.Round, data.Step, 100)

	case tmtypes.EventDataRoundState:
		return constypes.NewConsensusEvent(data.Height, data.Round, data.Step, 100)

	case tmtypes.EventDataCompleteProposal:
		return constypes.NewConsensusEvent(data.Height, data.Round, data.Step, 100)

	case tmtypes.EventDataVote:
		return mapEventDataVote(data, cp)

	default:
		return nil
	}
}

func mapEventDataVote(event tmtypes.EventDataVote, cp *client.Proxy) *constypes.ConsensusEvent {
	data := event.Vote

	consData, err := cp.ConsensusState()
	if err != nil {
		log.Error().Str("module", "consensus").Err(err).Msg("error while getting consensus data")
		return nil
	}

	type roundVote struct {
		Round              int32    `json:"round"`
		Prevotes           []string `json:"prevotes"`
		PrevotesBitArray   string   `json:"prevotes_bit_array"`
		Precommits         []string `json:"precommits"`
		PrecommitsBitArray string   `json:"precommits_bit_array"`
	}

	var votes []roundVote
	err = tmjson.Unmarshal(consData.Votes, &votes)
	if err != nil {
		log.Error().Str("module", "consensus").Err(err).Msg("error while getting consensus round votes")
		return nil
	}

	var vote *roundVote
	for _, rv := range votes {
		if rv.Round == data.Round {
			vote = &rv
		}
	}

	if vote == nil {
		log.Error().Str("module", "consensus").Int32("round", data.Round).Msg("round votes not found")
		return nil
	}

	var bAString string
	var voted, total int64
	var fracVoted float64
	_, err = fmt.Sscanf(vote.PrevotesBitArray, "%s %d/%d = %f", &bAString, &voted, &total, &fracVoted)
	if err != nil {
		log.Error().Str("module", "consensus").Err(err).
			Int32("round", data.Round).
			Msg("error while parsing prevotes bytearray")
		return nil
	}

	return constypes.NewConsensusEvent(data.Height, data.Round, tmtypes.EventVote, fracVoted*100)
}
