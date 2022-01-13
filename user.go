package goroomlib

type User struct {
	id             int
	userId         int
	name           string
	joinedRooms    map[string]Room
	userProperties map[string]interface{}
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

func (u *User) GetJoinedRooms() map[string]Room {
	return u.joinedRooms
}

func (u *User) AddRoom(r Room) map[string]Room {
	u.joinedRooms[r.GetRoomName()] = r
	return u.joinedRooms
}

func (u *User) RemoveRoom(r Room) map[string]Room {
	delete(u.GetJoinedRooms(), r.GetRoomName())
	return u.joinedRooms
}

func (u *User) GetJoinedRoomByName(roomName string) (Room, bool) {
	room, ok := u.GetJoinedRooms()[roomName]
	return room, ok
}

func (u *User) Remove() {
	u.DisconnectUser()
}

func (u *User) DisconnectUser() {
	for _, room := range u.joinedRooms {
		room.RemoveUserFromRoom(*u)
	}
}
