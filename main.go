package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/peterh/liner"
)

const (
	version         = `0.1.0-alpha`
	versionDate     = `2015-02-20`
	historyFilename = "/tmp/.galapagos_liner_history"
)

var (
	// TODO add functionality for these missing commands
	commandCompletions = []string{"q", "p", "h"}
	// TODO wordCompletions    = []string{"def", "update!"}
)

func configureLiner(linerState *liner.State) {
	linerState.SetCtrlCAborts(true)

	linerState.SetCompleter(func(line string) (c []string) {
		for _, n := range commandCompletions {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	/* TODO
	// WordCompleter takes the currently edited line with the cursor position and
	// returns the completion candidates for the partial word to be completed. If
	// the line is "Hello, wo!!!" and the cursor is before the first '!',
	// ("Hello, wo!!!", 9) is passed to the completer which may returns
	// ("Hello, ", {"world", "Word"}, "!!!") to have "Hello, world!!!".
	linerState.SetWordCompleter(func(line string, pos int) (head string, completions []string, tail string) {
		for _, n := range wordCompletions {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})
	*/
}

func openLinerHistory(line *liner.State) {
	if f, err := os.Open(historyFilename); err == nil {
		line.ReadHistory(f)
		f.Close()
	}
}

func writeLinerHistory(line *liner.State) {
	if f, err := os.Create(historyFilename); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}

func main() {

	showHelp := flag.Bool("help", false, "show the help")
	flag.Parse()

	if showHelp != nil && *showHelp {
		fmt.Printf("Usage of galapagos:\n")
		flag.PrintDefaults()
		return
	}

	// Setup liner

	line := liner.NewLiner()
	defer line.Close()
	openLinerHistory(line)
	configureLiner(line)

	// Initialize

	fmt.Printf("Galapagos %s (%s)\n", version, versionDate)
	fmt.Printf("(Press Ctrl+C or q to exit, h for help)\n\n")

	seedArena()

	// REPL

	for {
		input, err := line.Prompt("> ")

		if err != nil {
			if err.Error() == "prompt aborted" {
				fmt.Printf("Quiting...\n")
			} else {
				fmt.Printf("Prompt error: %s\n", err)
			}
			return
		}

		line.AppendHistory(input)
		writeLinerHistory(line)

		switch {
		case strings.HasPrefix(input, "o"):
			i, err := strconv.Atoi(input[1:])
			if err == nil {
				fmt.Println("Fetching original competitor with strainId", i)
				fmt.Println(originalArena[i])
			} else {
				fmt.Println("Error reading number")
			}
		case input == "q":
			return
		case input == "s":
			fmt.Println(stats())
		case input == "r":
			seedArena()
			fmt.Println("Arena has been reseeded")
		case input == "h":
			fmt.Println("COMMAND  RESULT")
			fmt.Println("-------  ------------------------")
			fmt.Println("<enter>  run 1 round + show stats")
			fmt.Println("#        run # rounds + show stats")
			fmt.Println("o#       show original competitor with strainId")
			fmt.Println("s        show stats")
			fmt.Println("r        reseed arena from scratch")
			fmt.Println("q        quit")
		default:
			i, err := strconv.Atoi(input)
			if err == nil {
				for round := 0; round < i; round++ {
					runRound(round)
					// Show stats for every tenth round, starting with one
					if round%10 == 1 {
						fmt.Println("ROUND #", round)
						fmt.Println(stats())
					}
					pruneAndReproduce()
				}
			} else {
				runRound(1)
				fmt.Println(stats())
				pruneAndReproduce()
			}
		}
	}
}
