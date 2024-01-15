package main

import (
	"github.com/DanilCodeGit/loyalty_system/go-musthave-diploma-tpl/internal/conf"
	"log"
)

func main() {
	err := conf.InitConfig()
	if err != nil {
		log.Fatalf("init config: %s", err)
	}
}
