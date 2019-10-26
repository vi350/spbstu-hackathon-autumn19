package Model

type User struct {
	Uniqueid     string
	Name       string
	Rate       int8
	Skills     string
	Favourites string
	Ignored    string
	Busy       bool
}

// user struct for select fro
type UserS struct {
	Id int
	Uniqueid     string
	Name       string
	Rate       int8
	Skills     string
	Favourites string
	Ignored    string
	Busy       bool
}

type DataInUser struct {
	Id   string  `json:"id"  binding:"required"`
	FirstName     string  `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Username       string `json:"username" binding:"required"`
	PhotoUrl string  `json:"photo_url" binding:"required"`
	AuthDate   string  `json:"auth_date" binding:"required"`
	Hash       string  `json:"hash" binding:"required"`
}


