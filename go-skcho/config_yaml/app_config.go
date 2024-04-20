package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const configPath = "config.yaml"

type IosConfig struct {
	LogLevel struct {
		Console string `yaml:"console"`
		File    string `yaml:"log_file"`
	} `yaml:"log_level"`
	ProcInfo struct {
		NATS_URL string `yaml:"nats_url"`
		MY_HOST  string `yaml:"my_host"`
	} `yaml:"proc_conf"`
}

var Cfg IosConfig

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("os.Open(%v) fail: %v\n", configPath, err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		fmt.Printf("decoder.Decode() fail: %v\n", err)
	}
}

func main() {

	ReadConfig()

	fmt.Printf("Cfg = [%v]\n", Cfg)
	fmt.Printf("nats_url = [%v]\n", Cfg.ProcInfo.NATS_URL)

}
