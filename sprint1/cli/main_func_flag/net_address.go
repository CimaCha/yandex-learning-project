package main

import (
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (net *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", net.Host, net.Port)
}

// Set связывает переменную типа со значением флага
// и устанавливает правила парсинга для пользовательского типа.
func (net *NetAddress) Set(flagValue string) error {
	net.Host = strings.Split(flagValue, ":")[0]
	port, err := strconv.ParseInt(strings.Split(flagValue, ":")[1], 10, 64)
	if err != nil {
		return err
	}
	net.Port = int(port)
	return nil
}
