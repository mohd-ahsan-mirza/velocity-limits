package internal

// Service interface
type Service interface {
	LoadFunds(string) ([]byte, error)
}
