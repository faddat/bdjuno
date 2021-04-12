package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/liquidity/x/liquidity/types"
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
	ID                    uint64   // index of this liquidity pool
	TypeID                uint32   // pool type of this liquidity pool
	ReserveCoinDenoms     []string // list of reserve coin denoms for this liquidity pool
	ReserveAccountAddress string   // reserve account address for this liquidity pool to store reserve coins
	PoolCoinDenom         string   // denom of pool coin for this liquidity pool
}

type DepositMsgState struct {
	MsgHeight  uint64 // height where this message is appended to the batch
	MsgIndex   uint64 // index of this deposit message in this liquidity pool
	Executed   bool   // true if executed on this batch, false if not executed yet
	Succeeded  bool   // true if executed successfully on this batch, false if failed
	ToBeDelete bool   // true if ready to be deleted on kvstore, false if not ready to be deleted
	Msg        types.MsgDepositWithinBatch
}

type WithdrawMsgState struct {
	MsgHeight  uint64 // height where this message is appended to the batch
	MsgIndex   uint64 // index of this withdraw message in this liquidity pool
	Executed   bool   // true if executed on this batch, false if not executed yet
	Succeeded  bool   // true if executed successfully on this batch, false if failed
	ToBeDelete bool   // true if ready to be deleted on kvstore, false if not ready to be deleted
	Msg        types.MsgWithdrawWithinBatch
}

type SwapMsgState struct {
	MsgHeight          uint64   // height where this message is appended to the batch
	MsgIndex           uint64   // index of this swap message in this liquidity pool
	Executed           bool     // true if executed on this batch, false if not executed yet
	Succeeded          bool     // true if executed successfully on this batch, false if failed
	ToBeDelete         bool     // true if ready to be deleted on kvstore, false if not ready to be deleted
	OrderExpiryHeight  int64    // swap orders are cancelled when current height is equal or higher than ExpiryHeight
	ExchangedOfferCoin sdk.Coin // offer coin exchanged until now
	RemainingOfferCoin sdk.Coin // offer coin currently remaining to be exchanged
	Msg                types.MsgSwapWithinBatch
}

type BatchResult struct {
	MatchType      MatchType
	PriceDirection types.PriceDirection
	SwapPrice      sdk.Dec
	EX             sdk.Dec
	EY             sdk.Dec
	OriginalEX     sdk.Int
	OriginalEY     sdk.Int
	PoolX          sdk.Dec
	PoolY          sdk.Dec
	TransactAmt    sdk.Dec
}

/*
type MsgDepositWithinBatch struct {
	DepositorAddress string `protobuf:"bytes,1,opt,name=depositor_address,json=depositorAddress,proto3" json:"depositor_address,omitempty" yaml:"depositor_address"`
	// id of the target pool
	PoolId uint64 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id" yaml:"pool_id"`
	// reserve coin pair of the pool to deposit
	DepositCoins sdk.Coins `protobuf:"bytes,3,rep,name=deposit_coins,json=depositCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"deposit_coins" yaml:"deposit_coins"`
}

type MsgSwapWithinBatch struct {
	// address of swap requester
	SwapRequesterAddress string `protobuf:"bytes,1,opt,name=swap_requester_address,json=swapRequesterAddress,proto3" json:"swap_requester_address,omitempty" yaml:"swap_requester_address"`
	// id of the target pool
	PoolId uint64 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id" yaml:"pool_id"`
	// id of swap type. Must match the value in the pool.
	SwapTypeId uint32 `protobuf:"varint,3,opt,name=swap_type_id,json=swapTypeId,proto3" json:"swap_type_id,omitempty" yaml:"swap_type_id"`
	// offer sdk.coin for the swap request, must match the denom in the pool.
	OfferCoin types.Coin `protobuf:"bytes,4,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin" yaml:"offer_coin"`
	// denom of demand coin to be exchanged on the swap request, must match the denom in the pool.
	DemandCoinDenom string `protobuf:"bytes,5,opt,name=demand_coin_denom,json=demandCoinDenom,proto3" json:"demand_coin_denom,omitempty" yaml:"demand_coin_denom"`
	// half of offer coin amount * params.swap_fee_rate for reservation to pay fees
	OfferCoinFee types.Coin `protobuf:"bytes,6,opt,name=offer_coin_fee,json=offerCoinFee,proto3" json:"offer_coin_fee" yaml:"offer_coin_fee"`
	// limit order price for the order, the price is the exchange ratio of X/Y where X is the amount of the first coin and
	// Y is the amount of the second coin when their denoms are sorted alphabetically
	OrderPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=order_price,json=orderPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"order_price" yaml:"order_price"`
}

func NewMsgDepositWithinBatch(
	depositor sdk.AccAddress,
	poolId uint64,
	depositCoins sdk.Coins,
) *MsgDepositWithinBatch {
	return &MsgDepositWithinBatch{
		DepositorAddress: depositor.String(),
		PoolId:           poolId,
		DepositCoins:     depositCoins,
	}
}
*/
