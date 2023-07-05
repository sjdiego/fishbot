package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var conn net.Conn

func connect() net.Conn {
	log.Printf("Trying connection to %s ...", config.Hostname)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Hostname, config.Port))

	if err != nil {
		panic(err)
	}

	log.Printf("Established connection to %s", config.Hostname)

	return conn
}

func disconnect() {
	send("QUIT :Finished")
	log.Printf("Closing connection...")
	defer conn.Close()
	os.Exit(0)
}

func send(message string) {
	log.Printf("-> %s", message)
	fmt.Fprintf(conn, "%s\r\n", message)
}

func read(message string) {
	raw := getRaw(message)
	if raw > 0 {
		parseRaw(raw, message)
	} else {
		parseCommand(message)
	}

	if strings.HasPrefix(message, "PING") {
		replyPing(message)
	}
}

func join(channel string) {
	log.Printf("Joining channel #%s ...", channel)
	send(fmt.Sprintf("JOIN #%s", channel))
}

func getRaw(message string) int {
	// :stirling.chathispano.com 433 * :El nick n está en uso.
	if !strings.HasPrefix(message, ":") {
		return 0
	}

	parts := strings.Split(message, " ")

	raw, err := strconv.Atoi(parts[1])

	if err != nil {
		return 0
	}

	return raw
}

func parseRaw(raw int, message string) {
	switch raw {
	case 001:
		go autoModes()
		go joinChannels()
	case 433, // Nick in use
		451, // Not registered
		464: // Invalid password
		setNick(randomString(8))
	case 372, 375, 376: // Ignore MOTD
		return
	}

	log.Printf("<- %s", message)
}

func parseCommand(message string) {
	parts := strings.Split(message, " ")
	if len(parts) < 4 {
		return
	}

	command := parts[1]

	if command == "PRIVMSG" && strings.HasPrefix(parts[2], "#") {
		nick := strings.Split(strings.TrimPrefix(parts[0], ":"), "!")
		text := strings.TrimPrefix(strings.Join(parts[3:], " "), ":")

		replyToChannel(parts[2], text, nick[0])
	} else if command == "INVITE" {
		nick := strings.Split(strings.TrimPrefix(parts[0], ":"), "!")
		channel := strings.TrimPrefix(strings.Join(parts[3:], " "), ":")

		replyInvite(channel, nick[0])
	}

	log.Printf("<- %s", message)
}

func replyPing(message string) {
	pongCode := strings.Split(message, ":")
	send(fmt.Sprintf("PONG :%s", pongCode[1]))
}

func setNick(nick string) {
	nickName := strings.Split(nick, ":")
	log.Printf("Changing nick to %s", nickName[0])
	send(fmt.Sprintf("NICK %s", nick))
}

func autoModes() {
	nickName := strings.Split(config.Nickname, ":")
	log.Printf("Settings %s user modes to %s", nickName[0], config.Automodes)
	send(fmt.Sprintf("MODE %s %s", nickName[0], config.Automodes))
}

func joinChannels() {
	if len(config.DefaultChannel) == 0 {
		return
	}

	join(config.DefaultChannel)
}

func replyInvite(channel string, nick string) {
	log.Printf("%s invited me to %s", nick, channel)

	if config.AllowInvites {
		join(strings.TrimPrefix(channel, "#"))
		time.Sleep(3 * time.Second)
		send(fmt.Sprintf("PRIVMSG %s :\x01ACTION m00s contentedly at %s\x01", channel, nick))
	}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmonpqrstuvwxyz")

	s := make([]rune, n)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
