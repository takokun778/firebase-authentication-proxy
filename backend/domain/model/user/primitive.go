package user

type Primitive struct {
	id ID
}

func (p Primitive) ID() ID {
	return p.id
}

type Values struct {
	id string
}

func NewPrimitive(id ID) Primitive {
	return Primitive{
		id: id,
	}
}

func (p Primitive) Values() Values {
	return Values{
		id: p.id.String(),
	}
}
