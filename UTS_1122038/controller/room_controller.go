package controller

import (
	m "UTS_1122038/model"
	"database/sql"
	"encoding/json"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT ID, Room_Name FROM Rooms`

	detailRoomsRow, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 500, "StatusBadRequest")
		return
	}

	var rooms []m.Rooms
	for detailRoomsRow.Next() {

		var room m.Rooms

		if err := detailRoomsRow.Scan(&room.ID, &room.Room_Name); err != nil {
			SendErrorResponse(w, 500, "StatusBadRequest")
			return
		}
		rooms = append(rooms, room)
	}

	if len(rooms) == 0 {
		SendErrorResponse(w, 400, "Room Tidak Ditemukan")
		return
	}

	SendSuccessAllRoomsResponse(w, rooms)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT r.id, r.room_name, p.id, a.id, a.username 
	FROM Rooms r JOIN Participants p ON r.id = p.id_room JOIN Accounts a ON p.id_account = a.id`

	detailRoomsRow, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 500, "Kesalahan pada query ")
		return
	}

	var detailAllRoom m.DetailRooms
	var detailAllRooms []m.DetailRooms
	for detailRoomsRow.Next() {

		if err := detailRoomsRow.Scan(&detailAllRoom.Rooms.ID, &detailAllRoom.Rooms.Room_Name, &detailAllRoom.Prticipans.ID, &detailAllRoom.Prticipans.Id_Account, &detailAllRoom.Prticipans.Username); err != nil {
			SendErrorResponse(w, 500, "Gagal Mendapatkan Data ")
			return
		} else {
			detailAllRooms = append(detailAllRooms, detailAllRoom)
		}
	}

	if len(detailAllRooms) == 0 {
		SendErrorResponse(w, 500, "Detail Rooms Tidak Ditemukan")
		return
	}

	SendSuccessDetailRoomsResponse(w, detailAllRooms)

}

