package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
}

func NewRouter(r *gin.Engine) Router {
	return &router{r: r}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	// Example
	//db := database.NewDatabase[domain.Seller](getDatabaseFile("seller"))
	// repo := seller.NewRepository(db)
	// service := seller.NewService(repo)
	// handler := handler.NewSeller(service)
	// r.r.GET("/seller", handler.GetAll)
}

func (r *router) buildProductRoutes() {}

func (r *router) buildSectionRoutes() {}

func (r *router) buildWarehouseRoutes() {}

func (r *router) buildEmployeeRoutes() {}

func (r *router) buildBuyerRoutes() {}

// DO NOT MODIFY
func getDatabaseFile(entity string) string {
	return fmt.Sprintf("./databases/%s.json", entity)
}
