package exe

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/zichouu/go-pkg/color"
)

func Run(dir string, aenv []string, command ...string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		arg := append([]string{"/c"}, command...)
		cmd = exec.Command("cmd", arg...)
	} else {
		arg := append([]string{"-c"}, command...)
		cmd = exec.Command("sh", arg...)
	}
	if len(aenv) > 0 {
		env := os.Environ()
		for _, v := range aenv {
			env = append(env, v)
		}
		cmd.Env = env
	}
	cmd.Dir = dir
	errColor := color.BgGreen
	fmt.Println(color.BgBlue, "执行", dir, strings.Join(command, " "), color.Reset)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errColor = color.BgRed
	}
	fmt.Println(errColor, "完成", dir, strings.Join(command, " "), color.Reset)
	if err != nil {
		fmt.Println(err)
	}
	if len(out) > 0 {
		fmt.Println(string(out))
	}
	return err
}

func IfExist(path string, filename string, aenv []string, command ...string) error {
	join := filepath.Join(path, filename)
	_, err := os.Stat(join)
	if err == nil {
		err := Run(path, aenv, command...)
		return err
	}
	return nil
}
