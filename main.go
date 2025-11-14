package main

func main() {
	err := Server("localhost", "8080")
	if err != nil {
		return
	}
}
