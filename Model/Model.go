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

