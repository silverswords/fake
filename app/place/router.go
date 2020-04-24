package place

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/place/page"
)

//Router is to run city router
func Router() {
	router := gin.Default()

	page.Place(router)
	page.Country(router)
	page.City(router)
	page.Coord(router)
	page.ZipCode(router)

	router.Run(":8004")
}
