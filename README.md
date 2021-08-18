# cliui
tiny library to make simple text based menus

two functions used to set up the menu and run it

Add(name string, handler func([]string)) adds a command to the possible pool of commands. The handler is a pointer to a function that will get called when the command is run, getting passed an array of strings consisting of the arguments to the command. Note that the command itself is not part of the arguments array.

Run() prints a menu and ask for a command, running the command if it is valid, printing help if no command is specified or if the command is help.
