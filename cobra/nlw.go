/*
Copyright © 2022 Nazif Enes Kızılcin enes.kizilcin@gmail.com
*/

package cobra

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nlw called")
	},
}

func init() {
	rootCmd.AddCommand(nlwCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nlwCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nlwCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
