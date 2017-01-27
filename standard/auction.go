package standard

import (
	"errors"

	"github.com/hinshun/powergrid"
)

var (
	ErrPlayersNotSufficient = errors.New("not enough players to start an auction")
	ErrBidNotSufficient     = errors.New("bid must be at least higher than highest bid")
	ErrAuctionEnded         = errors.New("auction has already ended")
)

type auction struct {
	powerPlant powergrid.PowerPlant
	players    []powergrid.Player
	highestBid powergrid.Elektro
}

var _ = powergrid.Auction((*auction)(nil))

func NewStandardsAuction(powerPlant powergrid.PowerPlant, players []powergrid.Player) (powergrid.Auction, error) {
	if len(players) < 2 {
		return nil, ErrPlayersNotSufficient
	}

	return &auction{
		powerPlant: powerPlant,
		players:    players,
	}, nil
}

func (a *auction) Item() powergrid.PowerPlant {
	return a.powerPlant
}

func (a *auction) Bidders() []powergrid.Player {
	return a.players
}

func (a *auction) CurrentBidder() powergrid.Player {
	return a.players[0]
}

func (a *auction) HighestBid() powergrid.Elektro {
	return a.highestBid
}

func (a *auction) Bid(bid powergrid.Elektro) error {
	if bid <= a.HighestBid() {
		return ErrBidNotSufficient
	}

	a.highestBid = bid
	a.players = append(a.players[1:], a.players[0])
	return nil
}

func (a *auction) Pass() (bool, error) {
	if len(a.players) == 1 {
		return false, ErrAuctionEnded
	}

	a.players = a.players[1:]
	return len(a.players) == 1, nil
}
