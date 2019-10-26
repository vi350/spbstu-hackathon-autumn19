package DB

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	_ "github.com/go-pg/pg/orm"
	"github.com/vi350/spbstu-hackathon-autumn19/Model"
	"github.com/vi350/spbstu-hackathon-autumn19/Passwords"
	"log"
)

var DB *pg.DB

// подключение к бд
func ConnectDB () {
	DB = pg.Connect(&pg.Options{
		Addr:     "0.tcp.ngrok.io:16384",
		User:     "postgres",
		Password: Passwords.DBPass,
		Database: "hackathon",
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
        uniqueid text,
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

	var m Model.User = Model.User{"qwergh","sname",5,jsoniseStrs([]string{"go","gin"}),jsoniseInts([]int{5}),jsoniseInts([]int{2}),true}

	err := db.Insert(&m)
	if err !=nil{
		fmt.Println("bleat")
		fmt.Println(err)
	}

	/* language=PostgreSQL */
	//s := `INSERT INTO users (uniqueid,name,rate,skills,favourites,ignored,busy) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	//_,err := db.Exec(s,"idunc","somename",23,[]string{"go", "gin"},[]int{5},[]int{2},true)
	//if err != nil{
	//  fmt.Println(err)
	//}

}


func jsoniseStrs(arr []string)string  {
	slc,_ := json.Marshal(arr)
	return string(slc)

}

func jsoniseInts(arr []int)string  {
	slc,_ := json.Marshal(arr)
	return string(slc)
}