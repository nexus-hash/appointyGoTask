package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	Takes user id as url path parameter and returns the user details
	For Respose Data Refer README.md
*/

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	// Thread-safe
	lock.Lock()
	defer lock.Unlock()

	// Get user id from url path parameter
	uid := strings.TrimPrefix(r.URL.Path, "/users/")
	if uid == ""{
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Convert user id from hex string to ObjectId
	uuid,err:= primitive.ObjectIDFromHex(uid);
	if err!=nil{
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Get user details from database
	getUserResult:=DB.Collection("users").FindOne(ctx, bson.M{"_id":uuid})
	if getUserResult.Err() != nil {
		http.Error(w, getUserResult.Err().Error(), http.StatusInternalServerError)
		return
	}

	var user User
	getUserResult.Decode(&user)

	// Not returning user password for security reasons although it is reverse engineering safe
	user.Password = ""

	// Convert user details to json
	json_data2,err:=json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Return user details
	fmt.Println(string(json_data2))
	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_data2)
	return ;
}