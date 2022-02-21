package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
)

var (
	outFile string
	fileOutCmd = &cobra.Command{
		Use:   "fileOut",
		Short: "fileOut",
		Long: `fileOut`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			content := []byte(param)
			err := ioutil.WriteFile(outFile, content, 0644)
			if err != nil {
				panic(err)
			}
		},
	}
)

func init(){
	fileOutCmd.PersistentFlags().StringVar(&outFile, "fo", "out file", "file need print")
	rootCmd.AddCommand(fileOutCmd)
}



