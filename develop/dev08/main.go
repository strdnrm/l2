package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Split(input, " ")

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println(err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)
		case "echo":
			if len(args) < 2 {
				fmt.Println("Usage: echo <message>")
				continue
			}
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Usage: kill <pid>")
				continue
			}
			pid := args[1]
			cmd := exec.Command("kill", pid)
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		case "ps":
			cmd := exec.Command("ps")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(output))
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
