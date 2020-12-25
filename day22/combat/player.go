package combat

type Player struct {
	id    int
	cards []int
}

func (p *Player) RemoveTop() int {
	top := p.cards[len(p.cards)-1]
	p.cards = p.cards[0 : len(p.cards)-1]
	return top
}

func (p *Player) GetCards() []int {
	return p.cards
}

func (p *Player) HasCards() bool {
	if len(p.cards) > 0 {
		return true
	}
	return false
}

func (p *Player) AddCard(c int) {
	p.cards = append([]int{c}, p.cards...)
}

func (p *Player) AddCards(cs []int) {
	p.cards = append(cs, p.cards...)
}

func NewPlayer(id int) *Player {
	p := &Player{}
	p.cards = make([]int, 0, 10)
	return p
}
