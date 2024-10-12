package internal

import "os"

func MustGetPwdPath() string {
	path, err := os.Getwd()

}
