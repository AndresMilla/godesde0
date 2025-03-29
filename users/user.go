package users

import (
	"fmt"
	"time"

	"github.com/AndresMilla/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Andrés", time.Now(), true)
	fmt.Println(u)
}
