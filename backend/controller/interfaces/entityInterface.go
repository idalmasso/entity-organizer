package interfaces

import "github.com/idalmasso/entity-organizer/backend/common"

type EntityTypeInterface interface{
	GetName() string
	RetrieveAllFields() (map[string]EntityFieldTypeInterface, error)
	AddSingleField(string, common.DataType) error
	//AddMultipleFields([]EntityFieldTypeInterface) error
	RemoveField(string) error
}
