package bkex

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"
)

/*The Base Endpoint URL*/
const (
	API_URL = "https://api.bkex.vip"
)

/*API Base Knowledge
Path: API function. Usually after the base endpoint URL
Method:
	Get - Call a URL, API return a response
	Post - Call a URL & send a request, API return a response
Public API:
	It doesn't need authorization/signature , can be called by browser to get response.
	using exchange.HttpGetRequest/exchange.HttpPostRequest
Private API:
	Authorization/Signature is requried. The signature request should look at Exchange API Document.
	using ApiKeyGet/ApiKeyPost
Response:
	Response is a json structure.
	Copy the json to https://transform.now.sh/json-to-go/ convert to go Struct.
	Add the go Struct to model.go

ex. Get /api/v1/depth
Get - Method
/api/v1/depth - Path*/

/*************** Public API ***************/
/*Get Coins Information (If API provide)
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Add Model of API Response
Step 3: Modify API Path(strRequestPath)*/
func (e *Bkex) GetCoinsData() error {
	jsonResponse := &JsonResponse{}
	exchangeData := ExchangeData{}

	strRequestPath := "/v1/exchangeInfo"
	strUrl := API_URL + strRequestPath

	jsonReturn := exchange.HttpGetRequest(strUrl, nil)
	if err := json.Unmarshal([]byte(jsonReturn), &jsonResponse); err != nil {
		return fmt.Errorf("%s Get Coins Json Unmarshal Err: %v %v", e.GetName(), err, jsonReturn)
	} else if jsonResponse.Code != 0 {
		return fmt.Errorf("%s Get Coins Failed: %v", e.GetName(), jsonResponse.Msg)
	}
	if err := json.Unmarshal(jsonResponse.Data, &exchangeData); err != nil {
		return fmt.Errorf("%s Get Coins Result Unmarshal Err: %v %s", e.GetName(), err, jsonResponse.Data)
	}

	for _, data := range exchangeData.CoinTypes {
		if data.SupportTrade {
			c := &coin.Coin{}
			switch e.Source {
			case exchange.EXCHANGE_API:
				c = coin.GetCoin(data.CoinType)
				if c == nil {
					c = &coin.Coin{
						Code: data.CoinType,
					}
					coin.AddCoin(c)
				}
			case exchange.JSON_FILE:
				c = e.GetCoinBySymbol(data.CoinType)
			}

			if c != nil {
				coinConstraint := &exchange.CoinConstraint{
					CoinID:       c.ID,
					Coin:         c,
					ExSymbol:     data.CoinType,
					ChainType:    exchange.MAINNET,
					TxFee:        data.WithdrawFee,
					Withdraw:     data.SupportWithdraw,
					Deposit:      data.SupportDeposit,
					Confirmation: DEFAULT_CONFIRMATION,
					Listed:       DEFAULT_LISTED,
				}

				e.SetCoinConstraint(coinConstraint)
			}
		}
	}
	return nil
}

/* GetPairsData - Get Pairs Information (If API provide)
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Add Model of API Response
Step 3: Modify API Path(strRequestUrl)*/
func (e *Bkex) GetPairsData() error {
	jsonResponse := &JsonResponse{}
	exchangeData := ExchangeData{}

	strRequestPath := "/v1/exchangeInfo"
	strUrl := API_URL + strRequestPath

	jsonReturn := exchange.HttpGetRequest(strUrl, nil)
	if err := json.Unmarshal([]byte(jsonReturn), &jsonResponse); err != nil {
		return fmt.Errorf("%s Get Coins Json Unmarshal Err: %v %v", e.GetName(), err, jsonReturn)
	} else if jsonResponse.Code != 0 {
		return fmt.Errorf("%s Get Coins Failed: %v", e.GetName(), jsonResponse.Msg)
	}
	if err := json.Unmarshal(jsonResponse.Data, &exchangeData); err != nil {
		return fmt.Errorf("%s Get Coins Result Unmarshal Err: %v %s", e.GetName(), err, jsonResponse.Data)
	}

	for _, data := range exchangeData.Pairs {
		pairStrs := strings.Split(data.Pair, "_")
		p := &pair.Pair{}
		switch e.Source {
		case exchange.EXCHANGE_API:
			base := coin.GetCoin(pairStrs[1])
			target := coin.GetCoin(pairStrs[0])
			if base != nil && target != nil {
				p = pair.GetPair(base, target)
			}
		case exchange.JSON_FILE:
			p = e.GetPairBySymbol(data.Pair)
		}
		if p != nil {
			lotSize := math.Pow10(-1 * data.AmountPrecision)
			priceFilter := math.Pow10(-1 * data.DefaultPrecision)
			pairConstraint := &exchange.PairConstraint{
				PairID:      p.ID,
				Pair:        p,
				ExSymbol:    data.Pair,
				MakerFee:    DEFAULT_MAKER_FEE,
				TakerFee:    DEFAULT_TAKER_FEE,
				LotSize:     lotSize,
				PriceFilter: priceFilter,
				Listed:      true,
			}
			e.SetPairConstraint(pairConstraint)
		}
	}
	return nil
}

