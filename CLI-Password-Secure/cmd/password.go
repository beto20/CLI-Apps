package password

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

type PasswordStruct struct {
	Id                int32
	Key               string
	PasswordEncrypted string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type PasswordInput struct {
	Key      string
	Password string
}

type EditRow struct {
	Id       string
	Key      string
	Password string
}

var Array = []PasswordStruct{}

func (ps *PasswordStruct) RecievePassword(input PasswordInput) {
	encrypted, err := encrypt(input.Password)

	if err != nil {
		fmt.Println("Error when recieving password")
	}

	pssw := PasswordStruct{
		Id:                generateRandomId(),
		Key:               input.Key,
		PasswordEncrypted: encrypted,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	Array = append(Array, pssw)
}

func (ps *PasswordStruct) Save(filename string) error {
	data, err := json.Marshal(Array)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, os.FileMode(0644))
}

func (ps *PasswordStruct) GetPasswords() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Key"},
			{Align: simpletable.AlignCenter, Text: "Password"},
			{Align: simpletable.AlignCenter, Text: "Created at"},
			{Align: simpletable.AlignCenter, Text: "Updated at"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, item := range Array {

		content := []*simpletable.Cell{
			{Text: strconv.Itoa(int(item.Id))},
			{Text: item.Key},
			{Text: item.PasswordEncrypted},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.UpdatedAt.Format(time.RFC822)},
		}

		cells = append(cells, content)
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("Rows registered: %d", len(Array))},
	}}

	table.SetStyle(simpletable.StyleRounded)
	table.Println()
}

func (ps *PasswordStruct) LoadPasswords(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, &Array)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PasswordStruct) GetPasswordById(id string, filename string) (PasswordStruct, error) {
	num32, _ := stringToInt32(id)

	for _, item := range Array {
		if item.Id == num32 {
			password, err := decrypt(item.PasswordEncrypted)

			if err != nil {
				return PasswordStruct{}, errors.New("error when get password")
			}

			return PasswordStruct{
				Id:                item.Id,
				Key:               item.Key,
				PasswordEncrypted: password,
				CreatedAt:         item.CreatedAt,
				UpdatedAt:         item.UpdatedAt,
			}, nil
		}
	}

	return PasswordStruct{}, errors.New("id not found or is invalid")
}

func (ps *PasswordStruct) EditRowById(filename string, editRow EditRow) {
	id, _ := stringToInt32(editRow.Id)

	for i, item := range Array {
		if item.Id == id {
			password, err := encrypt(editRow.Password)

			if err != nil {
				errors.New("error when get password")
			}

			rowEdited := PasswordStruct{
				Id:                item.Id,
				Key:               editRow.Key,
				PasswordEncrypted: password,
				CreatedAt:         item.CreatedAt,
				UpdatedAt:         time.Now(),
			}

			Array[i] = rowEdited
			ps.Save(filename)
		}
	}
}

func (ps *PasswordStruct) DeleteRow(id string, filename string) {
	num, _ := stringToInt32(id)

	index := -1

	for i, item := range Array {
		if item.Id == num {
			index = i
			break
		}
	}

	if index != -1 {
		Array = append(Array[:index], Array[index+1:]...)
		ps.Save(filename)
	}
}

func stringToInt32(value string) (int32, error) {
	num, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		fmt.Println("error when converting", err)
		return int32(0), errors.New("error when converting")
	}
	return int32(num), nil
}

func generateRandomId() int32 {
	code := rand.Int31()

	return code
}

func encrypt(password string) (string, error) {
	key := []byte("AES256Key-32Characters1234567890")
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	plaintextBytes := []byte(password)
	plaintextBytes = pKCS7Padding(plaintextBytes, aes.BlockSize)

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertext, plaintextBytes)

	encrypted := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func pKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func decrypt(passwordEncrypted string) (string, error) {
	key := []byte("AES256Key-32Characters1234567890")
	ciphertextBytes, err := base64.StdEncoding.DecodeString(passwordEncrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertextBytes))
	mode.CryptBlocks(decrypted, ciphertextBytes)

	decrypted = pKCS7Trimming(decrypted)

	return string(decrypted), nil
}

func pKCS7Trimming(data []byte) []byte {
	length := len(data)
	padding := int(data[length-1])
	return data[:(length - padding)]
}
