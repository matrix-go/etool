package cmd

import (
	"github.com/matrix-go/etool/cmd/gen"
	"os"

	"github.com/spf13/cobra"
)

// root represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:   "etool",
	Short: "A tool for develop software conveniently",
	//	Long: `A tool for develop software conveniently. For example:
	//SubCmd:
	//	- gen: generate code for usually use.
	//	- help: print help information.
	//	- version: print version.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(root *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the root.
func Execute() {
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// sub command
	root.AddCommand(gen.Cmd)
}
