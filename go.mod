module forum-in-golang

go 1.20

replace forum-in-golang/logger => ./logger

replace forum-in-golang/filelogger => ./filelogger

require (
	forum-in-golang/filelogger v0.0.0
	forum-in-golang/logger v0.0.0
)

require github.com/mattn/go-sqlite3 v1.14.17

require golang.org/x/crypto v0.14.0
