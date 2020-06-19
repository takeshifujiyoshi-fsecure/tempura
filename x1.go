package main
// #cgo CFLAGS: -fplugin=./attack.so 
import "C"
import "fmt"
func main() {
    fmt.Println("Hello tempura");
}
