package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/jsvisa/jsonrpc/backend"
	"github.com/jsvisa/jsonrpc/node"
	"github.com/jsvisa/jsonrpc/service/eth"
	"github.com/jsvisa/jsonrpc/service/trace"
)

const (
	clientIdentifier = "jsonrpc" // Client identifier to advertise over the network
)

var app = cli.NewApp()
var (
	chainFlag = &cli.StringFlag{
		Name:  "chain",
		Usage: "chain name",
		Value: "ethereum",
	}
	upstreamJSONRPCFlag = &cli.StringFlag{
		Name:  "upstream.jsonrpc",
		Usage: "upstream JSONRPC HTTP address with port",
		Value: "http://127.0.0.1:8545",
	}
	upstreamDBDSNFlag = &cli.StringFlag{
		Name:  "upstream.dbdsn",
		Usage: "upstream PostgreSQL connection DSN",
		Value: "postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable",
	}
	pprofFlag = &cli.BoolFlag{
		Name:  "pprof",
		Usage: "Enable the pprof HTTP server",
	}
	pprofPortFlag = &cli.IntFlag{
		Name:  "pprof.port",
		Usage: "pprof HTTP server listening port",
		Value: 6060,
	}
	pprofAddrFlag = &cli.StringFlag{
		Name:  "pprof.addr",
		Usage: "pprof HTTP server listening interface",
		Value: "127.0.0.1",
	}
	memprofilerateFlag = &cli.IntFlag{
		Name:  "pprof.memprofilerate",
		Usage: "Turn on memory profiling with the given rate",
		Value: runtime.MemProfileRate,
	}
	blockprofilerateFlag = &cli.IntFlag{
		Name:  "pprof.blockprofilerate",
		Usage: "Turn on block profiling with the given rate",
	}
	cpuprofileFlag = &cli.StringFlag{
		Name:  "pprof.cpuprofile",
		Usage: "Write CPU profile to the given file",
	}
)

func init() {
	// Initialize the CLI app and start web server
	app.Action = run
	app.Flags = []cli.Flag{
		utils.HTTPEnabledFlag,
		utils.HTTPListenAddrFlag,
		utils.HTTPPortFlag,
		utils.HTTPCORSDomainFlag,
		utils.HTTPVirtualHostsFlag,
		utils.HTTPApiFlag,
		utils.HTTPPathPrefixFlag,
		utils.WSEnabledFlag,
		utils.WSListenAddrFlag,
		utils.WSPortFlag,
		utils.WSApiFlag,
		utils.WSAllowedOriginsFlag,
		utils.WSPathPrefixFlag,
		chainFlag,
		upstreamJSONRPCFlag,
		upstreamDBDSNFlag,
		pprofFlag,
		pprofAddrFlag,
		pprofPortFlag,
		memprofilerateFlag,
		blockprofilerateFlag,
		cpuprofileFlag,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run is the main entry point into the system if no special subcommand is run.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
func run(ctx *cli.Context) error {
	if args := ctx.Args().Slice(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}

	cfg := loadBaseConfig(ctx)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		log.Crit("Failed to create the protocol stack", "err", err)
	}

	cctx := context.Background()
	backend, err := backend.NewMixinBackend(
		cctx,
		ctx.String(chainFlag.Name),
		ctx.String(upstreamJSONRPCFlag.Name),
		ctx.String(upstreamDBDSNFlag.Name),
	)
	if err != nil {
		log.Crit("Failed to register the Ethereum service", "err", err)
	}
	stack.RegisterAPIs(trace.APIs(backend))
	stack.RegisterAPIs(eth.APIs(backend))
	defer stack.Close()

	if err := stack.Start(); err != nil {
		log.Crit("Error starting protocol stack", "err", err)
	}
	stack.Wait()
	return nil
}

var (
	glogger *log.GlogHandler
)

func init() {
	glogger = log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogger)
}
