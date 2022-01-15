package goroomlib

type App struct {
	roomService RoomService
	userService UserService
}

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
