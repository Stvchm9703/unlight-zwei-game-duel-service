package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	Cf "ULZGameDuelService/pkg/config"

	"github.com/spf13/cobra"
)

var initCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
	noRedis      bool
}{}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init the server config file",
	Long:  `write the server config file, for the startup`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("grpc server Generator v0.9 -- HEAD")
		if len(args) > 0 {
			fmt.Println(args)
		}

		var configPoint *Cf.ConfTmp

		// configPoint
		if strings.Contains(runCMDInput.cfPath, ".yaml") {
			Cf.CreateConfigYaml(runCMDInput.cfPath, configPoint)
		}

	},
}

func init() {
	callPath, _ := os.Getwd()

	initCmd.Flags().StringVarP(
		&initCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start server with specific config file")

	rootCmd.AddCommand(initCmd)
}
