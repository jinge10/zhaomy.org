package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	_ "zhaomy.org/controller"
)

func init() {
}
func main() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		log.Fatalln("my.ini load fail")
	}
	fmt.Println("base.app  ->   ", cfg.Section("").Key("version"))
	//router.GetRouter().Run(":8080")
}
