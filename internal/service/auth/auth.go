package auth

type authRepository interface{}

type Service struct {
	repo authRepository
}

func NewService(repo authRepository) *Service {
	return &Service{repo: repo}
}
