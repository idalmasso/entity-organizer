package interfaces

type DatabaseEntityTypeInterface interface{
	CreateEntityType(string) (EntityTypeInterface,error)
	RemoveEntityType(string) error
	FindEntityType(string) (EntityTypeInterface,error)
	Init() error
}
