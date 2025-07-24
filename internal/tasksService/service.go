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

func (s *Service) CreateTask(text string, isDone bool, userID uint) (Task, error) {
	t := Task{
		Text:   text,
		IsDone: isDone,
		UserID: &userID,
	}
	return s.repo.Create(t)
}

func (s *Service) UpdateTask(task Task) (Task, error) {
	return s.repo.Update(task)
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
