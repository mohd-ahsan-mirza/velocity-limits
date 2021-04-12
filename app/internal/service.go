package internal

// Service interface
type Service interface {
	LoadFunds(string) (bool, []byte, error)
}
