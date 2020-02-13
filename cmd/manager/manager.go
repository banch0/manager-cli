package main

import (
	"database/sql"
	"fmt"
	"github.com/banch0/managers-cli/cmd/internal/dbinit"
	"github.com/banch0/managers-core/pkg/core"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

// TODO: для тех кто
// type ManagerCLI struct {
//db *sql.DB
//}

// find interface ctrl + alt + b
func main() {
	// os.Stdin, os.Strout, os.Stderr, File
	file, err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Print("start application")
	log.Print("open db")
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("can't open db: %v", err)
	}
	defer func() {
		log.Print("close db")
		if err := db.Close(); err != nil {
			log.Fatalf("can't close db: %v", err)
		}
	}()
	err = dbinit.Init(db)
	if err != nil {
		log.Fatalf("can't init db: %v", err)
	}

	fmt.Println("Добро пожаловать в наше приложение")
	log.Print("start operations loop")
	operationsLoop(db, unauthorizedOperations, unauthorizedOperationsLoop)
	log.Print("finish operations loop")
	log.Print("finish application")
}

func operationsLoop(db *sql.DB, commands string, loop func(db *sql.DB, cmd string) bool) {
	for {
		fmt.Println(commands)
		var cmd string
		_, err := fmt.Scan(&cmd)
		if err != nil {
			log.Fatalf("Can't read input: %v", err) // %v - natural ...
		}
		if exit := loop(db, strings.TrimSpace(cmd)); exit {
			return
		}
	}
}

func unauthorizedOperationsLoop(db *sql.DB, cmd string) (exit bool) {
	switch cmd {
	case "1":
		ok, err := handleLogin(db)
		if err != nil {
			log.Printf("can't handle login: %v", err)
		}
		if !ok {
			fmt.Println("Неправильно введён логин или пароль. Попробуйте ещё раз.")
			return false
		}
		operationsLoop(db, authorizedOperations, authorizedOperationsLoop)
	case "q":
		return true
	default:
		fmt.Printf("Вы выбрали неверную команду: %s\n", cmd)
	}

	return false
}

func authorizedOperationsLoop(db *sql.DB, cmd string) (exit bool) {
	switch cmd {
	case "1":
		fmt.Println("Полный список продуктов: ")
		prod, err := core.ShowAllProducts(db)
		if err != nil {
			return
		}
		printProduct(prod)
	case "2":
		// TODO: add sale
		fmt.Println("Тут будет добавление продаж")
	case "q":
		return true
	default:
		fmt.Printf("Вы выбрали неверную команду: %s\n", cmd)
	}
	return false
}

func printProduct(prod []core.Products) {
	for _, p := range prod {
		fmt.Printf("Name: %s  Price: %v  Quantity: %v  \n",
			p.Name,
			p.Price,
			p.Qty,
			)
	}
}

func handleLogin(db *sql.DB) (ok bool, err error) {
	fmt.Println("Введите ваш логин и пароль")
	var login string
	fmt.Print("Логин: ")
	_, err = fmt.Scan(&login)
	if err != nil {
		return false, err
	}
	var password string
	fmt.Print("Пароль: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return false, err
	}

	ok, err = core.Login(login, password, db)
	if err != nil {
		return false, err
	}

	return ok, err
}

func handleSale(db *sql.DB) (err error) {
	fmt.Println("Введите ваш логин и пароль")
	var id int64
	fmt.Print("Id of products")
	_, err = fmt.Scan(&id)
	if err != nil {
		return  err
	}
	var qty int64
	fmt.Print("quantity:")
	_, err = fmt.Scan(&qty)
	if err != nil {
		return  err
	}

	err = core.Sale(id, qty, db)
	if err != nil {
		return  err
	}

	return nil
}