// Copyright 2017 The go-ethereum Authors
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

package dpor

import (
	"bitbucket.org/cpchain/chain/consensus"
	"bitbucket.org/cpchain/chain/rpc"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-authority scheme.
type API struct {
	chain consensus.ChainReader
	dpor  *Dpor
	dh    *defaultDporHelper
}

// GetSnapshot retrieves the state Snapshot at a given block.
func (api *API) GetSnapshot(number *rpc.BlockNumber) (*DporSnapshot, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return its Snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.dpor.dh.snapshot(api.dpor, api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetSnapshotAtHash retrieves the state Snapshot at a given block.
func (api *API) GetSnapshotAtHash(hash common.Hash) (*DporSnapshot, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.dpor.dh.snapshot(api.dpor, api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetSigners retrieves the list of authorized signers at the specified block.
func (api *API) GetSigners(number *rpc.BlockNumber) ([]common.Address, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return the signers from its Snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.dpor.dh.snapshot(api.dpor, api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}

	return snap.signersOf(header.Number.Uint64()), nil
}

// GetSignersAtHash retrieves the state Snapshot at a given block.
func (api *API) GetSignersAtHash(hash common.Hash) ([]common.Address, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	number := header.Number.Uint64()
	snap, err := api.dpor.dh.snapshot(api.dpor, api.chain, number, hash, nil)
	if err != nil {
		return nil, err
	}
	return snap.signersOf(number), nil
}
