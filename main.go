package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sony/sonyflake"
)

var db *gorm.DB
var sf *sonyflake.Sonyflake

func init() {
	// setup snowflake
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return 1, nil
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
	//open a db connection
	var err error
	db, err = gorm.Open("sqlite3", "./vehicle.db")
	if err != nil {
		panic("failed to connect database")
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "vehicleservice_" + defaultTableName
	}

	//Migrate the schema
	db.AutoMigrate(&VehicleType{}, &Vehicle{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/vehicle")
	RegisterVehicleApiController(v1)

	router.Run()
}

/*
func createVehicle(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	var vehicleForm VehicleForm
	err := c.BindJSON(&vehicleForm)
	if err != nil {
		c.AbortWithError(500, err)
	}

	db.Save(vehicleForm)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Vehicle created", "resourceId": vehicleForm})
}
*/
