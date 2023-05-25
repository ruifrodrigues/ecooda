package challenge

import (
	"github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/config"
	"gorm.io/datatypes"
)

var Seeders = config.Seeds{
	&Category{}:           seedCategoriesTable,
	&Challenge{}:          seedChallengeTable,
	&CategoryChallenges{}: seedCategoryChallengesTable,
}

func seedCategoriesTable(db config.Database) {
	var c = []Category{
		{
			ID:   1,
			UUID: uuid.New(),
			Name: "culture",
		},
		{
			ID:   2,
			UUID: uuid.New(),
			Name: "nature",
		},
		{
			ID:   3,
			UUID: uuid.New(),
			Name: "food",
		},
		{
			ID:   4,
			UUID: uuid.New(),
			Name: "sports",
		},
		{
			ID:   5,
			UUID: uuid.New(),
			Name: "health",
		},
	}

	db.Create(&c)
}

func seedChallengeTable(db config.Database) {
	var c = []Challenge{
		{
			ID:          1,
			UUID:        uuid.New(),
			Name:        "Museu Colecção Berardo",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          2,
			UUID:        uuid.New(),
			Name:        "Museu Nacional de História Natural e da Ciência",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          3,
			UUID:        uuid.New(),
			Name:        "MAAT - Museu de Arte, Arquitetura e Tecnologia",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          4,
			UUID:        uuid.New(),
			Name:        "Casa Fernando Pessoa",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          5,
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Desporto",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          6,
			UUID:        uuid.New(),
			Name:        "Museu Arqueológico do Carmo",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          7,
			UUID:        uuid.New(),
			Name:        "Museu do Aljube",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          8,
			UUID:        uuid.New(),
			Name:        "Casa das Histórias",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          9,
			UUID:        uuid.New(),
			Name:        "Fundação Arpad Szenes - Vieira da Silva",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          10,
			UUID:        uuid.New(),
			Name:        "Museu do Dinheiro",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          11,
			UUID:        uuid.New(),
			Name:        "Museu Calouste Gulbenkian - Colecção do Fundador",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          12,
			UUID:        uuid.New(),
			Name:        "Museu do Oriente",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          13,
			UUID:        uuid.New(),
			Name:        "Atelier-Museu Júlio Pomar",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          14,
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Arqueologia",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          15,
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Azulejo",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          16,
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Arte Antiga",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          17,
			UUID:        uuid.New(),
			Name:        "Museu da Marioneta",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          18,
			UUID:        uuid.New(),
			Name:        "Museu Militar",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          19,
			UUID:        uuid.New(),
			Name:        "Museu da Carris",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          20,
			UUID:        uuid.New(),
			Name:        "Museu da Música",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          21,
			UUID:        uuid.New(),
			Name:        "Museu da Presidência da República",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          22,
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa - Santo António",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          23,
			UUID:        uuid.New(),
			Name:        "Museu do Fado",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          24,
			UUID:        uuid.New(),
			Name:        "Museu da Água",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          25,
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa – Teatro Romano",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          26,
			UUID:        uuid.New(),
			Name:        "Casa Museu Dr. Anastácio Gonçalves",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          27,
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa - Palácio Pimenta",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          28,
			UUID:        uuid.New(),
			Name:        "Museu de Arte Popular",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          29,
			UUID:        uuid.New(),
			Name:        "Museu da Farmácia",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          30,
			UUID:        uuid.New(),
			Name:        "Museu Nacional dos Coches",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          31,
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Arte Contemporânea do Chiado",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          32,
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Teatro e da Dança",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          33,
			UUID:        uuid.New(),
			Name:        "Museu das Comunicações",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          34,
			UUID:        uuid.New(),
			Name:        "Museu da Saúde",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          35,
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Etnologia",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          36,
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa - Torreão Poente",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          37,
			UUID:        uuid.New(),
			Name:        "Museu Rafael Bordalo Pinheiro",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          38,
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Traje",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          39,
			UUID:        uuid.New(),
			Name:        "Museu Benfica - Cosme Damião",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          40,
			UUID:        uuid.New(),
			Name:        "NewsMuseum",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          41,
			UUID:        uuid.New(),
			Name:        "Museu do Palácio Nacional da Ajuda",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
		{
			ID:          42,
			UUID:        uuid.New(),
			Name:        "Museu de Macau",
			Description: "description",
			Latitude:    39.5572,
			Longitude:   -7.85366,
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
               "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
             ]`),
		},
	}

	db.Create(&c)
}

func seedCategoryChallengesTable(db config.Database) {
	var cc = []CategoryChallenges{
		{
			CategoryID:  1,
			ChallengeID: 1,
		},
		{
			CategoryID:  1,
			ChallengeID: 2,
		},
		{
			CategoryID:  1,
			ChallengeID: 3,
		},
		{
			CategoryID:  1,
			ChallengeID: 4,
		},
		{
			CategoryID:  1,
			ChallengeID: 5,
		},
		{
			CategoryID:  1,
			ChallengeID: 6,
		},
		{
			CategoryID:  1,
			ChallengeID: 7,
		},
		{
			CategoryID:  1,
			ChallengeID: 8,
		},
		{
			CategoryID:  1,
			ChallengeID: 9,
		},
		{
			CategoryID:  1,
			ChallengeID: 10,
		},
		{
			CategoryID:  1,
			ChallengeID: 11,
		},
		{
			CategoryID:  1,
			ChallengeID: 12,
		},
		{
			CategoryID:  1,
			ChallengeID: 13,
		},
		{
			CategoryID:  1,
			ChallengeID: 14,
		},
		{
			CategoryID:  1,
			ChallengeID: 15,
		},
		{
			CategoryID:  1,
			ChallengeID: 16,
		},
		{
			CategoryID:  1,
			ChallengeID: 17,
		},
		{
			CategoryID:  1,
			ChallengeID: 18,
		},
		{
			CategoryID:  1,
			ChallengeID: 19,
		},
		{
			CategoryID:  1,
			ChallengeID: 20,
		},
		{
			CategoryID:  1,
			ChallengeID: 21,
		},
		{
			CategoryID:  1,
			ChallengeID: 22,
		},
		{
			CategoryID:  1,
			ChallengeID: 23,
		},
		{
			CategoryID:  1,
			ChallengeID: 24,
		},
		{
			CategoryID:  1,
			ChallengeID: 25,
		},
		{
			CategoryID:  1,
			ChallengeID: 26,
		},
		{
			CategoryID:  1,
			ChallengeID: 27,
		},
		{
			CategoryID:  1,
			ChallengeID: 28,
		},
		{
			CategoryID:  1,
			ChallengeID: 29,
		},
		{
			CategoryID:  1,
			ChallengeID: 30,
		},
		{
			CategoryID:  1,
			ChallengeID: 31,
		},
		{
			CategoryID:  1,
			ChallengeID: 32,
		},
		{
			CategoryID:  1,
			ChallengeID: 33,
		},
		{
			CategoryID:  1,
			ChallengeID: 34,
		},
		{
			CategoryID:  1,
			ChallengeID: 35,
		},
		{
			CategoryID:  1,
			ChallengeID: 36,
		},
		{
			CategoryID:  1,
			ChallengeID: 37,
		},
		{
			CategoryID:  1,
			ChallengeID: 38,
		},
		{
			CategoryID:  1,
			ChallengeID: 39,
		},
		{
			CategoryID:  1,
			ChallengeID: 40,
		},
		{
			CategoryID:  1,
			ChallengeID: 41,
		},
		{
			CategoryID:  1,
			ChallengeID: 42,
		},
	}

	db.Create(&cc)
}
