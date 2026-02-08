package questions

type Service struct {
	Repo *Repository
}

func (s *Service) Create(q *Question) error {
	return s.Repo.Create(q)
}

func (s *Service) List(tag string, difficulty int) ([]Question, error) {
	return s.Repo.List(tag, difficulty)
}

func (s *Service) Get(id int) (*Question, error) {
	return s.Repo.Get(id)
}

func (s *Service) Update(id int, q *Question) error {
	return s.Repo.Update(id, q)
}

func (s *Service) Delete(id int) error {
	return s.Repo.Delete(id)
}

func (s *Service) Random(limit int) ([]Question, error) {
	return s.Repo.Random(limit)
}
