package file

import (
	"io"
	"os"
)

func Copy(src string, dst string) error {
	srcOpen, err := os.Open(src)
	if err != nil {
		return err
	}
	dstOpen, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(dstOpen, srcOpen)
	_ = srcOpen.Close()
	_ = dstOpen.Close()
	return err
}
