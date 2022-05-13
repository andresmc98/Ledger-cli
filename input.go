package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type LedgerFileType string

const (
	IncomeLedgerFileType   = "Income"
	UnknownLedgerFileType  = "Unknown"
	ExpensesLedgerFileType = "Expenses"
)

var fileScanner map[string]*bufio.Scanner

func init() {
	fileScanner = make(map[string]*bufio.Scanner)
}

func OpenLedgerFile(path string) (LedgerFileType, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return UnknownLedgerFileType, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return UnknownLedgerFileType, err
	}
	text := scanner.Text()
	switch text {
	case "; Income":
		fileScanner[path] = scanner
		return IncomeLedgerFileType, nil
	case "; Expenses":
		fileScanner[path] = scanner
		return ExpensesLedgerFileType, nil
	default:
		return UnknownLedgerFileType, err
	}
}

type Accountable interface {
	GetData() string
}

type Income struct {
	Data string
}
type Expense struct {
	Data string
}

func (i *Income) GetData() string {
	return "soy un income muy sexy"
}

func readIncomes(path string) []*Income {
	scanner := fileScanner[path]
	incomesOutput := []*Income{}
	for scanner.Scan() {
		if strings.TrimSpace(string(scanner.Text()[0])) != "" {
			incomesOutput = append(incomesOutput, &Income{})
		}
	}
	return incomesOutput
}
func readExpenses(path string) []*Expense {
	scanner := fileScanner[path]
	expensesOutput := []*Expense{}
	for scanner.Scan() {
		if strings.TrimSpace(string(scanner.Text()[0])) != "" {
			expensesOutput = append(expensesOutput, &Expense{})
		}
	}
	return expensesOutput
}
func read(path string) {
	dataBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	dataStr := string(dataBytes)
	fmt.Println(dataStr)
}
