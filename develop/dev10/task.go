package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type TelnetConfig struct {
	Timeout *time.Duration
	Host    string
	Port    string
}

var config = &TelnetConfig{}

func initTelnetConfig() {
	config.Timeout = flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Usage = func() {
		fmt.Println("go-telnet [OPTIONS.... (--timeout=...)] [HOST] [PORT]")
	}
	flag.Parse()

	if !flag.Parsed() {
		flag.Usage()
		os.Exit(1)
	}

	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	config.Host = flag.Arg(0)
	config.Port = flag.Arg(1)
}

func init() {
	initTelnetConfig()
}

func main() {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(config.Host, config.Port), *config.Timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	wg := &sync.WaitGroup{}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := io.Copy(os.Stdout, conn) // Там под капотом бесконечный цикл, пока не дойдет до EOF.
		if err != nil {
			fmt.Println("Error in reading: ", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Error in writing: ", err)
		} // Не знаю как описать, пусть будет как ошибка.
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigs
		conn.Close()
	}()
	wg.Wait()
}
