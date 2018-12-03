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

package main

```
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