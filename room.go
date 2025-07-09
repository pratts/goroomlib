package goroomlib

import "sync"

// Room represents a chat/game room. All access to usersMap is protected by mu for thread safety.
type Room struct {
	mu           sync.RWMutex // protects usersMap
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

// GetUserMap returns a copy of the users map to avoid exposing internal state. Safe for concurrent use.
func (r *Room) GetUserMap() map[string]*User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	copyMap := make(map[string]*User, len(r.usersMap))
	for k, v := range r.usersMap {
		copyMap[k] = v
	}
	return copyMap
}

// GetUserNames returns a slice of user names in the room.
func (r *Room) GetUserNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.usersMap))
	for name := range r.usersMap {
		names = append(names, name)
	}
	return names
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

// addUserToRoom is unexported; atomic add should be done via RoomService
func (r *Room) addUserToRoom(u *User) map[string]*User {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.usersMap[u.name] = u
	u.addRoom(r)
	return r.usersMap
}

// removeUserFromRoom is unexported; atomic remove should be done via RoomService
func (r *Room) removeUserFromRoom(u *User) map[string]*User {
	r.mu.Lock()
	defer r.mu.Unlock()
	u.removeRoom(r)
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
		r.removeUserFromRoom(user)
	}
}
