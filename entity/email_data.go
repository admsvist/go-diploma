package entity

type EmailData struct {
	Country      string // alpha-2 — код страны;
	Provider     string // провайдер;
	DeliveryTime int    // среднее время доставки писем в миллисекундах.
}
