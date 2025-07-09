package goroomlib

type App struct {
	roomService RoomService
	userService UserService
}

func InitApp() *App {
	app := App{
		roomService: RoomService{},
		userService: UserService{},
	}
	app.roomService.Init()
	app.userService.Init()

	return &app
}

func (app *App) GetRoomService() *RoomService {
	return &app.roomService
}

func (app *App) GetUserService() *UserService {
	return &app.userService
}
