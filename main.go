package main

import (
	"fmt"
	"jiang-go-study/routers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	// 注册路由
	r := routers.CreateHttpHandler()
	fmt.Println(r)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}

	var dsn = "root:jiangjunfeng@tcp(127.0.0.1:3306)/jiangTest?charset=utf8mb4&parseTime=True&loc=Local&timeout=100s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println(db)
	}
	user := User{Name: "Jinzhu"}

	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(result)
}
