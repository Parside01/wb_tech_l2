package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
)

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
