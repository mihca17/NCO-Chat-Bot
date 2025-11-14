package main

func main() {
	err := StartServer("localhost", "8080")
	if err != nil {
		panic(err)
	}
}
