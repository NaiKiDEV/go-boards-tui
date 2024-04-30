package main

import (
	"fmt"

	"github.com/naikidev/go-boards-tui/boards"
	"github.com/naikidev/go-boards-tui/terminal"
);

func main() {
	fmt.Println("This is a boards app!")

	terminal.RenderModel(boards.InitialModel());	

}