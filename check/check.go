package check

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/zichouu/go-pkg/color"
)

func Path(commandList []string) (canUseList []string, err_if_len0 error) {
	canUseList = []string{}
	for _, v := range commandList {
		name := strings.Split(v, " ")[0]
		_, err := exec.LookPath(name)
		if err != nil {
			fmt.Println(color.BgRed, fmt.Sprintf("找不到 %v", name), color.Reset)
			fmt.Println(err)
		} else {
			canUseList = append(canUseList, v)
		}
	}
	if len(canUseList) == 0 {
		return nil, errors.New("len = 0")
	}
	return canUseList, nil
}
