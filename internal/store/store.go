package store

type Store struct {
	FilePath string
}

func NewStore(filePath string) *Store {
	return &Store{FilePath: filePath}
}
