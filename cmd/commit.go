package cmd

import (
	"errors"
	"fmt"
	"strings"
	"zkfmapf123/wip/internal"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

const (
	DEFAULT_COMMIT_MESSAGE = "wip"
)

var (
	wipCmd = &cobra.Command{
		Use:   "c",
		Short: "Commit => [0] 브랜치이름 [1] Commit Message",
		Long:  "Commit => [0] 브랜치이름 [1] Commit Message",
		Run: func(cmd *cobra.Command, args []string) {

			branch, commitMessage, err := checkArgs(args)
			if err != nil {
				internal.PanicError(err)
			}

			path, err := internal.GetGitPath()
			if err != nil {
				internal.PanicError(err)
			}

			gitAdd := []string{"git", "add", "."}
			gitCommit := []string{"git", "commit", "-m", commitMessage}
			gitPush := []string{"git", "push", "origin", branch}

			for _, arr := range [][]string{gitAdd, gitCommit, gitPush} {
				_, err := internal.MustRunCommandInPath(path, arr[0], arr[1:]...)
				if err != nil {
					internal.PanicError(err)
				}
			}

			fmt.Printf("Branch : %s\nCommitMessage : %s", branch, commitMessage)
		},
	}
)

func checkArgs(args []string) (branch string, commitMessage string, err error) {

	if len(args) == 0 {
		return "", "", errors.New("not Exists branch and commitMessage")
	}

	if len(args) == 1 {
		return args[0], fmt.Sprintf("%s-%s", DEFAULT_COMMIT_MESSAGE, getUUid()), nil
	}

	return args[0], fmt.Sprintf("%s-%s", args[1], getUUid()), nil
}

func getUUid() string {
	return strings.Split(uuid.New().String(), "-")[0]
}

func init() {
	rootCmd.AddCommand(wipCmd)
}
