package etData

type exchangeStorage interface {
	Add(key string, value interface{}) error
	Get() (interface{}, error)
	Delete(key string) error
}

func  loadData()