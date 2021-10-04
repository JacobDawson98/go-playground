package main

import "testing"

func TestGetSplitPayments_ShouldThrowErrorIfNoPersonsGiven(t *testing.T) {
	var personIns []PersonIn

	_, err := GetSplitPayments(personIns)

	if err == nil {
		t.Fail()
	}
}

func TestGetSplitPayments_ShouldThrowErrorIfOnlyGivenOnePerson(t *testing.T) {
	personIns := []PersonIn{
		PersonIn{ Name: "Bob", AmountPaid: 13.75 },
	}

	_, err := GetSplitPayments(personIns)

	if err == nil {
		t.Fail()
	}
}

func TestGetSplitPayments_ShouldSplitPaymentsBetweenTwoPeopleWhenOnlyOnePersonPaid(t *testing.T) {
	personIns := []PersonIn{
		PersonIn{Name: "Bob", AmountPaid: 0},
		PersonIn{Name: "Alice", AmountPaid: 12.50},
	}

	personOuts, err := GetSplitPayments(personIns)

	if err != nil {
		t.Fail()
	}
	if len(personOuts) != 2 {
		t.Fail()
	}
	var bobOut PersonOut
	var aliceOut PersonOut
	for _, personOut := range personOuts {
		if personOut.Name == "Alice" {
			aliceOut = personOut
		}
		if personOut.Name == "Bob" {
			bobOut = personOut
		}
	}

	if len(bobOut.Debts) != 1 {
		t.Fail()
	}
	if bobOut.Debts[0].Amount != 6.25 || bobOut.Debts[0].OwedTo != "Alice" {
		t.Fail()
	}
	if len(aliceOut.Debts) != 0 {
		t.Fail()
	}
}
