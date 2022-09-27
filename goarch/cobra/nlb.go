package cobra

import (
	nlayered_backend "github.com/eneskzlcn/goarch/goarch/nlayered-backend"
	"github.com/spf13/cobra"
)

// nlbCmd represents the ms command
var nlbCmd = &cobra.Command{
	Use:   "nlb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nlb := nlayered_backend.New(".")
		return nlb.Create()
	},
}

func init() {
	rootCmd.AddCommand(nlbCmd)

}