package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

// Inspired to solve https://open.kattis.com/problems/jobexpenses

type ExpenseData struct {
	Entry []int `json:"Expense"`
}

type Expenses struct {
	NumEnt  int         `json:"NumberOfEntries"`
	Entries ExpenseData `json:"Entries"`
	Total   int         `json:"Total"`
}

// wg creates a waitgroup to increment before each goRoutine. Outside a function so it has file scope
var wg sync.WaitGroup

func main() {
	// md is a []string that gets holds the multiple JSON values
	md := GetData()

	// range over the []string creating a go routine for each one and a waitgroup
	for i, s := range md {
		wg.Add(1)

		go CreateExpenseReport(s, i)
	}

	wg.Wait()
}

// CreateExpenseReport maps the JSON to a struct, calculates the expenses
func CreateExpenseReport(data string, ind int) {
	var ed1 ExpenseData

	UnpackData(data, &ed1)

	e := GetExpenses(&ed1)

	t := CalculateExpenses(*e)

	e.Total = t

	f := e.PackData()

	fmt.Println(f)

	wg.Done()
}

// UnpackData unmarshals the JSON into a ExpenseData struct
func UnpackData(si string, e *ExpenseData) {
	bs := []byte(si)

	err := json.Unmarshal(bs, &e)
	if err != nil {
		fmt.Printf("Error in ProcessData unmarshalling data. Err message: %v", err)
	}

}

// PackData returns a JSON of total expenses
func (e Expenses) PackData() string {
	bs, err := json.Marshal(e)
	if err != nil {
		fmt.Printf("Error in PackData() when marshalling. Err : %v", err)
	}

	return string(bs)
}

// GetExpenses counts number of expenses and maps the expenseData to the Expense struct
func GetExpenses(ed *ExpenseData) *Expenses {
	n := ed.CountIncExp()

	e := Expenses{
		n,
		*ed,
		0,
	}

	return &e
}

// CountExpenses loops through counting both income and expenses returning total number of expenses
func (ed *ExpenseData) CountIncExp() int {
	var c int

	for i := range ed.Entry {
		c = i + 1
	}

	return c
}

// CalculateExpenses calculates the total of expenses, returns the total
func CalculateExpenses(e Expenses) int {
	var t int

	for _, v := range e.Entries.Entry {
		if v <= 0 {
			t += v
		}

	}

	return t
}
