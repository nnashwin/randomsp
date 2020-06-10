package randomsp

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tlboright/go-rint"
	"net/http"
	"strings"
)

func getDaxStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/DAX")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tables := doc.Find(".mw-parser-output > table")
	tables.Filter("#constituents").Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(4)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getFinancialTimesStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/FTSE_100_Index")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("#constituents > tbody")
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(2)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getItalianFinancialTimesStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/FTSE_MIB")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("#constituents > tbody")
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(2)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getNasdaqStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/NASDAQ-100")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("#constituents > tbody")
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(2)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getNikkeiStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/Nikkei_225")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	// the list of stock names underneath the titles
	doc.Find("h3 + ul").Each(func(i int, s *goquery.Selection) {
		lis := s.Find("li")
		lis.Each(func(i int, s *goquery.Selection) {
			name := s.Find("a:nth-child(1)")
			symbol := s.Find("a:nth-child(3)")

			if name.Text() != "" && symbol.Text() != "" {
				stocks = append(stocks, fmt.Sprintf("%s %s\n", name.Text(), symbol.Text()))
			}
		})
	})

	return
}

func getStandardPoorsStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("#constituents > tbody")
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:first-child")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getRandomString(ss []string) string {
	randInt := rint.GenRange(1, len(ss))
	return ss[randInt:(randInt + 1)][0]
}

func GetRandomDaxStock() (stock Stock, err error) {
	rint.Init()
	stockSlice, err := getDaxStocks()
	if err != nil {
		return
	}

	stock = Stock{strings.TrimSpace(getRandomString(stockSlice)), "Dax"}
	return
}

func GetRandomFinancialTimesStock() (stock Stock, err error) {
	rint.Init()
	stockSlice, err := getFinancialTimesStocks()
	if err != nil {
		return
	}

	stock = Stock{strings.TrimSpace(getRandomString(stockSlice)), "Financial Times"}
	return
}

func GetRandomItalianFinancialTimesStock() (stock Stock, err error) {
	rint.Init()
	stockSlice, err := getItalianFinancialTimesStocks()
	if err != nil {
		return
	}

	stock = Stock{strings.TrimSpace(getRandomString(stockSlice)), "Italian Financial Times"}
	return
}

func GetRandomNasdaqStock() (stock Stock, err error) {
	rint.Init()
	stockSlice, err := getNasdaqStocks()
	if err != nil {
		return
	}

	stock = Stock{strings.TrimSpace(getRandomString(stockSlice)), "Nasdaq"}

	return
}

func GetRandomNikkeiStock() (stock Stock, err error) {
	rint.Init()
	stockSlice, err := getNikkeiStocks()
	if err != nil {
		return
	}

	stock = Stock{strings.TrimSpace(getRandomString(stockSlice)), "Nikkei"}

	return
}

func GetRandomSPStock() (stock Stock, err error) {
	rint.Init()
	stockSlice, err := getStandardPoorsStocks()
	if err != nil {
		return
	}

	stock = Stock{strings.TrimSpace(getRandomString(stockSlice)), "S&P 500"}

	return
}

func GetRandomIndexStock() (stock Stock, err error) {
	rint.Init()
	stockFuncs := []func() (Stock, error){GetRandomNasdaqStock, GetRandomSPStock, GetRandomNikkeiStock, GetRandomFinancialTimesStock, GetRandomItalianFinancialTimesStock, GetRandomDaxStock}
	stock, err = stockFuncs[rint.GenRange(0, len(stockFuncs))]()
	return
}
