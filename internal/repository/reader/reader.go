package reader

import (
	"context"
	"math/big"
	"russian-roulette/bindings"
	"russian-roulette/internal/entities/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ blockchain.EthereumReader = (*EthereumReaderImpl)(nil)

type EthereumReaderImpl struct {
	contract *bindings.BindingsCaller
}

func NewEthereumReader(rpcURL string, contractAddressStr string) (blockchain.EthereumReader, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(contractAddressStr) // Конвертация строки в common.Address

	boundContract, err := bindings.NewBindingsCaller(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &EthereumReaderImpl{contract: boundContract}, nil
}

func (e *EthereumReaderImpl) ActiveTgGroups(ctx context.Context, index int) (int64, error) {
	groupID, err := e.contract.ActiveTgGroups(nil, big.NewInt(int64(index)))
	if err != nil {
		return 0, err
	}
	return groupID, nil
}

func (e *EthereumReaderImpl) BettingToken(ctx context.Context) (common.Address, error) {
	return e.contract.BettingToken(nil)
}

func (e *EthereumReaderImpl) BurnBps(ctx context.Context) (*big.Int, error) {
	return e.contract.BurnBps(nil)
}

func (e *EthereumReaderImpl) Games(ctx context.Context, tgChatId int64) (blockchain.GameData, error) {
	game, err := e.contract.Games(nil, tgChatId)
	if err != nil {
		return blockchain.GameData{}, err
	}
	return blockchain.GameData{
		RevolverSize:             game.RevolverSize,
		MinBet:                   game.MinBet,
		HashedBulletChamberIndex: game.HashedBulletChamberIndex,
		InProgress:               game.InProgress,
		Loser:                    game.Loser,
	}, nil
}

func (e *EthereumReaderImpl) IsGameInProgress(ctx context.Context, tgChatId int64) (bool, error) {
	return e.contract.IsGameInProgress(nil, tgChatId)
}

func (e *EthereumReaderImpl) MinimumBet(ctx context.Context) (*big.Int, error) {
	return e.contract.MinimumBet(nil)
}

func (e *EthereumReaderImpl) Owner(ctx context.Context) (common.Address, error) {
	return e.contract.Owner(nil)
}

func (e *EthereumReaderImpl) RevenueBps(ctx context.Context) (*big.Int, error) {
	return e.contract.RevenueBps(nil)
}

func (e *EthereumReaderImpl) RevenueWallet(ctx context.Context) (common.Address, error) {
	return e.contract.RevenueWallet(nil)
}
