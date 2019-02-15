package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/**
隔离各个进程看到的挂载点视图
在不同namespace的进程中，看到的文件系统层次是不同的。在Mount Namespace中调用mount()和umount()仅仅只会
影响当前Namespace内的文件系统，而对全局的文件系统是没有影响的。
*/
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
