package main

import "AuthenticationService/app"

func main() {
	application := app.NewApplication()

	if err := application.Run(); err != nil {

	}
}
