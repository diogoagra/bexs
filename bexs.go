package bexs

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Bexs -
type Bexs struct {
	Exchange  string
	APIKey    string
	APISecret string
	Debug     bool
	Lock      *sync.RWMutex
}

// Exchanges -
var exchanges = map[string]string{
	"bleutrade":      "https://bleutrade.com",
	"exccripto":      "https://trade.exccripto.com",
	"bitrecife":      "https://exchange.bitrecife.com.br",
	"bomesp":         "https://exchange.bomesp.com.br",
	"bomespg":        "https://exchange.bomesp.com",
	"bullgain":       "https://trade.bullgain.com",
	"comprarbitcoin": "https://app.comprarbitcoin.com.br",
	"dev":            "https://dev.bleu.com.br",
	"staging":        "https://staging.bleu.com.br",
	"localhost":      "http://localhost",
}

// New - Exchanges avaliable (bleutrade, exccripto, bitrecife, bomesp, bomespg, bullgain)
func New(exchange, apikey, apisecret string, debug bool) *Bexs {
	return &Bexs{
		Exchange:  exchange,
		APIKey:    apikey,
		APISecret: apisecret,
		Debug:     debug,
		Lock:      &sync.RWMutex{},
	}
}

// getURL -
func (c *Bexs) getURL(endpoint string, private bool) (response []byte, err error) {
	var uri, apisign string

	if private {
		c.Lock.Lock()
		defer c.Lock.Unlock()

		params := make(url.Values)
		params.Add("apikey", c.APIKey)
		params.Add("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

		uri = fmt.Sprintf("%s%s&%s", exchanges[c.Exchange], endpoint, params.Encode())
		apisignhmac512 := hmac.New(sha512.New, []byte(c.APISecret))
		_, err = apisignhmac512.Write([]byte(uri))
		if err != nil {
			return
		}
		apisign = hex.EncodeToString(apisignhmac512.Sum(nil))
	} else {
		uri = exchanges[c.Exchange] + endpoint
	}

	if c.Debug {
		log.Println("[debug] request:", uri)
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		if c.Debug {
			log.Println("[debug] error:", err)
		}
		return
	}

	req.Header.Set("User-Agent", "BexsClient/Golang")

	if private {
		req.Header.Set("apisign", apisign)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if c.Debug {
			log.Println("[debug] error:", err)
		}
		return
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.Debug {
			log.Println("[debug] error:", err)
		}
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("Status code %v - %s", resp.StatusCode, uri)
		if c.Debug {
			log.Println("[debug] error:", err)
		}
		return
	}

	return
}

// parseResult -
func (c *Bexs) parseResult(result []byte) (response []byte, err error) {
	var defaultReturn DefaultReturn
	err = json.Unmarshal(result, &defaultReturn)
	if err != nil {
		if c.Debug {
			log.Println("[debug] error:", err)
		}
		return
	}

	if !defaultReturn.Success {
		err = errors.New(defaultReturn.Message)
		if c.Debug {
			log.Println("[debug] error:", err)
		}
		return
	}

	response, err = json.Marshal(defaultReturn.Result)
	if err != nil {
		if c.Debug {
			log.Println("[debug] error:", err)
		}
		return
	}

	if c.Debug {
		log.Println("[debug] response:", string(response))
	}

	return
}
