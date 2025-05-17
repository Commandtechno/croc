//go:build !windows

package utils

import (
	"errors"
	"time"
)

func SetFileTime(path string, atime, ctime, mtime time.Time) (err error) {
	return errors.New("not implemented")
}
