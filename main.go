package main

import "os"

func main() {

	application := App{}
	application.Init(os.Getenv("dbusername"), os.Getenv("dbpassword"), os.Getenv("dbname"))

	application.Run(":5678")
}
