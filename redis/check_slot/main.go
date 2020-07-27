package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func main() {
	checkSlotByCli("10.96.76.21:6000")
}

func checkSlotByCli(IPPort string) error {
	cli := "../bin/redis-cli"

	cmd := exec.Command(cli, "--cluster", "check", IPPort)
	out, err := cmd.Output()
	if err != nil {
        fmt.Println(err)
	}
	tmp := bufio.NewReader(bytes.NewReader(out))
	for {
		line, _, err2:= tmp.ReadLine()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatal(err2)
		}
		if output(string(line)) {
			fmt.Println(string(line))
		}
	}
	return err
}

func output(s string) bool {
	if strings.Contains(s, "[OK]") {
		return true
	}
	if strings.Contains(s, "[ERR]") {
		return true
	}
	if strings.Contains(s, "[OK]") {
		return true
	}
	if strings.Contains(s, "[WARNING]") {
		return true
	}
	if strings.Contains(s, ">>>") {
		return true
	}
	return false
}
