package e

import (
	"fmt"
	"testing"
)

func TestE运行_mac(t *testing.T) {
	//command := "ping www.baidu.com"
	command := "ls -al"
	wait := true
	output := X运行_mac(command, wait, func(line string) {
		t.Log("??x", line)
	})
	t.Log("??", output)
	//  add  more  tests  here
}

func TestE运行_win(t *testing.T) {
	return
	command := "ping"
	wait := true
	output := X运行_win(command, wait)
	fmt.Println(output)
	//  add  more  tests  here
}
