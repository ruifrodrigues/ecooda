package config

type (
	Seeds map[interface{}]func(Database)
)

func RunFactories(db IDatabase, seeds Seeds) {
	if db == nil {
		panic("Database must be enabled in NewConfig")
	}

	sqlDB, err := db.OpenConnection()
	if err != nil {
		panic("DB Connection failed")
	}

	for model, seed := range seeds {
		var count int64

		sqlDB.Model(model).Count(&count)
		if count == 0 {
			seed(sqlDB)
		}
	}

	err = sqlDB.CloseConnection()
	if err != nil {
		panic("DB connection not closed")
	}
}
