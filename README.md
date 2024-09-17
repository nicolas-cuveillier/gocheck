# GOCHECK

A Go-based CLI application for checking the strength of passwords and 
attempting to crack them using brute-force and dictionary attacks. 
This tool is designed for both educational purposes and to demonstrate 
the risks associated with weak passwords.

## Features

- **Password Strength Checker**: Evaluate the security of passwords based
on length, character diversity, dictionary attacks, and common patterns.
- **Password Cracker**: Use brute-force or dictionary-based techniques 
to attempt cracking weak passwords.
- **Multi-threaded Cracking**: Speed up password cracking with concurrent
Goroutines.
- **Entropy Calculation**: Measure the unpredictability of passwords.
- **Leaked Password Check**: Integrate with external APIs (e.g., 
HaveIBeenPwned) to check if a password has been compromised in data 
breaches.
- **Detailed Reporting**: Get comprehensive reports on password strength
evaluations and cracking attempts.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.23 or higher

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/nicolas-cuveillier/gocheck.git
   cd gocheck
   ```

2. Build the project:

   ```bash
   go build
   ```

3. Run the CLI tool:

   ```bash
   ./gocheck
   ```

## Usage

The Password Strength Checker & Cracker supports multiple commands and flags for different operations. You can check the strength of a password or attempt to crack it using various methods.

### Commands

#### 1. Password Generator

Generate a password

![process](https://raw.githubusercontent.com/nicolas-cuveillier/gocheck/main/.github/gocheck_generate.png)

#### 2. Password Strength Checker

Check the strength of a given password:

![process](https://raw.githubusercontent.com/nicolas-cuveillier/gocheck/main/.github/gocheck_check.png)

#### 2. Password Cracker

##### Brute-Force & Dictionary Attack

Attempt to crack a password using brute-force and a dictionary:

![process](https://raw.githubusercontent.com/nicolas-cuveillier/gocheck/main/.github/gocheck_crack.png)

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -m "Add new feature"`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

## Acknowledgements

- [Go Programming Language](https://golang.org/)
- [HaveIBeenPwned](https://haveibeenpwned.com/)

---

Feel free to reach out if you have any questions or suggestions. Happy hacking!

