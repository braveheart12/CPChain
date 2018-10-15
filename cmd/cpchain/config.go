package main

import (
	"os"
	"path/filepath"

	"bitbucket.org/cpchain/chain/accounts"
	"bitbucket.org/cpchain/chain/accounts/keystore"
	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/core"
	"bitbucket.org/cpchain/chain/eth"
	"bitbucket.org/cpchain/chain/node"
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/urfave/cli"
)

type config struct {
	Eth  eth.Config
	Node node.Config
}

// begin node configs ********************************************************************88

func updateDataDirFlag(ctx *cli.Context, cfg *node.Config) {
	if ctx.IsSet("datadir") {
		cfg.DataDir = ctx.String("datadir")
	}
}

func updateNodeGeneralConfig(ctx *cli.Context, cfg *node.Config) {
	updateDataDirFlag(ctx, cfg)
	// update identity
	if ctx.IsSet("identity") {
		cfg.UserIdent = ctx.String("identity")
	}
}

func updateP2pConfig(ctx *cli.Context, cfg *p2p.Config) {
}

// TODO @sangh
func updateRpcConfig(ctx *cli.Context, cfg *node.Config) {
}

func updateNodeConfig(ctx *cli.Context, cfg *node.Config) {
	updateNodeGeneralConfig(ctx, cfg)
	updateP2pConfig(ctx, &cfg.P2P)
	updateRpcConfig(ctx, cfg)
}

// begin chain configs ********************************************************************88

// Updates the account for cfg.Etherbase
func updateBaseAccount(ctx *cli.Context, ks *keystore.KeyStore, cfg *eth.Config) {
	if ctx.IsSet("account") {
		val := ctx.String("account")
		if !common.IsHexAddress(val) {
			log.Fatalf("Invalid account hex address: %v", val)
		}
		account := accounts.Account{Address: common.HexToAddress(val)}
		cfg.Etherbase = account.Address
	} else {
		// fall back on the first account
		accs := ks.Accounts()
		if len(accs) > 0 {
			account := accs[0].Address
			cfg.Etherbase = account
			log.Warnf("Use account %v as the default account.", account.String())
		} else {
			log.Warn("No default account to use.")
		}
	}
}

// Updates transaction pool configurations
func updateTxPool(ctx *cli.Context, cfg *core.TxPoolConfig) {

}

// Updates chain configuration
func updateChainConfig(ctx *cli.Context, cfg *eth.Config, n *node.Node) {
	ks := n.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	updateBaseAccount(ctx, ks, cfg)
	// setGPO(ctx, &cfg.GPO)
	updateTxPool(ctx, &cfg.TxPool)

	// network id setup
	if ctx.IsSet("networkid") {
		cfg.NetworkId = ctx.Uint64("networkid")
	}
}

// Updates config from --config file
func updateConfigFromFile(ctx *cli.Context, cfg *config) {
	var path string
	if ctx.GlobalIsSet("config") {
		p := ctx.GlobalString("config")
		if _, err := os.Stat(p); os.IsNotExist(err) {
			log.Fatalf("Config file doesn't exist: %v", p)
		}
		path = p
	} else {
		// try to read from the datadir/config.toml
		p := filepath.Join(cfg.Node.DataDir, "config.toml")
		if _, err := os.Stat(p); !os.IsNotExist(err) {
			path = p
		}
	}

	if path != "" {
		log.Infof("Load config file from: %v", path)
		if _, err := toml.DecodeFile(path, &cfg); err != nil {
			log.Fatalf("Invalid TOML config file: %v", err)
		}
	}
}

// Creates a config and a node
func newConfigNode(ctx *cli.Context) (config, *node.Node) {
	// default
	cfg := config{
		Eth:  eth.DefaultConfig,
		Node: node.DefaultConfig,
	}
	updateConfigFromFile(ctx, &cfg)

	// now update from command line arguments
	updateNodeConfig(ctx, &cfg.Node)
	// create node
	n, err := node.New(&cfg.Node)
	if err != nil {
		log.Fatalf("Node creation failed: %v", err)
	}
	// update chain config
	updateChainConfig(ctx, &cfg.Eth, n)
	return cfg, n
}
