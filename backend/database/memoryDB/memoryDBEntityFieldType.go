package memoryDB

import "github.com/idalmasso/entity-organizer/backend/common"

type entityField struct {
	name string
	dataType common.DataType

}

func (e *entityField)	GetName() string{
	return e.name
}
func (e *entityField)GetDataType() (common.DataType , error){
	return e.dataType,nil
}
func (e *entityField)SetDataType(dataType common.DataType) error{
	e.dataType = dataType
	return nil
}

