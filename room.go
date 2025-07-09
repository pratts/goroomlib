package goroomlib

import "sync"

type Room struct {
	mu           sync.RWMutex
	id           int
	name         string
	usersMap     map[string]*User
	maxUserCount int
}

func (r *Room) Init() {
	r.createUserMap()
}

func (r *Room) GetId() int {
	return r.id
}

func (r *Room) GetRoomName() string {
	return r.name
}

func (r *Room) createUserMap() map[string]*User {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.usersMap = make(map[string]*User)
	return r.usersMap
}

func (r *Room) GetUserMap() map[string]*User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.usersMap
}

func (r *Room) GetMaxUserCount() int {
	return r.maxUserCount
}

func (r *Room) GetUserByName(userName string) (*User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.usersMap[userName]
	return user, ok
}

func (r *Room) AddUserToRoom(u *User) map[string]*User {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.usersMap[u.name] = u
	u.AddRoom(r)
	return r.usersMap
}

func (r *Room) RemoveUserFromRoom(u *User) map[string]*User {
	r.mu.Lock()
	defer r.mu.Unlock()
	u.RemoveRoom(r)
	delete(r.usersMap, u.GetName())
	return r.usersMap
}

func (r *Room) ClearUsers() map[string]*User {
	return r.createUserMap()
}

func (r *Room) RemoveAllUsers() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, user := range r.usersMap {
		r.RemoveUserFromRoom(user)
	}
}
