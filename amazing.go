// amazing.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var sensor_script *[]command_record

type command_record struct {
	Source    string
	What      string
	Condition string
	Value     string
	Host      string
	Action    string
}

func main() {
	file, e := ioutil.ReadFile("./sensor.json")
	if e != nil {
		fmt.Printf("File error %v\n", e)
		os.Exit(1)
	}
	json.Unmarshal(file, &sensor_script)
	for _, v := range *sensor_script {
		fmt.Println(v.Source)
		fmt.Println(v.What)
		fmt.Println(v.Condition)
		fmt.Println(v.Value)
		fmt.Println(v.Host)
		fmt.Println(v.Action)
		fmt.Println("=========")
	}
}
