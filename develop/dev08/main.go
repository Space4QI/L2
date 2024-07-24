package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	// Обработка сигналов для корректного завершения работы шелла
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("\nReceived an interrupt, stopping shell...")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		// Вывод приглашения
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Удаление лишних пробелов и символов новой строки
		input = strings.TrimSpace(input)

		// Разделение ввода на аргументы
		args := strings.Split(input, " ")

		// Проверка и выполнение команд
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("cd: missing argument")
			} else {
				changeDirectory(args[1])
			}
		case "pwd":
			printWorkingDirectory()
		case "echo":
			echo(args[1:])
		case "kill":
			if len(args) < 2 {
				fmt.Println("kill: missing argument")
			} else {
				killProcess(args[1])
			}
		case "ps":
			listProcesses()
		case "\\quit":
			fmt.Println("Exiting shell...")
			return
		default:
			executeCommand(args)
		}
	}
}

func changeDirectory(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Println("cd:", err)
	}
}

func printWorkingDirectory() {
	if dir, err := os.Getwd(); err == nil {
		fmt.Println(dir)
	} else {
		fmt.Println("pwd:", err)
	}
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func killProcess(pid string) {
	p, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("kill: invalid pid", pid)
		return
	}

	process, err := os.FindProcess(p)
	if err != nil {
		fmt.Println("kill:", err)
		return
	}

	if err := process.Kill(); err != nil {
		fmt.Println("kill:", err)
	}
}

func listProcesses() {
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("ps:", err)
		return
	}
	fmt.Print(string(output))
}

func executeCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}
