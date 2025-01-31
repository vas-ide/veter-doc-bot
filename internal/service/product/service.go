package product



type Service struct {}


func NewService() *Service{
	return &Service{}
}

func (s *Service) LstService() []Product {
	return allProducts
}

func (s Service) GetProducts (ind int) (*Product, error) {
	return &allProducts[ind], nil
}
