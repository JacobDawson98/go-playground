package main

import "errors"

type PersonIn struct {
	Name string
	AmountPaid float32	
}

type Debt struct {
	OwedTo string	
	Amount float32
}

type PersonOut struct {
	Name string
	Debts []Debt	
	differenceOwed float32
	totalOwedToSurplusPayers float32
	surplusPayerPercentage float32
}

func GetSplitPayments(persons []PersonIn) ([]PersonOut, error) {
	if len(persons) < 2 {
		return []PersonOut{}, errors.New("must provide more than one person") 
	}
	var totalPaid float32
	for _, person := range persons {
		totalPaid += person.AmountPaid
	}
	var totalPersons = float32(len(persons))
	totalCostPerPerson := totalPaid / totalPersons 
	var personOuts []PersonOut
	for _, person := range persons {
		personOuts = append(personOuts, PersonOut{ Name: person.Name })
	}
	/**
	difference owed = total cost paid - total cost per person 

	total owed to surplus payers = sum of negative difference owed (using abs val)

	surplus payer percentage of total owed = difference owed / total owed to surplus payers

	alice owes $32, total owed = $500, bob's surplus payer % is 50
	**/

	return 
}

/**
{
	"splitPayments": [
		{
			"name": "Bob",
			"debts": [
				{
					"owedTo": "Alice",
					"amount": 13.41
				}
			]
		}
	]
}
**/
func main() {

}
