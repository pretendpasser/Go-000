package main

import (
	"fmt"
	"week02/service"
)

func main() {
	s := service.NewService()
	_, err := s.GetUsernameByUserById(10);
	fmt.Printf("%+v\n", err)
}
