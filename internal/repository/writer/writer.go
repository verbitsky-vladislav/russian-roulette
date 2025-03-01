package writer

import (
	"context"
	"math/big"
	"russian-roulette/bindings"
	"russian-roulette/internal/entities/blockchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ blockchain.EthereumWriter = (*EthereumWriterImpl)(nil)

type EthereumWriterImpl struct {
	contract *bindings.BindingsTransactor
	auth     *bind.TransactOpts
}

func NewEthereumWriter(rpcURL string, contractAddressStr string, privateKey string) (blockchain.EthereumWriter, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(contractAddressStr)

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	// Укажите правильный chainID для вашей сети, например 1 для Mainnet
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	auth.GasLimit = uint64(3000000)

	boundContract, err := bindings.NewBindingsTransactor(contractAddress, client)
	if err != nil {
		return nil, err
	}

	return &EthereumWriterImpl{
		contract: boundContract,
		auth:     auth,
	}, nil
}

func (e *EthereumWriterImpl) AbortAllGames(ctx context.Context) error {
	tx, err := e.contract.AbortAllGames(e.auth)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(ctx, nil, tx)
	return err
}

func (e *EthereumWriterImpl) AbortGame(ctx context.Context, tgChatId int64) error {
	tx, err := e.contract.AbortGame(e.auth, tgChatId)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(ctx, nil, tx)
	return err
}

func (e *EthereumWriterImpl) EndGame(ctx context.Context, tgChatId int64, loser uint16, data []string) error {
	tx, err := e.contract.EndGame(e.auth, tgChatId, loser, data)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(ctx, nil, tx)
	return err
}

func (e *EthereumWriterImpl) NewGame(ctx context.Context, tgChatId int64, revolverSize int, minBet int, hashedBulletChamberIndex [32]byte, players []common.Address, bets []*big.Int) ([]*big.Int, error) {
	tx, err := e.contract.NewGame(e.auth, tgChatId, big.NewInt(int64(revolverSize)), big.NewInt(int64(minBet)), hashedBulletChamberIndex, players, bets)
	if err != nil {
		return nil, err
	}
	_, err = bind.WaitMined(ctx, nil, tx)
	if err != nil {
		return nil, err
	}
	return bets, nil
}

func (e *EthereumWriterImpl) RenounceOwnership(ctx context.Context) error {
	tx, err := e.contract.RenounceOwnership(e.auth)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(ctx, nil, tx)
	return err
}

func (e *EthereumWriterImpl) TransferOwnership(ctx context.Context, newOwner common.Address) error {
	tx, err := e.contract.TransferOwnership(e.auth, newOwner)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(ctx, nil, tx)
	return err
}
