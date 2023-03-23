package account

type Repository interface {
	Save(ac *Account) error
	FindById(accountId int64) (*Account, error)
}
