package main

func main() {
	a := App{}
	a.InitDB()
	a.StartGRPC()
	a.InitHandler()
	a.StartHTTP()
}
