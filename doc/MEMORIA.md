# Memoria TFG

Este documento representa la memoria como un archivo donde plasmar las ideas iniciales y como base a la redacción final de la memoria a entregar.

## Introducción

Se realizará un CLI* en Golang debido a su alto rendimiento a la vez que por su familiaridad con otros lenguajes de programación. La versión utilizada es la más actualizada con fecha de este proyecto, la versión `1.22`.

Existen muchos CLIs creados con Go, algunos [ejemplos](https://go.dev/solutions/clis) son herramientas de las siguientes empresas / programas: Github, Comcast, Hugo, Kubernetes, MongoDB, Netflix, Stripe, Uber.

El CLI de este proyecto creará un proyecto nuevo hecho en Go. Este nuevo proyecto creado a raíz del CLI será una API REST que creará los endpoints CRUD a partir de las tablas de una base de datos.

## Proyectos similares existentes

No se ha encontrado algo exactamente igual pero sí servicios que ofrecen programas para manejar APIs a partir de una interfaz gráfica y conectar con una base de datos:

- https://www.cdata.com/apiserver/
- https://strapi.io
- https://www.gravitee.io
- https://guide.dreamfactory.com/docs/

## Alcance

Sí está en el alcance de este proyecto:

- Creación de una API con un CRUD a partir de una tabla SQL.
- Conectar con una base de datos SQL para leer sus tablas.

No se tendrá en cuenta, ya que no está al alcance de la magnitud de este proyecto:

- Bases de datos no relacionales.
- Comprobar las relaciones de las tablas.
- Diseños que no cumplan las 4 formas normales de un diseño relacional.
- Actualizaciones / Migraciones de la base de datos utilizada.
- Todos los tipos de datos de las bases de datos.

## Requerimientos

Como primer paso se definen los requerimientos de este programa a raíz de los casos de uso.

Se definirá como CRUD a la cración de 5 endpoints uno para leer todos los recursos, otro para leer un recurso, otro para crear un recurso, otro para actualizar un recurso y otro para eliminar un recurso.

### Casos de uso

1. Un usuario crea una API con un CRUD de una tabla escogida de una base de datos Postgresql ya existente.

2. Un usuario crea una API con un CRUD por cada una de las tablas de una base de datos Postgresql ya existente.

3. Un usuario puede un CRUD a una API ya existente creada con este CLI.

### Requerimientos funcionales

1. El usuario debe escoger qué tipo base de datos se va a utilizar a través de un selector. Como primera iteración la única selección será Postgresql.

2. El usuario deberá informar de los datos necesarios para conectar a la base de datos. Para Postgresql:
  - `host`
  - `port`
  - `user`
  - `password` (ofuscado)

3. El usuario puede interrumpir el programa con ctrl + C en cualquier momento y no se deberá haber aplicado ningún cambio

### Requerimientos técnicos

1. Se utiliza el lenguaje Golang

2. El programa se maneja por línea de comandos (CLI)

## Herramientas utilizadas

- [Cobra](https://cobra.dev)
- [Fiber](https://gofiber.io)
- [Github Copilot](https://github.com/features/copilot)

## Proyecto a crear

- Se crea una API Rest con dependencias como el framework Go Fiber
