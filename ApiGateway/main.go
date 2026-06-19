package main

import "ApiGateway/app"

func main() {
	application := app.NewApplication()

	err:=application.Run()
	if err!=nil{
		
	}
}
