package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Could not start the FTP server, error: %s\n", err)
	}
	fmt.Println("Starting Golang ftp server ...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go serverProcessInterpreter(conn)
	}
}

const (

	// Ok
	REPLY_SERVICE_READY  = 220
	REPLY_USER_NAME_OKAY = 331

	// Errors
	REPLY_SYNTAX_ERROR = 500
)

const (
	CMD_UNKNOWN = iota
	CMD_USER
	CMD_PASS
	CMD_ACCT
	CMD_CWD
	CMD_CDUP
	CMD_SMNT
	CMD_QUIT
	CMD_REIN
	CMD_PORT
	CMD_PASV
	CMD_TYPE
	CMD_STRU
	CMD_MODE
	CMD_RETR
	CMD_STOR
	CMD_STOU
	CMD_APPE
	CMD_ALLO
	CMD_REST
	CMD_RNFR
	CMD_RNTO
	CMD_ABOR
	CMD_DELE
	CMD_RMD
	CMD_MKD
	CMD_PWD
	CMD_LIST
	CMD_NLST
	CMD_SITE
	CMD_SYST
	CMD_STAT
	CMD_HELP
	CMD_NOOP
)

type Command struct {
	Code           int
	Args           []string
	ErrorRawString string
}

type Reply struct {
	Code    int
	Message string
}

func (r Reply) String() string {
	return fmt.Sprintf("%d %s\n", r.Code, r.Message)
}

func read(s string) Command {
	parts := strings.Split(s, " ")
	if len(parts) <= 0 {
		return Command{
			Code:           CMD_UNKNOWN,
			ErrorRawString: "empty string",
		}
	}
	switch code := parts[0]; code {
	case "USER":
		return Command{
			Code: CMD_USER,
		}
	default:
		return Command{
			Code:           CMD_UNKNOWN,
			ErrorRawString: s,
		}
	}

}

func eval(c Command) Reply {
	switch code := c.Code; code {
	case CMD_USER:
		return Reply{
			Code:    REPLY_USER_NAME_OKAY,
			Message: "User name okay, need password",
		}
	default:
		return Reply{
			Code:    REPLY_SYNTAX_ERROR,
			Message: fmt.Sprintf("Syntax error, command unrecognized: `%s`", c.ErrorRawString),
		}
	}
}

func serverProcessInterpreter(c net.Conn) {
	defer c.Close()
	io.WriteString(c, Reply{
		Code:    REPLY_SERVICE_READY,
		Message: "Service ready for new user",
	}.String())
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		text := scanner.Text()
		io.WriteString(c, eval(read(text)).String())
		continue
	}
	err := scanner.Err()
	if err != nil {
		fmt.Println("error scanning: ", err)
	} else {
		fmt.Println("Closing connection")
	}
}
