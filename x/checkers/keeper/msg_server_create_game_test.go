package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/checkers"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "1",
	},
		*createResponse)
}

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreate1GameHasSaved(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId: 2,
	}, systemInfo)

	game1, found1 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")

	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, game1)

}

func TestCreate3Game(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "1",
	}, *createResponse)

	createResponse2, err2 := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err2)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "2",
	}, *createResponse2)

}

func TestCreateGameRedAddressBad(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     "badaddress",
	})
	require.Nil(t, createResponse)
	require.NotNil(t, err)
	require.Equal(t,
		"red address is invalid: badaddress: decoding bech32 failed: invalid separator index -1",
		err.Error())
}

func TestCreateGameEmptyAddress(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     "",
	})
	require.Nil(t, createResponse)
	require.NotNil(t, err)
	require.Equal(t,
		"red address is invalid: : empty address string is not allowed",
		err.Error())
}

func TestCreate1GameEmitted(t *testing.T) {
	msg, _, context := setupMsgServerCreateGame(t)
	msg.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "new-game-created",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: alice},
			{Key: "game-index", Value: "1"},
			{Key: "black", Value: bob},
			{Key: "red", Value: carol},
		},
	}, event)
}
