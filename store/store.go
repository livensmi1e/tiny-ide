package store

type Store struct {
	driver Driver
}

func New(driver Driver) *Store {
	return &Store{driver: driver}
}

func (s *Store) Migrate() error {
	return s.driver.Migrate()
}
