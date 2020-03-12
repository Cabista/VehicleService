package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
)

var db *gorm.DB
var sf *sonyflake.Sonyflake

func init() {

	viper.SetDefault("MachineID", uint16(rand.Int31n(65534)))
	viper.SetDefault("DBDialect", "sqlite3")
	viper.SetDefault("DBConnection", "./service.db")
	viper.SetDefault("HostPort", 8080)

	viper.SetEnvPrefix("srv")
	viper.AutomaticEnv()

	// setup snowflake
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return uint16(viper.GetUint32("MachineID")), nil
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}

	fmt.Println("Using database dialect " + viper.GetString("DBDialect") + " with connection string " + viper.GetString("DBConnection"))
	//open a db connection
	var err error
	db, err = gorm.Open(viper.GetString("DBDialect"), viper.GetString("DBConnection"))
	if err != nil {
		panic(err.Error())
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "vehicleservice_" + defaultTableName
	}

	//Migrate the schema
	db.AutoMigrate(&Vehicle{})
	db.AutoMigrate(&VehicleType{})
}

func main() {
	router := gin.Default()

	vehicleV1 := router.Group("/api/v1/vehicle")
	RegisterVehicleApiController(vehicleV1)
	vehicleV2 := router.Group("/api/v1/vehicletype")
	RegisterVehicleTypeController(vehicleV2)

	router.Run(":" + strconv.FormatUint(uint64(viper.GetUint("HostPort")), 10))
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
