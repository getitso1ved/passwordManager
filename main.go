package main

import (
	"basic_golang/passwordManager/account"
	"basic_golang/passwordManager/files"
	"fmt"
)

func main() {
	files.WriteFile("Привет! Я файл!!", "file.txt")
	files.ReadFile()
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	myAccount.OutputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
