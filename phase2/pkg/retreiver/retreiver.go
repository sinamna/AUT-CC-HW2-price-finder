package retreiver

import (
	"encoding/json"
	"fmt"
	"github.com/ccloud/phase2/pkg/config"
	"github.com/ccloud/phase2/pkg/model"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func RetrievePrice(currencyID string) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://rest.coinapi.io/v1/assets/%s", currencyID), nil)
	if err != nil {
		return 0.0, errors.New("could not create request")
	}
	req.Header.Set("X-CoinAPI-Key", config.Conf.Server.ApiKey)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0.0, errors.Wrap(err, "could not make request")
	}
	rspBody, _ := ioutil.ReadAll(rsp.Body)

	var results []model.RetrieverAssets
	err = json.Unmarshal(rspBody, &results)
	if err != nil {
		return 0.0, err
	}
	return results[0].PriceUSD, err
}
