// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethtracers "github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jsvisa/hdt/backend"
)

// API is the collection of tracing APIs exposed over the private debugging endpoint.
type API struct {
	backend backend.Backend
}

// NewAPI creates a new API definition for the tracing methods of the Ethereum service.
func NewAPI(backend backend.Backend) *API {
	return &API{backend: backend}
}

// blockByNumber is the wrapper of the chain access function offered by the backend.
// It will return an error if the block is not found.
func (api *API) blockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	block, err := api.backend.BlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block #%d not found", number)
	}
	return block, nil
}

// TraceBlockByNumber returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
func (api *API) Block(ctx context.Context, number rpc.BlockNumber, config *ethtracers.TraceConfig) ([]*backend.CallFrame, error) {
	return api.backend.TraceBlock(ctx, number)
}

// Transaction returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
func (api *API) Transaction(ctx context.Context, hash common.Hash, config *ethtracers.TraceConfig) ([]*backend.CallFrame, error) {
	return api.backend.TraceTransaction(ctx, hash)
}

// APIs return the collection of RPC services the tracer package offers.
func APIs(backend backend.Backend) []rpc.API {
	// Append all the local APIs and return
	return []rpc.API{
		{
			Namespace: "trace",
			Service:   NewAPI(backend),
		},
	}
}
