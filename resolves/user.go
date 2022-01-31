package resolves

import 
(
	"log"
	"errors"
	"encoding/json"
	"github.com/baemestrada-source/bitsports/models" 
	"github.com/baemestrada-source/bitsports/middlew" 
	"github.com/baemestrada-source/bitsports/db" 
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var TokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "token",
		Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},

		},
	},
)

func GetUserByKey() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}
	return &graphql.Field{
		Name:        "Insert Todo",
		Type:        userType,
		Description: "Get user by id",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			username, ok := p.Args["username"].(string)
			if ok {
				// Find user
				err, respuesta := db.FindUser(username)
				if err == nil {
					//log.Println(respuesta)
					return respuesta, nil
				}
			}
			return nil, nil
		},
	}
}

//CheckIn registra usuario
func CheckIn() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}
	return &graphql.Field{
		Name:        "CheckIn",
		Type:        userType,
		Description: "registrar el usuario",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user := models.User{}
			
			body, err := json.Marshal(p.Args)
			if err != nil {
				return user, err
			}
			err = json.Unmarshal(body, &user)
			if err != nil {
				return user, err
			}

			//log.Println("Paso 2: ",user)
			
			err, user = db.CreateUser(user)
			if err != nil {
				return user, err
			}


			return user,nil
		},
	}
}


//Login registra usuario
func Login() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}
	return &graphql.Field{
		Name:        "Login",
		Type:        TokenType,
		Description: "Ingresa el usuario",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			token := models.Token{}
			
			body, err := json.Marshal(p.Args)
			if err != nil {
				return token, err
			}
			err = json.Unmarshal(body, &token)
			if err != nil {
				return token, err
			}

			//log.Println("Paso 3: ",token)

			if len(token.Username) == 0 || len(token.Password) == 0 {
				return token,  errors.New("El usuario y/o clave son requeridos ")
			}

			existe := db.TryLogin(token.Username, token.Password)
			if existe == false {
				return token,  errors.New("Usuario y/o Contraseña inválidos ")
			}
			
			log.Println("paso aqui: ", existe )
			
			jwtKey, err := middlew.GeneroJWT(token)
			if err != nil {
				return token, err
			}
					
			resp := models.Token{
				Username: token.Username,
				Password: token.Password,
				Token: jwtKey,
			}
		
			return resp, nil
		},
	}
}