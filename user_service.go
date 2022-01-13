package goroomlib

type UserService struct {
	usernameMap map[string]User
	userIdMap   map[int]User
	lastUserId  int
}

func (us *UserService) initUserMaps() {
	us.usernameMap = make(map[string]User)
	us.userIdMap = make(map[int]User)
}

func (us *UserService) createUser(userId int, name string, userProperties map[string]interface{}) User {
	us.lastUserId += 1
	user := User{id: us.lastUserId, userId: userId, name: name, userProperties: userProperties}
	user.joinedRooms = make(map[string]Room)

	us.usernameMap[name] = user
	us.userIdMap[us.lastUserId] = user
	return user
}

func (us *UserService) RemoveUserByName(name string) {
	user, err := us.usernameMap[name]
	if err == false {
		user.Remove()
	}
}

func (us *UserService) RemoveUserById(id int) {
	user, err := us.userIdMap[id]
	if err == false {
		user.Remove()
	}
}

func (us *UserService) DisconnectUserByName(name string) {
	us.RemoveUserByName(name)
}

func (us *UserService) DisconnectUserById(id int) {
	us.RemoveUserById(id)
}
