package liquidity

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/desmos-labs/juno/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"google.golang.org/grpc"

	"github.com/desmos-labs/juno/modules"
	"github.com/faddat/bdjuno/database"
	lmtypes "github.com/tendermint/liquidity/x/liquidity/types"
)

// Module represent x/liquidity module
type Module struct {
	encodingConfig *params.EncodingConfig
	lmClient       lmtypes.QueryClient
	authClient     authtypes.QueryClient
	bankClient     banktypes.QueryClient
	db             *database.BigDipperDb
}

// NewModule returns a new Module instance
func NewModule(encodingConfig *params.EncodingConfig, grpcConnection *grpc.ClientConn, db *database.BigDipperDb) *Module {
	return &Module{
		encodingConfig: encodingConfig,
		lmClient:       lmtypes.NewQueryClient(grpcConnection),
		authClient:     authtypes.NewQueryClient(grpcConnection),
		bankClient:     banktypes.NewQueryClient(grpcConnection),
		db:             db,
	}
}

var _ modules.Module = &Module{}

// Name implements modules.Module (returns name)
func (m *Module) Name() string {
	return "liquidity"
}

// RunAdditionalOperations implements modules.Module
func (m *Module) RunAdditionalOperations() error {
	return nil
}

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(_ *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	return HandleGenesis(appState, m.encodingConfig.Marshaler, m.lmClient, m.db)
}

// HandleBlock implements modules.Module
func (m *Module) HandleBlock(*tmctypes.ResultBlock, []*types.Tx, *tmctypes.ResultValidators) error {
	return nil
}

// HandleTx implements modules.Module
func (m *Module) HandleTx(*types.Tx) error {
	return nil
}

/*
// HandleMsg implements modules.Module
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *lmtypes.Tx) error {
	return HandleMsg(tx, msg, m.lmClient, m.authClient, m.bankClient, m.encodingConfig.Marshaler, m.db)
}
*/
