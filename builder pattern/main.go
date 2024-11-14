package main

import (
	"fmt"
)

type User struct {
	Name        string
	Email       string
	Preferences map[string]string
}

func NewUser() *User {
	return &User{}
}

type UserBuilder struct{
	user *User
}

func NewUserBuilder() *UserBuilder{
	return &UserBuilder{&User{}}
}

func (b *UserBuilder) SetName(name string) *UserBuilder{
	b.user.Name =name
	return b
}

func (b *UserBuilder) SetEmail(email  string) *UserBuilder{
	b.user.Email = email
	return b
}

func (b *UserBuilder) SetPreferences(preferences map[string]string) *UserBuilder{
	b.user.Preferences = preferences
	return b
}

func (b *UserBuilder) Build()*User{
	return b.user
}

func main() {
	builder := NewUserBuilder()

	user1 := builder.SetName("bilal").SetEmail("bilal@gmail.com").SetPreferences(map[string]string{ "age": "21", "aa":"pp"}).Build()

	preferences := map[string]string{"theme": "dark"}
	builder1 := NewUserBuilder()
	user2 := builder1.SetName("adam").SetEmail("adam@gmail.com").SetPreferences(preferences).Build()

	
	fmt.Println(user1)
	fmt.Println(user2)

}
