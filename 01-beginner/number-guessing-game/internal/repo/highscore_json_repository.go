package repo

import "go-backend-labs/01-beginner/number-guessing-game/internal/domain"

type HighScoreRepository interface {
	Load() (domain.HighScores, error)
	Save(domain.HighScores) error
}

type jsonHighScoreRepo struct {
	inner *JSONRepository[domain.HighScores]
}

func NewJSONHighScoreRepository(filename string) HighScoreRepository {
	return &jsonHighScoreRepo{
		inner: NewJSONRepository[domain.HighScores](filename),
	}
}

func (r *jsonHighScoreRepo) Load() (domain.HighScores, error) {
	return r.inner.Load()
}

func (r *jsonHighScoreRepo) Save(data domain.HighScores) error {
	return r.inner.Save(data)
}
