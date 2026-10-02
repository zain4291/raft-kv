package main

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: map[string]string{},
	}
}
func (s *Store) Get(key string) (string, bool) {
	value, ok := s.data[key]
	return value, ok

}

func (s *Store) Put(key, value string) {

}

func (s *Store) Delete(key string) {

}
