# Prueba tecnica Bitsports
Crear un API en GraphQL donde puedas autenticar por medio de JWT, una
vez autenticado poder acceder a los endpoints de products y categories y hacer
el respectivo CRUD de cada una de las entidades, debes guardar esto en una
base de datos PostgreSQL y realizar las pruebas unitarias.
Modelo
USERS
CATEGORIES
PRODUCTS

## Comenzando 🚀

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

Mira **Deployment** para conocer como desplegar el proyecto.

### Pre-requisitos 📋

Instalacion de Docker

### Instalación 🔧

Primeramenente debo de clonar el proyecto posterior se incluyo un archivo de la imagen en docker utilizada de la base de datos postgres en el archivo docker-compose.yml del proyecto esta tiene ya la configuracion que se necesita para el proyecto
```
docker-compose up -d
```

Ya con la base de datos montada, puede ejecutar el proyecto en el directorio del mismo, el puerto escucha es el 8080

```
go run main.go
```

## Ejecutando las pruebas ⚙️

Para ejecutar el proyecto primero se debe registrar un usuario, puede utilizar postman metodo POST, un ejemplo

```
http://localhost:8080/user?query=mutation+_{checkin(username:"bestrada",name:"byron moreira",email:"bestrada111@ss.com", password: "12345"){username,name,email,password}}
```

esto registrara el usuario al proyecto deberia obtener algo como esto, la clave se guarda encriptada

```
{
    "data": {
        "checkin": {
            "email": "bestrada111@ss.com",
            "name": "byron moreira",
            "password": "$2a$08$iEUelyv0HBAIOFoGn/8VX.oCSXd9iBNcCDYDWSYx3zG5NzUiuDXvy",
            "username": "bestrada"
        }
    }
}
```


Postermente ya con este usuario registrado puede obtener el token generado en JWT con la siguiente ruta

```
http://localhost:8080/user?query=mutation+_{login(username:"bestrada",password:"12345"){token}}
```

con esto obtendra el token, un ejemplo de respuesta

```
{
    "data": {
        "login": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDc2NjMsInBhc3N3b3JkIjoiMTIzNDUiLCJ1c2VybmFtZSI6InBydWViYTEifQ.u8_inUx13xlE1rsGW-P2uO_d2XfXGralbZKFdASWAa0"
        }
    }
}
```

Ahora estamos listos puede utilizar siempre la utilidad de /graphql o las rutas directas que son /product, /categorie

## CRUD Categorias

Puede probar la creacion de una categoria de 2 formas, 

1. Ejemplo en /graphql

```
http://localhost:8080/graphql?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM1ODg1NDksInBhc3N3b3JkIjoiMTIzNDUiLCJ1c2VybmFtZSI6InBydWViYTEifQ.xGSh6wS8_Hq7tZvWvf9zLMwQSHxQ91sv4cLutnv5ezo
```

```
mutation {
  createCategorie(
         name: "otros", 
      ) 
  {
    createCount
    result{
      id
      name
     }
  }
}
```

2. Ruta directa: 

```
http://localhost:8080/categorie?query=mutation+_{createCategorie(name:"prueba"){createCount result{id,name}}}
```
aqui debe poner el token en una variable header de la siguiente manera

```
Key:Authorization 
Value:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM2MDc2NjMsInBhc3N3b3JkIjoiMTIzNDUiLCJ1c2VybmFtZSI6InBydWViYTEifQ.u8_inUx13xlE1rsGW-P2uO_d2XfXGralbZKFdASWAa0
```
esto es con todos los endpoints asi funcionan

daremos el resto utilizando /graphql pero se puede siempre de las 2 formas,

### Crear categoria ejemplo

```
mutation {
  createCategorie(
         name: "otros", 
      ) 
  {
    createCount
    result{
      id
      name
     }
  }
}
```

resultado como este donde me dice si fue creada correctamente,

```
{
    "data": {
        "createCategorie": {
            "createCount": 1,
            "result": {
                "id": 5,
                "name": "otros"
            }
        }
    }
}
```


### Actualizar Categoria ejemplo

```
mutation {
  updateCategorie(
         id: 2
         name: "Farmaceuticos"
    ) 
   {
    modifiedCount
    result{
      name
    }
  }
}
```

resultado como este donde me dice si fue modificada correctamente,

```
{
    "data": {
        "updateCategorie": {
            "modifiedCount": 1,
            "result": {
                "name": "Farmaceuticos"
            }
        }
    }
}
```

### Eliminar Categoria ejemplo

```
mutation {
  deleteCategorie(id: 1)
  {
    deletedCount
  }
}
```

