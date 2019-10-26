package DB

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	//"github.com/go-pg/pg/orm"
	_ "github.com/go-pg/pg/orm"
	"github.com/vi350/spbstu-hackathon-autumn19/Model"
	"github.com/vi350/spbstu-hackathon-autumn19/Passwords"
	"log"
)

var DB *pg.DB

// подключение к бд
func ConnectDB () {
	DB = pg.Connect(&pg.Options{
		Addr:     "89.208.196.56:5432",
		User:     "dima",
		Password: Passwords.DBPass,
		Database: "keyzu",
	})
	Status()
}

// проверка подключения к бд
func Status() bool{
	var status bool
	_, err := DB.Exec("SELECT 1")
	if err != nil {
		status = false
		fmt.Println("Подключение к базе данных не удалось")
		log.Println(err)
	} else {
		status = true
		fmt.Println("Подключение к базе данных успешно")
	}
	return status
}

// создание таблиц и полей
func CreateTables () {
	db := DB
	qs := []string{
		/* language=PostgreSQL */
		`CREATE TABLE IF NOT EXISTS users(
         id SERIAL PRIMARY KEY,
         uniqueid text ,
    name text,
    rate numeric,
    skills text,
    favourites text,
    ignored text,
    busy bool
  )`,

	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			fmt.Println("Таблица не создана")
			log.Panic(err)
		} else {
			fmt.Println("Таблица успешно создана или уже есть")
		}
	}

	var m Model.User = Model.User{"pdrs","oleg",5,jsoniseStrs([]string{"c","arduino"}),jsoniseInts([]int{2,6}),jsoniseInts([]int{5,7,9}),false}

	err := db.Insert(&m)
	if err !=nil{
		fmt.Println("bleat")
		fmt.Println(err)
	}


}


func jsoniseStrs(arr []string)string  {
	slc,_ := json.Marshal(arr)
	return string(slc)

}

func jsoniseInts(arr []int)string  {
	slc,_ := json.Marshal(arr)
	return string(slc)
}



func SelectBySkills(skills []string,id string){

	var model []Model.User

	/* language=PostgreSQL */
	err := DB.Model(&model).Column("ignored").Where(`uniqueid = ?`,id).Select()
	if err!=nil{
		fmt.Println("SELECT FAILED!")
		fmt.Println(err)
	}else {
		fmt.Println("smthng selected")
		fmt.Println(model[0].Ignored)
	}

	var ignored []int

	err = json.Unmarshal([]byte(model[0].Ignored),&ignored)
	if err != nil{
		fmt.Println("unmarshal ignored failed",err)
	}



	var allUsers []Model.UserS
	err = DB.Model(&allUsers).Select()
	if err!=nil{
		fmt.Println(err)
	}

	var users []Model.UserS

	for _,i := range allUsers{
		for _,j := range ignored{
			if i.Id != j{
				users = append(users,i)
			}
		}
	}









}