/*Get Pair Market Depth
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Add Model of API Response
Step 3: Get Exchange Pair Code ex. symbol := e.GetSymbolByPair(p)
Step 4: Modify API Path(strRequestUrl)
Step 5: Add Params - Depend on API request
Step 6: Convert the response to Standard Maker struct*/
func (e *Bkex) OrderBook(p *pair.Pair) (*exchange.Maker, error) {
	jsonResponse := &JsonResponse{}
	orderBook := OrderBook{}
	symbol := e.GetSymbolByPair(p)

	mapParams := make(map[string]string)
	mapParams["pair"] = symbol

	strRequestPath := "/v1/q/depth"
	strUrl := API_URL + strRequestPath

	maker := &exchange.Maker{
		WorkerIP:        exchange.GetExternalIP(),
		Source:          exchange.EXCHANGE_API,
		BeforeTimestamp: float64(time.Now().UnixNano() / 1e6),
	}

	jsonOrderbook := exchange.HttpGetRequest(strUrl, mapParams)
	if err := json.Unmarshal([]byte(jsonOrderbook), &jsonResponse); err != nil {
		return nil, fmt.Errorf("%s Get Orderbook Json Unmarshal Err: %v %v", e.GetName(), err, jsonOrderbook)
	} else if jsonResponse.Code != 0 {
		return nil, fmt.Errorf("%s Get Orderbook Failed: %v", e.GetName(), jsonResponse.Msg)
	}
	if err := json.Unmarshal(jsonResponse.Data, &orderBook); err != nil {
		return nil, fmt.Errorf("%s Get Orderbook Result Unmarshal Err: %v %s", e.GetName(), err, jsonResponse.Data)
	}

	maker.AfterTimestamp = float64(time.Now().UnixNano() / 1e6)

	var err error
	for _, bid := range orderBook.Bids {
		buydata := exchange.Order{}
		buydata.Quantity = bid.Amt
		buydata.Rate = bid.Price
		maker.Bids = append(maker.Bids, buydata)
	}

	for _, ask := range orderBook.Asks {
		selldata := exchange.Order{}
		selldata.Quantity = ask.Amt
		selldata.Rate = ask.Price
		maker.Asks = append(maker.Asks, selldata)
	}

	return maker, err
}

/*************** Private API ***************/
func (e *Bkex) UpdateAllBalances() {
	if e.API_KEY == "" || e.API_SECRET == "" {
		log.Printf("%s API Key or Secret Key are nil.", e.GetName())
		return
	}

	jsonResponse := &JsonResponse{}
	accountBalance := AccountBalances{}

	strRequestPath := "/v1/u/wallet/balance"

	jsonBalanceReturn := e.ApiKeyGet(strRequestPath, make(map[string]string))
	if err := json.Unmarshal([]byte(jsonBalanceReturn), &jsonResponse); err != nil {
		log.Printf("%s UpdateAllBalances Json Unmarshal Err: %v %v", e.GetName(), err, jsonBalanceReturn)
		return
	} else if jsonResponse.Code != 0 {
		log.Printf("%s UpdateAllBalances Failed: %v", e.GetName(), jsonResponse.Msg)
		return
	}
	if err := json.Unmarshal(jsonResponse.Data, &accountBalance); err != nil {
		log.Printf("%s UpdateAllBalances Result Unmarshal Err: %v %s", e.GetName(), err, jsonResponse.Data)
		return
	}

	for _, balance := range accountBalance {
		c := e.GetCoinBySymbol(balance.CoinType)
		if c != nil {
			balanceMap.Set(c.Code, balance.Available)
		}
	}
}

