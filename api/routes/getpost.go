package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	pid := strings.TrimPrefix(r.URL.Path, "/posts/")
	if pid == ""{
		http.Error(w, "Invalid request No Post Id Found", http.StatusBadRequest)
		return
	} 
	id,_ := strconv.ParseInt(pid, 10, 64)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	getUserResult:=DB.Collection("Post").FindOne(ctx, bson.M{"_id":id})
	if getUserResult.Err() != nil {
		http.Error(w, getUserResult.Err().Error(), http.StatusBadRequest)
		return
	}
	var post Post
	getUserResult.Decode(&post)
	json_data2, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_data2)
	return
}