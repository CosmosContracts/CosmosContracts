package interchaintest

import (
	feesharetypes "github.com/CosmosContracts/juno/v13/x/feeshare/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
)

var (
	junoRepo, junoVersion = GetDockerImageInfo()

	junoImage = ibc.DockerImage{
		Repository: junoRepo,
		Version:    junoVersion,
		UidGid:     "1025:1025",
	}

	junoConfig = ibc.ChainConfig{
		Type:                "cosmos",
		Name:                "juno",
		ChainID:             "juno-2",
		Images:              []ibc.DockerImage{junoImage},
		Bin:                 "junod",
		Bech32Prefix:        "juno",
		Denom:               "ujuno",
		CoinType:            "118",
		GasPrices:           "0.003ujuno",
		GasAdjustment:       1.1,
		TrustingPeriod:      "112h",
		NoHostMount:         false,
		SkipGenTx:           false,
		PreGenesis:          nil,
		ModifyGenesis:       nil,
		ConfigFileOverrides: nil,
		EncodingConfig:      junoEncoding(),
	}

	pathJunoGaia        = "juno-gaia"
	genesisWalletAmount = int64(10_000_000)
)

// junoEncoding registers the Juno specific module codecs so that the associated types and msgs
// will be supported when writing to the blocksdb sqlite database.
func junoEncoding() *simappparams.EncodingConfig {
	cfg := cosmos.DefaultEncoding()

	// register custom types
	feesharetypes.RegisterInterfaces(cfg.InterfaceRegistry)

	return &cfg
}
