package main

import "github.com/paw1a/ecommerce-api/internal/app"

type Person struct {
	name     string
	age      int
	children []Person
}

func Clone() {
	return
}

type Cloneable interface {
	Clone()
}

func main() {
	app.Run("config/config.yml")
}
