package cli

import (
	"fmt"
	"github.com/nicolas-cuveillier/gocheck/core/utils"
	"io"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func CmdCheck(args []string) {
	fmt.Println()

	if len(args) == 0 {
		fmt.Println("You must provide a valid password")
		return
	}

	password := args[0]
	if len(strings.TrimSpace(password)) <= 0 {
		fmt.Println("You must provide a valid password")
		return
	}

	fmt.Printf("Assessing the strength of the password : %s", args[0])
	fmt.Println()

	// Check password length
	if len(password) < 8 {
		color.Red("Length score %10s", "0/10")
	} else if len(password) >= 8 && len(password) < 12 {
		color.Yellow("Length score %10s", "2/10")
	} else if len(password) >= 12 && len(password) < 16 {
		color.Yellow("Length score %10s", "4/10")
	} else if len(password) >= 16 && len(password) < 20 {
		color.Green("Length score %10s", "6/10")
	} else if len(password) >= 20 && len(password) < 24 {
		color.Green("Length score %10s", "8/10")
	} else if len(password) >= 24 {
		color.Green("Length score %15s", "10/10")
	}

	// Check password entropy & craking time
	s := 0
	if match, _ := regexp.MatchString("[a-z]", password); match {
		s += 26
	} else if match, _ := regexp.MatchString("[A-Z]", password); match {
		s += 26
	} else if match, _ := regexp.MatchString("[0-9]", password); match {
		s += 10
	} else if match, _ := regexp.MatchString(`[^\w\s]`, password); match {
		s += 43
	}

	entropy := math.Log2(math.Pow(float64(s), float64(len(password))))
	guaranteedCrackingTime := (math.Pow(2, entropy) / 100000000000000)
	if entropy < 40 {
		color.Red("Entropy score %14f", entropy)
		color.Red("Guaranteed cracking time: %.2f seconds", guaranteedCrackingTime)
	} else if entropy >= 40 && entropy < 50 {
		color.Red("Entropy score %14f", entropy)
		color.Red("Guaranteed cracking time: %.2f seconds", guaranteedCrackingTime)
	} else if entropy >= 50 && entropy < 60 {
		color.Red("Entropy score %14f", entropy)
		color.Red("Guaranteed cracking time: %.2f minutes", guaranteedCrackingTime/60)
	} else if entropy >= 60 && entropy < 70 {
		color.Yellow("Entropy score %14f", entropy)
		color.Yellow("Guaranteed cracking time: %.2f days", guaranteedCrackingTime/86400)
	} else if entropy >= 70 && entropy < 80 {
		color.Yellow("Entropy score %14f", entropy)
		color.Yellow("Guaranteed cracking time: %.2f days", guaranteedCrackingTime/86400)
	} else if entropy >= 80 {
		color.Green("Entropy score %14f", entropy)
		color.Green("Guaranteed cracking time: %.2f years", guaranteedCrackingTime/31536000)
	}

	// Leak password in haveibeenpwned
	hashedPassword := utils.HashPassword(password)
	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + string(hashedPassword[:5]))
	if err != nil {
		color.Red("Error while checking password in haveibeenpwned")
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("Error while checking password in haveibeenpwned")
		return
	}

	hashList := strings.Split(string(body), "\r\n")
	passwordSeenCount := 0
	for _, s := range hashList {
		if strings.ToUpper(hashedPassword[5:]) == strings.Split(s, ":")[0] {
			passwordSeenCount, _ = strconv.Atoi(strings.Split(s, ":")[1])
		}
	}

	if passwordSeenCount == 0 {
		color.Green("Password not found in haveibeenpwned")
	} else {
		color.Red("Password found %d times in haveibeenpwned", passwordSeenCount)
	}

	fmt.Println()
}

func CmdGenerate(args []string) {
	fmt.Println()

	// Default values
	length := 20
	possibleValueTypes := []string{"number"}

	// Parse the arguments
	for i := 0; i < len(args); i += 1 {
		if len(args[i]) == 2 {

			switch args[i] {
			case "-l":
				if len(args) <= i+1 {
					fmt.Printf("Missing value for option %s \n", args[i])
					return
				}

				length, _ = strconv.Atoi(args[i+1])
				i += 1
			case "-c":
				possibleValueTypes = append(possibleValueTypes, "char")
			case "-u":
				possibleValueTypes = append(possibleValueTypes, "special")
			case "-s":
				possibleValueTypes = append(possibleValueTypes, "special")
			default:
				fmt.Println("Unknown option")
			}
		} else if len(args[i]) > 2 {
			// handle multiple parameters
			param := strings.Trim(args[i], "-")
			params := strings.Split(param, "")

			for j := 0; j < len(params); j += 1 {
				switch params[j] {
				case "c":
					possibleValueTypes = append(possibleValueTypes, "char")
				case "s":
					possibleValueTypes = append(possibleValueTypes, "special")
				case "u":
					possibleValueTypes = append(possibleValueTypes, "uppercase")
				default:
					fmt.Printf("Wrong parameter %s \n", params[j])
					return
				}
			}

		} else {
			fmt.Printf("Wrong parameter %s \n", args[i])
			return
		}
	}

	password := ""

	for i := 0; i < length; i += 1 {
		valueType := possibleValueTypes[rand.Intn(len(possibleValueTypes))]

		switch valueType {
		case "number":
			password = password + strconv.Itoa(rand.Intn(10))
		case "char":
			password = password + string('a'+rune(rand.Intn(26)))
		case "special":
			specialChars := []string{"!", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "]", "^", "_", "{", "}", "|", "~"}
			password = password + specialChars[rand.Intn(len(specialChars))]
		case "uppercase":
			password = password + string('A'+rune(rand.Intn(26)))
		default:
			return
		}

	}

	fmt.Printf("Generated password: %s", password)

	fmt.Println()
}

func CmdCheckHelp() {
	fmt.Println()

	fmt.Printf("%-15s %s", "help", "Show the list of available commands")
	fmt.Println()

	fmt.Printf("%-15s %s", "generate", "Generate a password")
	fmt.Println()

	fmt.Printf("%-15s %s", "check", "Check the strengh of a password")
	fmt.Println()

	fmt.Printf("%-15s %s", "exit", "Exit the program, come back to the main menu")

	fmt.Println()
}
