/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		port := viper.GetInt("port")
		if verbose {
			fmt.Printf("Starting server on port %d\n", port)
			fmt.Printf("Verbose mode enabled\n")
			fmt.Println("Configuration validated successfully")
		} else {
			fmt.Printf("Starting server on port %d\n", port)
		}
		// in real app, you would start server here
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	// define local flag for serve command
	// flag, shorthand, default value, description
	serveCmd.Flags().IntP("port", "p", 8080, "port to run the server on") // Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
