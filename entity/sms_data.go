package entity

type SMSData struct {
	Сountry      string // alpha-2 — код страны;
	Bandwidth    string // пропускная способность канала от 0 до 100%;
	ResponseTime string // среднее время ответа в миллисекундах;
	Provider     string // название компании-провайдера.
}
