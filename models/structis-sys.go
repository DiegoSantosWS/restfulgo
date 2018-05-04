package models

//Products Struct of products
type Products struct {
	ID          int     `json:"id" db:"id"`
	Nome        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Stock       float64 `json:"stock" db:"stock"`
	Width       float64 `json:"width" db:"width"`
	Height      float64 `json:"heigth" db:"height"`
	Amount      float64 `json:"amount" db:"amount"`
	Weight      float64 `json:"weight" db:"weight"`
	Price       float64 `json:"price" db:"price"`
	Discount    float64 `json:"discount" db:"discount"`
	Promotion   float64 `json:"promotion" db:"promotion"`
}

//Message struct of the message
type Message struct {
	Executed bool
	Message  string
	COD      int
}

//AddressClients to listing clients
type AddressClients struct {
	Address      string `json:"address,omitempty" db:"address"`
	Number       string `json:"number,omitempty" db:"number"`
	City         string `json:"city,omitempty" db:"city"`
	Neighborhood string `json:"neighborhood,omitempty" db:"neighborhood"`
	Country      string `json:"country,omitempty" db:"country"`
	State        string `json:"state,omitempty" db:"state"`
}

//Clients to listing clients
type Clients struct {
	ID      int    `json:"id,omitempty" db:"id"`
	Name    string `json:"name,omitempty" db:"name"`
	Email   string `json:"email,omitempty" db:"email"`
	Phone   string `json:"phone,omitempty" db:"phone"`
	Status  string `json:"status,omitempty" db:"status"`
	Date    string `json:"date,omitempty" db:"date"`
	*AddressClients
}
