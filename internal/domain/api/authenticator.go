package api

type IAuthServicePort interface {
	PasswordEncoder(password string) (string, error)
}
