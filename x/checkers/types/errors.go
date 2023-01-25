package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)

var (
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1102, "black address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1103, "game cannot be parsed")

	ErrGameNotFound     = sdkerrors.Register(ModuleName, 1104, "game by id not found")
	ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1105, "message creator is not a player")
	ErrNotPlayerTurn    = sdkerrors.Register(ModuleName, 1106, "player tried to play out of turn")
	ErrWrongMove        = sdkerrors.Register(ModuleName, 1107, "wrong move")

	ErrBlackAlreadyPlayed = sdkerrors.Register(ModuleName, 1108, "black player has already played")
	ErrRedAlreadyPlayed   = sdkerrors.Register(ModuleName, 1109, "red player has already played")

	ErrInvalidDeadline = sdkerrors.Register(ModuleName, 1110, "deadline cannot be parsed: %s")

	ErrGameFinished = sdkerrors.Register(ModuleName, 1011, "game is already finished")

	ErrCannotFindWinnerByColor = sdkerrors.Register(ModuleName, 1111, "cannot find winner by color: %s")
)

var (
	ErrBlackCannotPay    = sdkerrors.Register(ModuleName, 1112, "black cannot pay the wager")
	ErrRedCannotPay      = sdkerrors.Register(ModuleName, 1113, "red cannot pay the wager")
	ErrNothingToPay      = sdkerrors.Register(ModuleName, 1114, "there is nothing to pay, should not have been called")
	ErrCannotRefundWager = sdkerrors.Register(ModuleName, 1115, "cannot refund wager to: %s")
	ErrCannotPayWinnings = sdkerrors.Register(ModuleName, 1116, "cannot pay winnings to winner: %s")
	ErrNotInRefundState  = sdkerrors.Register(ModuleName, 1117, "game is not in a state to refund, move count: %d")
)
