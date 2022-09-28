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
	Short: "Creates a n-layered web application architecture.",
	Long: `
Call that command when you want to start from a 
template n-layered web application to continue 
work from the template structure.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nlw := nlayered_web.New(".")
		return nlw.Create()
	},
}

func init() {
	rootCmd.AddCommand(nlwCmd)

}
