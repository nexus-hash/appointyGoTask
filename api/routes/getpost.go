package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"go.mongodb.org/mongo-driver/bson"
)

/*
	Takes a post id as url path variable and returns the post details
	For Respose Data Refer README.md
*/

func GetPostHandler(w http.ResponseWriter, r *http.Request) {

	// Get the post id from the url path variable
	pid := strings.TrimPrefix(r.URL.Path, "/posts/")
	if pid == ""{
		http.Error(w, "Invalid request No Post Id Found", http.StatusBadRequest)
		return
	}

	// Get the post details from the database
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	getUserResult:=DB.Collection("Post").FindOne(ctx, bson.M{"_id":pid})
	if getUserResult.Err() != nil {
		http.Error(w, getUserResult.Err().Error(), http.StatusBadRequest)
		return
	}

	// Decode the post details
	var post Post
	getUserResult.Decode(&post)

	// Send the post details in the response
	json_data2, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data2)
	return
}