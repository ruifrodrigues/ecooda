package main

func main() {
	go runGrpcServer()
	runHttpProxy()
}
