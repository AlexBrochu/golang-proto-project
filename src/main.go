package main

import (
	"fmt"

	mylogger "github.com/AlexBrochu/golang-package-tuto"
	"github.com/AlexBrochu/golang-package-tuto/vendors/github.com/alexbrochu/awesomepackage/corpus"
)

func main() {
	fmt.Println("hello world")
	mylogger.LogError("for fuck sake")
	mylogger.LogInfo(corpus.Corpus_name[0])
}

//https://developers.google.com/protocol-buffers/docs/gotutorial
