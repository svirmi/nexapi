package websocket

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/linstohu/nexapi/woox/websocket/types"
)

func (w *WooXWebsocketClient) Subscribe(topics []string) error {
	return w.subscribe(topics)
}

func (w *WooXWebsocketClient) UnSubscribe(topics []string) error {
	return w.unsubscribe(topics)
}

func (w *WooXWebsocketClient) handle(msg *types.SubscribedMessage) error {
	if w.debug {
		w.logger.Info(fmt.Sprintf("subscribed message, topic: %s, timestamp: %v", msg.Topic, msg.Timestamp))
	}

	switch {
	case strings.HasSuffix(msg.Topic, "@orderbook100") ||
		strings.HasSuffix(msg.Topic, "@orderbook"):
		var data types.Orderbook
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.HasSuffix(msg.Topic, "@trade"):
		var data types.Trade
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.HasSuffix(msg.Topic, "@ticker"):
		var data types.Ticker24H
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case msg.Topic == "tickers":
		var data types.Tickers
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.HasSuffix(msg.Topic, "@bbo"):
		var data types.BBO
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case msg.Topic == "bbos":
		var data types.AllBBO
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.Contains(msg.Topic, "@kline_"):
		var data types.Kline
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.HasSuffix(msg.Topic, "@indexprice"):
		var data types.IndexPrice
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.HasSuffix(msg.Topic, "@markprice"):
		var data types.MarkPrice
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case msg.Topic == "markprices":
		var data types.MarkPrices
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.HasSuffix(msg.Topic, "@openinterest"):
		var data types.OpenInterest
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	case strings.HasSuffix(msg.Topic, "@estfundingrate"):
		var data types.EstFundingRate
		err := json.Unmarshal(msg.OriginData, &data)
		if err != nil {
			return err
		}
		w.GetListeners(msg.Topic, &data)
	default:
		return fmt.Errorf("woox unknown message, topic: %s, timestamp: %v", msg.Topic, msg.Timestamp)
	}

	return nil
}
