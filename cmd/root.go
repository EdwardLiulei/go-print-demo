package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)



var (
	param string
	rootCmd = &cobra.Command{
	Use:   "print",
	Short: "print anything",
	Long: `print anything`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(param)
	},
}
)

func init() {

	rootCmd.PersistentFlags().StringVar(&param, "p", "parameter", "parameter need print")

	//rootCmd.AddCommand(addCmd)
	//rootCmd.AddCommand(initCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}