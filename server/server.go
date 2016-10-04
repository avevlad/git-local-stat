package server

import "fmt"

type CommandLineOptions struct {
	Port int `short:"p" long:"port" description:"Http port" default:"1339"`
}

func Init(options CommandLineOptions) {
	fmt.Println(Commits[0].Author)
	fmt.Println(len(Commits))
	//const name  = Commits[0].Body
	//for _, element := range Commits {
	//	fmt.Println(element.Author)
	//	//fmt.Println(element)
	//	fmt.Println("--------------")
	//}
}
