# Goroomlib

This is a basic library with basic Room-User architecture in Golang.

To import:
  `go get github.com/pratts/goroomlib`


The library contains the following services:
  1. **User Service** : To maintain the users connected/joined into the system
  2. **Room Service** : To maintain the rooms created on the server.

Following are the models that library provides:
  1. **User** : Every user will have basic details like userId, name and properties attached to it. Apart from these details, a user can have a list of room it has joined
  2. **Room** : Every room will have an auto-generated id, name, map of users who joined that room and custom properties attached to it.

Library Usage:
1. Importing

  `
    package main

    import (
      "fmt"
      "github.com/pratts/goroomlib"
    )

    func main() {
      app := goroomlib.InitApp()
    }
`