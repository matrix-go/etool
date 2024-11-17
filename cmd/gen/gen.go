package gen

import (
	"github.com/matrix-go/etool/cmd/gen/dao"
	"github.com/spf13/cobra"
)

// Cmd represents the gen command
var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "gen code with command",
	Long:  `A tool collection for gen code`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//ddlPath = Cmd.PersistentFlags().StringP("in", "i", "", "ddl sql path")
	//outPath = Cmd.PersistentFlags().StringP("out", "o", "stdout", "stdout or file, default stdout")
	//prefix = Cmd.PersistentFlags().StringP("prefix", "p", "t", "prefix of table, default t")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// sub commands
	Cmd.AddCommand(dao.Cmd)
}
