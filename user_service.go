package goroomlib

type UserService struct {
	UsernameMap map[string]User
	UserIdMap   map[int]User
	LastUserId  int
}

func (us *UserService) Init() {
	us.UsernameMap = make(map[string]User)
	us.UserIdMap = make(map[int]User)
}

func (us *UserService) CreateUser(userId int, name string, userProperties map[string]interface{}) User {
	us.LastUserId += 1
	user := User{Id: us.LastUserId, UserId: userId, Name: name}
	user.Init(userProperties)

	us.UsernameMap[name] = user
	us.UserIdMap[us.LastUserId] = user
	return user
}

func (us *UserService) GetUserByName(name string) (User, bool) {
	user, err := us.UsernameMap[name]
	return user, err
}

func (us *UserService) RemoveUserByName(name string) {
	user, err := us.UsernameMap[name]
	if err == false {
		user.Remove()
	}
}

func (us *UserService) RemoveUserById(id int) {
	user, err := us.UserIdMap[id]
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
