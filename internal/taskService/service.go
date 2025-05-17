package taskService

type TaskService struct {
	repo TaskRepository
}

func (s *TaskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

/*
func (s *TaskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}
*/

func (s *TaskService) UpdateTaskByID(id uint, task interface{}) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) (res int, err error) {
	return s.repo.DeleteTaskByID(id)
}

func (s *TaskService) GetTasksByUserID(userId uint) (res []Task, err error) {
	return s.repo.GetTasksByUserID(userId)
}

func (s *TaskService) PostTask(task Task) (Task, error) {
	return s.repo.PostTask(task)
}
