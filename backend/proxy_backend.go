package backend

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mixinBackend struct {
	ec *ethclient.Client
	db *gorm.DB
}

func NewMixinBackend(ctx context.Context, rawurl string, dbdsn string) (*mixinBackend, error) {
	ec, err := ethclient.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dbdsn, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &mixinBackend{
		ec: ec,
		db: db,
	}, nil
}

func (p *mixinBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	return nil, nil
}
func (p *mixinBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	return nil, nil
}
func (p *mixinBackend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	return nil, common.Hash{}, 0, 0, nil
}
func (p *mixinBackend) TraceBlock(ctx context.Context, number rpc.BlockNumber) ([]FlatCallFrame, error) {
	return nil, nil
}
func (p *mixinBackend) TraceTransaction(ctx context.Context, txHash common.Hash) ([]FlatCallFrame, error) {
	return nil, nil
}
