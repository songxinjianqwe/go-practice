package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	// 这里的/proc/self/exe 调用
	// 其中/proc/self指的是当前运行进程自己的环境，exec其实就是自己调用了自己，我们使用这种方式实现对创建出来的进程进行初始化
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d\n", syscall.Getpid())
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 异步重新调用main函数，之后可以执行添加cgroup操作，此时必须知道自己的pid
	// 在重新调用中来执行真正的command
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Printf("%v\n", cmd.Process.Pid)
		cgroupName := "test-memory-limit-in-go"
		os.Mkdir(path.Join(cgroupMemoryHierarchyMount, cgroupName), 0755)
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, cgroupName, "tasks"),
			[]byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, cgroupName, "memory.limit_in_bytes"), []byte("100m"), 0644)
	}
	cmd.Process.Wait()
}
