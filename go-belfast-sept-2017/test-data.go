package main

import "os"

func getDir() {
	wd, _ := os.Getwd()
	return wd
}

func dankFixture() {
	f, _ := os.Open("testdata/somefixture.json")
	return f
}
