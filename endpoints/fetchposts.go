package endpoints

import(
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"tutorials/backendwebdev/gopherface/models/socialmedia"
)

func FetchPostsEndPoint(w http.ResponseWriter, r *http.Request){
	v := mux.Vars(r)

	if v["username"] == "username" {
		mockPosts := make([]socialmedia.Post, 3)

		post1:= socialmedia.NewPost("username", socialmedia.Moods["thrilled"], "Go is neat!", "Check out the Go website!", "https://golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post2:=socialmedia.NewPost("username", socialmedia.Moods["thrilled"], "Go is neat!", "Check out the Go website!", "https://golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
		post3:=socialmedia.NewPost("username", socialmedia.Moods["thrilled"], "Go is neat!", "Check out the Go website!", "https://golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})

		mockPosts = append(mockPosts, *post1)
		mockPosts = append(mockPosts, *post2)
		mockPosts = append(mockPosts, *post3)
		json.NewEncoder(w).Encode(mockPosts)
	} else{
		json.NewEncoder(w).Encode(nil)
	}
}