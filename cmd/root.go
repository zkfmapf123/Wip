package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Wip",
		Short: "Wip is Simple Git Add -> Commit -> Push CLI",
		Long:  "Wip is Simple Git Add -> Commit -> Push CLI",
		Run: func(cmd *cobra.Command, args []string) {
			// Hello world
		},
	}
)

func initConfig() {

}

func MustExecute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
