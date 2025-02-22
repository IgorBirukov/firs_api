package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}

func (s *TaskService) UpdateTaskByID(id string, task interface{}) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id string) (res int, err error) {
	return s.repo.DeleteTaskByID(id)
}
