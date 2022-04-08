package main

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

var logger = logrus.StandardLogger()

type Property struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Properties []Property
type Resource map[string]Properties
type Service map[string]Resource
type Cloud map[string]Service
type Meta map[string]Cloud

func (meta Meta) process() {
	writeToFile(meta, []string{})
	for k, v := range meta {
		v.process([]string{k})
	}
}

func (cloud Cloud) process(keys []string) {
	writeToFile(cloud, keys)
	for k, v := range cloud {
		v.process(append(keys, k))
	}
}

func (service Service) process(keys []string) {
	writeToFile(service, keys)
	for k, v := range service {
		v.process(append(keys, k))
	}
}

func (resource Resource) process(keys []string) {
	writeToFile(resource, keys)
}

func main() {
	inp := readJson("./input.json")
	var meta Meta
	err := json.Unmarshal(inp, &meta)
	if err != nil {
		logger.Fatal(err)
	}
	meta.process()
	logger.Info("success!")
}
