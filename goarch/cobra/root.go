/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cobra

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:        "goarch",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "Gon",
	Args:       cobra.ArbitraryArgs,
	Long: `Goarch is a CLI tool for Go that creates a directory
structure for chosen architecture with some ready library templates inside.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Example)
	},
	Example: `
"goarch ms": For domain driven microservice architecture
"goarch nlb": For n-layered backend architecture
"goarch nlw": For n-layered web application architecture`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
