package firebase

type RefreshToken string

func NewRefreshToken(value string) (RefreshToken, error) {
	v := RefreshToken(value)
	if err := v.validate(); err != nil {
		return RefreshToken(""), err
	}

	return v, nil
}

func (t RefreshToken) Value() string {
	return string(t)
}

func (t RefreshToken) validate() error {
	return nil
}
