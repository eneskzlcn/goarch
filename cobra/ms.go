/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cobra

import (
	"github.com/eneskzlcn/goarch/microservice"
	"github.com/eneskzlcn/goarch/tech"
	"github.com/spf13/cobra"
)

// msCmd represents the ms command
var msCmd = &cobra.Command{
	Use:   "ms",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		options := tech.OptionsFromStr(args[0])
		return microservice.CreateArchitecture(options)
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
