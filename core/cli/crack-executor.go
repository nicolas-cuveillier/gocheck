package cli

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/nicolas-cuveillier/gocheck/core/utils"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	uppercase     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits        = "0123456789"
	special       = "!@#$%^&*()-_=+[]{}|;:,.<>?/\\~`"
	allCharacters = strings.ToLower(uppercase) + uppercase + digits + special
)

func CmdCrack(args []string) {
	if len(args) == 0 {
		fmt.Println("You must provide a valid password")
		return
	}

	password := args[0]
	if len(strings.TrimSpace(password)) <= 0 {
		color.Red("You must provide a valid password")
		return
	}

	bfFlag := false
	daFlag := false
	maxLength := 2
	if len(args) > 1 {
		for k, arg := range args {
			params := strings.Split(arg, "")
			if params[0] == "-" {
				for _, param := range params[1:] {
					switch param {
					case "d":
						daFlag = true
					case "b":
						bfFlag = true
					case "l":
						if k+1 < len(args) {
							maxLen, err := strconv.Atoi(args[k+1])
							if err != nil {
								color.Red("You must provide a valid max length after the -l flag")
								return
							}
							maxLength = maxLen
						}
					default:
						color.Red("Invalid flag: %s", param)
						return
					}
				}
			}
		}
	}

	hashedPassword := utils.HashPassword(password)
	fmt.Printf("Cracking the password : %s (hash value : %s)\n", password, hashedPassword)

	if !daFlag && !bfFlag {
		color.Yellow("Provide a flag to launch a specific attack")
		//CmdHelp()
		return
	}

	///////////////////////// HASHED BRUTE-FORCE ////////////////////////////////////
	if bfFlag {
		fmt.Println("\n============= HASHED BRUTE-FORCE =============")
		totTries := uint64(math.Pow(float64(len(allCharacters)), float64(maxLength)))
		nbTries := 0
		start := time.Now()
		t := time.Now()
		foundPasswordBF := ""

		data := make(chan string)
		isDone := false

		go func() {
			queue := strings.Split(allCharacters, "")
			for i := 0; i < maxLength-1; i++ {
				for _, item := range queue {
					for _, v := range allCharacters {
						newItem := item + string(v)
						data <- newItem
						queue = append(queue, newItem)
					}
				}
			}

			isDone = true
		}()

		for !isDone && nbTries < int(totTries) {
			pw := <-data
			pwHash := utils.HashPassword(pw)
			if pwHash == hashedPassword {
				foundPasswordBF = pw
				isDone = true
				t = time.Now()
				break
			}

			nbTries++
			percent := float64(nbTries) / float64(totTries) * 100
			fmt.Printf("\rnumber of tries: %d/%d (%.2f%%)", nbTries, totTries, percent)
		}

		elapsed := t.Sub(start)
		if foundPasswordBF != "" {
			color.Red("\nFound matching password : %s in %s \n", foundPasswordBF, elapsed)
		} else {
			color.Green("\nNo matching password found after %s \n", time.Since(start))
		}
	}

	///////////////////////// HASHED DICTIONARY ////////////////////////////////////
	if daFlag {
		fmt.Println("\n============= DICTIONARY ATTACK =============")
		foundPasswordDA := ""

		file, err := utils.OpenAssetsFile("10-million-password-list-top-1000000.txt")
		if err != nil {
			log.Fatal(err)
		}

		// metrics
		fmt.Println("Testing most common used passwords...")
		bar := pb.StartNew(1000000)
		start := time.Now()
		t := time.Now()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := scanner.Text()
			hashedWord := utils.HashPassword(word)

			if hashedWord == hashedPassword {
				foundPasswordDA = word
				t = time.Now()
				break
			}
			bar.Increment()
		}
		bar.Finish()

		// Close the file at the end of the program
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}

		elapsed := t.Sub(start)
		if foundPasswordDA != "" {
			color.Red("Found matching password : %s in %s \n", foundPasswordDA, elapsed)
		} else {
			color.Green("No matching password found \n")
		}
	}
}
