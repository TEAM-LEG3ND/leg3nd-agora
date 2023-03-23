package account

type Account struct {
	Email         string
	Nickname      string
	FullName      string
	OAuthProvider OAuthProvider
}

type OAuthProvider string

const (
	Google = OAuthProvider("google")
	GitHub = OAuthProvider("github")
)

func NewAccount(email string, nickname string, fullName string, oAuthProvider OAuthProvider) *Account {
	return &Account{
		Email:         email,
		Nickname:      nickname,
		FullName:      fullName,
		OAuthProvider: oAuthProvider,
	}
}
