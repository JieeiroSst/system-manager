package models


type User struct{
	id string `json:"id" from:"_"`
	UserName string `json:"user_name" from:"user_name"`
	Password string `json:"password" from:"password"`
}

