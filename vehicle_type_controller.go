package main

import (
	"fmt"

	ginhelpers "github.com/Cabista/X/GinHelpers"
	"github.com/gin-gonic/gin"
)

func RegisterVehicleTypeController(group *gin.RouterGroup) {
	group.POST("/", createVehicle)
	group.PUT("/:ID", putVehicle)
	group.GET("/:ID", getVehicleType)
}

func getVehicleType(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)

	id := c.Param("ID")
	if id == "" {
		ginCtx.AbortFailure(400, "id param was not provided as such this request was aborted")
		return
	}

	var vehicleType VehicleType
	err := db.Where("ID = ?", id).First(&vehicleType).Error
	if err != nil {
		ginCtx.AbortFailureErr(404, err)
		return
	}
	// if vehicle == nil {
	// 	fmt.Println("vehicle")
	// 	ginCtx.AbortFailure(404, "a resource by the provided id could not be found")
	// }

	c.JSON(200, vehicleType)
}

func createVehicleType(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)
	var vehicleType VehicleType
	err := c.BindJSON(&vehicleType)

	if err != nil {
		fmt.Println(err.Error())
		ginCtx.AbortFailureErr(500, err)
		return
	}

	err = db.Save(&vehicleType).Error
	if err != nil {
		ginCtx.AbortFailureErr(500, err)
		return
	}

	ginCtx.Created("Resource created", vehicleType.ID)
}

func putVehicleType(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)

	id := c.Param("id")

	if id == "" {
		ginCtx.AbortFailure(400, "id param was not provided as such this request was aborted")
		return
	}

	var vehicleTypeUpdate *VehicleType
	err := c.BindJSON(&vehicleTypeUpdate)
	if err != nil {
		ginCtx.AbortFailureErr(500, err)
		return
	}

	var vehicleType VehicleType
	err = db.Where("ID = ?", id).First(&vehicleType).Error
	if err != nil {
		ginCtx.AbortFailureErr(500, err)
		return
	}

	if vehicleTypeUpdate.Make == "" {
		vehicleType.Make = vehicleTypeUpdate.Make
	}

	if vehicleTypeUpdate.Model == "" {
		vehicleType.Model = vehicleTypeUpdate.Model
	}

	if vehicleTypeUpdate.Color == "" {
		vehicleType.Color = vehicleTypeUpdate.Color
	}

	if vehicleTypeUpdate.YearOfManufacture == 0 {
		vehicleType.YearOfManufacture = vehicleTypeUpdate.YearOfManufacture
	}

	if vehicleTypeUpdate.PassengerCount == 0 {
		vehicleType.PassengerCount = vehicleTypeUpdate.PassengerCount
	}

	if vehicleTypeUpdate.CO2Emissions == 0 {
		vehicleType.CO2Emissions = vehicleTypeUpdate.CO2Emissions
	}

	db.Save(&vehicleType)

	ginCtx.Accepted("The modifcations to the vehicle have been accepted and will all being well be processed shortly")
}
