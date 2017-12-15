/*
=====================Golang======================
	Author 	: Audy Vaksi Pranata
	Link	: https://github.com/vaksi
	Email	: vaksipranata@gmail.com
==================================================
*/
package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     []Role `json:role`
}

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Users []User
