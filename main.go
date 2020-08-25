package main

import (
	"yeeep.cn/yepgo/router"
)

func main() {
	//db := gdo.InitDB()
	//defer db.Close()
	r := router.InitRouter()
	_ = r.Run()
}
