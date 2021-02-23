package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func yesno(message string, defaultYes bool) bool {
	fmt.Print(message)
	if defaultYes {
		fmt.Print("[Y/n]")
	} else {
		fmt.Print("[y/N]")
	}
	defer fmt.Println()

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if char == 'y' {
			return true
		}
		if char == 'n' {
			return false
		}
		if key == keyboard.KeyEnter {
			return defaultYes
		}
	}
	return defaultYes
}
