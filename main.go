package main

import (
	"github.com/mrstexx/webservice/controllers"
	"net/http"
)

//
//func main() {
//	u := models.User{
//		ID:        11,
//		FirstName: "Stefan",
//		LastName:  "Miljevic",
//	}
//	fmt.Println(u)
//}

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
