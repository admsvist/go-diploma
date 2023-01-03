package entity

type VoiceCallData struct {
	Country             string  // alpha-2 — код страны;
	Bandwidth           int     // текущая нагрузка в процентах;
	ResponseTime        int     // среднее время ответа;
	Provider            string  // провайдер;
	ConnectionStability float32 // стабильность соединения;
	TTFB                int     // ?
	VoicePurity         int     // чистота TTFB-связи;
	MedianOfCallsTime   int     // медиана длительности звонка.
}
