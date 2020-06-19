package main
// #cgo LDFLAGS: ./pingtest.so
// void pingtest(void);
import "C"
func main() {
    C.pingtest();
}
