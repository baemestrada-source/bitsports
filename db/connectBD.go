package db

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/baemestrada-source/bitsports/ent"
)

/*posgresCN es el objeto de conexión a la BD */
var PosgresCN = ConectarBD()

/*ConectarBD es la función que me permite conectar la BD */
func ConectarBD() *ent.Client {
	client, err := ent.Open("postgres","host=localhost port=5432 user=postgres dbname=postgres password=secure_pass_here sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")

	return client
}