package main

import "testing"

func TestWriteToFile(t *testing.T) {
	writeToFile("xxx", []string{"aws", "vpc", "vpc-aaa"})
}
