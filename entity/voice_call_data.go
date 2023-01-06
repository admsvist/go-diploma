package entity

type VoiceCallData struct {
	Country             string  `json:"country"`             // alpha-2 — код страны;
	Bandwidth           int     `json:"bandwidth"`           // текущая нагрузка в процентах;
	ResponseTime        int     `json:"responseTime"`        // среднее время ответа;
	Provider            string  `json:"provider"`            // провайдер;
	ConnectionStability float32 `json:"connectionStability"` // стабильность соединения;
	TTFB                int     `json:"TTFB"`                // ?
	VoicePurity         int     `json:"voicePurity"`         // чистота TTFB-связи;
	MedianOfCallsTime   int     `json:"medianOfCallsTime"`   // медиана длительности звонка.
}
