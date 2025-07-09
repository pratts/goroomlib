package goroomlib

type UserService struct {
	usernameMap map[string]User
	userIdMap   map[int]User
	lastUserId  int
}

func (us *UserService) Init() {
	us.usernameMap = make(map[string]User)
	us.userIdMap = make(map[int]User)
}

func (us *UserService) CreateUser(userId int, name string) User {
	us.lastUserId += 1
	user := User{id: us.lastUserId, userId: userId, name: name}
	user.Init()

	us.usernameMap[name] = user
	us.userIdMap[us.lastUserId] = user
	return user
}

func (us *UserService) GetUserByName(name string) (User, bool) {
	user, err := us.usernameMap[name]
	return user, err
}

func (us *UserService) RemoveUserByName(name string) {
	user, err := us.usernameMap[name]
	if !err {
		user.Remove()
	}
}

func (us *UserService) RemoveUserById(id int) {
	user, err := us.userIdMap[id]
	if !err {
		user.Remove()
	}
}

func (us *UserService) DisconnectUserByName(name string) {
	us.RemoveUserByName(name)
}

func (us *UserService) DisconnectUserById(id int) {
	us.RemoveUserById(id)
}
