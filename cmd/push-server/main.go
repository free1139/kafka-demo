package main

func main() {
	addr := ":5001"
	svc, err := NewService()
	if err != nil {
		panic(err)
	}
	svc.(*pushSvc).listen(addr)
}
