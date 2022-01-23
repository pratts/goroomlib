package goroomlib

import "encoding/json"

type User struct {
	Id          int
	UserId      int
	Name        string
	JoinedRooms map[string](bool)
	Properties  UserProperties
	IsConnected bool
}

type UserProperties struct {
	IsAdmin bool `json:"isAdmin"`
}

func (u *User) Init(Properties map[string]interface{}) {
	u.JoinedRooms = make(map[string]bool)
	if Properties == nil || len(Properties) == 0 {
		Properties = make(map[string]interface{})
		Properties["IsAdmin"] = false
	}
	jsonbody, err := json.Marshal(Properties)
	if err != nil {
		return
	}

	u.Properties = UserProperties{}
	json.Unmarshal(jsonbody, &(u.Properties))
}

func (u *User) GetId() int {
	return u.Id
}

func (u *User) GetUserId() int {
	return u.UserId
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetIsConnected(IsConnected bool) {
	u.IsConnected = IsConnected
}

func (u *User) GetIsConnected() bool {
	return u.IsConnected
}

func (u *User) GetJoinedRooms() map[string]bool {
	return u.JoinedRooms
}

func (u *User) AddRoom(r Room) map[string]bool {
	u.JoinedRooms[r.GetRoomName()] = true
	return u.JoinedRooms
}

func (u *User) RemoveRoom(r Room) map[string]bool {
	delete(u.GetJoinedRooms(), r.GetRoomName())
	return u.JoinedRooms
}

func (u *User) GetJoinedRoomByName(roomName string) (bool, bool) {
	room, ok := u.GetJoinedRooms()[roomName]
	return room, ok
}

func (u *User) Remove() {
	u.DisconnectUser()
}

func (u *User) DisconnectUser() {
	for roomName, _ := range u.JoinedRooms {
		roomService := GetAppInstance().GetRoomService()
		room, err := roomService.GetRoomByName(roomName)
		if err == true {
			(&room).RemoveUserFromRoom(*u)
		}
	}
}
