package cache

type Cache interface {
	Get(key interface{}) (interface{},error)
	Insert(key ,data interface{}) (interface{},error)
	Update(key , data interface{}) error
	Remove(key interface{}) error
}

type DefaultCache struct {
	getFunc func(key interface{}) (interface{},error)
	insertFunc func(key ,data interface{}) (interface{},error)
	updateFunc func(key , data interface{}) error
	removeFunc func(key interface{}) error
}

func (d *DefaultCache) Get(key interface{}) (interface{}, error) {
	if d.getFunc == nil {
		panic("please implement Get !")
	}
	return d.getFunc(key)
}

func (d *DefaultCache) Insert(key, data interface{}) (interface{}, error) {
	if d.insertFunc == nil {
		panic("please implement Insert !")
	}
	return d.insertFunc(key,data)
}

func (d *DefaultCache) Update(key, data interface{}) error {
	if d.updateFunc == nil {
		panic("please implement Update !")
	}
	return d.updateFunc(key,data)
}

func (d *DefaultCache) Remove(key interface{}) error {
	if d.removeFunc == nil {
		panic("please implement Remove !")
	}
	return d.removeFunc(key)
}



type Builder interface {
	GetFunc(getFunc func(key interface{}) (interface{},error)) Builder
	InsertFunc(insertFunc func(key ,data interface{}) (interface{},error)) Builder
	UpdateFunc(updateFunc func(key , data interface{}) error) Builder
	RemoveFunc(removeFunc func(key interface{}) error) Builder
	Build() Cache
}

