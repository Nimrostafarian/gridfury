package app

import (
	"github.com/gridnet/gridhub/types"
)

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() types.GenesisState {
	encCfg := MakeEncodingConfig()
	return ModuleBasics.DefaultGenesis(encCfg.Marshaler)
}
