package controller

import (
	"github.com/idalmasso/entity-organizer/backend/controller/internal/interfaces"
)

type serviceCtrl struct {
	db interfaces.DatabaseEntityTypeInterface
	name string
}

//GetName returns the actual name of the service
func (s *serviceCtrl) GetName() string{
	return s.name
}
//Init initialize the service controller
func (s *serviceCtrl) Init() error{
	err:=s.db.Init()
	return err
}

//NewServiceController returns a new service controller by name 
func NewServiceController(name string, db interfaces.DatabaseEntityTypeInterface) (*serviceCtrl,error) {
	srv:=serviceCtrl{db: db}
	srv.Init()
	return &srv, nil
}