/* Withdraw(coin *coin.Coin, quantity float64, addr, tag string) */
func (e *Bkex) Withdraw(coin *coin.Coin, quantity float64, addr, tag string) bool {
	if e.API_KEY == "" || e.API_SECRET == "" {
		log.Printf("%s API Key or Secret Key are nil.", e.GetName())
		return false
	}

	jsonResponse := &JsonResponse{}
	withdraw := WithdrawResponse{}
	strRequestPath := "/v1/u/wallet/withdraw"

	mapParams := make(map[string]string)
	mapParams["coinType"] = e.GetSymbolByCoin(coin)
	mapParams["txAddress"] = addr
	mapParams["password"] = e.TradePassword
	mapParams["amount"] = strconv.FormatFloat(quantity, 'f', -1, 64)

	jsonSubmitWithdraw := e.ApiKeyGet(strRequestPath, mapParams)
	if err := json.Unmarshal([]byte(jsonSubmitWithdraw), &jsonResponse); err != nil {
		log.Printf("%s Withdraw Json Unmarshal Err: %v %v", e.GetName(), err, jsonSubmitWithdraw)
		return false
	} else if jsonResponse.Code != 0 {
		log.Printf("%s Withdraw Failed: %v", e.GetName(), jsonResponse.Msg)
		return false
	}
	if err := json.Unmarshal(jsonResponse.Data, &withdraw); err != nil {
		log.Printf("%s Withdraw Result Unmarshal Err: %v %s", e.GetName(), err, jsonResponse.Data)
		return false
	}

	return true
}

func (e *Bkex) LimitSell(pair *pair.Pair, quantity, rate float64) (*exchange.Order, error) {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return nil, fmt.Errorf("%s API Key or Secret Key are nil.", e.GetName())
	}

	jsonResponse := &JsonResponse{}
	strRequestPath := "/v1/u/trade/order/create"

	mapParams := make(map[string]string)
	mapParams["pair"] = e.GetSymbolByPair(pair)
	mapParams["direction"] = "ASK"
	mapParams["price"] = strconv.FormatFloat(rate, 'f', -1, 64)
	mapParams["amount"] = strconv.FormatFloat(quantity, 'f', -1, 64)

	jsonPlaceReturn := e.ApiKeyRequest("POST", strRequestPath, mapParams)
	if err := json.Unmarshal([]byte(jsonPlaceReturn), &jsonResponse); err != nil {
		return nil, fmt.Errorf("%s LimitSell Json Unmarshal Err: %v %v", e.GetName(), err, jsonPlaceReturn)
	} else if jsonResponse.Code != 0 {
		return nil, fmt.Errorf("%s LimitSell Failed: %v", e.GetName(), jsonResponse.Msg)
	}
	orderId := fmt.Sprint("%s", jsonResponse.Data)

	order := &exchange.Order{
		Pair:         pair,
		OrderID:      orderId,
		Rate:         rate,
		Quantity:     quantity,
		Side:         "Sell",
		Status:       exchange.New,
		JsonResponse: jsonPlaceReturn,
	}
	return order, nil
}

func (e *Bkex) LimitBuy(pair *pair.Pair, quantity, rate float64) (*exchange.Order, error) {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return nil, fmt.Errorf("%s API Key or Secret Key are nil.", e.GetName())
	}

	jsonResponse := &JsonResponse{}
	strRequestPath := "/v1/u/trade/order/create"

	mapParams := make(map[string]string)
	mapParams["pair"] = e.GetSymbolByPair(pair)
	mapParams["direction"] = "ASK"
	mapParams["price"] = strconv.FormatFloat(rate, 'f', -1, 64)
	mapParams["amount"] = strconv.FormatFloat(quantity, 'f', -1, 64)

	jsonPlaceReturn := e.ApiKeyRequest("POST", strRequestPath, mapParams)
	if err := json.Unmarshal([]byte(jsonPlaceReturn), &jsonResponse); err != nil {
		return nil, fmt.Errorf("%s LimitBuy Json Unmarshal Err: %v %v", e.GetName(), err, jsonPlaceReturn)
	} else if jsonResponse.Code != 0 {
		return nil, fmt.Errorf("%s LimitBuy Failed: %v", e.GetName(), jsonResponse.Msg)
	}
	orderId := fmt.Sprint("%s", jsonResponse.Data)

	order := &exchange.Order{
		Pair:         pair,
		OrderID:      orderId,
		Rate:         rate,
		Quantity:     quantity,
		Side:         "Buy",
		Status:       exchange.New,
		JsonResponse: jsonPlaceReturn,
	}
	return order, nil
}

func (e *Bkex) OrderStatus(order *exchange.Order) error {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return fmt.Errorf("%s API Key or Secret Key are nil.", e.GetName())
	}

	jsonResponse := &JsonResponse{}
	orderStatus := OrderItem{}
	strRequestPath := "/v1/u/trade/order/unfinished/detail"

	mapParams := make(map[string]string)
	mapParams["pair"] = e.GetSymbolByPair(order.Pair)
	mapParams["orderNo"] = order.OrderID

	jsonOrderStatus := e.ApiKeyGet(strRequestPath, mapParams)
	if err := json.Unmarshal([]byte(jsonOrderStatus), &jsonResponse); err != nil {
		return fmt.Errorf("%s OrderStatus Json Unmarshal Err: %v %v", e.GetName(), err, jsonOrderStatus)
	} else if jsonResponse.Code != 0 {
		return fmt.Errorf("%s OrderStatus Failed: %v", e.GetName(), jsonResponse.Msg)
	}
	if err := json.Unmarshal(jsonResponse.Data, &orderStatus); err != nil {
		return fmt.Errorf("%s OrderStatus Result Unmarshal Err: %v %s", e.GetName(), err, jsonResponse.Data)
	}

	order.DealRate = orderStatus.Price
	order.DealQuantity = orderStatus.Amt

	return nil
}

