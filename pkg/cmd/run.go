package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	cm "ULZGameDuelService/pkg/common"
	Cf "ULZGameDuelService/pkg/config"

	wbs "ULZGameDuelService/pkg/serverCtl"
	// wb "ULZGameDuelService/pkg/serverCtlNoRedis"

	"github.com/spf13/cobra"
)

var runCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
	noRedis      bool
}{}

var runCmd = &cobra.Command{
	Use:   "start",
	Short: "start the server of grpc server",
	Long:  `grpc server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc server  Generator v0.9 -- HEAD")
		if len(args) > 0 {
			fmt.Println(args)
		}
		// , _ = os.Getwd()
		fmt.Println(runCMDInput.cfPath)

		var configPoint *Cf.ConfTmp
		var err error
		if strings.Contains(runCMDInput.cfPath, ".toml") {
			configPoint, err = Cf.OpenToml(runCMDInput.cfPath)
		} else if strings.Contains(runCMDInput.cfPath, ".yaml") {
			configPoint, err = Cf.OpenYaml(runCMDInput.cfPath)
		}
		fmt.Printf("%#v\n", configPoint)
		log.Println(runCMDInput.mode)

		if err == nil {
			// cm.
			// Wb.ServerMainProcess(configPoint, callPath, runCMDInput.mode)
			if runCMDInput.mode == "dev" || runCMDInput.mode == "test" {
				cm.DebugTestRun = true
			}
			cm.Mode = runCMDInput.mode
			wbs.ServerMainProcess(configPoint)
		} else {
			panic(err)
		}
	},
}

func init() {
	callPath, _ := os.Getwd()

	runCmd.Flags().StringVarP(
		&runCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start server with specific config file")

	runCmd.Flags().StringVarP(
		&runCMDInput.mode,
		"mode", "m",
		"prod",
		"server running mode [prod / dev / test]")

	runCmd.Flags().BoolVarP(
		&runCMDInput.noRedis,
		"no-redis", "R",
		false,
		"server run without redis")

	rootCmd.AddCommand(runCmd)
}
