package product

import (
	"pesto-product-manager/database"
	"pesto-product-manager/log"
	"time"

	"go.uber.org/zap"
)

type Product struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Manufacturer string    `json:"manufacturer"`
	Description  string    `json:"description"`
	Price        float32   `json:"price"`
	Origin       string    `json:"origin"`
	LastUpdated  time.Time `json:"lastUpdated"`
	CreatedAt    time.Time `json:"createdAt"`
}

// create
func (p *Product) Create() error {
	// validations

	// write in DB
	data, err := database.DB.Exec(`INSERT INTO products ( id, name, manufacturer, description, price,
	 origin, last_updated, createdd_at, category) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		p.Id, p.Name, p.Manufacturer, p.Description, p.Price, p.Origin, p.LastUpdated, p.CreatedAt, p.Category)
	if err != nil {
		log.Logger.Error("Failed storing product in DB : ", zap.Error(err))
		return err
	}
	log.Logger.Debug("Product inserted with details : ", zap.Any("product", data))
	return nil
}

// read
func (p *Product) Get() error {
	// validations

	// delete from DB
	data, err := database.DB.Query(`select * from products where id=$1`, p.Id)
	if err != nil {
		log.Logger.Error("Failed reading product from DB : ", zap.Error(err))
		return err
	}
	log.Logger.Debug("Product fetched with details : ", zap.Any("product", data))
	for data.Next() {
		err = data.Scan(&p.Id, &p.Name, &p.Category, &p.Manufacturer, &p.Description,
			&p.Price, &p.Origin, &p.LastUpdated, &p.CreatedAt)
		if err != nil {
			log.Logger.Error("Failed scanning product details : ", zap.Error(err))
			return err
		}
	}
	return nil
}

// update

// delete
func (p *Product) Delete() error {
	// validations

	// delete from DB
	data, err := database.DB.Exec(`delete from products where id=$1`, p.Id)
	if err != nil {
		log.Logger.Error("Failed deleting product from DB : ", zap.Error(err))
		return err
	}
	log.Logger.Debug("Product deleted with details : ", zap.Any("product", data))
	return nil
}
