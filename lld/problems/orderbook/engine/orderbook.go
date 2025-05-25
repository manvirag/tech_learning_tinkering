package engine

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"problem/types"
)

type OrderQueue []*types.Order

func (oq OrderQueue) Len() int { return len(oq) }

func (oq OrderQueue) Less(i, j int) bool {
	if oq[i].Price == oq[j].Price {
		return oq[i].Timestamp.Before(oq[j].Timestamp)
	}
	if oq[i].Type == types.Buy {
		return oq[i].Price > oq[j].Price // Max-heap
	}
	return oq[i].Price < oq[j].Price // Min-heap
}

func (oq OrderQueue) Swap(i, j int) {
	oq[i], oq[j] = oq[j], oq[i]
}

func (oq *OrderQueue) Push(x interface{}) {
	order := x.(*types.Order)
	*oq = append(*oq, order)
}

func (oq *OrderQueue) Pop() interface{} {
	old := *oq
	n := len(old)
	order := old[n-1]
	*oq = old[0 : n-1]
	return order
}

type OrderBook struct {
	buyOrders  OrderQueue
	sellOrders OrderQueue
	orderIDGen int64
	lock       sync.Mutex
}

func NewOrderBook() *OrderBook {
	buy := make(OrderQueue, 0)
	sell := make(OrderQueue, 0)
	heap.Init(&buy)
	heap.Init(&sell)
	return &OrderBook{
		buyOrders:  buy,
		sellOrders: sell,
	}
}

func (ob *OrderBook) SubmitOrder(order *types.Order) {
	ob.lock.Lock()
	defer ob.lock.Unlock()

	order.ID = ob.nextOrderID()
	order.Timestamp = time.Now()

	if order.Type == types.Buy {
		if order.IsMarket {
			ob.matchMarketBuy(order)
		} else {
			ob.matchLimitBuy(order)
		}
	} else {
		if order.IsMarket {
			ob.matchMarketSell(order)
		} else {
			ob.matchLimitSell(order)
		}
	}
}

func (ob *OrderBook) matchLimitBuy(order *types.Order) {
	for order.Quantity > 0 && ob.sellOrders.Len() > 0 {
		bestAsk := ob.sellOrders[0]
		if bestAsk.Price > order.Price {
			break
		}
		ob.execute(order, bestAsk)
		if bestAsk.Quantity == 0 {
			heap.Pop(&ob.sellOrders)
		}
	}
	if order.Quantity > 0 {
		heap.Push(&ob.buyOrders, order)
	}
}

func (ob *OrderBook) matchLimitSell(order *types.Order) {
	for order.Quantity > 0 && ob.buyOrders.Len() > 0 {
		bestBid := ob.buyOrders[0]
		if bestBid.Price < order.Price {
			break
		}
		ob.execute(order, bestBid)
		if bestBid.Quantity == 0 {
			heap.Pop(&ob.buyOrders)
		}
	}
	if order.Quantity > 0 {
		heap.Push(&ob.sellOrders, order)
	}
}

func (ob *OrderBook) matchMarketBuy(order *types.Order) {
	for order.Quantity > 0 && ob.sellOrders.Len() > 0 {
		bestAsk := ob.sellOrders[0]
		ob.execute(order, bestAsk)
		if bestAsk.Quantity == 0 {
			heap.Pop(&ob.sellOrders)
		}
	}
}

func (ob *OrderBook) matchMarketSell(order *types.Order) {
	for order.Quantity > 0 && ob.buyOrders.Len() > 0 {
		bestBid := ob.buyOrders[0]
		ob.execute(order, bestBid)
		if bestBid.Quantity == 0 {
			heap.Pop(&ob.buyOrders)
		}
	}
}

func (ob *OrderBook) execute(maker, taker *types.Order) {
	// Simple trade execution
	tradeQty := min(maker.Quantity, taker.Quantity)
	fmt.Printf("Trade: %s %d @ %.2f (makerID %d, takerID %d)\n",
		maker.Type, tradeQty, taker.Price, maker.ID, taker.ID)
	maker.Quantity -= tradeQty
	taker.Quantity -= tradeQty
}

func (ob *OrderBook) nextOrderID() int64 {
	ob.orderIDGen++
	return ob.orderIDGen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
