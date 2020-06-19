package main
import "fmt"
import "os/exec"

func main() {
    fmt.Println("Hello tempura")
    out, err := exec.Command("ping", "-c3", "lumpini.sobajorg.org").Output()
    if err != nil {
	fmt.Printf("%s", out);
    }
}
