#!/bin/bash
go build -ldflags="-w -s" -o builds/forum-app main.go main-natasha.go
chmod +x builds/forum-app