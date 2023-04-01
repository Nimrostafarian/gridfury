package tibc_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	migratetibc "github.com/gridnet/gridhub/migrate/tibc"
	"github.com/gridnet/gridhub/simapp"
)

func TestLoadClient(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1})
	migratetibc.CreateClient(ctx, app.AppCodec(), "v1.3", app.TIBCKeeper.ClientKeeper)
}
