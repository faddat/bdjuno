package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Proposal represent storing a gov.proposal
// For final tolly result, it stored in tally result as they share same proposal ID and VotingEndTime
type Proposal struct {
	Title           string
	Description     string
	ProposalRoute   string
	ProposalType    string
	ProposalID      uint64
	Status          gov.ProposalStatus
	SubmitTime      time.Time
	DepositEndTime  time.Time
	VotingStartTime time.Time
	VotingEndTime   time.Time
	Proposer        string
}

// NewProposal return a new Proposal instance
func NewProposal(
	title string,
	description string,
	proposalRoute string,
	proposalType string,
	proposalID uint64,
	status gov.ProposalStatus,
	submitTime time.Time,
	depositEndTime time.Time,
	votingStartTime time.Time,
	votingEndTime time.Time,
	proposer string,

) Proposal {
	return Proposal{
		Title:           title,
		Description:     description,
		ProposalRoute:   proposalRoute,
		ProposalType:    proposalType,
		ProposalID:      proposalID,
		Status:          status,
		SubmitTime:      submitTime,
		DepositEndTime:  depositEndTime,
		VotingStartTime: votingStartTime,
		VotingEndTime:   votingEndTime,
		Proposer:        proposer,
	}
}

//MsgVote
type TallyResult struct {
	ProposalID uint64
	Yes        int64
	Abstain    int64
	No         int64
	NoWithVeto int64
	Height     int64
}

// NewTallyResult return a new TallyResult instance
func NewTallyResult(
	proposalID uint64,
	yes int64,
	abstain int64,
	no int64,
	noWithVeto int64,
	height int64,
) TallyResult {
	return TallyResult{
		ProposalID: proposalID,
		Yes:        yes,
		Abstain:    abstain,
		No:         no,
		NoWithVeto: noWithVeto,
		Height:     height,
	}
}

// Vote describe a msgVote
type Vote struct {
	ProposalID uint64
	Voter      string
	Option     gov.VoteOption
	Height     int64
}

// NewVote return a new Vote instance
func NewVote(
	proposalID uint64,
	voter string,
	option gov.VoteOption,
	height int64,
) Vote {
	return Vote{
		ProposalID: proposalID,
		Voter:      voter,
		Option:     option,
		Height:     height,
	}
}

// Deposit represent a message that a user do deposit action
// Assume the entry with latest height get final total deposit
type Deposit struct {
	ProposalID uint64
	Depositor  string
	Amount     sdk.Coins
	Height     int64
}

//NewDeposit return a new Deposit instance
func NewDeposit(
	proposalID uint64,
	depositor string,
	amount sdk.Coins,
	height int64,
) Deposit {
	return Deposit{
		ProposalID: proposalID,
		Depositor:  depositor,
		Amount:     amount,
		Height:     height,
	}
}
