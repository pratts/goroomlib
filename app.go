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

func (app *App) GetRoomService() *RoomService {
	return &app.roomService
}

func (app *App) GetUserService() *UserService {
	return &app.userService
}
