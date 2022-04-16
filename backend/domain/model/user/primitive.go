package user

type Primitive struct {
	id Id
}

func (p Primitive) Id() Id {
	return p.id
}

type Values struct {
	id string
}

func NewPrimitive(id Id) Primitive {
	return Primitive{
		id: id,
	}
}

func (p Primitive) Values() Values {
	return Values{
		id: p.id.String(),
	}
}
