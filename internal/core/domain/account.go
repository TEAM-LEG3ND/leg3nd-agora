package domain

type Account struct {
	Id            int64
	Email         string
	Nickname      *string
	FullName      string
	OAuthProvider OAuthProvider
	Status        Status
}

type Status string

const (
	Draft     = Status("draft")
	Ok        = Status("ok")
	Suspended = Status("suspended")
	Withdraw  = Status("withdraw")
)

func (s Status) String() string {
	return string(s)
}

type OAuthProvider string

const (
	Google = OAuthProvider("google")
	GitHub = OAuthProvider("github")
)

func (oap OAuthProvider) String() string {
	return string(oap)
}

func NewAccount(email string, fullName string, oAuthProvider OAuthProvider) *Account {
	return &Account{
		Email:         email,
		FullName:      fullName,
		OAuthProvider: oAuthProvider,
	}
}
