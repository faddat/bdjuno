module github.com/faddat/bdjuno

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.3
	github.com/desmos-labs/juno v0.0.0-20210218075545-df64a128e096
	github.com/go-co-op/gocron v0.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.2.1-0.20200324155115-ee514944af4b
	github.com/lib/pq v1.9.0
	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
	github.com/rs/zerolog v1.20.0
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/liquidity v1.2.3
	github.com/tendermint/tendermint v0.34.8
	google.golang.org/grpc v1.35.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
