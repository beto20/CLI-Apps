package main

import (
    "bufio"
    "errors"
    "flag"
    "fmt"
	"os"
	"strings"

	password "github.com/beto20/CLI-Password-Encrypt/cmd"
)

const (
	  filename = "/Users/beto/Documents_local/1.AmbDesarrollo/data/my_passwords.json"
	  version_app = "0.0.1-beta"
)

func main() {
	add := flag.Bool("p", false, "register a user & password")
	list := flag.Bool("l", false, "list rows")
	getOne := flag.Bool("g", false, "get one row")
	edit := flag.Bool("e", false, "edit user & password")
	delete := flag.Bool("d", false, "delete row")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	pps := &password.PasswordStruct{}
	err := loadData(filename, pps)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		registerPassword(pps)

	case *list:
		listPasswords(pps)

	case *getOne:
		getPassword(filename, pps)

	case *edit:
		editRow(filename, pps)

	case *delete:
		deleteRow(filename, pps)

	case *version:
		fmt.Println("current version:", version_app)

	default:
		fmt.Fprintln(os.Stdout, "Invalid command, user -help to know the commands")
		os.Exit(0)
	}
}

func registerPassword(pps *password.PasswordStruct) {
	data, err := getInput(os.Stdin, flag.Args())

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	input := strings.Split(data, ",")

	ppi := password.PasswordInput{
		Key:      input[0],
		Password: input[1],
	}

	pps.RecievePassword(ppi)
	err = pps.Save(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("registering a password")
}

func getInput(reader *os.File, args []string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, ""), nil
	}

	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	err := scanner.Err()

	if err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty value is not allowed")
	}

	return text, nil
}

func listPasswords(pps *password.PasswordStruct) {
	pps.GetPasswords()
}

func loadData(filename string, pps *password.PasswordStruct) error {
	return pps.LoadPasswords(filename)
}

func getPassword(filename string, pps *password.PasswordStruct) {
	data, err := getInput(os.Stdin, flag.Args())

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	resp, err := pps.GetPasswordById(data, filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	if resp.Id != 0 {
		fmt.Printf("id: %v username: %v, password: %v", resp.Id, resp.Key, resp.PasswordEncrypted)
	}
}

func editRow(filename string, pps *password.PasswordStruct) {
	data, err := getInput(os.Stdin, flag.Args())

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	input := strings.Split(data, ",")

	er := password.EditRow{
		Id:       input[0],
		Key:      input[1],
		Password: input[2],
	}

	pps.EditRowById(filename, er)
	fmt.Println("Edited")
}

func deleteRow(filename string, pps *password.PasswordStruct) {
	data, err := getInput(os.Stdin, flag.Args())

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	pps.DeleteRow(data, filename)
	fmt.Println("Deleted")
}
