package main

import "fmt"

type Invoice struct {
	State string  `json:"state"`
	Value float32 `json:"value"`
}

type Representation struct {
	State      string  `json:"state"`
	Percentage float32 `json:"percentage"`
}

func representationPercentage(invoices []Invoice) []Representation {
	var total float32
	for _, invoice := range invoices {
		total += invoice.Value
	}

	representations := make([]Representation, 0)
	for _, invoice := range invoices {
		representations = append(representations, Representation{
			State:      invoice.State,
			Percentage: (invoice.Value * 100) / total,
		})
	}

	return representations
}

func main() {
	invoices := []Invoice{
		{
			State: "SP",
			Value: 67836.43,
		},
		{
			State: "RJ",
			Value: 36678.66,
		},
		{
			State: "MG",
			Value: 29229.88,
		},
		{
			State: "ES",
			Value: 27165.48,
		},
		{
			State: "Outros",
			Value: 19849.53,
		},
	}

	representations := representationPercentage(invoices)

	for _, representation := range representations {
		fmt.Printf("Percentual de representação -> Estado: %s - Valor: %.2f%% \n", representation.State, representation.Percentage)
	}
}
