package goroomlib

type RoomService struct {
	roomMap         map[string]Room
	numRoomsCreated int
}

func (rs *RoomService) Init() {
	rs.roomMap = make(map[string]Room)
	rs.numRoomsCreated = 0
}

func (rs *RoomService) GetRoomMap() map[string]Room {
	return rs.roomMap
}

func (rs *RoomService) CreateRoom(roomName string, maxUsers int) Room {
	room := Room{id: rs.numRoomsCreated + 1, name: roomName, maxUserCount: maxUsers}
	room.Init()
	rs.numRoomsCreated += 1
	rs.addRoom(room)
	return room
}

func (rs *RoomService) addRoom(room Room) {
	rs.roomMap[room.GetRoomName()] = room
}

func (rs *RoomService) RemoveRoom(roomName string) {
	room, ok := rs.roomMap[roomName]
	if ok {
		room.RemoveAllUsers()
	}
	delete(rs.roomMap, roomName)
}

func (rs *RoomService) GetRoomByName(roomName string) (Room, bool) {
	room, ok := rs.roomMap[roomName]
	return room, ok
}

func (rs *RoomService) GetUserForRoom(roomName string) map[string]User {
	room, ok := rs.roomMap[roomName]
	if ok {
		return room.GetUserMap()
	}
	return nil
}

func (rs *RoomService) AddUserToRoom(user User, roomName string) bool {
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

func (rs *RoomService) RemoveUserFromRoom(user User, roomName string) bool {
	room, ok := rs.roomMap[roomName]
	if ok {
		room.RemoveUserFromRoom(user)
		return true
	}
	return false
}
