package util

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var CONFIG Config

type Config struct {
	CONNECTIONS Connections
	RUN_PARAMS  Run_Params
}

type Connections struct {
	DB           *gorm.DB
	API_TOKEN    string
	DSN_STRING   string
	API_URL_BASE string
}

var (
	SEASON      int
	WEEK        int
	SEASON_TYPE string
)

type Run_Params struct {
	PSCD            []string
	START_SEASON    int
	END_SEASON      int
	GET_SEASON      bool
	GET_WEEKLY      bool
	GET_OFFSEASON   bool
	GET_ONE_OFFS    bool
	GET_FULL_SEASON bool
	INSERT_CAL      bool
}

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	CONFIG = Config{
		CONNECTIONS: Connections{
			API_TOKEN:    getEnv("API_TOKEN"),
			DSN_STRING:   getEnv("DSN_STRING"),
			API_URL_BASE: "https://api.collegefootballdata.com/",
		},
		RUN_PARAMS: Run_Params{
			PSCD:            strings.Split(getEnv("PSCD"), ","),
			START_SEASON:    getEnvAsInt("START_SEASON"),
			END_SEASON:      getEnvAsInt("END_SEASON"),
			GET_SEASON:      getEnvAsBool("GET_SEASON"),
			GET_WEEKLY:      getEnvAsBool("GET_WEEKLY"),
			GET_OFFSEASON:   getEnvAsBool("GET_OFFSEASON"),
			GET_ONE_OFFS:    getEnvAsBool("GET_ONE_OFFS"),
			GET_FULL_SEASON: getEnvAsBool("GET_FULL_SEASON"),
			INSERT_CAL:      getEnvAsBool("INSERT_CAL"),
		},
	}

	SEASON = getEnvAsInt("SEASON")
	WEEK = getEnvAsInt("WEEK")
	SEASON_TYPE = getEnv("SEASON_TYPE")
}

func getEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(val) == "" {
		panic("Missing required environment variable: " + key)
	}
	return val
}

func getEnvAsInt(key string) int {
	valStr := getEnv(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		panic("Invalid int for env variable " + key + ": " + valStr)
	}
	return val
}

func getEnvAsBool(key string) bool {
	valStr := getEnv(key)
	val, err := strconv.ParseBool(valStr)
	if err != nil {
		panic("Invalid bool for env variable " + key + ": " + valStr)
	}
	return val
}
