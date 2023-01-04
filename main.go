package main

import (
	_ "Video/routers"
	"github.com/astaxie/beego"
)

func main() {
	//juge := License()
	//if juge == "true" {
	//	beego.Run()
	//} else {
	//	fmt.Println("--------------------License error--------------------")
	//}
	beego.Run()
}

//func License() string {
//	licenseData, err := ioutil.ReadFile("./License_master/License")
//	if err != nil {
//		log.Fatal(err)
//	}
//	license := string(licenseData)
//	if license != "true" {
//		return "false"
//	}
//
//	return "true"
//}
