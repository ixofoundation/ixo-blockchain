module github.com/ixofoundation/ixo-blockchain

go 1.15

require (
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.40.0 // latest
	github.com/cosmos/go-bip39 v1.0.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.1 //latest
	github.com/tendermint/tm-db v0.6.3
	golang.org/x/crypto v0.0.0-20210503195802-e9a32991a82e
	google.golang.org/genproto v0.0.0-20210503173045-b96a97608f20
	google.golang.org/grpc v1.36.1
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
