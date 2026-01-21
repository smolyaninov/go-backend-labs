package repo

type BudgetRepository interface {
	Load() (map[string]float64, error)
	Save(map[string]float64) error
}

type jsonBudgetRepo struct {
	inner *JSONRepository[map[string]float64]
}

func NewJSONBudgetRepository(filename string) BudgetRepository {
	return &jsonBudgetRepo{
		inner: NewJSONRepository[map[string]float64](filename),
	}
}

func (r *jsonBudgetRepo) Load() (map[string]float64, error) {
	return r.inner.Load()
}

func (r *jsonBudgetRepo) Save(data map[string]float64) error {
	return r.inner.Save(data)
}
