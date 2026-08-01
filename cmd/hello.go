package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "prints hello [name] or world if no name",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := "world"
		if len(args) > 0 {
			name = args[0]
		}
		fmt.Printf("Hello %s!\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringP("name", "n", "", "name to greet")
}
