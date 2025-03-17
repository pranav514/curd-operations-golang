package main

import (
	"github.com/pranav514/curd-operations-golang/initializers"
	 "github.com/pranav514/curd-operations-golang/models"
)


func init(){
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}