package interfaces

import (
	"github.com/idalmasso/entity-organizer/backend/common"
)

type EntityFieldTypeInterface interface{
	GetName() string
	GetDataType() (common.DataType , error)
	SetDataType(common.DataType) error
}

