package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/bdjuno/x/liquidity"
)

type MatchType int


type MatchResult struct {
	OrderExpiryHeight      int64
	OrderMsgIndex          uint64
	OrderPrice             sdk.Dec
	OfferCoinAmt           sdk.Dec
	TransactedCoinAmt      sdk.Dec
	ExchangedDemandCoinAmt sdk.Dec
	OfferCoinFeeAmt        sdk.Dec
	ExchangedCoinFeeAmt    sdk.Dec
	BatchMsg               *SwapMsgState
}

type Order struct {
	Price         sdk.Dec
	BuyOfferAmt   sdk.Int
	SellOfferAmt  sdk.Int
	SwapMsgStates []*SwapMsgState
}

type Pool struct {
	Id                     uint64         // index of this liquidity pool
	TypeId                 uint32         // pool type of this liquidity pool
	ReserveCoinDenoms      []string       // list of reserve coin denoms for this liquidity pool
	ReserveAccountAddress  string         // reserve account address for this liquidity pool to store reserve coins
	PoolCoinDenom          string         // denom of pool coin for this liquidity pool
}

type DepositMsgState struct {
	MsgHeight  uint64 // height where this message is appended to the batch
	MsgIndex   uint64 // index of this deposit message in this liquidity pool
	Executed   bool   // true if executed on this batch, false if not executed yet
	Succeeded  bool   // true if executed successfully on this batch, false if failed
	ToBeDelete bool   // true if ready to be deleted on kvstore, false if not ready to be deleted
	Msg        MsgDepositWithinBatch
}

type WithdrawMsgState struct {
	MsgHeight  uint64 // height where this message is appended to the batch
	MsgIndex   uint64 // index of this withdraw message in this liquidity pool
	Executed   bool   // true if executed on this batch, false if not executed yet
	Succeeded  bool   // true if executed successfully on this batch, false if failed
	ToBeDelete bool   // true if ready to be deleted on kvstore, false if not ready to be deleted
	Msg        MsgWithdrawWithinBatch
}

type SwapMsgState struct {
	MsgHeight          uint64 // height where this message is appended to the batch
	MsgIndex           uint64 // index of this swap message in this liquidity pool
	Executed           bool   // true if executed on this batch, false if not executed yet
	Succeeded          bool   // true if executed successfully on this batch, false if failed
	ToBeDelete         bool   // true if ready to be deleted on kvstore, false if not ready to be deleted
	OrderExpiryHeight  int64  // swap orders are cancelled when current height is equal or higher than ExpiryHeight
	ExchangedOfferCoin sdk.Coin // offer coin exchanged until now
	RemainingOfferCoin sdk.Coin // offer coin currently remaining to be exchanged
	Msg                MsgSwapWithinBatch
}

type BatchResult struct {
	MatchType      MatchType
	PriceDirection PriceDirection
	SwapPrice      sdk.Dec
	EX             sdk.Dec
	EY             sdk.Dec
	OriginalEX     sdk.Int
	OriginalEY     sdk.Int
	PoolX          sdk.Dec
	PoolY          sdk.Dec
	TransactAmt    sdk.Dec
}


