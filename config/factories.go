package config

type (
	Seeds map[interface{}]func(Database)
)

func RunFactories(db DB, seeds Seeds) {
	if db == nil {
		panic("Database must be enabled in NewConfig")
	}

	if err := db.OpenConnection(); err != nil {
		panic("DB Connection failed")
	}

	for model, seed := range seeds {
		var count int64

		db.Ctx().Model(model).Count(&count)
		if count == 0 {
			seed(db.Ctx())
		}
	}

	if err := db.CloseConnection(); err != nil {
		panic("DB connection not closed")
	}
}
