package model

type Account struct {
	Username string
	Email    string
	Password string
}

func NewAccount() Account {
	return Account{}
}

func (account *Account) Register() {

}

func (account *Account) Login() {

}

func (account *Account) Logout() {

}
