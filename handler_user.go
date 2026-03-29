
package main

import(
	"errors"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("a username is required")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
    		return err
	}
	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println("User was set")
	return nil
}
func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("a username is required")
	}
	name := cmd.Args[0]
	user, err := s.db.CreateUser(
	context.Background(),
    	database.CreateUserParams{
        	ID: uuid.New(),
        	CreatedAt: time.Now(),
        	UpdatedAt: time.Now(),
        	Name: name,
    	},
	)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Println("user created")
	fmt.Println(user)
	return nil
}
func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
		fmt.Printf("* %v (current)\n", user.Name)
	} else {
		fmt.Printf("* %v\n", user.Name)
	}
	}
	return nil
}
