package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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


