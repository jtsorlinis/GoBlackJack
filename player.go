package main

import (
	"fmt"
)

var playerNumCount int32
var maxSplits int32 = 10

// Player class
type Player struct {
	MPlayerNum  string
	MHand       []*Card
	MValue      int32
	MEarnings   float32
	MAces       int32
	MIsSoft     bool
	MSplitCount int32
	MIsDone     bool
	MSplitFrom  *Player
	MBetMult    float32
	MHasNatural bool
	MTable      *Table
	MInitialBet int32
}

// NewPlayer constructor
func NewPlayer(table *Table, split *Player) *Player {
	p := new(Player)
	p.MTable = table
	p.MInitialBet = p.MTable.MBetSize
	p.MBetMult = 1
	p.MHand = make([]*Card, 0, 5)
	if split != nil {
		p.MHand = append(p.MHand, split.MHand[1])
		p.MSplitCount++
		p.MPlayerNum = split.MPlayerNum + "S"
		p.MInitialBet = split.MInitialBet
		p.MSplitFrom = split
	} else {
		playerNumCount++
		p.MPlayerNum = fmt.Sprint(playerNumCount)
	}
	return p
}

// DoubleBet doubles the player's bet
func (p *Player) DoubleBet() {
	p.MBetMult = 2
}

// ResetHand resets the player's hand
func (p *Player) ResetHand() {
	p.MHand = p.MHand[:0]
	p.MValue = 0
	p.MAces = 0
	p.MIsSoft = false
	p.MSplitCount = 0
	p.MIsDone = false
	p.MBetMult = 1
	p.MHasNatural = false
	p.MInitialBet = p.MTable.MBetSize
}

// CanSplit checks if the player can split
func (p *Player) CanSplit() int32 {
	if len(p.MHand) == 2 && p.MHand[0].MRank[0] == p.MHand[1].MRank[0] && p.MSplitCount < maxSplits {
		return p.MHand[0].MValue
	}
	return 0
}

// Win increases player earnings
func (p *Player) Win(mult float32) {
	x := float32(p.MInitialBet) * p.MBetMult * mult
	p.MEarnings += x
	p.MTable.MCasinoEarnings -= x

}

// Lose decreases player earnings
func (p *Player) Lose() {
	x := float32(p.MInitialBet) * p.MBetMult
	p.MEarnings -= x
	p.MTable.MCasinoEarnings += x
}

// Print prints the players number and hand
func (p *Player) Print() string {
	output := "Player " + fmt.Sprint(p.MPlayerNum) + ": "
	for _, card := range p.MHand {
		output += card.Print() + " "
	}
	for i := len(p.MHand); i < 5; i++ {
		output += "  "
	}
	output += "\tScore: " + fmt.Sprint(p.MValue)
	if p.MValue > 21 {
		output += " (Bust) "
	} else {
		output += "       "
	}
	output += "\tBet: "
	output += fmt.Sprint(float32(p.MInitialBet) * p.MBetMult)
	return output
}

// Evaluate evaluates the player's hand
func (p *Player) Evaluate() {
	p.MAces = 0
	p.MValue = 0
	for _, card := range p.MHand {
		p.MValue += card.MValue
		// check for ace
		if card.MIsAce {
			p.MAces++
			p.MIsSoft = true
		}
	}

	for p.MValue > 21 && p.MAces > 0 {
		p.MValue -= 10
		p.MAces--
	}

	if p.MAces == 0 {
		p.MIsSoft = false
	}
}
