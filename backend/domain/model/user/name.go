package user

type Name string

func NewName(value string) (Name, error) {
	v := Name(value)
	if err := v.validate(); err != nil {
		return Name(""), err
	}
	return Name(value), nil
}

func (n Name) Value() string {
	return string(n)
}

func (n Name) validate() error {
	return nil
}
