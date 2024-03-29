package DB

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/vi350/spbstu-hackathon-autumn19/Passwords"

	//"github.com/go-pg/pg/orm"
	_ "github.com/go-pg/pg/orm"
	"github.com/vi350/spbstu-hackathon-autumn19/Model"
	"log"
)

var DB *pg.DB

// подключение к бд
func ConnectDB() {
	DB = pg.Connect(&pg.Options{
		Addr:     "89.208.196.56:5432",
		User:     "dima",
		Password: Passwords.DBPass, //you need manually add package Passwords ("Passwords/Passwords.go") and create string DBPass with your pass
		Database: "keyzu",
	})
	Status()
}

// проверка подключения к бд
func Status() bool {
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
func CreateTables() {
	db := DB
	qs := []string{
		/* language=PostgreSQL */
		`CREATE TABLE IF NOT EXISTS users(
        id SERIAL PRIMARY KEY,
        uniqueid text ,
        token text,
    	name text,
    	username text,
    	photourl text,
    	rating numeric,
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
}

func SelectUsers(c *gin.Context) {
	type DataInUs struct {
		Id        string `json:"id"  binding:"required"`
		FirstName string `json:"first_name" binding:"required"`
	}
	type skillsFor struct {
		Skills   []string `json:"skills" binding:"required"`
		Uniqueid string   `json:"uniqueid" binding:"required"`
	}
	var data skillsFor

	status := 200
	message := "ok"

	err := c.ShouldBindJSON(&data)
	if err != nil {
		status = 400
		message = "json not binded"
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(data.Skills)
	var users []Model.UserS
	users = SelectBySkills(data.Skills, data.Uniqueid)

	type sendUsers struct {
		Id         int
		Uniqueid   string
		Name       string
		Rating     int8
		Skills     []string
		Favourites []int
		Ignored    []int
		Busy       bool
	}

	var outUsers []sendUsers

	for _, i := range users {
		var outUser sendUsers
		outUser.Id = i.Id
		outUser.Busy = i.Busy
		outUser.Name = i.Name
		outUser.Uniqueid = i.Uniqueid
		var ignored, favourites []int
		var skills []string
		_ = json.Unmarshal([]byte(i.Favourites), &favourites)
		_ = json.Unmarshal([]byte(i.Ignored), &ignored)
		_ = json.Unmarshal([]byte(i.Skills), &skills)
		outUser.Ignored = ignored
		outUser.Favourites = favourites
		outUser.Skills = skills
		outUsers = append(outUsers, outUser)
	}

	c.JSON(status, gin.H{
		"message": message,
		"users":   outUsers,
	})
}

//noinspection ALL
func SelectBySkills(skills []string, id string) []Model.UserS {

	var model []Model.User

	err := DB.Model(&model).Column("ignored").Where("uniqueid = ?", id).Select()
	if err != nil {
		fmt.Println("SELECT FAILED!")
		fmt.Println(err)
	} else {
		fmt.Println("smthng selected")
		fmt.Println(model[0].Ignored)
	}

	var ignored []int

	err = json.Unmarshal([]byte(model[0].Ignored), &ignored)
	if err != nil {
		fmt.Println("unmarshal ignored failed", err)
	}

	fmt.Println("ignored:", ignored)

	var allUsers []Model.UserS
	err = DB.Model(&allUsers).Select()
	if err != nil {
		fmt.Println(err)
	}

	var users []Model.UserS

	for _, i := range allUsers {
		adding := true
		for _, j := range ignored {
			if i.Id == j {
				adding = false
			}
		}
		if adding {
			users = append(users, i)
		}
	}

	// нужные юзеры фильтрованные по нужному нам стеку
	var needingUsers []Model.UserS
	// кол во совпадающих технологий
	var coincidences []int

	for _, usr := range users {
		var itskills []string
		_ = json.Unmarshal([]byte(usr.Skills), &itskills)
		adding := false
		count := 0
		for _, skill := range itskills {
			for _, sk := range skills {
				if sk == skill {
					adding = true
					count++
				}
			}
		}
		if adding {
			needingUsers = append(needingUsers, usr)
			coincidences = append(coincidences, count)
		}
	}

	fmt.Println("++++++++++")
	fmt.Println(needingUsers)
	fmt.Println(coincidences)

	// sort users and coincidenses
	for i := range coincidences {
		biggest := coincidences[i]
		index := i

		for j := i; j < len(coincidences); j++ {
			if biggest < coincidences[j] {
				index = j
				biggest = coincidences[j]
			}
		}
		var m Model.UserS
		var num int
		m = needingUsers[i]
		num = coincidences[i]
		needingUsers[i] = needingUsers[index]
		coincidences[i] = coincidences[index]
		needingUsers[index] = m
		coincidences[index] = num

	}

	fmt.Println("++++++++++")
	fmt.Println(needingUsers)
	fmt.Println(coincidences)

	return needingUsers
}
