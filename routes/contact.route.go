package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hudarashid/golang_rest/controllers"
)

type ContactRoutes struct {
	contactController controllers.ContactController
}

func NewRouteContact(contactController controllers.ContactController) ContactRoutes {
	return ContactRoutes{contactController}
}

func (cr *ContactRoutes) ContactRoute(rg *gin.RouterGroup) {

	router := rg.Group("contacts")
	router.POST("/", cr.contactController.CreateContact)
	router.GET("/", cr.contactController.GetAllContacts)
	router.PATCH("/:contactId", cr.contactController.UpdateContact)
	router.GET("/:contactId", cr.contactController.GetContactById)
	router.DELETE("/:contactId", cr.contactController.DeleteContactById)
}
