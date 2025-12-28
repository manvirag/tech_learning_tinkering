package models

type Piece interface {
	GetColor() string
	GetType() string
}

type Bishop struct {
	Color string
}

func (b Bishop) GetColor() string {
	return b.Color
}

func (b Bishop) GetType() string {
	return "Bishop"
}

type Rook struct {
	Color string
}

func (r Rook) GetColor() string {
	return r.Color
}

func (r Rook) GetType() string {
	return "Rook"
}

type Knight struct {
	Color string
}

func (k Knight) GetColor() string {
	return k.Color
}

func (k Knight) GetType() string {
	return "Knight"
}

type Queen struct {
	Color string
}

func (q Queen) GetColor() string {
	return q.Color
}

func (q Queen) GetType() string {
	return "Queen"
}

type King struct {
	Color string
}

func (k King) GetColor() string {
	return k.Color
}

func (k King) GetType() string {
	return "King"
}

type Pawn struct {
	Color string
}

func (p Pawn) GetColor() string {
	return p.Color
}

func (p Pawn) GetType() string {
	return "Pawn"
}