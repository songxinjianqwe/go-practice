package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func anonymousPipeline(cmd1 *exec.Cmd, cmd2 *exec.Cmd) {
	if cmd1 == nil || cmd2 == nil {
		fmt.Printf("cmd1 or cmd2 couldn't be nil")
	}
	var cmd1Output bytes.Buffer
	cmd1.Stdout = &cmd1Output
	err := cmd1.Run()
	if err != nil {
		fmt.Printf("Error: Couldn't execute command No.0: %s\n", err)
		return
	}
	//fmt.Printf("cmd1: %s\n", string(cmd1Output.String()))
	cmd2.Stdin = &cmd1Output
	var cmd2Output bytes.Buffer
	cmd2.Stdout = &cmd2Output
	err = cmd2.Run()
	if err != nil {
		fmt.Printf("Error: Couldn't execute command No.1: %s\n", err)
		return
	}
	fmt.Printf("cmd2: %s\n", string(cmd2Output.String()))
}


func fileBasedPipe() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("Error: Couldn't create the named pipe: %s\n", err)
	}
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}

func inMemorySyncPipe() {
	reader, writer := io.Pipe()
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}

func main() {
	//cmd1 := exec.Command("ps", "-ef")
	//cmd2 := exec.Command("grep","go")
	//anonymousPipeline(cmd1, cmd2)
	fileBasedPipe()
	inMemorySyncPipe()
}