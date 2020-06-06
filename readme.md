# randomsp
> A library to pull and return a random stock from a smattering of stock indices

Based on some research through computer simulations, it was found that perhaps a more random strategy to investing in the stock market was more stable over the long term than using any particular strategy.

[Link to forbes article that explains the study](https://www.forbes.com/sites/alexknapp/2013/03/22/computer-simulation-suggests-that-the-best-investment-strategy-is-a-random-one/#2189846a5136)

In the same vein of thought, randomsp was born.

## Indices
The study from the article used 5 different stock indices to make their predictions: the UK FTSE, the MIB FTSE (Italian stock exchange), the Nikkei 225 (Japanese market), the DAX (German market), and the S&P 500.

This library includes those 5 indices plus the Nasdaq stock exchange.

## Notes
- The Nikkei 225 stocks have symbols, but more often than not are denoted by a number value in their index.  Because of this, the name of the company and number are returned instead of the symbol acryonym.

## Install
```
$ go get github.com/ru-lai/randomsp
```

## Usage
```
import "github.com/ru-lai/randomsp"

func main() {
	stock, err := randomsp.GetRandomIndexStock()
	// ZTS

	stock, err := randomsp.GetRandomNasdaqStock()
	// ATVI

	stock, err := randomsp.GetRandomFinancialTimesStock()
	// G4S

	stock, err := randomsp.GetRandomSPStock()
	// WRK

	stock, err := randomsp.GetRandomItalianFinancialTimesStock()
	// BPE

        stock, err := randomsp.GetRandomNikkeiStock()
        // Tokyu Land 3289

	stock, err := randomsp.GetRandomDaxStock()
	// BMW

        // each function call should have error handling
        if err != nil {
            // handle error here
            log.Fatal(err)
        }
}
```

## API
Each method will return a string of a random stock from either the index you called or a random index and an error if applicable

## License

MIT © Tyler Boright
