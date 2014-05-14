package main

func CreateUser(user User) (result string) {
	user.Save()
	return
}
