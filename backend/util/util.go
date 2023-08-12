package util

import (
	"encoding/json"
	"errors"
	"math/big"
	"math/rand"
	"net/http"
	"time"
)

func GetETHExchangeRate() (float64, error) {
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd")
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	var result map[string]map[string]float64
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	ethUSD := result["ethereum"]["usd"]
	return ethUSD, nil
}

func ConvertEthToWei(ethValue float64) (*big.Int, error) {
	// Convert Ether to Wei using 10^18 as the conversion factor
	weiValue := new(big.Int)
	weiValue, ok := weiValue.SetString(big.NewFloat(ethValue).Mul(big.NewFloat(1e18), big.NewFloat(1e18)).Text('f', 0), 10)
	if !ok {
		return nil, errors.New("conversion error")
	}

	return weiValue, nil
}

const customEpoch = 1300000000000

func GenerateRowID(shardID int) uint {
	ts := uint(time.Now().UnixNano()/int64(time.Millisecond) - customEpoch)
	randID := uint(rand.Intn(512))
	ts = (ts << 6)
	ts = ts + uint(shardID)
	return (ts << 9) + randID
}

func ConvertUSDToETH(usdPrice float64) (float64, error) {
	ethRate, err := GetETHExchangeRate()
	if err != nil {
		return 0, err
	}

	ethPrice := usdPrice / ethRate
	return ethPrice, nil
}
