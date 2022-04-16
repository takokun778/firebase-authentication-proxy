package firebase

type Primitive struct {
	password Password
}

type Values struct {
	password string
}

func NewPrimitive(password Password) Primitive {
	return Primitive{
		password: password,
	}
}

func (p Primitive) Values() Values {
	return Values{
		password: p.password.Value(),
	}
}
