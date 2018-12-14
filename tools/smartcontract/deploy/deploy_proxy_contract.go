// Copyright 2018 The cpchain authors
// This file is part of the cpchain library.
//
// The cpchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cpchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cpchain library. If not, see <http://www.gnu.org/licenses/>.

package deploy

import (
	"context"
	"fmt"
	"math/big"

	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/contracts/proxy/proxy_contract"
	"bitbucket.org/cpchain/chain/tools/smartcontract/config"
	"github.com/ethereum/go-ethereum/common"
)

func ProxyContractRegister(password string) common.Address {
	client, err, privateKey, _, fromAddress := config.Connect(password)
	printBalance(client, fromAddress)
	// Launch contract deploy transaction.
	auth := newAuth(client, privateKey, fromAddress)
	contractAddress, tx, _, err := contract.DeployProxyContractRegister(auth, client)
	if err != nil {
		log.Fatal(err.Error())
	}
	printTx(tx, err, client, contractAddress)
	return contractAddress
}

func DeployProxy(password string) common.Address {
	client, err, privateKey, _, fromAddress := config.Connect(password)
	printBalance(client, fromAddress)
	// Launch contract deploy transaction.
	auth := bind.NewKeyedTransactor(privateKey)
	contractAddress, tx, _, err := contract.DeployProxy(auth, client)
	if err != nil {
		log.Fatal(err.Error())
	}
	printTx(tx, err, client, contractAddress)
	return contractAddress
}

func RegisterProxyAddress(proxyContractAddress, realAddress common.Address, password string) common.Address {
	proxyAddress := DeployProxy(password)
	success := UpdateRegisterProxyAddress(proxyContractAddress, proxyAddress, realAddress, password)
	if success {
		return proxyAddress
	} else {
		return common.Address{}
	}
}

func UpdateRegisterProxyAddress(proxyContractAddress, proxyAddress, realAddress common.Address, password string) bool {
	FormatPrint("register proxy address")

	PrintContract(proxyAddress)
	fmt.Println("Register proxy contract proxy -> real:" + realAddress.Hex())
	client, err, privateKey, _, fromAddress := config.Connect(password)
	if err != nil {
		fmt.Println("failed")
		log.Fatal(err.Error())
		return false

	}
	printBalance(client, fromAddress)
	proxyContractRegister, _ := contract.NewProxyContractRegister(proxyContractAddress, client)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(500)
	auth.GasLimit = 3000000
	gasPrice, err := client.SuggestGasPrice(context.Background())
	// fmt.Println("gasPrice:", gasPrice)
	auth.GasPrice = gasPrice

	transaction, err := proxyContractRegister.RegisterProxyContract(auth, proxyAddress, realAddress)
	if err != nil {
		fmt.Println("failed")
		log.Fatal(err.Error())
		return false
	}
	receipt, err := bind.WaitMined(context.Background(), client, transaction)
	if err != nil {
		fmt.Println("failed")
		log.Fatalf("failed to deploy contact when mining :%v", err)
		return false
	}
	// fmt.Println("receipt.Status:", receipt.Status)
	if receipt.Status == 1 {
		fmt.Println("success")
		return true
	} else {
		fmt.Println("failed")
		return false
	}
}
