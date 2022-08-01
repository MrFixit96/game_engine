package main // import "github.com/mrfixit96/game_engine

import (
	"os"

	"github.com/mrfixit96/game_engine/cmd"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}
