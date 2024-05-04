package product

import "github.com/EronAlves1996/services/products/db"

func GetAll() ([]Product, error) {
	db := db.GetConnection()
	rows, err := db.Query("SELECT * FROM product")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ps := make([]Product, 0)

	for rows.Next() {
		var id uint32
		var name string
		var price int
		var image_url string

		if err := rows.Scan(&id, &name, &price, &image_url); err != nil {
			return nil, err
		}

		ps = append(ps, Product{
			Id:       id,
			Name:     name,
			Price:    float64(price) / 100,
			ImageUrl: image_url,
		})
	}

	return ps, nil
}
