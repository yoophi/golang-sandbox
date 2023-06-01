package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	cfg, err := ini.Load("setting.ini")
	if err != nil {
		log.Fatal(err)
	}

	foo := cfg.Section("test").Key("foo").String()
	fmt.Println("foo:", foo)

	port, err := cfg.Section("server").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	host := cfg.Section("server").Key("host").String()
	fmt.Println("server.port:", port)
	fmt.Println("server.host:", host)
}
