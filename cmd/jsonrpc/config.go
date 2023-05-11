package main

import (
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/jsvisa/jsonrpc/node"
)

type gethConfig struct {
	Eth  ethconfig.Config
	Node node.Config
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.HTTPModules = append(cfg.HTTPModules, "eth")
	cfg.WSModules = append(cfg.WSModules, "eth")
	cfg.IPCPath = "geth.ipc"
	return cfg
}

// loadBaseConfig loads the gethConfig based on the given command line
// parameters and config file.
func loadBaseConfig(ctx *cli.Context) gethConfig {
	// Load defaults.
	cfg := gethConfig{
		Eth:  ethconfig.Defaults,
		Node: defaultNodeConfig(),
	}

	// Apply flags.
	setHTTP(ctx, &cfg.Node)
	return cfg
}

func setHTTP(ctx *cli.Context, cfg *node.Config) {
	if ctx.Bool(utils.HTTPEnabledFlag.Name) && cfg.HTTPHost == "" {
		cfg.HTTPHost = "127.0.0.1"
		if ctx.IsSet(utils.HTTPListenAddrFlag.Name) {
			cfg.HTTPHost = ctx.String(utils.HTTPListenAddrFlag.Name)
		}
	}

	if ctx.IsSet(utils.HTTPPortFlag.Name) {
		cfg.HTTPPort = ctx.Int(utils.HTTPPortFlag.Name)
	}

	if ctx.IsSet(utils.HTTPCORSDomainFlag.Name) {
		cfg.HTTPCors = utils.SplitAndTrim(ctx.String(utils.HTTPCORSDomainFlag.Name))
	}

	if ctx.IsSet(utils.HTTPApiFlag.Name) {
		cfg.HTTPModules = utils.SplitAndTrim(ctx.String(utils.HTTPApiFlag.Name))
	}

	if ctx.IsSet(utils.HTTPVirtualHostsFlag.Name) {
		cfg.HTTPVirtualHosts = utils.SplitAndTrim(ctx.String(utils.HTTPVirtualHostsFlag.Name))
	}

	if ctx.IsSet(utils.HTTPPathPrefixFlag.Name) {
		cfg.HTTPPathPrefix = ctx.String(utils.HTTPPathPrefixFlag.Name)
	}
}
