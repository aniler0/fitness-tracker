package app

type ServiceInterface interface {
	GetGreeting() string
}

type Service struct {}

// NewService creates a new instance of Service
func NewService() *Service {
    return &Service{}
}

// GetGreeting returns a greeting message
func (s *Service) GetGreeting() string {
    return "Hello, World!"
}