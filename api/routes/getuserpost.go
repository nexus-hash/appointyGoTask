package routes

// Neccessary imports for the function
import(
	"net/http"
	"strings"
	"encoding/json"
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// Request Body Format
type GetUserPost struct{
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

// Get Response Format
type Result struct{
	Id          string                `json:"_id"`
	Caption     string                `json:"caption"`
	ImageUrl    string                `json:"imageurl"`
	PostedAt    time.Time             `json:"postedat"`
	UserId      primitive.ObjectID    `json:"userid"`
}

func  GetUserPostsHandler(w http.ResponseWriter, r *http.Request){

	/*
		Get User Posts
		Takes Page No. and Limit as input
		Returns the list of posts of the user for the given page no.
		For Respose Data Refer README.md
	*/


	// Extract user id from the url
	uid := strings.TrimPrefix(r.URL.Path, "/posts/users/")
	fmt.Println(uid)
	if uid == ""{
		http.Error(w, "Invalid User Id", http.StatusNotFound)
		return
	}

	// Decode request Body
	var details GetUserPost
	if err:= json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Convert user id from hex string to ObjectId
	uuid,err:= primitive.ObjectIDFromHex(uid);
	if err != nil {
		http.Error(w, "Invalid User Id", http.StatusNotFound)
		return
	}

	// Set findOptions
	findOptions:= options.Find();

	// Set limit and skip
	findOptions.SetLimit(int64(details.Limit));
	findOptions.SetSkip(int64((details.Page-1) * details.Limit));

	// Set sort order
	findOptions.SetSort(bson.M{"postedat": -1});

	// filter by user id
	filter:= bson.M{"userid": uuid}; 
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Find the posts
	queryResult,err:= DB.Collection("Post").Find(ctx, filter, findOptions);
	if err != nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var results []Result

	// Iterate over the results decoding into struct and append to results array
	for queryResult.Next(ctx){
		var result Result
		err:= queryResult.Decode(&result);
		if err != nil{
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		results = append(results, result)
	}

	// If no results found return No More Posts
	if results == nil{
		http.Error(w, "No Posts Found", http.StatusNotFound)
		return
	}

	// Encode the results into json
	json_data2,error := json.Marshal(results)
	if error != nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(error)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_data2)
	return
}