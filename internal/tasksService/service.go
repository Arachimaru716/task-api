package tasksService

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateTask(task Task) (Task, error) {
	return s.repo.Create(task)
}

func (s *Service) UpdateTask(task Task) (Task, error) {
	return s.repo.Update(task)
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}