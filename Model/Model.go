package Model

type User struct {
	Uniqueid     string
	Name       string
	Rate       int8
	Skills     []string
	Favourites []int32
	Ignored    []int32
	Busy       bool
}