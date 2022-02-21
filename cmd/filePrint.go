package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	printFile string
	filePrintCmd = &cobra.Command{
	Use:   "filePrint",
	Short: "filePrint",
	Long: `filePrint`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		file, err := os.Open(printFile)
		if err != nil {
			fmt.Println(err)
		}
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println(err)
			}
		}()


		b, err := ioutil.ReadAll(file)
		fmt.Print(string(b))
	},
}
)

func init(){
	filePrintCmd.PersistentFlags().StringVar(&printFile, "fp", "print file", "file need print")
	rootCmd.AddCommand(filePrintCmd)
}



