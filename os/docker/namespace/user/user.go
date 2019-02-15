package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/**
	隔离用户的用户组ID
    一个进程的User ID和Group ID在User namespace内外可以是不同的。
	比较常用的是，在宿主机上以一个非root用户运行创建一个User namespace，然后在User namespace中映射为root用户。
*/
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
	}
	cmd.SysProcAttr.Credential = &syscall.Credential{
		Uid: uint32(1),
		Gid: uint32(1),
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
