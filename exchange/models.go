package main

type (
	RateResponse struct {
		Success   bool   `json:"success"`
		TimeStamp int    `json:"timestamp"`
		Base      string `json:"base"`
		Date      string `json:"date"`
		Rates     Rate   `json:"rate"`
	}

	Rate struct {
		GBP float64 `json:"GBP"`
	}

	HistoricalRateResponse struct {
		Success    bool   `json:"success"`
		TimeStamp  int    `json:"timestamp"`
		Historical bool   `json:"historical"`
		Base       string `json:"base"`
		Date       string `json:"date"`
		Rates      Rate   `json:"rate"`
	}
)
