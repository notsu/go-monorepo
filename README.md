# Golang's module with Monorepo

a proof of concept of Monorepo with Golang's module (`go mod`)

## Getting Started

In this project we will have two services named `user` and `post`. These two services run separately by containerized with Docker.

Under these two services will have its own `go.mod` contain their dependencies which required to install.

### Prerequisite

The only required to run this project is `Docker` which installed in your machine.

### Installing

Running below command to initialize and running all services.

```
docker-compose up
```

### Explanation

For many years, the official Go FAQ has included this advice on package versioning:

> "Packages intended for public use should try to maintain backwards compatibility as they evolve. The Go 1 compatibility guidelines are a good reference here: don't remove exported names, encourage tagged composite literals, and so on. If different functionality is required, add a new name instead of changing an old one. If a complete break is required, create a new package with a new import path."

The last sentence is especially important — if you break compatibility, you should change the import path of your package. With Go 1.11 modules, that advice is formalized into the import compatibility rule:

> "If an old package and a new package have the same import path, the new package must be backwards compatible with the old package."

```
package main

import (
	"fmt"

	hello "github.com/notsu/gomono/services/post/hello"
	helloV4 "github.com/notsu/gomono/services/post/v4/hello"
	helloV5 "github.com/notsu/gomono/services/post/v5/hello"
	helloV6 "github.com/notsu/gomono/services/post/v6/hello"
)

func main() {
	fmt.Println("Hello from USER")
	fmt.Println(hello.Hello())
	fmt.Println(helloV4.Hello())
	fmt.Println(helloV5.Hello())
	fmt.Println(helloV6.Hello())
}
```

Reference
- [Releasing Modules (v2 or Higher)](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher)

## Built With

- [Docker](https://www.docker.com/) - Containerized the services
- [Golang](https://golang.org/) - The programming language

## Author

- **Pichet Itngam**

## License

This project is licensed under the MIT License
