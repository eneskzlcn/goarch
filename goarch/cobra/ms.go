/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cobra

import (
	"github.com/eneskzlcn/goarch/goarch/microservice"
	"github.com/spf13/cobra"
)

// msCmd represents the ms command
var msCmd = &cobra.Command{
	Use:   "ms",
	Short: "Creates a domain driven designed microservice.",
	Long: `
Call that command when you want to create a new microservice.
Prefer to use it as a template before start to write any code,
and then overwrite the template to make your application 
work.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ms := microservice.New(".")
		return ms.Create()
	},
}

func init() {
	rootCmd.AddCommand(msCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// msCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// msCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
