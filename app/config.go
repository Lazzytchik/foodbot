package app

type Config struct {
	FoodApi ServiceUrl
}

type ServiceUrl struct {
	Host string
	Port string
}
