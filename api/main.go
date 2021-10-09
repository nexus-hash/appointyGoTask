package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"social/api/routes"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func checkMethodType(method string,handler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w,r)
	}
}



func main() {
	mux:= http.NewServeMux()

	mux.HandleFunc("/users",checkMethodType("POST",routes.UserHandler));
	mux.HandleFunc("/users/:id",checkMethodType("GET",routes.GetUserHandler));

	mux.HandleFunc("/posts",checkMethodType("POST",routes.PostHandler));
	mux.HandleFunc("/posts/:id",checkMethodType("GET",routes.GetPostHandler));
	mux.HandleFunc("/posts/users/:id",checkMethodType("GET",routes.GetUserPostsHandler));

	db,err := DatabaseConnection();
	if err != nil {
		log.Fatal("Database Connection Error $s",err)
	}
	fmt.Println("Database connection Success!")

	UserCollection := db.Collection("user")
	fmt.Println("User Collection Success!",UserCollection);

	error := http.ListenAndServe(os.Getenv("PORT"), mux);
	if error != nil {
		log.Fatal(error)
	}
}

func DatabaseConnection()(*mongo.Database,error){
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println(os.Getenv("MONGO_URI"));
	client,err := mongo.Connect(ctx,options.Client().ApplyURI(os.Getenv("MONGO_URI")));

	if err != nil {
		return nil,err
	}
	
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(os.Getenv("DATABASE_NAME"))

	return database,nil
}