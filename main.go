package main

import "hacktiv8-chapter-two-session-two-challenge/routers"

func main() {
	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}
