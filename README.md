# LABORA API: Personas

**Introducción:**

En esta guía aprenderás a crear una API REST CRUD (Create, Read, Update, Delete) para gestionar "Personas" usando el lenguaje de programación Go sin usar una base de datos, almacenando la información en memoria.

Se podrá registrar una persona junto con su país en abreviatura, pero al momento de devolverlas, deberá hacerlo con la información completa de su país:

Nombre completo del país, Timezone y tipo de moneda.

Para realizar esto consultaremos a [restcountries](https://restcountries.com/).

<aside>
📑 **Comportamiento**: La `Persona` no puede tener los datos sin valores.

</aside>

![Proyecto.jpg](LABORA%20API%20Personas%20b0fbb776db4049669a396b2bd17c58cc/Proyecto.jpg)

# Hitos

Podemos llevar a cabo nuestra aplicación en los siguientes hitos:

- **Creación del core**, nuestro modelo persona junto con su comportamiento.
- **Creación del servicios**, que sería información de la lógica de nuestra aplicación (operaciones CRUD y llamadas externas).
- **Creación de infraestructura**, como vamos a consumir externamente nuestra aplicación, en esste caso servidor junto con sus handlers.

Esto de forma muy sencilla es DDD:

![Untitled-Project.jpg](LABORA%20API%20Personas%20b0fbb776db4049669a396b2bd17c58cc/Untitled-Project.jpg)

**Pasos:**

**1. Crear el proyecto:**

- Abre tu terminal y navega a la carpeta donde deseas crear el proyecto.
- Ejecuta el siguiente comando:

```go
go mod init labora-persona
```

**2. Definir la estructura de la persona:**

- Crea un archivo `models/persona.go` y define la siguiente estructura:

```go
type Persona struct {
  ID int
  Nombre string
  Apellido string
  Edad int
	País string
}

// Comportamiento de la Persona..
func (p Persona) Validate() bool 
```

**3. Crear un slice para almacenar las personas:**

- En el archivo `main.go` crea un slice de tipo `Persona` para almacenar las personas:

```go
var PersonasDb []Persona
```

### 4. Podremos comenzar con `createPersona` de nuestro servicio

```go
func insertarPersonaEnLaDb(persona Persona) int {
	// Incrementar el ID de la persona
	persona.ID = len(PersonasDB) + 1
	PersonasDB = append(PersonasDB, persona)
	return persona.ID
}

// Resto de operaciones crud
func editarPersonaEnLaDb(persona Persona) err
//...
```

### 5. Llamada a servicio externo

Aquí irá la magía, deberiamos crear una struct para contemplar este nueva “Persona Extendida Con Su País” en donde deberá contar con la información ya mencionada.

1. **Crear handlers en un archivo `handlers.go`**

<aside>
💡 No vamos a necesitar una librería externa, en go1.22 [tenemos sorpresas!](https://tip.golang.org/doc/go1.22#enhanced_routing_patterns)

</aside>

```go
func createPersona(w http.ResponseWriter, r *http.Request) {
	// Aquí implementa la lógica para crear una persona específica
}

func getPersona(w http.ResponseWriter, r *http.Request) {
	// Aquí implementa la lógica para obtener una persona específica
}

func updatePersona(w http.ResponseWriter, r *http.Request) {
	// Aquí implementa la lógica para actualizar una persona
}

func deletePersona(w http.ResponseWriter, r *http.Request) {
	// Aquí implementa la lógica para eliminar una persona
}
```