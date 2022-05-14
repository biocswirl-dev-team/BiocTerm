package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

// YamlConfig is exported.
type YamlConfig struct {
	Connection struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
}

func main() {
    fmt.Println("Parsing YAML file")

    var fileName string
    flag.StringVar(&fileName, "f", "", "YAML file to parse.")
    flag.Parse()

    if fileName == "" {
        fmt.Println("Please provide yaml file by using -f option")
        return
    }

    yamlFile, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Printf("Error reading YAML file: %s\n", err)
        return
    }

    var yamlConfig YamlConfig
    err = yaml.Unmarshal(yamlFile, &yamlConfig)
    if err != nil {
        fmt.Printf("Error parsing YAML file: %s\n", err)
    }

    fmt.Printf("Result: %v\n", yamlConfig)
}
