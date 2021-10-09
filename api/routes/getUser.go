package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	Takes user id as url path parameter and returns the user details
*/

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	uid := strings.TrimPrefix(r.URL.Path, "/users/")
	if uid == ""{
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	id,_ := strconv.ParseInt(uid, 10, 64)
	fmt.Println(id)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	getUserResult:=DB.Collection("users").FindOne(ctx, bson.M{"_id":id})
	if getUserResult.Err() != nil {
		http.Error(w, getUserResult.Err().Error(), http.StatusInternalServerError)
		return
	}
	var user User
	getUserResult.Decode(&user)

	// Not returning user password for security reasons although it is reverse engineering safe

	user.Password = ""
	json_data2,err:=json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	fmt.Println(string(json_data2))
	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_data2)
	return;
	
}