package internal

import (
	"errors"
	"os"
)

func MustGetPwdPath() string {

	path := WrapMustFunc(os.Getwd)
	return path
}

func MustGetHomePath() string {
	path := WrapMustFunc(os.UserHomeDir)
	return path
}

func GetGitPath() (string, error) {

	homePath := MustGetHomePath()
	for {

		currentDir := MustGetPwdPath()
		_, err := os.Stat(".git")

		// .git을 찾았을 때
		if err == nil {
			return MustGetPwdPath(), nil
		}

		// home과 같아짐...
		if currentDir == homePath {
			break
		}

		// 상위 Directory로 이동
		if err := os.Chdir(".."); err != nil {
			return "", err
		}
	}

	return "", errors.New("not found .git")
}
