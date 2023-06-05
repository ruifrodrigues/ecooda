package challenge

import (
	"github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/config"
	"github.com/ruifrodrigues/ecooda/utils"
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
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](1),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](2),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](3),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](4),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](5),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](6),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](7),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](8),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](9),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](10),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](11),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](12),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](13),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](14),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](15),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](16),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](17),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](18),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](19),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](20),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](21),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](22),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](23),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](24),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](25),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](26),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](27),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](28),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](29),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](30),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](31),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](32),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](33),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](34),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](35),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](36),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](37),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](38),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](39),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](40),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](41),
		},
		{
			CategoryID:  utils.Pointer[uint](1),
			ChallengeID: utils.Pointer[uint](42),
		},
	}

	db.Create(&cc)
}
