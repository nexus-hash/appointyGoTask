package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	Id 				int64 		`bson:"_id" json:"_id"`
	Caption 	string		`json:"caption"`	
	ImageURL 	string		`json:"imageurl"`
	PostedAt 	time.Time `json:"-"`
	UserId 		int64 		`json:"userid"`
}

func PostHandler (w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.PostedAt = time.Now()

	postDetails := bson.D{
		{Key: "_id",Value: post.Id},
		{Key: "caption",Value: post.Caption},
		{Key: "imageurl",Value: post.ImageURL},
		{Key: "postedat",Value: post.PostedAt},
		{Key: "userid",Value: post.UserId},
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	insertResult, err := DB.Collection("Post").InsertOne(ctx, postDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(insertResult.InsertedID)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"Post created\"}"))
	return;
}