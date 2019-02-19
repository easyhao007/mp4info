package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"mp4info/mov"
)

func main() {
	fmt.Println("vim-go")
	pfile , err := os.Open("F:\\video\\file6.mp4")
	if err != nil{
		log.Panic(err.Error())
		return
	}

	buf , err := ioutil.ReadAll(pfile)
	if err != nil{
		log.Panic(err.Error())
		return
	}

	for {
		boxLen , err := mov.ParseBox(buf)
		if err != nil{
			log.Println(err.Error())
		}
		buf = buf[boxLen:]
		if len(buf) == 0{
			break
		}
	}
}
