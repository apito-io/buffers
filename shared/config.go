package shared

type CommonDatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type CacheDBConfig struct {
	DB          *CommonDatabaseConfig
	CacheDriver string
	CacheTTL    string
}

type SystemDBConfig struct {
	DB          *CommonDatabaseConfig
	LocalDBPath string
}
