package goroomlib

type RoomService struct {
	RoomMap         map[string]Room
	NumRoomsCreated int
}

func (rs *RoomService) Init() {
	rs.RoomMap = make(map[string]Room)
	rs.NumRoomsCreated = 0
}

func (rs *RoomService) GetRoomMap() map[string]Room {
	return rs.RoomMap
}

func (rs *RoomService) CreateRoom(roomName string, maxUsers int, roomProperties map[string]interface{}) Room {
	room := Room{Id: rs.NumRoomsCreated + 1, Name: roomName, MaxUserCount: maxUsers}
	room.Init(roomProperties)
	rs.NumRoomsCreated += 1
	rs.addRoom(room)
	return room
}

func (rs *RoomService) addRoom(room Room) {
	rs.RoomMap[room.GetRoomName()] = room
}

func (rs *RoomService) RemoveRoom(roomName string) {
	room, ok := rs.RoomMap[roomName]
	if ok == true {
		room.RemoveAllUsers()
	}
	delete(rs.RoomMap, roomName)
}

func (rs *RoomService) GetRoomByName(roomName string) (Room, bool) {
	room, ok := rs.RoomMap[roomName]
	return room, ok
}

func (rs *RoomService) GetUserForRoom(roomName string) map[string]User {
	room, ok := rs.RoomMap[roomName]
	if ok == true {
		return room.GetUserMap()
	}
	return nil
}

func (rs *RoomService) AddUserToRoom(user User, roomName string) bool {
	room, ok := rs.RoomMap[roomName]
	if ok == true {
		userCount := len(room.UsersMap)
		if userCount == room.MaxUserCount {
			return false
		}
		room.AddUserToRoom(user)
		return true
	}
	return false
}

func (rs *RoomService) RemoveUserFromRoom(user User, roomName string) bool {
	room, ok := rs.RoomMap[roomName]
	if ok == true {
		room.RemoveUserFromRoom(user)
		return true
	}
	return false
}
