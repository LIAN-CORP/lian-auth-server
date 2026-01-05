package api

type ArgonHasher interface {
	GenerateCryptographicSalt(password string) (string, error)
	HashPasswordSecure(password string) (string, error)
}
