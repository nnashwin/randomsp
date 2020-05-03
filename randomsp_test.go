package randomsp

import (
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestUnitGetRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ss := strings.Split(str, "")

	if getRandomString(ss) == getRandomString(ss) {
		t.Error("The Get Rand String function should get different strings from the slice")
	}
}

type testStockIndex struct {
	IndexName string
	Stocks    []string
}

// a test to hit all stock data sources and verify that acronyms are returned
func TestIntegrationAcronymsReturned(t *testing.T) {
	var stocks []testStockIndex

	dax, err := getDaxStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed with the following error: %s", err)
	}

	daxIdx := testStockIndex{
		"DAX",
		dax,
	}

	ft, err := getFinancialTimesStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed with the following error: %s", err)
	}

	ftIdx := testStockIndex{
		"FinancialTimes",
		ft,
	}

	ift, err := getItalianFinancialTimesStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed with the following error: %s", err)
	}

	iftIdx := testStockIndex{
		"ItalianFinancialTimes",
		ift,
	}

	n, err := getNasdaqStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed with the following error: %s", err)
	}

	nIdx := testStockIndex{
		"Nasdaq",
		n,
	}

	sp, err := getStandardPoorsStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed with the following error: %s", err)
	}

	spIdx := testStockIndex{
		"StandardPoor",
		sp,
	}

	stocks = append(stocks, daxIdx)
	stocks = append(stocks, ftIdx)
	stocks = append(stocks, iftIdx)
	stocks = append(stocks, nIdx)
	stocks = append(stocks, spIdx)

	// Create Regex To Test that inputs have sequences of capital letters (prove that it is a stock quote acryonym
	stockRegex := regexp.MustCompile("[A-Z]+")

	for _, idxSlice := range stocks {
		for _, stock := range idxSlice.Stocks {
			if stockRegex.FindStringIndex(stock) == nil {
				t.Fatalf("the stock returned does not match the stock regex and is probably not a stock.  \nIndex: %q\nStock: %q", idxSlice.IndexName, stock)
			}
		}
	}
}
