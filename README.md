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
    // Initializing the App
	app := goroomlib.InitApp()

	/*
		Creating a Room with following properties:
			1. Name : "test"
			2. Max users allowed : 10
			3. Room Properties : Is password protected with "test" password
	*/
	roomService := app.GetRoomService()
	roomProperties := make(map[string]interface{})
	roomProperties["isPasswordProtected"] = true
	roomProperties["password"] = "test"
	(&roomService).CreateRoom("test", 10, roomProperties)

	// Fetching a room by name
	room, isFound := (&roomService).GetRoomByName("test")

	/*
		Creating a Room with following properties:
			1. User ID : 1
			2. Name : "pratts"
			3. User Properties : With isAdmin flag to true
	*/
	userService := app.GetUserService()
	userProperties := make(map[string]interface{})
	userProperties["isAdmin"] = true
	user := (&userService).CreateUser(1, "pratts", userProperties)

	// Adding a user to a room
	roomService.AddUserToRoom(user, "test")

	// Removing a room and all users in it
	(&roomService).RemoveRoom("test")

	// Fetching all the users in a room
	users := (&roomService).GetUserForRoom("test")

    // Removing user from  a room
    (&roomService).GetUserForRoom(user, "test")
`