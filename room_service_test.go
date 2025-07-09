package goroomlib

import (
	"testing"
)

func TestRoomService_CreateRoomAndAddUser(t *testing.T) {
	rs := &RoomService{}
	rs.Init()

	r := rs.CreateRoom("testroom", 2)
	if r == nil {
		t.Fatal("expected room to be created")
	}

	u := &User{name: "alice"}
	u.Init()

	ok := rs.AddUserToRoom(u, "testroom")
	if !ok {
		t.Fatal("expected user to be added to room")
	}

	room, found := rs.GetRoomByName("testroom")
	if !found {
		t.Fatal("room should exist")
	}
	if _, exists := room.GetUserMap()["alice"]; !exists {
		t.Error("user should be in room's user map")
	}
	if _, exists := u.GetJoinedRooms()["testroom"]; !exists {
		t.Error("room should be in user's joined rooms")
	}
}

func TestRoomService_AddUserToFullRoom(t *testing.T) {
	rs := &RoomService{}
	rs.Init()

	_ = rs.CreateRoom("fullroom", 1)
	u1 := &User{name: "bob"}
	u1.Init()
	u2 := &User{name: "carol"}
	u2.Init()

	ok1 := rs.AddUserToRoom(u1, "fullroom")
	ok2 := rs.AddUserToRoom(u2, "fullroom")

	if !ok1 {
		t.Error("first user should be added")
	}
	if ok2 {
		t.Error("second user should not be added to full room")
	}
}

func TestRoomService_RemoveUserFromRoom(t *testing.T) {
	rs := &RoomService{}
	rs.Init()

	_ = rs.CreateRoom("deleteroom", 2)
	u := &User{name: "dave"}
	u.Init()

	rs.AddUserToRoom(u, "deleteroom")
	ok := rs.RemoveUserFromRoom(u, "deleteroom")
	if !ok {
		t.Fatal("expected user to be removed from room")
	}
	room, _ := rs.GetRoomByName("deleteroom")
	if _, exists := room.GetUserMap()["dave"]; exists {
		t.Error("user should not be in room's user map after removal")
	}
	if _, exists := u.GetJoinedRooms()["deleteroom"]; exists {
		t.Error("room should not be in user's joined rooms after removal")
	}
}

func TestRoomService_RemoveNonExistentUser(t *testing.T) {
	rs := &RoomService{}
	rs.Init()

	_ = rs.CreateRoom("emptyroom", 2)
	u := &User{name: "eve"}
	u.Init()

	ok := rs.RemoveUserFromRoom(u, "emptyroom")
	if ok {
		t.Error("should not remove user who is not in the room")
	}
}

func TestRoomService_ConcurrentAddRemove(t *testing.T) {
	rs := &RoomService{}
	rs.Init()

	room := rs.CreateRoom("concurrent", 100)
	users := make([]*User, 100)
	for i := 0; i < 100; i++ {
		users[i] = &User{name: string(rune('a'+i%26)) + string(rune('A'+(i/26)))}
		users[i].Init()
	}

	done := make(chan bool)

	// Concurrently add users
	for _, u := range users {
		u := u
		go func() {
			ok := rs.AddUserToRoom(u, "concurrent")
			if !ok {
				t.Errorf("failed to add user %s", u.name)
			}
			done <- true
		}()
	}
	for range users {
		<-done
	}

	// Concurrently remove users
	for _, u := range users {
		u := u
		go func() {
			ok := rs.RemoveUserFromRoom(u, "concurrent")
			if !ok {
				t.Errorf("failed to remove user %s", u.name)
			}
			done <- true
		}()
	}
	for range users {
		<-done
	}

	if len(room.GetUserMap()) != 0 {
		t.Errorf("expected all users to be removed, got %d", len(room.GetUserMap()))
	}
}
