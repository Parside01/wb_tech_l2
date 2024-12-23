package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/process"
	"os"
	"os/signal"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps

- cd <args> — смена директории (в качестве аргумента могут быть то-то и то);

- pwd — показать путь до текущего каталога;

- echo <args> — вывод аргумента в STDOUT;

- kill <args> — «убить» процесс, переданный в качестве аргумента (пример: такой-то пример);

- ps — выводит общую информацию по запущенным процессам в формате такой-то формат.

поддержать fork/exec команды
конвейер на пайпах

Реализовать утилиту netcat (nc) клиент // Этого в задании на платформе нет, так что пока не буду реализовывать.
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	shell := MustNewShell()
	signal.Ignore(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	shell.Run()
}

type Command func(args []string) error

type Shell struct {
	currentDir string
	commands   map[string]Command
}

func MustNewShell() *Shell {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	shell := &Shell{
		currentDir: dir,
	}

	shell.commands = map[string]Command{
		"cd":   shell.cd,
		"pwd":  shell.pwd,
		"echo": shell.echo,
		"kill": shell.kill,
		"ps":   shell.ps,
		//"exec": exec,
		//"fork": fork,
	}

	return shell
}

func (s *Shell) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	s.printCurrentDir()
	for scanner.Scan() {
		input := scanner.Text()

		if input == "\\quit" || input == "\\exit" || input == "\\q" {
			break
		}

		commands := strings.Split(input, "|")
		for i := range commands {
			cmd, args := s.getCommandAndArgs(commands[i])

			f, ok := s.commands[cmd]
			if !ok {
				s.printError(fmt.Errorf("unknown command %s", cmd))
				continue
			}
			if err := f(args); err != nil {
				s.printError(err)
				continue
			}

		}
		s.printCurrentDir()
	}
}

func (s *Shell) printError(err error) {
	fmt.Println(color.RedString("Error: %v", err))
}

func (s *Shell) printCurrentDir() {
	fmt.Print(color.YellowString("%s> ", s.currentDir))
}

func (s *Shell) printText(format string, args ...interface{}) {
	fmt.Printf(color.GreenString(fmt.Sprintf(format, args...)))
}

func (s *Shell) getCommandAndArgs(input string) (string, []string) {
	var cmd string
	var args []string

	heap := strings.Fields(input)
	if len(heap) == 0 {
		return cmd, args
	}

	cmd = heap[0]
	if len(heap) == 1 {
		return cmd, args
	}

	args = slices.Clone(heap[1:])
	return cmd, args
}

// =========================================================>>>>>>

var (
	ErrManyArgs = errors.New("error: many arguments")
)

func (s *Shell) pwd(args []string) error {
	if len(args) != 0 {
		return ErrManyArgs
	}

	s.printText("\nPath\n----\n%s\n\n", s.currentDir)
	return nil
}

func (s *Shell) cd(args []string) error {
	if len(args) == 0 {
		return nil
	}
	if len(args) > 1 {
		return ErrManyArgs
	}

	if strings.HasPrefix(args[0], "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		args[0] = filepath.Join(homeDir, args[0][1:])
	}

	if err := os.Chdir(args[0]); err != nil {
		return err
	}

	currDir, err := os.Getwd()
	s.currentDir = currDir
	return err
}

func (s *Shell) echo(args []string) error {
	s.printText("%s\n", strings.Join(args, " "))
	return nil
}

func (s *Shell) kill(args []string) error {
	var killProcess []*os.Process
	for _, arg := range args {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		killProcess = append(killProcess, proc)
	}

	for _, proc := range killProcess {
		if err := proc.Kill(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Shell) ps(args []string) error {
	if len(args) != 0 {
		return ErrManyArgs
	}

	processes, err := process.Processes()
	if err != nil {
		return err
	}

	s.printText("PID USER CPU MEM VSZ RSS TTY STAT START TIME COMMAND\n")
	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			continue
		}
		pid := proc.Pid

		user, err := proc.Username()
		if err != nil {
			user = "N/A"
		}

		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		memInfo, err := proc.MemoryInfo()
		if err != nil {
			memInfo = nil
		}
		if memInfo == nil {
			continue
		}

		status, err := proc.Status()
		if err != nil {
			status = "N/A"
		}

		s.printText("%d %s %.1f %.1f %d %d %s %s %s %s %s\n",
			pid, user, cpuPercent, float64(memInfo.RSS)/1024, memInfo.VMS/1024, memInfo.RSS/1024, "N/A", status, "N/A", "N/A", name)
	}
	return nil
}

func (s *Shell) fork(args []string) error {
	// Должно быть чего-то такое, но на винде не работает, а замены не нашел(.
	//pid, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	//if err != nil {
	//	return err
	//}
	//s.printText("Fork process with pid %v", pid)
	return nil
}

func (s *Shell) exec(args []string) error {
	if len(args) == 0 {
		return nil
	}

	cmd := args[0]
	if len(args) == 1 {
		return syscall.Exec(cmd, nil, nil)
	}
	return syscall.Exec(cmd, args[1:], nil)
}
