/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

// THE PATH TO CFG FILE SET HERE

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobruh",
	Short: "demo app for cobra and viper",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initConfig(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobruh.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig(cmd *cobra.Command) error {
	// allow viper to use env vars
	viper.SetEnvPrefix("cobruh")
	// allow for nested key in env vars, e.g. COBRUH_SERVER_PORT=8080
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()

	// handle config file
	if cfgFile != "" {
		// use config file from flag
		viper.SetConfigFile(cfgFile)
	} else {
		// find home directory elsewhere
		home, err := os.UserHomeDir()
		// panic if we can't find home dir
		cobra.CheckErr(err)

		// search for config in home dir with name ".cobruh" (w/o extensions)
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + "./cobruh")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// read config file if found
	// use robust error check to handle file not found errors, panic on other errors
	if err := viper.ReadInConfig(); err != nil {
		// its ok if config file does not exist
		var configFileNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFound) {
			return err
		}
	}

	// bind cobra flags to viper
	// makes flag values available to through viper
	// flag set of the command that is passed in is binded
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	//useful debugging info
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	return nil
}
