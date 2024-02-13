package Configuration

const (
	// System Configuration
	ConstProcessorNumber = 4

	// Application Property
	ConstAppName = "Web Client Auth RPC"

	// Logging Queue Name
	ConstOnlineLogQueueName = "OnlineLog"

	// Database POSTGRESQL Configuration
	ConstPSQLHost = "172.31.28.208"
	ConstPSQLPort = "5432"
	ConstPSQLUser = "pawitra"
	ConstPSQLPass = "Eliandri3"
	ConstPSQLName = "krapoex_auth"

	// Database MONGODB
	ConstMongoHost          = "172.31.28.208"
	ConstMongoPort          = 27017
	ConstMongoUser          = "chandra"
	ConstMongoPass          = "Eliandri3"
	ConstMongoAuditrailDB   = "krapoex"
	ConstAuditrailMongoColl = "web_client_auditrail"

	// RabbitMQ
	ConstRabbitHost  = "172.31.28.208"
	ConstRabbitPort  = "5672"
	ConstRabbitUser  = "chandra"
	ConstRabbitPass  = "Eliandri3"
	ConstRabbitVHost = "krapoex"
	// MSI RAKAI
	// ConstRabbitHost  = "172.31.28.208"
	// ConstRabbitPort  = "5672"
	// ConstRabbitUser  = "biznet"
	// ConstRabbitPass  = "SayaBisa123!"
	// ConstRabbitVHost = "KRAPOEX"

	// Redis
	// ConstRedisHost            = "localhost"
	// ConstRedisPort            = "6333"
	// ConstRedisPass            = "Eliandri3"
	// ConstRedisWebClientAuthDB = 1
	// MSI RAKAI
	ConstRedisHost            = "172.31.28.208"
	ConstRedisPort            = "6363"
	ConstRedisPass            = "SayaBisa123!"
	ConstRedisWebClientAuthDB = 13

	// ConstRedisHost            = "127.0.0.1"
	// ConstRedisPort            = "6379"
	// ConstRedisPass            = "123456"
	// ConstRedisWebClientAuthDB = 13

	// RPC WebClient Auth
	ConstRPCWebClientAuthHost = "localhost"
	ConstRPCWebClientAuthPort = "18000"

	// General SALT encryption
	ConstDefaultSaltEncryption = "_Semoga_Semua_Mahluk_Berbahagia_"

	ConstJwtSecret = "jwtsecret"
)
