package cmd

import (
	"fmt"
	"zkfmapf123/wip/internal"

	"github.com/spf13/cobra"
)

var (
	commitCmd = &cobra.Command{
		Use:   "c",
		Short: "Commit => [0] 브랜치이름 [1] Commit Message",
		Long:  "Commit => [0] 브랜치이름 [1] Commit Message",
		Run: func(cmd *cobra.Command, args []string) {

			_, commitMessage, err := internal.CheckArgs(args)
			if err != nil {
				internal.PanicError(err)
			}

			path, err := internal.GetGitPath()
			if err != nil {
				internal.PanicError(err)
			}

			gCmd := internal.NewGitCmd()
			for _, arr := range [][]string{gCmd.GetAdd(), gCmd.GetCommit(commitMessage)} {
				_, err := internal.MustRunCommandInPath(path, arr[0], arr[1:]...)
				if err != nil {
					internal.PanicError(err)
				}
			}

			fmt.Printf("nCommitMessage : %s", commitMessage)
		},
	}
)

func init() {
	rootCmd.AddCommand(commitCmd)
}
