package rcombat

import "fmt"

type Player struct {
	id    int
	cards []int
}

func (p *Player) String() string {
	return fmt.Sprintf("{%d:%v}", p.id, p.cards)
}

func (p *Player) GetID() int {
	return p.id
}

func (p *Player) RemoveTop() int {
	top := p.cards[len(p.cards)-1]
	p.cards = p.cards[0 : len(p.cards)-1]
	return top
}

func (p *Player) GetCards() []int {
	return p.cards
}

func (p *Player) GetTopCards(n int) []int {
	cs := make([]int, n, n)
	for i, j := n-1, len(p.cards)-1; i >= 0; i, j = i-1, j-1 {
		cs[i] = p.cards[j]
	}
	return cs
}

func (p *Player) HasCards() bool {
	if len(p.cards) > 0 {
		return true
	}
	return false
}

func (p *Player) NumCards() int {
	return len(p.cards)
}

func (p *Player) AddCard(c int) {
	p.cards = append([]int{c}, p.cards...)
}

func (p *Player) AddCards(cs []int) {
	p.cards = append(cs, p.cards...)
}

func NewPlayer(id int) *Player {
	p := &Player{}
	p.id = id
	p.cards = make([]int, 0, 10)
	return p
}
