package cliui

import (
	"bufio"
	"fmt"
	"os"
)

type UI struct {
	nArr []string
	fArr []func([]string)
	Name string
}

// Add adds a function to the list of possible commands
func (u *UI) Add(name string, handler func([]string)) {
	//these should never have a different amount of args
	if u.nArr == nil || u.fArr == nil {
		u.nArr = make([]string, 0)
		u.fArr = make([]func([]string), 0)
	}
	u.nArr = append(u.nArr, name)
	u.fArr = append(u.fArr, handler)
}

// Run displays the menu and calls the appropriate handle functions when a valid command is entered. If the command is not valid, an error prints. Both "help" and "?" print available commands, but can both be overridden by an added command.
func (u *UI) Run() {
	for {
	start:
		var fc = string(interrogate(u.Name + "> "))

		if fc == "" {
			u.help()
			goto start
		}

		var cmd = parseCmd(fc)

		var validCommand = false
		for i := 0; i < len(u.nArr); i++ {
			if cmd[0] == u.nArr[i] {
				u.fArr[i](cmd[1:])
				validCommand = true
				break
			}
		}
		if !validCommand {
			if cmd[0] == "help" || cmd[0] == "?" {
				u.help()
			} else {
				fmt.Println("Invalid command.")
			}
		}
	}
}

func (u *UI) help() {
	fmt.Println("Available commands:")
	for _, c := range u.nArr {
		fmt.Println("\t" + c)
	}
}

func interrogate(q string) []byte {
	var sc = bufio.NewReader(os.Stdin)
	fmt.Print(q)
	str, err := sc.ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading user input")
		fmt.Println(err)
		os.Exit(1)
	}
	return str[:len(str)-1]
}

func parseCmd(s string) []string {
	var ret = make([]string, 0)
	var x = -1
	var rs = append([]rune(s), ' ')
	for i, r := range rs {
		if r != ' ' && x == -1 {
			x = i
		} else if r == ' ' && x != -1 {
			ret = append(ret, s[x:i])
			x = -1
		}
	}
	return ret
}
