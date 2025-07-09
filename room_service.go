package goroomlib

import "sync"

type RoomService struct {
	mu              sync.RWMutex
	roomMap         map[string]*Room
	numRoomsCreated int
}

func (rs *RoomService) Init() {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.roomMap = make(map[string]*Room)
	rs.numRoomsCreated = 0
}

func (rs *RoomService) GetRoomMap() map[string]*Room {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return rs.roomMap
}

func (rs *RoomService) CreateRoom(roomName string, maxUsers int) *Room {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	room := &Room{id: rs.numRoomsCreated + 1, name: roomName, maxUserCount: maxUsers}
	room.Init()
	rs.numRoomsCreated += 1
	rs.addRoom(room)
	return room
}

func (rs *RoomService) addRoom(room *Room) {
	rs.roomMap[room.GetRoomName()] = room
}

func (rs *RoomService) RemoveRoom(roomName string) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	room, ok := rs.roomMap[roomName]
	if ok {
		room.RemoveAllUsers()
	}
	delete(rs.roomMap, roomName)
}

func (rs *RoomService) GetRoomByName(roomName string) (*Room, bool) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	room, ok := rs.roomMap[roomName]
	return room, ok
}

func (rs *RoomService) GetUserForRoom(roomName string) map[string]*User {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	room, ok := rs.roomMap[roomName]
	if ok {
		return room.GetUserMap()
	}
	return nil
}

// AddUserToRoom atomically adds a user to a room and the room to the user's joinedRooms.
// Locks Room first, then User to avoid deadlocks.
func (rs *RoomService) AddUserToRoom(user *User, roomName string) bool {
	rs.mu.RLock() // Only reading roomMap
	room, ok := rs.roomMap[roomName]
	rs.mu.RUnlock()
	if !ok {
		return false
	}

	// Lock Room first, then User (consistent order)
	room.mu.Lock()
	defer room.mu.Unlock()
	user.mu.Lock()
	defer user.mu.Unlock()

	userCount := len(room.usersMap)
	if userCount == room.maxUserCount {
		return false
	}
	// Only add if not already present
	if _, exists := room.usersMap[user.name]; exists {
		return false
	}
	room.usersMap[user.name] = user
	user.joinedRooms[room.name] = room
	return true
}

// RemoveUserFromRoom atomically removes a user from a room and the room from the user's joinedRooms.
// Locks Room first, then User to avoid deadlocks.
func (rs *RoomService) RemoveUserFromRoom(user *User, roomName string) bool {
	rs.mu.RLock() // Only reading roomMap
	room, ok := rs.roomMap[roomName]
	rs.mu.RUnlock()
	if !ok {
		return false
	}

	room.mu.Lock()
	defer room.mu.Unlock()
	user.mu.Lock()
	defer user.mu.Unlock()

	if _, exists := room.usersMap[user.name]; !exists {
		return false
	}
	delete(room.usersMap, user.name)
	delete(user.joinedRooms, room.name)
	return true
}
