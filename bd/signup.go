package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/maalpu/zeiboxuser/models"
	"github.com/maalpu/zeiboxuser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sencencia := "INSERT INTO users (User_Email, UserUUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sencencia)

	_, err = Db.Exec(sencencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución Exitosa")
	return nil
}
