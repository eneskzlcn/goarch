/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cobra

import (
	nlayered_backend "github.com/eneskzlcn/goarch/goarch/nlayered-backend"
	"github.com/spf13/cobra"
)

// nlbCmd represents the nlb command
var nlbCmd = &cobra.Command{
	Use:   "nlb",
	Short: "Creates n-layered backend service architecture.",
	Long: `
Use that command when you want to create a backend service
that has an n-layered directory and application structure.
Prefer to call it before start the project and overwrite 
the created templates and directories to make your 
application work.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nlb := nlayered_backend.New(".")
		return nlb.Create()
	},
}

func init() {
	rootCmd.AddCommand(nlbCmd)

}
