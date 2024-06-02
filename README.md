# CLI hecho en GO

Este CLI consiste en un proyecto de final de carrera.

La versión de Go utilizada es la 1.22.

```bash
go version
> 1.22.0
```

Consiste en un pequeño CLI hecho con el framework de [Cobra](https://github.com/spf13/cobra/tree/main)

## Uso

Este CLI crea una aplicación API REST hecha con [Go Fiber](https://gofiber.io) que tiene endpoints CRUD por tabla de una base de datos Postgresql.

### Instalación

Instalar el CLI con el comando:

```bash
go install github.com/eddyvy/tfg-go-cli
```

Comprobar que se ha instalado correctamente:

```bash
tfg-go-cli version
# Version: 0.0.1
```

### Crear un proyecto nuevo

Utilizar el comando `new` para crear un proyecto nuevo:

```bash
tfg-go-cli new example_app
```

Esto generará una carpeta `example_app` con el proyecto.

### Agregar endpoint a proyecto

Utilizar el comando `add` para agregar los endpoints de una tabla de la base de datos con la que se creó el proyecto:

```bash
tfg-go-cli add example_app
```

Nota: Ejecutar el comando en la carpeta padre del proyecto e indicar el nombre del proyecto.
