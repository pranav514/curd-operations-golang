package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pranav514/curd-operations-golang/controllers"
	"github.com/pranav514/curd-operations-golang/initializers"
)
func init(){
	initializers.LoadEnv()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()
	r.POST("/createpost", controllers.CreatePost)
	r.GET("/getpost" , controllers.GetPost)
	r.GET("/getspecific/:id" , controllers.GetSpecific);
	r.PUT("/update/:id" , controllers.Update);
	r.DELETE("/delete/:id" , controllers.Delete);
	
	port := os.Getenv("PORT")
	fmt.Println(port)
	r.Run(port) 
}