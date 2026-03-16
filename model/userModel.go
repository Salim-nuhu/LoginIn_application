package model
import(
	"database/sql"
)

type User struct{
	ID       int    `json:"id"  swaggerignore:"true"`
	Email    string `json:"email"`
	Password string `json:"password"`
} 

type ResponseError struct{
	Error string  `json:"err"`
}

type InitDB struct{
	DB *sql.DB
}