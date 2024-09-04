package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Invoice struct {
	Dia   int     `json:"dia"`
	Valor float32 `json:"valor"`
}

func ReadJson(path string) []Invoice {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	file, err := os.Open(basepath + "/" + path)

	if err != nil {
		log.Fatal("error to open json file:", err)
	}

	defer file.Close()

	bv, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("error to read json file:", err)
	}

	var invoices []Invoice

	if err := json.Unmarshal(bv, &invoices); err != nil {
		log.Fatal("error parsing json data:", err)
	}

	return invoices
}

func minInvoiceOfMonth(invoices []Invoice) Invoice {
	min := invoices[0].Valor
	minIndex := 0
	for i := 1; i < len(invoices)-1; i++ {
		if invoices[i].Valor != 0 && invoices[i].Valor < min {
			min = invoices[i].Valor
			minIndex = i
		}
	}
	return invoices[minIndex]
}

func maxInvoiceOfMonth(invoices []Invoice) Invoice {
	max := invoices[0].Valor
	maxIndex := 0
	for i := 1; i < len(invoices)-1; i++ {
		if invoices[i].Valor > max {
			max = invoices[i].Valor
			maxIndex = i
		}
	}
	return invoices[maxIndex]
}

func daysHigherThanMonthlyAverage(invoices []Invoice) int {
	var sum, count float32
	for _, invoice := range invoices {
		if invoice.Valor > 0 {
			sum += invoice.Valor
			count += 1
		}
	}

	mean := sum / count

	days := 0

	for _, invoice := range invoices {
		if invoice.Valor > mean {
			days++
		}
	}

	return days
}

func main() {
	invoices := ReadJson("invoice.json")
	fmt.Printf("O menor valor de faturamento ocorrido em um dia do mês: %+v\n", minInvoiceOfMonth(invoices))
	fmt.Printf("O maior valor de faturamento ocorrido em um dia do mês: %+v\n", maxInvoiceOfMonth(invoices))
	fmt.Printf("Número de dias em que o valor de faturamento diário foi superior à média mensal: %d\n", daysHigherThanMonthlyAverage(invoices))
}
