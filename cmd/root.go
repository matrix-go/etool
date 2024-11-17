/*
Package root
Copyright © 2024 chaorendexiaoneiku@icloud.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.etool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//root.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// sub command
	root.AddCommand(gen.Cmd)
}
