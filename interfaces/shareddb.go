package interfaces

type SharedDB interface {
	Get(id string) (interface{}, error)
	Set(id string, data interface{}) error
}
