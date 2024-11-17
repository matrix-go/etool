/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dao

import (
	"fmt"
	"github.com/matrix-go/etool/internal/gen/dao/parser/mysql"

	"github.com/spf13/cobra"
)

var (
	ddlPath *string
	outPath *string
	prefix  *string
)

// Cmd represents the dao command
var Cmd = &cobra.Command{
	Use:   "dao",
	Short: "gen dao with ddl sql",
	Long: `A tool for generate dao code with ddl sql, for example:

  1. generate dao code with user.sql and write to path internal:
    etool gen dao --in tests/testdata/user.sql --out ./internal
  2. generate dao code with short command:
    etool gen dao -i tests/testdata/user.sql -o ./internal
  3. generate dao code with prefix set:
    etool gen dao -i tests/testdata/user.sql -o ./internal -p t`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := mysql.NewParser(*ddlPath, *prefix)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := p.Write(*outPath); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cmd.PersistentFlags().String("foo", "", "A help for foo")
	//ddlPath = genCmd.PersistentFlags().StringP("in", "i", "", "ddl sql path")
	//outPath = genCmd.PersistentFlags().StringP("out", "o", "stdout", "stdout or file, default stdout")
	//prefix = genCmd.PersistentFlags().StringP("prefix", "p", "t", "prefix of table, default t")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ddlPath = Cmd.Flags().StringP("in", "i", "", "ddl sql path")
	outPath = Cmd.Flags().StringP("out", "o", "stdout", "stdout or file, default stdout")
	prefix = Cmd.Flags().StringP("prefix", "p", "t", "prefix of table, default t")
}
