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
}

type Users []User