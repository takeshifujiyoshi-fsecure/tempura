package main
// #cgo CFLAGS: -fplugin=./pingtest.so
import "C"
import "fmt"

func main() {
    fmt.Println("Hello tempura")
}
