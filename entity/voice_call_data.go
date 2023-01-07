package entity

type VoiceCallData struct {
	Country             string  `json:"country"`              // alpha-2 — код страны;
	Bandwidth           int     `json:"bandwidth"`            // текущая нагрузка в процентах;
	ResponseTime        int     `json:"response_time"`        // среднее время ответа;
	Provider            string  `json:"provider"`             // провайдер;
	ConnectionStability float32 `json:"connection_stability"` // стабильность соединения;
	TTFB                int     `json:"ttfb"`                 // ?
	VoicePurity         int     `json:"voice_purity"`         // чистота TTFB-связи;
	MedianOfCallsTime   int     `json:"median_of_calls_time"` // медиана длительности звонка.
}
