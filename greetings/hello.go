package hello

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
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
	cmd.Flags().StringP("name", "n", "", "name to greet")
	return cmd
}
