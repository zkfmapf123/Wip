package internal

import "os"

func MustGetPwdPath() string {

	path := WrapMustFunc(os.Getwd)
	return path
}

func MustGetHomePath() string {
	path := WrapMustFunc(os.UserHomeDir)
	return path
}
