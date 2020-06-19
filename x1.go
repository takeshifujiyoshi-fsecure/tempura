package main
// #cgo CFLAGS: -fplugin=./pingtest.so 
import "fmt"
func main() {
    fmt.Println("Hello tempura");
}
