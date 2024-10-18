package cmd

import (
	"fmt"
	"zkfmapf123/wip/internal"

	"github.com/spf13/cobra"
)

var (
	commitPushCmd = &cobra.Command{
		Use:   "cp",
		Short: "Commit && Push => [0] 브랜치이름 [1] Commit Message && Push",
		Long:  "Commit && Push => [0] 브랜치이름 [1] Commit Message && Push",
		Run: func(cmd *cobra.Command, args []string) {

			branch, commitMessage, err := internal.CheckArgs(args)
			if err != nil {
				internal.PanicError(err)
			}

			path, err := internal.GetGitPath()
			if err != nil {
				internal.PanicError(err)
			}

			gCMD := internal.NewGitCmd()

			for _, arr := range [][]string{gCMD.GetAdd(), gCMD.GetCommit(commitMessage), gCMD.GetPush(branch)} {
				_, err := internal.MustRunCommandInPath(path, arr[0], arr[1:]...)
				if err != nil {
					internal.PanicError(err)
				}
			}

			fmt.Printf("Branch : %s\nCommitMessage : %s", branch, commitMessage)
		},
	}
)

func init() {
	rootCmd.AddCommand(commitPushCmd)
}
