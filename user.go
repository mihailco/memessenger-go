package meme

type User struct {
	Id        int    `json: "id"`
	Password  string `json: "password"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
	Phone     string `json: "phone"`
	Email     string `json: "email"`
	Status    string `json: "status"`
	ImageURL  string `json: "imageURL"`
	Username  string `json: "username"`
	IsOnline  bool   `json: "isOnline"`
}
