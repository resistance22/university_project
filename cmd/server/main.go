package main

import "fmt"

func Run() error {
	fmt.Println("Starting App")
	return nil
}

func main() {
	fmt.Println("Go Rest")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
