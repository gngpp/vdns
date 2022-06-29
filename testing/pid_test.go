package testing

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestPid(t *testing.T) {
	cmd := exec.Command("/bin/zsh", "-c", "ps -ef | grep vdns")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	outBytes, _ := ioutil.ReadAll(stdout)
	err = stdout.Close()
	if err != nil {
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}

	fmt.Println("Execute finished:\n" + string(outBytes))
}
