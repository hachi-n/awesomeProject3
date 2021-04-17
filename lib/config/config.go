package config

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/awesomeProject3"
	"os"
)

type ConfigJson struct {
	AAA string `"aaa":string`
	BBB string `"bbb":string`
	CCC string `"ccc":string`
}

var filePath = "config2/sample.json"


func LoadConfig() *ConfigJson {
	configJson := new(ConfigJson)
	b, err := awesomeProject3.EmbedFiles.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := json.Unmarshal(b, configJson); err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	return configJson
}
