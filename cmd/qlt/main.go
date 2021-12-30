package main

import "io/ioutil"

func main() {
	ioutil.WriteFile("./aa.txt", []byte("测试"), 0777)
}
