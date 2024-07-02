package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Mediator interface
type Mediator interface {
	bid(bidder *Bidder, amount int)
}

// AuctionMediator is the concrete mediator
type AuctionMediator struct {
	bidders []*Bidder
	mutex   sync.Mutex
}

func (m *AuctionMediator) addBidder(bidder *Bidder) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.bidders = append(m.bidders, bidder)
}

func (m *AuctionMediator) bid(bidder *Bidder, amount int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, b := range m.bidders {
		if b != bidder {
			b.receiveBid(amount)
		}
	}
}

// Bidder represents a participant in the auction
type Bidder struct {
	name     string
	mediator Mediator
}

func (b *Bidder) placeBid(amount int) {
	fmt.Printf("%s placed a bid of $%d\n", b.name, amount)
	b.mediator.bid(b, amount)
}

func (b *Bidder) receiveBid(amount int) {
	fmt.Printf("%s received a bid of $%d\n", b.name, amount)
}

func main() {
	mediator := &AuctionMediator{}
	bidder1 := &Bidder{name: "Bidder 1", mediator: mediator}
	bidder2 := &Bidder{name: "Bidder 2", mediator: mediator}

	mediator.addBidder(bidder1)
	mediator.addBidder(bidder2)

	rand.Seed(time.Now().UnixNano())

	// Simulate bidding process
	for i := 0; i < 5; i++ {
		amount := rand.Intn(100)
		bidder1.placeBid(amount)
		time.Sleep(time.Second)
		amount = rand.Intn(100)
		bidder2.placeBid(amount)
		time.Sleep(time.Second)
	}
}
