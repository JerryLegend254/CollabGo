package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{config: cfg}

	mux := app.mount()

	log.Printf("server listening at %s", app.config.addr)
	log.Fatal(app.run(mux))

}
