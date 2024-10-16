package main

import (
	"go_learn/cmd/app/go_learn/pkg"
	"go_learn/cmd/app/go_learn/server"
	"log"
)

func main() {
	// exit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	// signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	for {
	// 		log.Printf("Try connect to database")
	// 		select {
	// 		case <-exit:
	// 			fmt.Println("Database disconnected")
	// 			return
	// 		case <-time.After(1 * time.Minute):
	// 			fmt.Println("Database is connected")
	// 		}

	// 	}

	// }()

	pkg.Database_conn()
	log.Printf("Server is started")
	server.Server()
}
