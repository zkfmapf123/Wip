package cmd

import (
	"errors"
	"fmt"
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
			branch, commitMessage := args[0], args[1]

			// check to branch
			if branch == "" {
				internal.PanicError(errors.New("not Exists Branch"))
			}

			// check to commit
			if commitMessage == "" {
				commitMessage = fmt.Sprintf("%s-%s", DEFAULT_COMMIT_MESSAGE, uuid.New().String())
			}

			fmt.Println(branch, commitMessage)

		},
	}
)

func init() {
	rootCmd.AddCommand(wipCmd)
}
