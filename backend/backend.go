package backend

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type FlatCallFrame struct {
	Action              FlatCallAction  `json:"action"`
	BlockHash           *common.Hash    `json:"blockHash"`
	BlockNumber         uint64          `json:"blockNumber"`
	Error               string          `json:"error,omitempty"`
	Result              *FlatCallResult `json:"result,omitempty"`
	Subtraces           int             `json:"subtraces"`
	TraceAddress        []int           `json:"traceAddress"`
	TransactionHash     *common.Hash    `json:"transactionHash"`
	TransactionPosition uint64          `json:"transactionPosition"`
	Type                string          `json:"type"`
}

type FlatCallAction struct {
	Author         *common.Address `json:"author,omitempty"`
	RewardType     string          `json:"rewardType,omitempty"`
	SelfDestructed *common.Address `json:"address,omitempty"`
	Balance        *big.Int        `json:"balance,omitempty"`
	CallType       string          `json:"callType,omitempty"`
	CreationMethod string          `json:"creationMethod,omitempty"`
	From           *common.Address `json:"from,omitempty"`
	Gas            *uint64         `json:"gas,omitempty"`
	Init           *[]byte         `json:"init,omitempty"`
	Input          *[]byte         `json:"input,omitempty"`
	RefundAddress  *common.Address `json:"refundAddress,omitempty"`
	To             *common.Address `json:"to,omitempty"`
	Value          *big.Int        `json:"value,omitempty"`
}

type FlatCallResult struct {
	Address *common.Address `json:"address,omitempty"`
	Code    *[]byte         `json:"code,omitempty"`
	GasUsed *uint64         `json:"gasUsed,omitempty"`
	Output  *[]byte         `json:"output,omitempty"`
}

// Backend interface provides the common API services
type Backend interface {
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error)
	GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
	TraceBlock(ctx context.Context, number rpc.BlockNumber) ([]FlatCallFrame, error)
	TraceTransaction(ctx context.Context, txHash common.Hash) ([]FlatCallFrame, error)
}
