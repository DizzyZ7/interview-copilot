package quiz

import (
	"errors"
	"sync"

	"interview-copilot/backend/internal/questions"
)

type Service struct {
	Questions *questions.Service
	sessions  map[int]*QuizSession
	mu        sync.Mutex
}

func NewService(q *questions.Service) *Service {
	return &Service{
		Questions: q,
		sessions:  make(map[int]*QuizSession),
	}
}

func (s *Service) Start(userID int, limit int) ([]questions.Question, error) {
	qs, err := s.Questions.Random(limit)
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(qs))
	for i, q := range qs {
		ids[i] = q.ID
	}

	s.mu.Lock()
	s.sessions[userID] = &QuizSession{
		UserID:    userID,
		Questions: ids,
		Current:   0,
	}
	s.mu.Unlock()

	return qs, nil
}

func (s *Service) Answer(userID int, correct bool) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, ok := s.sessions[userID]
	if !ok {
		return 0, errors.New("no active session")
	}

	session.Current++
	if session.Current >= len(session.Questions) {
		delete(s.sessions, userID)
		return 0, nil
	}
	return session.Questions[session.Current], nil
}