func (e *Bkex) ListOrders() ([]*exchange.Order, error) {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return nil, fmt.Errorf("%s API Key or Secret Key are nil.\n", e.GetName())
	}

	strRequestPath := "/v1/u/trade/order/listUnfinished"
	mapParams := make(map[string]string)
	var res []*exchange.Order

	jsonResponse := &JsonResponse{}
	orders := OrdersPage{}

	pairs := e.GetPairs()
	if len(pairs) > 0 {
		for _, pairIt := range pairs {
			mapParams["pair"] = e.GetSymbolByPair(pairIt)

			jsonOrders := e.ApiKeyGet(strRequestPath, mapParams)
			if err := json.Unmarshal([]byte(jsonOrders), &jsonResponse); err != nil {
				return nil, fmt.Errorf("%s ListOrders Json Unmarshal Err: %v %v", e.GetName(), err, jsonOrders)
			} else if jsonResponse.Code != 0 {
				return nil, fmt.Errorf("%s ListOrders Failed: %v", e.GetName(), jsonResponse.Msg)
			}
			if err := json.Unmarshal(jsonResponse.Data, &orders); err != nil {
				return nil, fmt.Errorf("%s ListOrders Result Unmarshal Err: %v %s", e.GetName(), err, jsonResponse.Data)
			}

			for _, orderItem := range orders.Data {
				tmpPair := e.GetPairBySymbol(orderItem.Pair)

				order := &exchange.Order{
					Pair:         tmpPair,
					OrderID:      fmt.Sprintf("%d", orderItem.Id),
					Rate:         orderItem.Price,
					Quantity:     orderItem.DealAmount,
					Side:         orderItem.Direction,
					Status:       exchange.New,
					JsonResponse: jsonOrders,
				}
				res = append(res, order)
			}
		}
	}

	return res, nil
}

func (e *Bkex) CancelOrder(order *exchange.Order) error {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return fmt.Errorf("%s API Key or Secret Key are nil.", e.GetName())
	}

	jsonResponse := &JsonResponse{}
	strRequestPath := "/v1/u/trade/order/cancel"

	mapParams := make(map[string]string)
	mapParams["pair"] = e.GetSymbolByPair(order.Pair)
	mapParams["orderNo"] = order.OrderID

	jsonCancelOrder := e.ApiKeyRequest("POST", strRequestPath, mapParams)
	if err := json.Unmarshal([]byte(jsonCancelOrder), &jsonResponse); err != nil {
		return fmt.Errorf("%s CancelOrder Json Unmarshal Err: %v %v", e.GetName(), err, jsonCancelOrder)
	} else if jsonResponse.Code != 0 {
		return fmt.Errorf("%s CancelOrder Failed: %v", e.GetName(), jsonResponse.Msg)
	}

	order.Status = exchange.Canceling
	order.CancelStatus = jsonCancelOrder

	return nil
}

func (e *Bkex) CancelAllOrder() error {
	return nil
}

/*************** Signature Http Request ***************/
/*Method: API Get Request and Signature is required
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Create mapParams Depend on API Signature request
Step 3: Add HttpGetRequest below strUrl if API has different requests*/
func (e *Bkex) ApiKeyGet(strRequestPath string, mapParams map[string]string) string {
	res := e.ApiKeyRequest("GET", strRequestPath, mapParams)
	return res
}

/*Method: API Request and Signature is required
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Create mapParams Depend on API Signature request*/
func (e *Bkex) ApiKeyRequest(strMethod, strRequestPath string, mapParams map[string]string) string {
	strMethod = strings.ToUpper(strMethod)
	queryStr := exchange.Map2UrlQuery(mapParams)
	signature := exchange.ComputeHmac256NoDecode(queryStr, e.API_SECRET)
	strUrl := API_URL + strRequestPath + "?" + queryStr

	request, err := http.NewRequest(strMethod, strUrl, strings.NewReader(queryStr))
	if nil != err {
		return err.Error()
	}
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("X_ACCESS_KEY", e.API_KEY)
	request.Header.Add("X_SIGNATURE", signature)

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}
