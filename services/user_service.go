package services



type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var Users = []User{
    {ID: "1", Name: "Yamm"},
    {ID: "2", Name: "All"},
}

func GetAllUsers() []User {
	return Users
}