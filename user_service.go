package goroomlib

import "sync"

type UserService struct {
	mu          sync.RWMutex
	usernameMap map[string]*User
	userIdMap   map[int]*User
	lastUserId  int
}

func (us *UserService) Init() {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.usernameMap = make(map[string]*User)
	us.userIdMap = make(map[int]*User)
}

func (us *UserService) CreateUser(userId int, name string) *User {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.lastUserId += 1
	user := &User{id: us.lastUserId, userId: userId, name: name}
	user.Init()

	us.usernameMap[name] = user
	us.userIdMap[us.lastUserId] = user
	return user
}

func (us *UserService) GetUserByName(name string) (*User, bool) {
	us.mu.RLock()
	defer us.mu.RUnlock()
	user, err := us.usernameMap[name]
	return user, err
}

func (us *UserService) RemoveUserByName(name string) {
	us.mu.Lock()
	defer us.mu.Unlock()
	user, err := us.usernameMap[name]
	if !err {
		user.Remove()
	}
}

func (us *UserService) RemoveUserById(id int) {
	us.mu.Lock()
	defer us.mu.Unlock()
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
