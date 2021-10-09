package routes

import (
	"fmt"
	"net/http"
)


func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("id"))
	w.Write([]byte("GetUserHandler"))
	return;
	
}