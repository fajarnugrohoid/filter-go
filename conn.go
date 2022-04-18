package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v3"
	"os"
)

const configPath = "config.yml"
const CONNECTED = "Successfully connected to database: %v"

type Cfg struct {
	DB string `yaml:"db"`
}

var AppConfig *Cfg

type MongoDatastore struct {
	db      *mongo.Database
	Session *mongo.Client
	logger  *logrus.Logger
}

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	fmt.Println(decoder)
	err = decoder.Decode(&AppConfig)

	fmt.Println(AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}

/*
func NewDatastore(config config.GeneralConfig, logger *logrus.Logger) *MongoDatastore {

	var mongoDataStore *MongoDatastore
	db, session := connect(config, logger)
	if db != nil && session != nil {

		// log statements here as well

		mongoDataStore = new(MongoDatastore)
		mongoDataStore.db = db
		mongoDataStore.logger = logger
		mongoDataStore.Session = session
		return mongoDataStore
	}

	logger.Fatalf("Failed to connect to database: %v", config.DatabaseName)

	return nil
}

func connect(generalConfig config.GeneralConfig, logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(generalConfig, logger)
	})

	return db, session
}

func connectToMongo(generalConfig config.GeneralConfig, logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {

	var err error
	session, err := mongo.NewClient(generalConfig.DatabaseHost)
	if err != nil {
		logger.Fatal(err)
	}
	session.Connect(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	var DB = session.Database(generalConfig.DatabaseName)
	logger.Info(CONNECTED, generalConfig.DatabaseName)

	return DB, session
}*/
