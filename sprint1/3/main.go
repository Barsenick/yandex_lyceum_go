package main

type User struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(args ...interface{}) User {
	if len(args) == 0 {
		return User{"", 18, true}
	}
	if len(args) == 1 {
		name := args[0].(string)
		return User{name, 18, true}
	}
	if len(args) == 2 {
		name := args[0].(string)
		age := args[1].(int)
		return User{name, age, true}
	}
	if len(args) == 3 {
		name := args[0].(string)
		age := args[1].(int)
		active := args[2].(bool)
		return User{name, age, active}
	}
	return User{"", 18, true}
}
