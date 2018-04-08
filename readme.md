# randomsp
> A library to pull and return a random stock from a smattering of stock indices

Based on some research through computer simulations, it was found that perhaps a more random strategy to investing in the stock market was more stable over the long term than using any particular strategy.

[Link to forbes article that explains the study](https://www.forbes.com/sites/alexknapp/2013/03/22/computer-simulation-suggests-that-the-best-investment-strategy-is-a-random-one/#2189846a5136)

In the same vein of thought, randomsp was born.

## Indices
The study from the article used 4 different stock indices to make their predictions: the UK FTSE, the MIB FTSE (Italian stock exchange), the DAX (German market), and the S&P 500.

This library will include those plus the Nasdaq stock exchange.

## Install
```
$ go get github.com/ru-lai/randomsp
```

## Usage
```
import "github.com/ru-lai/randomsp"

func main() {
	randomsp.GetRandomIndexStock()
	// ZTS

	randomsp.GetRandomNasdaqStock()
	// ATVI

	randomsp.GetRandomFinancialTimesStock()
	// G4S

	randomsp.GetRandomSPStock()
	// WRK

	randomsp.GetRandomItalianFinancialTimesStock()
	// BPE
}
```

## API
Each method will return a string of a random stock from either the index you called or a random index 

## License

MIT © Tyler Boright
