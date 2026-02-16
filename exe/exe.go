package exe

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/zichouu/go-pkg/color"
)

func Run(dir string, command string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Dir = dir
	errColor := color.BgGreen
	fmt.Println(color.BgBlue, "执行", dir, command, color.Reset)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errColor = color.BgRed
	}
	fmt.Println(errColor, "完成", dir, command, color.Reset)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	return err
}

func IfExist(path string, filename string, command string) error {
	join := filepath.Join(path, filename)
	_, err := os.Stat(join)
	if err == nil {
		err := Run(path, command)
		return err
	}
	return nil
}
