package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	DEFAULT_COMMIT_MESSAGE = "wip"
)

type GitCommand struct {
}

func NewGitCmd() GitCommand {
	return GitCommand{}
}

func (g GitCommand) GetAdd() []string {
	return []string{"git", "add", "."}
}

func (g GitCommand) GetCommit(msg string) []string {
	return []string{"git", "commit", "-m", msg}
}

func (g GitCommand) GetPush(branch string) []string {
	return []string{"git", "push", "origin", branch}
}

func CheckArgs(args []string) (branch string, commitMessage string, err error) {

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
