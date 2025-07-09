package goroomlib

import "sync"

type User struct {
	mu          sync.RWMutex
	id          int
	userId      int
	name        string
	joinedRooms map[string]*Room
	isConnected bool
}

func (u *User) Init() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.joinedRooms = make(map[string]*Room)
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetUserId() int {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetIsConnected(isConnected bool) {
	u.isConnected = isConnected
}

func (u *User) GetIsConnected() bool {
	return u.isConnected
}

// GetJoinedRooms returns a copy of the joinedRooms map to avoid exposing internal state.
func (u *User) GetJoinedRooms() map[string]*Room {
	u.mu.RLock()
	defer u.mu.RUnlock()
	copyMap := make(map[string]*Room, len(u.joinedRooms))
	for k, v := range u.joinedRooms {
		copyMap[k] = v
	}
	return copyMap
}

// GetJoinedRoomNames returns a slice of room names the user has joined.
func (u *User) GetJoinedRoomNames() []string {
	u.mu.RLock()
	defer u.mu.RUnlock()
	names := make([]string, 0, len(u.joinedRooms))
	for name := range u.joinedRooms {
		names = append(names, name)
	}
	return names
}

// addRoom is unexported; atomic add should be done via RoomService
func (u *User) addRoom(r *Room) map[string]*Room {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.joinedRooms[r.GetRoomName()] = r
	return u.joinedRooms
}

// removeRoom is unexported; atomic remove should be done via RoomService
func (u *User) removeRoom(r *Room) map[string]*Room {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.joinedRooms, r.GetRoomName())
	return u.joinedRooms
}

func (u *User) GetJoinedRoomByName(roomName string) (*Room, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	room, ok := u.joinedRooms[roomName]
	return room, ok
}

func (u *User) Remove() {
	u.DisconnectUser()
}

func (u *User) DisconnectUser() {
	u.mu.Lock()
	defer u.mu.Unlock()
	for _, room := range u.joinedRooms {
		room.removeUserFromRoom(u)
	}
}
