package main

import (
	"flag"
	"fmt"
	"log"
	"tiny-blockchain-app/app/config"
	"tiny-blockchain-app/app/pkg/restapi"
)

func main() {

	blockchainConfig := loadConfig().BlockChain()
	fmt.Println(blockchainConfig)

	restapi.NewServer()

}

func loadConfig() *config.Config {
	var configFileName, configFilePath string
	flag.StringVar(&configFileName, "conf-file", "config", "Config File Name")
	flag.StringVar(&configFilePath, "conf-path", "./config", "Config File Path")
	flag.Parse()

	// load configs from file
	conf, err := config.New(configFilePath, configFileName, "yaml")
	if err != nil {
		log.Fatalln("failed to load config file - ", err.Error())
	}
	return conf
}
