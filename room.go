package goroomlib

import (
	"encoding/json"
)

type Room struct {
	Id           int
	Name         string
	UsersMap     map[string]User
	MaxUserCount int
	Properties   RoomProperties
}

type RoomProperties struct {
	IsPassword bool   `json:"isPasswordProtected"`
	Password   string `json:"password"`
}

func (r *Room) Init(Properties map[string]interface{}) {
	r.createUserMap()
	if Properties == nil || len(Properties) == 0 {
		Properties = make(map[string]interface{})
		Properties["IsPassword"] = false
	}
	jsonbody, err := json.Marshal(Properties)
	if err != nil {
		return
	}

	r.Properties = RoomProperties{}
	json.Unmarshal(jsonbody, &(r.Properties))
}

func (r *Room) GetId() int {
	return r.Id
}

func (r *Room) GetRoomName() string {
	return r.Name
}

func (r *Room) createUserMap() map[string]User {
	r.UsersMap = make(map[string]User)
	return r.UsersMap
}

func (r *Room) GetUserMap() map[string]User {
	return r.UsersMap
}

func (r *Room) GetMaxUserCount() int {
	return r.MaxUserCount
}

func (r *Room) GetUserByName(userName string) (User, bool) {
	user, ok := r.UsersMap[userName]
	return user, ok
}

func (r *Room) AddUserToRoom(u User) map[string]User {
	r.UsersMap[u.Name] = u
	u.AddRoom(*r)
	return r.UsersMap
}

func (r *Room) RemoveUserFromRoom(u User) map[string]User {
	u.RemoveRoom(*r)
	delete(r.UsersMap, u.GetName())
	return r.UsersMap
}

func (r *Room) ClearUsers() map[string]User {
	return r.createUserMap()
}

func (r *Room) RemoveAllUsers() {
	for _, user := range r.UsersMap {
		r.RemoveUserFromRoom(user)
	}
}
