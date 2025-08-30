package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Addr
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
		case "/join":
		case "/rooms":
		case "/msg":
		case "/quit":
		default:
			c.err(fmt.Errorf("unknown command %s", cmd))
		}
	}
}
