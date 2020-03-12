package main

import (
	"fmt"

	ginhelpers "github.com/Cabista/X/GinHelpers"
	"github.com/gin-gonic/gin"
)

func RegisterVehicleApiController(group *gin.RouterGroup) {
	group.POST("/", createVehicle)
	group.PUT("/:ID", putVehicle)
	group.GET("/:ID", getVehicle)
}

func getVehicle(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)

	id := c.Param("ID")
	if id == "" {
		ginCtx.AbortFailure(400, "id param was not provided as such this request was aborted")
		return
	}

	var vehicle Vehicle
	err := db.Where("ID = ?", id).First(&vehicle).Error
	if err != nil {
		ginCtx.AbortFailureErr(404, err)
		return
	}
	// if vehicle == nil {
	// 	fmt.Println("vehicle")
	// 	ginCtx.AbortFailure(404, "a resource by the provided id could not be found")
	// }

	c.JSON(200, vehicle)

}

func createVehicle(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)
	var vehicle Vehicle
	err := c.BindJSON(&vehicle)

	if err != nil {
		fmt.Println(err.Error())
		ginCtx.AbortFailureErr(500, err)
		return
	}

	db.Save(&vehicle)
	ginCtx.Created("Resource created", vehicle.ID)
}

func putVehicle(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)

	id := c.Param("id")

	if id == "" {
		ginCtx.AbortFailure(400, "id param was not provided as such this request was aborted")
		return
	}

	var vehicleUpdate *Vehicle
	err := c.BindJSON(&vehicleUpdate)
	if err != nil {
		ginCtx.AbortFailureErr(500, err)
		return
	}

	var vehicle *Vehicle
	db.Where("ID = ?", id).First(&vehicle)

	if vehicle == nil {
		ginCtx.AbortFailure(404, "a resource by the provided id could not be found")
	}

	if vehicleUpdate.DriverID != 0 {
		vehicle.DriverID = vehicleUpdate.DriverID
	}

	if vehicleUpdate.IsActive != vehicle.IsActive {
		vehicle.IsActive = vehicleUpdate.IsActive
	}

	if vehicleUpdate.MOTExpiry != nil {
		vehicle.MOTExpiry = vehicleUpdate.MOTExpiry
	}

	if vehicleUpdate.PlateExpiry != nil {
		vehicle.PlateExpiry = vehicleUpdate.PlateExpiry
	}

	if vehicleUpdate.InsuranceExpiry != nil {
		vehicle.InsuranceExpiry = vehicleUpdate.InsuranceExpiry
	}

	if vehicleUpdate.RoadTaxExpiry != nil {
		vehicle.RoadTaxExpiry = vehicleUpdate.RoadTaxExpiry
	}

	if vehicleUpdate.Registration != "" {
		vehicle.Registration = vehicleUpdate.Registration
	}

	if vehicleUpdate.VehicleTypeID != 0 {
		vehicle.VehicleTypeID = vehicleUpdate.VehicleTypeID
	}

	db.Save(&vehicle)

	ginCtx.Accepted("The modifcations to the vehicle have been accepted and will all being well be processed shortly")
}

//method to validate link to vehicle type
