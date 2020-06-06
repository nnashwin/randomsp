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
		t.Fatalf("acronym integration test failed getting Dax stocks with the following error: %s", err)
	}

	daxIdx := testStockIndex{
		"DAX",
		dax,
	}

	ft, err := getFinancialTimesStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed getting Financial Times stocks with the following error: %s", err)
	}

	ftIdx := testStockIndex{
		"FinancialTimes",
		ft,
	}

	ift, err := getItalianFinancialTimesStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed getting Italian Financial Times stocks with the following error: %s", err)
	}

	iftIdx := testStockIndex{
		"ItalianFinancialTimes",
		ift,
	}

	n, err := getNasdaqStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed getting Nasdaq stocks with the following error: %s", err)
	}

	nIdx := testStockIndex{
		"Nasdaq",
		n,
	}

	nik, err := getNikkeiStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed getting Nikkei stocks with the following error: %s", err)
	}

	nikIdx := testStockIndex{
		"Nikkei",
		nik,
	}

	sp, err := getStandardPoorsStocks()
	if err != nil {
		t.Fatalf("acronym integration test failed getting StandardPoors stocks with the following error: %s", err)
	}

	spIdx := testStockIndex{
		"StandardPoor",
		sp,
	}

	stocks = append(stocks, daxIdx)
	stocks = append(stocks, ftIdx)
	stocks = append(stocks, iftIdx)
	stocks = append(stocks, nIdx)
	stocks = append(stocks, nikIdx)
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

// create a struct to allow easier printing of the function name
type testStockFn struct {
	Name string
	Fn   func() (Stock, error)
}

func TestIntegrationStockStringsRemoveSpaces(t *testing.T) {
	stockFuncs := []testStockFn{
		testStockFn{Name: "GetRandomNasdaqStock", Fn: GetRandomNasdaqStock},
		testStockFn{Name: "GetRandomNikkeiStock", Fn: GetRandomNikkeiStock},
		testStockFn{Name: "GetRandomSPStock", Fn: GetRandomSPStock},
		testStockFn{Name: "GetRandomFinancialTimesStock", Fn: GetRandomFinancialTimesStock},
		testStockFn{Name: "GetRandomItalianFinancialTimesStock", Fn: GetRandomItalianFinancialTimesStock},
		testStockFn{Name: "GetRandomDaxStock", Fn: GetRandomDaxStock},
	}

	for _, stockFn := range stockFuncs {
		stock, err := stockFn.Fn()
		if err != nil {
			t.Fatalf("The stockFn did not return a stock correctly. Stock Function: %s\nError: %s\n", stockFn.Name, err)
		}

		sLen := len(stock.Symbol)
		if string(stock.Symbol[0]) == " " || string(stock.Symbol[sLen-1]) == " " {
			t.Fatalf("A stock was returned with an untrimmed Symbol. Symbol: %q\n, First Character: %s\n, Last Character: %s\n", stock.Symbol, string(stock.Symbol[0]), string(stock.Symbol[sLen-1]))
		}
	}
}