resultado como este:

```
{
    "data": {
        "deleteCategorie": {
            "deletedCount": 1
        }
    }
}
```

### Consultar por ID la categoria
```
query {
  Categorie(id:1) {  
   name
  }
}
```

resultado 

```
{
    "data": {
        "Categorie": {
            "name": "Prueba"
        }
    }
}
```

### Consultar todas las categorias

```
query {
    Categories{
        id,
        name
    }
}
```

resultado como este: 

```
{
    "data": {
        "Categories": [
            {
                "id": 1,
                "name": "alimentos"
            },
            {
                "id": 5,
                "name": "otros"
            },
            {
                "id": 4,
                "name": "acidos"
            },
            {
                "id": 3,
                "name": "maderas"
            },
            {
                "id": 2,
                "name": "Farmaceuticos"
            }
        ]
    }
}
```


## CRUD Productos

a travez de la forma en /graphql pero se puede siempre desde la ruta completa ejemplo http://localhost:8080/product?query={Products{name,info,price}}

### Crear producto ejemplo

```
mutation {
  createProduct(
         name: "jona2", 
    	   info: "son deliciosas",
         price: 20.25,
         categorie_id: 2
         ) 
  {
    createCount
    result{
      id
      name
      info
      price
      categorie_id
    }
  }
}
```

resultado como este donde me dice si fue creada correctamente,

{
    "data": {
        "createProduct": {
            "createCount": 1,
            "result": {
                "categorie_id": 2,
                "id": 14,
                "info": "son deliciosas",
                "name": "jona2",
                "price": 20.25
            }
        }
    }
}
```


### Actualizar producto ejemplo

```
mutation {
  updateProduct(
         id: 2
         name: "Mixtas", 
    	   info: "Muy ricas",
         price: 10,
         categorie_id: 1
         ) 
  {
    modifiedCount
    result{
      name
      info
      price
      categorie_id
    }
  }
}
```

resultado como este donde me dice si fue modificada correctamente,

```
{
    "data": {
        "updateProduct": {
            "modifiedCount": 1,
            "result": {
                "categorie_id": 1,
                "info": "Muy ricas",
                "name": "Mixtas",
                "price": 10
            }
        }
    }
}
```

### Eliminar producto ejemplo

```
mutation {
  deleteProduct(id: 9)
  {
    deletedCount
  }
}
```

resultado como este si en caso no existiera el producto, contador en 0 y el mensaje correspondiente

```
{
    "data": {
        "deleteProduct": {
            "deletedCount": 0
        }
    },
    "errors": [
        {
            "message": "ent: products not found",
            "locations": [
                {
                    "line": 2,
                    "column": 3
                }
            ],
            "path": [
                "deleteProduct"
            ]
        }
    ]
}
```

### Consultar por ID la producto
```
query {
  Product(id:2) {  
   name
   info
   price
   categorie_id
  }
}
```

resultado si el token estuviera incorrecto

```
Error en el Token ! token Inválido
```

### Consultar todos los productos

```
query {
    Products{
        name,
        info,
        price
    }
}
```

resultado como este: 

```
{
    "data": {
        "Products": [
            {
                "info": "son deliciosas",
                "name": "Tortas",
                "price": 20.25
            },
            {
                "info": "son deliciosas",
                "name": "Tortas",
                "price": 20.25
            },
            {
                "info": "es una bebida muy especial pero dañina",
                "name": "pepsi cola",
                "price": 5.12
            },
            {
                "info": "son deliciosas",
                "name": "Tortas",
                "price": 20.25
            },
            {
                "info": "son deliciosas",
                "name": "Tortas",
                "price": 20.25
            },
            {
                "info": "Muy ricas",
                "name": "Mixtas",
                "price": 10
            }
        ]
    }
}
```


## Construido con 🛠️

Se utilizo lenguaje GO y algunas librerias como 

http://entgo.io/  -- sirve como ORM de la base de datos 

https://github.com/graphql-go/graphql -- sirve generar API con grapql

https://github.com/lib/pq  -- paquete para conectar a posgresql

https://golang.org/x/crypto/bcrypt -- paquete para encriptar la clave del usuario esta se guarda en la BD encriptada

https://github.com/dgrijalva/jwt-go --paquete con el cual se genera el token al hacer login

## Base de datos  🖇️

Se dejo el archivo docker-compose.yml ya este puede crear una imagen de la base de datos en postgres con la conexion que se utiliza en el proyecto

## Autores ✒️

Byron Arturo Estrada Moreira

## Licencia 📄
OpenSource
