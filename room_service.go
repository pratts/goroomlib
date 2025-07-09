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

func (rs *RoomService) AddUserToRoom(user *User, roomName string) bool {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	room, ok := rs.roomMap[roomName]
	if ok {
		userCount := len(room.usersMap)
		if userCount == room.maxUserCount {
			return false
		}
		room.AddUserToRoom(user)
		return true
	}
	return false
}

func (rs *RoomService) RemoveUserFromRoom(user *User, roomName string) bool {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	room, ok := rs.roomMap[roomName]
	if ok {
		room.RemoveUserFromRoom(user)
		return true
	}
	return false
}
