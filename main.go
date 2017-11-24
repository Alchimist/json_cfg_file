package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"flag"
)


type Config struct {
	Applications []Application
}

type Application struct {
	Name interface{}
	Db   interface{}
}




func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "config file")
	jtest, err := ioutil.ReadFile(configFile)
	type test map[string]interface{}
	var   itest test
	err = json.Unmarshal(jtest, &itest)
	if err != nil {
		fmt.Println("error:", err )
	}
//	fmt.Println(itest)
	for update := range itest {
		fmt.Println(update)
		names := itest[update].(interface{})
//		fmt.Printf("%s\n", names)
//		fmt.Printf("%s\n", names)
		m := names.(map[string]interface{})
//		fmt.Printf("%s\n", m["coins"])
		f := m["coins"]
		for update1 := range f.(map[string]interface{}) {
			fmt.Printf("%s, %f\n", update1, f.(map[string]interface{})[update1])
		}
	}
}
