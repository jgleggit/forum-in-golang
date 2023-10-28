package main

import(
    "time"
)

type User struct {
    userUUID   	 string
    Name 	 string
    Email	 string
    Password     string
    PostsIDs     []int
}

type Post struct {
    postUUID      string
    titlePost     string
    contentPost   string
    userID  	  int
    isAuthor      bool //check if the viewer is author
    postDate      string //have a look @ the Go package "time", what data type is there for time?
    likesCount	  int
    dislikesCount int
    commentsCount int
    categories	  []string // one to many

}

type Comment struct {
    commentUUID  	string
    postID     		int
    userID     		int
    commentContent   	string
    commentDate   	string //have a look @ the Go package "time", what data type is there for time?
    likesCount	  	int
    dislikesCount 	int
    isAuthor      	bool //check if the viewer is author
}

type Session struct {
	sessionUUID    	string
	token  		string
        expiry   	time.Time // need to import "time"
	userUUID 	string
	IsAuthorized 	bool
}


type Error struct {
	ErrorMessage string
}
