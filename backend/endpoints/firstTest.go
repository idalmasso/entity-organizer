package endpoints

import (
	"github.com/idalmasso/entity-organizer/backend/controller"
	"github.com/idalmasso/entity-organizer/backend/database/memoryDB"
)





func TestEndpoint() {
	var newDB = memoryDB.NewDB()
	controller.NewServiceController("Test", newDB)
}

