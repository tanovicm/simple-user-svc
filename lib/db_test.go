package lib

import (
	"testing"

	"github.com/Kamva/mgm"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestSetupDefaultConnection(t *testing.T) {
	err := mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("mongodb://localhost:27017"))

	require.Nil(t, err)
}

func TestSetupWrongConnection(t *testing.T) {
	err := mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("wrong://wrong:wrong@localhost:27017"))

	require.NotNil(t, err)
}

func TestPanicOnGetCtx(t *testing.T) {
	mgm.ResetDefaultConfig()

	defer func() {
		require.NotNil(t, recover(), "Getting context before set default config must panic")
	}()

	_ = mgm.Ctx()
}
