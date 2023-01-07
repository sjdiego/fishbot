package main

import (
	"bufio"
	"fmt"
	"log"
	"net/textproto"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	parseConfig()
	catchSigterm()
	conn = connect()

	send(fmt.Sprintf("USER %s 0 * :%s", config.Ident, config.Realname))
	setNick(config.Nickname)

	tp := textproto.NewReader(bufio.NewReader(conn))

	for {
		status, err := tp.ReadLine()
		if err != nil {
			log.Print(err)
			os.Exit(0)
		}
		read(status)
	}
}

func catchSigterm() {
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		disconnect()
	}()
}
