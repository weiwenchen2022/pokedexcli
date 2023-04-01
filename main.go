package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"pokedexcli/internal/pokeapi"
)

type callbackFunc func(string) error

type cliCommand struct {
	name        string
	description string
	callback    callbackFunc
}

func (c *cliCommand) String() string {
	return fmt.Sprintf("%s: %s", c.name, c.description)
}

const cacheInterval = 5 * time.Second

var commands map[string]cliCommand
var config *pokeapi.Config

func init() {
	config = pokeapi.NewConfig(cacheInterval)

	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap(config),
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapb(config),
		},

		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokémon in a given area",
			callback:    commandExplore(config),
		},

		"catch": {
			name:        "catch",
			description: "Catching Pokemon adds them to the user's Pokedex",
			callback:    commandCatch(config),
		},

		"inspect": {
			name:        "inspece",
			description: "Displays details about a Pokemon if it have been caught before",
			callback:    commandInspect(config),
		},

		"pokedex": {
			name:        "pokedex",
			description: "Print a list of all the names of the Pokemon you has caught",
			callback:    commandPokedex(config),
		},
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")

	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		cmd := fields[0]
		var arg string
		if len(fields) > 1 {
			arg = fields[1]
		}

		if c, ok := commands[cmd]; !ok {
			fmt.Print("not found command")
		} else {
			err := c.callback(arg)
			if err != nil {
				log.Print(err)
			}
		}

		fmt.Print("Pokedex > ")
	}
}

func commandHelp(_ string) error {
	fmt.Printf("\nWelcome to the Pokedex!")
	fmt.Printf("\nUsage:\n\n")

	var b strings.Builder
	for _, v := range commands {
		fmt.Fprint(&b, v.String(), "\n")
	}
	fmt.Print(b.String())
	fmt.Println()

	return nil
}

func commandExit(_ string) error {
	os.Exit(0)
	return nil
}

func commandMap(c *pokeapi.Config) callbackFunc {
	return func(_ string) error {
		names, err := pokeapi.Next(c)
		if err == nil {
			fmt.Println(strings.Join(names, "\n"))
		}

		return err
	}
}

func commandMapb(c *pokeapi.Config) callbackFunc {
	return func(_ string) error {
		names, err := pokeapi.Previous(c)
		if err == nil {
			fmt.Println(strings.Join(names, "\n"))
		}

		return err
	}
}

func commandExplore(c *pokeapi.Config) callbackFunc {
	return func(areaName string) error {
		fmt.Print("Exploring ", areaName, "...\n")

		names, err := pokeapi.Explore(c, areaName)
		if err == nil {
			fmt.Print("Found Pokemon:")
			for _, n := range names {
				fmt.Print("\n - ", n)
			}
			fmt.Println()
		}

		return err
	}
}

func commandCatch(c *pokeapi.Config) callbackFunc {
	return func(name string) error {
		fmt.Print("Throwing a Pokeball at ", name, "...\n")

		ok, err := pokeapi.Catch(c, name)
		if err != nil {
			return err
		}

		if ok {
			fmt.Print(name, " wat caught!\n")
			fmt.Print("You may now inspect it with the inspect command.\n")
		} else {
			fmt.Print(name, " escaped!\n")
		}

		return nil
	}
}

func commandInspect(c *pokeapi.Config) callbackFunc {
	return func(name string) error {
		if pokemon, ok := c.Pokedex[name]; !ok {
			fmt.Println("you have not caught that pokemon")
		} else {
			fmt.Print(pokemon.String())
		}

		return nil
	}
}

func commandPokedex(c *pokeapi.Config) callbackFunc {
	return func(_ string) error {
		fmt.Println("Your Pokedex:")
		for k := range c.Pokedex {
			fmt.Printf(" - %s\n", k)
		}

		return nil
	}
}
