package auction

import "errors"

var (
	ErrPlayersNotSufficient = errors.New("not enough players to start an auction")
	ErrBidNotSufficient     = errors.New("bid must be at least higher than highest bid")
	ErrAuctionEnded         = errors.New("auction has already ended")
)

type standardAuction struct {
	powerPlant powergrid.PowerPlant
	players    []powergrid.Player
	highestBid powergrid.Elektro
}

var _ = powergrid.Auction((*standardAuction)(nil))

func NewStandardAuction(powerPlant powergrid.PowerPlant, players []powergrid.Player) (Auction, error) {
	if len(players) < 2 {
		return nil, ErrPlayersNotSufficient
	}

	return &standardAuction{
		powerPlant: powerPlant,
		players:    players,
	}, nil
}

func (sa *standardAuction) Item() powergrid.PowerPlant {
	return sa.powerPlant
}

func (sa *standardAuction) Bidders() []powergrid.Player {
	return sa.players
}

func (sa *standardAuction) CurrentBidder() powergrid.Player {
	return players[0]
}

func (sa *standardAuction) HighestBid() powergrid.Bid {
	return sa.highestBid
}

func (sa *standardAuction) Bid(bid powergrid.Elektro) error {
	if bid <= sa.HighestBid() {
		return ErrBidNotSufficient
	}

	sa.highestBid = bid
	sa.players = append(sa.players[1:], sa.players[0])
	return nil
}

func (sa *standardAuction) Pass() (bool, error) {
	if len(sa.players) == 1 {
		return false, ErrAuctionEnded
	}

	sa.players = sa.players[1:]
	return len(sa.players) == 1, nil
}
