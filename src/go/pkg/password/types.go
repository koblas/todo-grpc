package password

type Hasher interface {
	Hash(plain string) (string, error)
}

type Verifier interface {
	Verify(plain, hash string) (bool, error)
}
