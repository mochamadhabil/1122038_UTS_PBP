package model

type User struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
}

type Product struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Price int    `json:"Price"`
}

type Transaction struct {
	ID       int `json:"ID"`
	User     User
	Product  Product
	Quantity string `json:"Quantity"`
}

type TransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"Data"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResultUser struct {
	ID      int
	Name    string
	Age     int
	Address string
}

// BARU
type Account struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
}

type Games struct {
	ID        int    `json:"ID"`
	Name      string `json:"Name"`
	MaxPlayer int    `json:"Max Player"`
}

type Rooms struct {
	ID        int    `json:"ID"`
	Room_Name string `json:"Room Name"`
	Id_Game   int    `json:"Id Game"`
}

type Participants struct {
	ID         int    `json:"ID"`
	Id_Account int    `json:"ID Account"`
	Username   string `json:"Username"`
}

// /// 1
type AllRoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

// // 2
type DetailRooms struct {
	Rooms      Rooms        `json:"room"`
	Prticipans Participants `json:"participants"`
}

type DetailRoomsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []DetailRooms `json:"data"`
}
