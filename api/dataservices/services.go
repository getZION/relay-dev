package dataservices

//todo: do we need that anymore?
type Service interface {
	GetAll() (interface{}, error)
	Insert(interface{}) (interface{}, error)
}
