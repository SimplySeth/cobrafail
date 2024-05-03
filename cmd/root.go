package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ihsnow/cmd/gcpcmd"
	"ihsnow/cmd/sshcmd"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const VERSION = "0.0.1"

type Config struct {
	User     string
	Instance string
	Password string
}

var (
	ConfigObj  Config
	appVersion bool
	csvOutput  bool
	inActive   bool
	jsonOutput bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ihsnow",
	Short: "A command line toolffor interaction with Intuit's ServiceNow instance",
	Long: `ihsnow gives us the superpower to interact with access requests in ServiceNow.
  With this tool, we can:
   - annotate requests
   - approve requests
   - assign tasks
   - close tasks
   - get details of a request
   - get details of a task
   - list requests
   - list tasks
   - reject requests
   - view tasks `,
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if len(args) < 1 && !version {
			cmd.Help()
		} else {
			if version {
				fmt.Println(VERSION)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setConfigPath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	fullpath := filepath.Join(home, "."+filepath.Base(os.Args[0]))
	return fullpath
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(gcpcmd.GcpCmd)
	RootCmd.AddCommand(sshcmd.SshCmd)
	RootCmd.Flags().BoolVarP(&appVersion, "version", "V", false, "Print version number.")
	RootCmd.PersistentFlags().BoolVarP(&csvOutput, "csv", "c", false, "CSV output")
	RootCmd.PersistentFlags().BoolVar(&inActive, "inactive", false, "Include inactive records")
	RootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "JSON output")
	RootCmd.PersistentFlags().StringP("task", "t", "", "Task number ie TASK2691066")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fullpath := setConfigPath()
	v := viper.New()

	// we are not setting config from flags
	v.SetDefault("instance", "intuitdev01")
	v.SetDefault("user", "serviceaccount")
	v.SetDefault("password", "svcacctpassword")
	v.SetConfigFile(fullpath)
	v.SetConfigType("toml")

	err := v.ReadInConfig()
	if err != nil {
		v.SafeWriteConfigAs(fullpath)
		log.Fatal("Please populate the config file " + fullpath)
	}
	ConfigObj.User = v.GetString("user")
	ConfigObj.Instance = v.GetString("instance")
	ConfigObj.Password = v.GetString("password")
}
