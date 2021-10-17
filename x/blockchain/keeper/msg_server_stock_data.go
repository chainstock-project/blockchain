package keeper

import (
	"context"
	"fmt"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStockData(goCtx context.Context, msg *types.MsgCreateStockData) (*types.MsgCreateStockDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetStockData(ctx, msg.Date)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("date %v already set", msg.Date))
	}

	root_address := "cosmos1s3pzgpduvnq4r59mjx0vmdzfttqkhywwj7f8lk"
	if root_address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("creator not root %+v", msg.Creator))
	}

	var stockData = types.StockData{
		Date:    msg.Date,
		Creator: msg.Creator,
		Stocks:  msg.Stocks,
	}

	k.SetStockData(
		ctx,
		stockData,
	)
	return &types.MsgCreateStockDataResponse{}, nil
}

func (k msgServer) UpdateStockData(goCtx context.Context, msg *types.MsgUpdateStockData) (*types.MsgUpdateStockDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStockData(ctx, msg.Date)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("date %v not set", msg.Date))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var stockData = types.StockData{
		Date:    msg.Date,
		Creator: msg.Creator,
		Stocks:  msg.Stocks,
	}

	k.SetStockData(ctx, stockData)

	return &types.MsgUpdateStockDataResponse{}, nil
}

func (k msgServer) DeleteStockData(goCtx context.Context, msg *types.MsgDeleteStockData) (*types.MsgDeleteStockDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStockData(ctx, msg.Date)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("date %v not set", msg.Date))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStockData(ctx, msg.Date)

	return &types.MsgDeleteStockDataResponse{}, nil
}
