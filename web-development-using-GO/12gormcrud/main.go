package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	// ENTER the PostgreSQL database credentials
	dsn := "host=localhost user=postgres password=ajujacob dbname=new_test_db2 port=5432 sslmode=disable TimeZone=Asia/kolkata"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Automatically create the "users" table if it doesn't exist
	db.AutoMigrate(&User{})

	// Get action input from the user
	var action string
	fmt.Print("Enter action (add/edit/delete): ")
	fmt.Scanln(&action)

	switch action {
	case "add", "edit":
		// Get ID input from the user
		var id uint
		fmt.Print("Enter ID: ")
		fmt.Scanln(&id)

		// Check if a user with the given ID already exists
		var user User
		result := db.First(&user, id)
		if result.Error != nil {
			// If no user is found with the given ID, create a new user
			fmt.Println("User not found, creating a new user...")

			var name, email string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter email: ")
			fmt.Scanln(&email)

			user = User{ID: id, Name: name, Email: email}
			result = db.Create(&user)
			if result.Error != nil {
				panic(result.Error)
			}
		} else {
			// If a user is found with the given ID, update the user's name and email
			fmt.Println("User found, updating user...")

			var name, email string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter email: ")
			fmt.Scanln(&email)

			user.Name = name
			user.Email = email
			result = db.Save(&user)
			if result.Error != nil {
				panic(result.Error)
			}
		}
	case "delete":
		// Get ID input from the user
		var id uint
		fmt.Print("Enter ID: ")
		fmt.Scanln(&id)

		// Delete the user with the given ID
		result := db.Delete(&User{}, id)
		if result.Error != nil {
			panic(result.Error)
		}
	default:
		fmt.Println("Invalid action")
		return
	}

	// Retrieve all users from the "users" table
	var users []User
	db.Find(&users)
	fmt.Println(users)
}
