package challenge

import (
	"github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/config"
	"gorm.io/datatypes"
)

var Seeders = config.Seeds{
	&Category{}:  seedCategoriesTable,
	&Challenge{}: seedChallengeTable,
}

func seedCategoriesTable(db config.Database) {
	var c = []Category{
		{
			UUID: uuid.New(),
			Name: "culture",
		},
		{
			UUID: uuid.New(),
			Name: "nature",
		},
		{
			UUID: uuid.New(),
			Name: "food",
		},
		{
			UUID: uuid.New(),
			Name: "sports",
		},
		{
			UUID: uuid.New(),
			Name: "health",
		},
	}

	db.Create(&c)
}

func seedChallengeTable(db config.Database) {
	var c = []Challenge{
		{
			UUID:        uuid.New(),
			Name:        "Museu Colecção Berardo",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional de História Natural e da Ciência",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "MAAT - Museu de Arte, Arquitetura e Tecnologia",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Casa Fernando Pessoa",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Desporto",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Arqueológico do Carmo",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu do Aljube",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Casa das Histórias",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Fundação Arpad Szenes - Vieira da Silva",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu do Dinheiro",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Calouste Gulbenkian - Colecção do Fundador",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu do Oriente",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Atelier-Museu Júlio Pomar",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Arqueologia",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Azulejo",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Arte Antiga",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu da Marioneta",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Militar",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu da Carris",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu de Macau",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu da Música",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu da Presidência da República",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa - Santo António",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu do Fado",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu da Água",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa – Teatro Romano",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Casa Museu Dr. Anastácio Gonçalves",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa - Palácio Pimenta",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu de Arte Popular",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu da Farmácia",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional dos Coches",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Arte Contemporânea do Chiado",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Teatro e da Dança",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu das Comunicações",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu da Saúde",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional de Etnologia",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu de Lisboa - Torreão Poente",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Rafael Bordalo Pinheiro",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Nacional do Traje",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu Benfica - Cosme Damião",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "NewsMuseum",
			Description: "description",
			Thumbnail:   "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
			Gallery: datatypes.JSON(`[
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/4/47/Rijksmuseum_from_Museumplein_2523.jpg"
              ]`),
		},
		{
			UUID:        uuid.New(),
			Name:        "Museu do Palácio Nacional da Ajuda",
			Description: "description",
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
