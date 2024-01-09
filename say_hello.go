package gosayhello

import "fmt"

func SayHello(name string, address string) string {
	return fmt.Sprintf("Hello, nama saya %s dan saya berasal dari %s.", name, address)
}
