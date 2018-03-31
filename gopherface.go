package main

import (
	"net/http"
	"os"
	"log"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// "tutorials/backendwebdev/gopherface/endpoints"
	"tutorials/backendwebdev/gopherface/handlers"
	"tutorials/backendwebdev/gopherface/middleware"
	// "tutorials/backendwebdev/gopherface/common"
	// "tutorials/backendwebdev/gopherface/common/datastore"
)

const(
	WEBSERVERPORT = ":8443"
)

func main(){

	db, err := datastore.NewDatastore(datastore.MYSQL, "gopherface:gopherface@/gopherfacedb")

	if err != nil{
		log.Print(err)
	}
	defer db.Close()

	env := common.ENV{DB: db}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET,POST")
	r.HandleFunc("/login", handlers.LoginHandler(&env)).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	r.HandleFunc("/friends", handlers.FriendsHandler).Methods("GET")
	r.HandleFunc("/find", handlers.FindHandler).Methods("GET,POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/triggerpanic", handlers.TriggerPanicHandler).Methods("GET")
	r.HandleFunc("/foo", handlers.FooHandler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/postpreview", handlers.PostPreviewHandler).Methods("GET","POST")
	r.HandleFunc("/upload-image", handlers.UploadImageHandler).Methods("GET","POST")
	r.HandleFunc("/upload-video", handlers.UploadVideoHandler).Methods("GET","POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("restapi/socialmediapost/{username}", endpoints.FetchPostsEndpoint).Methods("GET")
	r.HandleFunc("restapi/socialmediapost/{postid}", endpoints.CreatePostEndpoint).Methods("POST")
	r.HandleFunc("restapi/socialmediapost/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
	r.HandleFunc("restapi/socialmediapost/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")
	
		// http.Handle("/", r)
		// http.Handle("/", ghandlers.LoggingHandler(os.Stdout, r))
		// http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))
		http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.PanicRecoveryHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	err = http.ListenAndServerTLS(WEBSERVERPORT, "certs/gopherfacecert.pem", "certs/gopherfacekey.pem", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}