package model

import "database/sql"

type User struct {
	ID       int
	Login    string
	Password string
	Email    sql.NullString // Специальный тип, который возвращает nil, если строки не было, а в обратном случае - строку
	// А вот проверка на не-null значение user.Email.Valid
}
