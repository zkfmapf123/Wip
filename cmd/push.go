package cmd

import (
	"fmt"
	"zkfmapf123/wip/internal"

	"github.com/spf13/cobra"
)

var (
	pushCmd = &cobra.Command{
		Use:   "p",
		Short: "push",
		Long:  "push",
		Run: func(cmd *cobra.Command, args []string) {

			branch, _, err := internal.CheckArgs(args)
			if err != nil {
				internal.PanicError(err)
			}

			path, err := internal.GetGitPath()
			if err != nil {
				internal.PanicError(err)
			}

			gCmd := internal.NewGitCmd()
			for _, arr := range [][]string{gCmd.GetPush(branch)} {
				_, err := internal.MustRunCommandInPath(path, arr[0], arr[1:]...)
				if err != nil {
					internal.PanicError(err)
				}
			}

			fmt.Printf("Branch : %s", branch)
		},
	}
)

func init() {
	rootCmd.AddCommand(pushCmd)
}
