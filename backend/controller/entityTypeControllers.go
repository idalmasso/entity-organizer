package controller

//CreateEntity creates an entity from the system or returns an error
//
//name: the name of the entity to create
func (s *serviceCtrl) CreateEntity(name string) error{
	_, err:=s.db.CreateEntityType(name)
	return err
}
//RemoveEntity removes an entity from the system or returns an error
//
//name: the name of the entity to remove
func (s *serviceCtrl) RemoveEntity(name string) error{
	err:=s.db.RemoveEntityType(name)
	return err
}
