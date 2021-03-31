package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ClearTerminal ...
func ClearTerminal() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

var in = bufio.NewReader(os.Stdin)

// ScanLine ....
func ScanLine(prompt string, errMessage string) string {
	fmt.Print(prompt)
	line, _ := in.ReadString('\n')
	line = strings.Replace(line, "\n", "", -1)
	line = strings.Replace(line, "\r", "", -1)
	line = strings.Trim(line, " ")
	if line == "" {
		return ScanLine(errMessage, errMessage)
	}
	return line
}
