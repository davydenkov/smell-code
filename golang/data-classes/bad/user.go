package main

type User struct {
	id    int
	name  string
	email string
	age   int
}

func NewUser(id int, name string, email string, age int) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
		age:   age,
	}
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) SetAge(age int) {
	u.age = age
}

// No behavior, just data and getters/setters - this is a data class smell!

func main() {
	user := NewUser(1, "John Doe", "john@example.com", 30)
	println("User ID:", user.GetId())
	println("User Name:", user.GetName())
	println("User Email:", user.GetEmail())
	println("User Age:", user.GetAge())
}
