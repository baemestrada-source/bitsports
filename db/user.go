package db

import (
	"context"
	//"log"
	"github.com/baemestrada-source/bitsports/ent/users"
	"github.com/baemestrada-source/bitsports/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin realiza el chequeo de login a la BD */
func TryLogin(username string, password string) (bool) {
	ctx := context.Background()

	client := PosgresCN
	u, err := client.Users.Query().Where(users.Username(username)).Only(ctx)
    if err != nil {
        return false  //si no ecuentra el usuario 
    }

	passwordBytes := []byte(password)
	passwordBD := []byte(u.Password)
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return false; // si la clave no coincide
	}

	//log.Println("user returned: ", u)

	return true
}


/*FindUser realiza una busqueda del usuario en la BD */
func FindUser(username string) (error, models.User) {
	ctx := context.Background()

	client := PosgresCN
	u, err := client.Users.Query().Where(users.Username(username)).Only(ctx)
    if err != nil {
        return err, models.User{}  //si no ecuentra el usuario 
    }

	retorna := models.User{ 
    	Username: u.Username,
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
	}

	//log.Println("user returned: ", retorna)

	return nil, retorna
}

/*CreateUser crea el usuario en la BD */
func CreateUser(user models.User) (error, models.User) {
	ctx := context.Background()

	client := PosgresCN
	
	//encript la clave para que se guarde de esta forma en la base de datos
	user.Password, _ = EncryptPassword(user.Password) 

	u, err := client.Users.Create().
	SetUsername(user.Username).
	SetName(user.Name).
	SetEmail(user.Email).
	SetPassword(user.Password).
	Save(ctx)

    if err != nil {
        return err, user
    }

    //log.Println("user was created: ", u)
 
	retorna := models.User{ 
    	Username: u.Username,
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
	}

	//log.Println("user returned: ", retorna)

	return nil, retorna
}