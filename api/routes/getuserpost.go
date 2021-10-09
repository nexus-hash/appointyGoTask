package routes

import(
	"net/http"
	"strings"
)

func  GetUserPostsHandler(w http.ResponseWriter, r *http.Request){
	uid := strings.TrimPrefix(r.URL.Path, "/posts/users/")
	if uid != ""{
		http.Error(w, "Invalid User Id", http.StatusNotFound)
		return
	}
	
	return
}