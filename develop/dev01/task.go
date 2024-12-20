package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func PrintCurrentTime(w io.Writer) error {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return err
	}
	if _, err = fmt.Fprintln(w, "Current time: ", time); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := PrintCurrentTime(os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error with ntp.Time: %s\n", err)
		os.Exit(1)
	}
}
