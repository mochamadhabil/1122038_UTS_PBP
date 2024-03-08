package controller

import (
	"encoding/json"
	"net/http"

	m "UTS_1122038/model"
)

func SendSuccessAllRoomsResponse(w http.ResponseWriter, data []m.Rooms) {
	w.Header().Set("Content-Type", "application/json")
	var response m.AllRoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = data

	json.NewEncoder(w).Encode(response)
}

func SendSuccessDetailRoomsResponse(w http.ResponseWriter, data []m.DetailRooms) {
	w.Header().Set("Content-Type", "application/json")
	var response m.DetailRoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = data

	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}
