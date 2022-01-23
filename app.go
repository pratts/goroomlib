package goroomlib

import (
	"fmt"
	"sync"
)

type App struct {
	roomService RoomService
	userService UserService
}

var lock = &sync.Mutex{}
var appInstance *App

func InitApp() *App {
	app := App{}
	roomService := RoomService{}
	roomService.Init()

	userService := UserService{}
	userService.Init()

	app.roomService = roomService
	app.userService = userService
	return &app
}

func (app *App) GetRoomService() RoomService {
	return app.roomService
}

func (app *App) GetUserService() UserService {
	return app.userService
}

func GetAppInstance() *App {
	if appInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if appInstance == nil {
			fmt.Println("Creting Single Instance Now")
			appInstance = &App{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return appInstance
}
