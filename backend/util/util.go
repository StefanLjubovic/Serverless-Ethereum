package util

import (
	"encoding/json"
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

func ConvertUSDToWei(usdPrice float64) (*big.Int, error) {
	ethRate, _ := GetETHExchangeRate()
	usdToEthPrice := usdPrice / ethRate
	ethInWei := big.NewFloat(usdToEthPrice * 1e18) // Convert ETH to wei (1 ETH = 10^18 wei)

	ethInWeiInt, _ := ethInWei.Int(nil) // Convert big float to big int

	return ethInWeiInt, nil
}
