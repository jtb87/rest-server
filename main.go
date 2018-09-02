package main

import ()

func main() {
	app := App{}
	port := "9090"
	app.initApp()
	app.Run(port)
	d := &dbStore{
		db: make(map[int]Todo),
	}
	InitStore(d)
}
