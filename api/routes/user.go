package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       string `bson:"_id" json:"_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserHandler (w http.ResponseWriter, r *http.Request) {
	
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	
	var user User
	if err:= json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"User created\"}"))
	fmt.Println(user.Username);
}