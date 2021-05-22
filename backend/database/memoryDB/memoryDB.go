package memoryDB

import "errors"

type memDB struct{
	entities map[string]entity
}

func (m *memDB)Init() error{
	m.entities=map[string]entity{}
	return nil
}

func (m *memDB) CreateEntityType(name string) (*entity,error){
	if _, ok:=m.entities[name]; ok{
		return nil,errors.New("already exists")
	}
	e:=entity{name: name}
	m.entities[name] = e
	return &e,nil
}
func (m *memDB) RemoveEntityType(name string) error {
	if _, ok:=m.entities[name]; !ok{
		return errors.New("not exists")
	}
	delete(m.entities, name)
	return nil
}

func (m *memDB) FindEntityType(name string) (*entity,error){
	e, ok:=m.entities[name]
	if !ok{
		return nil,errors.New("not exists")
	}
	return &e, nil
}
	

func NewDB() *memDB{
	db:=memDB{}
	return &db
}
