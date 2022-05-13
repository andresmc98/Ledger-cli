package main

import "testing"

func TestLedgerFileType(t *testing.T) {
	scenarios := []struct {
		path           string
		ledgerFileType LedgerFileType
	}{
		{
			path:           "./fixtures/income.ledger",
			ledgerFileType: IncomeLedgerFileType,
		},
		{
			path:           "./fixtures/income.ledger",
			ledgerFileType: IncomeLedgerFileType,
		},
		{
			path:           "./fixtures/unknown.ledger",
			ledgerFileType: UnknownLedgerFileType,
		},
	}
	for _, scenario := range scenarios {
		ledgerType, err := OpenLedgerFile(scenario.path)
		if err != nil {
			t.Error(err)
		}
		if ledgerType != scenario.ledgerFileType {
			t.Errorf("expected ledger type %s got %s", scenario.ledgerFileType, ledgerType)
		}
	}
}

func TestIncomFileType(t *testing.T) {
	_, err := OpenLedgerFile("./fixtures/income.ledger")
	if err != nil {
		t.Error(err)
	}
	incomes := readIncomes("./fixtures/income.ledger")
	if len(incomes) != 2 {
		t.Errorf("expected incomes length to be equal 2 got %d", len(incomes))
	}
}
