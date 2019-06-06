package initial

import (
	"sync"

	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/exchange/bibox"
	"github.com/bitontop/gored/exchange/bigone"
	"github.com/bitontop/gored/exchange/biki"
	"github.com/bitontop/gored/exchange/binance"
	"github.com/bitontop/gored/exchange/binancedex"
	"github.com/bitontop/gored/exchange/bitfinex"
	"github.com/bitontop/gored/exchange/bitforex"
	"github.com/bitontop/gored/exchange/bitmart"
	"github.com/bitontop/gored/exchange/bitmax"
	"github.com/bitontop/gored/exchange/bitmex"
	"github.com/bitontop/gored/exchange/bitrue"
	"github.com/bitontop/gored/exchange/bitstamp"
	"github.com/bitontop/gored/exchange/bittrex"
	"github.com/bitontop/gored/exchange/bitz"
	"github.com/bitontop/gored/exchange/coinbene"
	"github.com/bitontop/gored/exchange/coineal"
	"github.com/bitontop/gored/exchange/coinex"
	"github.com/bitontop/gored/exchange/cointiger"
	"github.com/bitontop/gored/exchange/dcoin"
	"github.com/bitontop/gored/exchange/dragonex"
	"github.com/bitontop/gored/exchange/gateio"
	"github.com/bitontop/gored/exchange/gemini"
	"github.com/bitontop/gored/exchange/hitbtc"
	"github.com/bitontop/gored/exchange/huobi"
	"github.com/bitontop/gored/exchange/huobiotc"
	"github.com/bitontop/gored/exchange/ibankdigital"
	"github.com/bitontop/gored/exchange/idex"
	"github.com/bitontop/gored/exchange/kraken"
	"github.com/bitontop/gored/exchange/kucoin"
	"github.com/bitontop/gored/exchange/lbank"
	"github.com/bitontop/gored/exchange/liquid"
	"github.com/bitontop/gored/exchange/mxc"
	"github.com/bitontop/gored/exchange/okex"
	"github.com/bitontop/gored/exchange/otcbtc"
	"github.com/bitontop/gored/exchange/poloniex"
	"github.com/bitontop/gored/exchange/stex"
	"github.com/bitontop/gored/exchange/tokok"
	"github.com/bitontop/gored/exchange/tradeogre"
	"github.com/bitontop/gored/exchange/tradesatoshi"
)

var instance *InitManager
var once sync.Once

type InitManager struct {
	exMan *exchange.ExchangeManager
}

func CreateInitManager() *InitManager {
	once.Do(func() {
		instance = &InitManager{
			exMan: exchange.CreateExchangeManager(),
		}
	})
	return instance
}

func (e *InitManager) Init(config *exchange.Config) exchange.Exchange {
	switch config.ExName {
	case exchange.BINANCE:
		ex := binance.CreateBinance(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITTREX:
		ex := bittrex.CreateBittrex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINEX:
		ex := coinex.CreateCoinex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.STEX:
		ex := stex.CreateStex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITMEX:
		ex := bitmex.CreateBitmex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.KUCOIN:
		ex := kucoin.CreateKucoin(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HUOBIOTC:
		ex := huobiotc.CreateHuobiOTC(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITMAX:
		ex := bitmax.CreateBitmax(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITSTAMP:
		ex := bitstamp.CreateBitstamp(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.OTCBTC:
		ex := otcbtc.CreateOtcbtc(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HUOBI:
		ex := huobi.CreateHuobi(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BIBOX:
		ex := bibox.CreateBibox(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.OKEX:
		ex := okex.CreateOkex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITZ:
		ex := bitz.CreateBitz(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HITBTC:
		ex := hitbtc.CreateHitbtc(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.DRAGONEX:
		ex := dragonex.CreateDragonex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BIGONE:
		ex := bigone.CreateBigone(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITFINEX:
		ex := bitfinex.CreateBitfinex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.GATEIO:
		ex := gateio.CreateGateio(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.IDEX:
		ex := idex.CreateIdex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.LIQUID:
		ex := liquid.CreateLiquid(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITFOREX:
		ex := bitforex.CreateBitforex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.TOKOK:
		ex := tokok.CreateTokok(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.MXC:
		ex := mxc.CreateMxc(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITRUE:
		ex := bitrue.CreateBitrue(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.TRADESATOSHI:
		ex := tradesatoshi.CreateTradeSatoshi(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.KRAKEN:
		ex := kraken.CreateKraken(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.POLONIEX:
		ex := poloniex.CreatePoloniex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINEAL:
		ex := coineal.CreateCoineal(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.TRADEOGRE:
		ex := tradeogre.CreateTradeogre(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINBENE:
		ex := coinbene.CreateCoinbene(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.IBANKDIGITAL:
		ex := ibankdigital.CreateIbankdigital(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.LBANK:
		ex := lbank.CreateLbank(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BINANCEDEX:
		ex := binancedex.CreateBinanceDex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITMART:
		ex := bitmart.CreateBitmart(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BIKI:
		ex := biki.CreateBiki(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.DCOIN:
		ex := dcoin.CreateDcoin(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINTIGER:
		ex := cointiger.CreateCointiger(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.GEMINI:
		ex := gemini.CreateGemini(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	}
	return nil
}
