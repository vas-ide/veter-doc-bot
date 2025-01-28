package product



type Service struct {}


func NewService() *Service{
	return &Service{}
}

func (s *Service) lstService() []Product {
	return allProducts
}
