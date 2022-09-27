/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cobra

import (
	nlayered_web "github.com/eneskzlcn/goarch/goarch/nlayered-web"
	"github.com/spf13/cobra"
)

// nlwCmd represents the nlw command
var nlwCmd = &cobra.Command{
	Use:   "nlw",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nlw := nlayered_web.New(".")
		return nlw.Create()
	},
}

func init() {
	rootCmd.AddCommand(nlwCmd)

}
