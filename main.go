package main

import "fmt"
import "os"
import "os/exec"

func exists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func getgo() {
	fmt.Println("Please install Go (https://golang.org/)")
	os.Exit(1)
}

func gofmt() {
	err := exec.Command("go", "fmt").Run()
	if err != nil {
		fmt.Println()
		return
	}
}

func goget(pkg string) {
	err := exec.Command("go", "get", pkg).Run()
	if err != nil {
		fmt.Println("There was an error trying to get ", pkg)
		fmt.Println("Please get it manually!")
		os.Exit(1)
	}
}

func lazygit() {
	if !exists("lazygit") {
		fmt.Println("Please add $GOPATH/bin to your $PATH!")
		os.Exit(1)
	}

	err := exec.Command("lazygit").Run()
	if err != nil {
		fmt.Println("lazygit encountered an error, is this a Git repo?")
		fmt.Println("If not, maybe you should report it... (https://github.com/jesseduffield/lazygit/issues)")
	}
}

func main() {
	if !exists("go") {
		getgo()
	}
	if !exists("lazygit") {
		goget("github.com/jesseduffield/lazygit")
	}

	gofmt()
	lazygit()
}
