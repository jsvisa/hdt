package main

import (
	"net/http"

	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"

	"github.com/jsvisa/hdt/pkg/db"
	"github.com/jsvisa/hdt/pkg/handlers"
)

var app = cli.NewApp()
var (
	upstreamDBDSNFlag = &cli.StringFlag{
		Name:    "upstream.dbdsn",
		Usage:   "upstream PostgreSQL connection DSN",
		Value:   "postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable",
		EnvVars: []string{"UPSTREAM_DBDSN"},
	}
	slackWebhookURLFlag = &cli.StringFlag{
		Name:    "notify.slack.webhook",
		Usage:   "Send an notify into Slack via this webhook",
		EnvVars: []string{"SLACK_WEBHOOK"},
	}
	slackChannelFlag = &cli.StringFlag{
		Name:    "notify.slack.channel",
		Usage:   "Send a notify into this Slack channel",
		EnvVars: []string{"SLACK_CHANNEL"},
	}
	slackSeverityFlag = &cli.StringFlag{
		Name:  "notify.slack.severity",
		Usage: "Slack severity threshold, a comma split list",
		Value: "HIGH,CRITICAL",
	}
)

func init() {
	// Initialize the CLI app and start web server
	app.Action = run
	app.Flags = []cli.Flag{
		utils.HTTPListenAddrFlag,
		utils.HTTPPortFlag,
		upstreamDBDSNFlag,
		slackWebhookURLFlag,
		slackChannelFlag,
		slackSeverityFlag,
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

	DB := db.Init(ctx.String(upstreamDBDSNFlag.Name))
	h := handlers.New(DB, ctx.String(slackWebhookURLFlag.Name), ctx.String(slackChannelFlag.Name), ctx.String(slackSeverityFlag.Name))
	router := mux.NewRouter()

	router.HandleFunc("/webhook/alerts", h.AddAlert).Methods(http.MethodPost)

	addr := fmt.Sprintf("%s:%d", ctx.String(utils.HTTPListenAddrFlag.Name), ctx.Int(utils.HTTPPortFlag.Name))
	log.Info("API is running!", "listen", addr)
	http.ListenAndServe(addr, router)
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
