package backend

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// Backend interface provides the common API services
type Backend interface {
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error)
	BlockTimestamp(ctx context.Context, number rpc.BlockNumber) (uint64, error)
	TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, uint64, uint64, error)
	TraceBlock(ctx context.Context, number rpc.BlockNumber) ([]*CallFrame, error)
	TraceTransaction(ctx context.Context, txHash common.Hash) ([]*CallFrame, error)
}
