package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/181192/ops-cli/pkg/cmd/cmdutils"
	"github.com/181192/ops-cli/pkg/cmd/completion"
	"github.com/181192/ops-cli/pkg/cmd/dashboard"
	"github.com/181192/ops-cli/pkg/cmd/version"
	cliConfig "github.com/181192/ops-cli/pkg/config"

	logger "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var loglevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd *cobra.Command

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.EnableCommandSorting = false
	flagGrouping := cmdutils.NewGrouping()
	rootCmd = dashboard.Command(flagGrouping)

	cliConfig.CfgName = "dashboard"

	cobra.OnInitialize(cliConfig.InitConfig)
	rootCmd.AddCommand(completion.Command(rootCmd))
	rootCmd.AddCommand(version.Command(flagGrouping))

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := setUpLogs(os.Stdout, loglevel); err != nil {
			return err
		}
		return nil
	}

	rootCmd.PersistentFlags().StringVar(&loglevel, "log-level", logger.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
	rootCmd.PersistentFlags().StringVar(&cliConfig.CfgFile, "config", "", fmt.Sprintf("Config file (default is %s/%s.[yaml|json|toml|properties])", cliConfig.GetConfigDirectory(), cliConfig.CfgName))

	rootCmd.SetUsageFunc(flagGrouping.Usage)

}

func setUpLogs(out io.Writer, level string) error {
	logger.SetOutput(out)
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})
	lvl, err := logger.ParseLevel(level)
	if err != nil {
		return err
	}
	logger.SetLevel(lvl)
	return nil
}
