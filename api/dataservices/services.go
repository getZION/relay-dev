package dataservices

type Service interface {
	GetAll() (interface{}, error)
	Insert([]byte) error
}
