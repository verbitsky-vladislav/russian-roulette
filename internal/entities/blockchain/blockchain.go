package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EthereumReader interface {
	ActiveTgGroups(ctx context.Context, index int) (int64, error)
	BettingToken(ctx context.Context) (common.Address, error)
	BurnBps(ctx context.Context) (*big.Int, error)
	Games(ctx context.Context, tgChatId int64) (GameData, error)
	IsGameInProgress(ctx context.Context, tgChatId int64) (bool, error)
	MinimumBet(ctx context.Context) (*big.Int, error)
	Owner(ctx context.Context) (common.Address, error)
	RevenueBps(ctx context.Context) (*big.Int, error)
	RevenueWallet(ctx context.Context) (common.Address, error)
}

type EthereumWriter interface {
	AbortAllGames(ctx context.Context) error
	AbortGame(ctx context.Context, tgChatId int64) error
	EndGame(ctx context.Context, tgChatId int64, loser uint16, data []string) error
	NewGame(ctx context.Context, tgChatId int64, revolverSize int, minBet int, hashedBulletChamberIndex [32]byte, players []common.Address, bets []*big.Int) ([]*big.Int, error)
	RenounceOwnership(ctx context.Context) error
	TransferOwnership(ctx context.Context, newOwner common.Address) error
}

type GameData struct {
	RevolverSize             *big.Int
	MinBet                   *big.Int
	HashedBulletChamberIndex [32]byte
	InProgress               bool
	Loser                    uint16
}
