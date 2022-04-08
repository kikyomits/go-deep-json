package main

import (
	"io/ioutil"
	"os"
)

func readJson(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) { _ = jsonFile.Close() }(file)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return data
}

//writeToFile: add parent keys to the data object and write to file
// in:
// 	data = [{"name": "xxx", "description": "aaa"}]
//  keys = ["aws", "vpc", "vpc-a"]
// out:
//  { "aws": { "vpc": { "vpc-a": [{"name": "xxx", "description": "aaa"}] } } }
func writeToFile(data interface{}, keys []string) {
	// if no parent keys, just write out the data
	if len(keys) == 0 {
		// TODO: write `data` to json file.
		return
	}

	// build the youngest child object first
	out := map[string]interface{}{
		keys[len(keys)-1]: data,
	}

	// run on reverse loop to add parent keys to the youngest child object
	for i := len(keys) - 2; i >= 0; i-- {
		out = map[string]interface{}{
			keys[i]: out,
		}
	}
	// TODO: write `out` to json file.
	return
}
