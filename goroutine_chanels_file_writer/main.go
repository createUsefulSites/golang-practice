package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type User struct {
	id string
	name string
	logs []TLogs
	role []string
}

func (u *User) returnUserInfo() string {
	userInfo := fmt.Sprintf("Пользователь: %s\n----------------\nРоли: ", u.name)
	for _, role := range u.role{
		userInfo += ("[" + role + "] ")
	}
	userInfo += "\n----------------\nЛоги пользователя:\n"
	for _, log := range u.logs{
		userInfo += fmt.Sprintf("%s. at [%s]\n", log.log, log.time)
	}

	return userInfo + "----------------\n"
}

type TLogs struct {
	log string
	time time.Time
}

func (log *TLogs) addTimeNow() {
	log.time = time.Now()
}

const (
	AdminRole = "admin"
	ModeratorRole = "moderator"
	UserRole = "user"

	LOGGED = "user was logged"
	DELETE = "user delete file"
	ADD = "user add file"
	UNLOGGED = "user was unlogged"
)

func main() {
	wg := &sync.WaitGroup{}

	timeStart := time.Now()
	users := generateUsers(10000)

	for i, user := range users {
		wg.Add(1)
		go func () {
			file, err := os.OpenFile(fmt.Sprintf("logs/user%v.txt", i+1), os.O_RDWR | os.O_CREATE, 0644)
			if err != nil {
				panic(fmt.Sprintf("cannot create file user%v.txt", i+1))
			}
	
			file.WriteString(user.returnUserInfo())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("time working: ", time.Since(timeStart))
}

func generateUsers(count int) []User {
	var result []User
	for i := 0; i < count; i++ {
		result = append(result, User{ id: fmt.Sprint(i+1), name: fmt.Sprintf("User%v", i+1), role: getRandomRoles(), logs: getRandomLogs() })
	}

	return result
}

func getRandomRoles() []string {
	var result []string
	count := 1 + rand.Intn(3)
	for i := 0; i < count; i++ {
		switch rand.Intn(3) {
		case 0:
			result = append(result, AdminRole)
		case 1:
			result = append(result, ModeratorRole)
		default:
			result = append(result, UserRole)
		}
	}

	return result
}

func getRandomLogs() []TLogs {
	var result []TLogs
	count := rand.Intn(4)

	for i := 0; i < count; i++ {
		var log TLogs
		log.addTimeNow()
		switch rand.Intn(4) {
		case 0:
			log.log = ADD
		case 1:
			log.log = DELETE
		case 2:
			log.log = UNLOGGED
		default:
			log.log = LOGGED
		}

		result = append(result, log)
	}

	return result
}
