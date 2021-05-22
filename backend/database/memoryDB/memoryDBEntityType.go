package memoryDB

import (
	"errors"

	"github.com/idalmasso/entity-organizer/backend/common"
)


type entity struct{
	name string
	fields map[string]entityField
}

func (e*entity) GetName() string{
	return e.name
}


func (e*entity) RetrieveAllFields() (map[string]entityField, error){
	return e.fields, nil
}
func (e*entity) AddSingleField(name string, dataType common.DataType) error{
	if _, ok:=e.fields[name]; ok{
		return errors.New("already exists")
	}
	field:=entityField{name: name, dataType: dataType}
	e.fields[name]=field
	return nil
}

func (e*entity) RemoveField(name string) error{
	if _, ok:=e.fields[name]; !ok{
		return errors.New("not exists")
	}
	
	delete(e.fields,name)
	return nil
}



