package filereader

import (
	"io"
	"net/http"
)

type UrlReader struct{}

func New() *UrlReader {
	return &UrlReader{}
}

func (r *UrlReader) ReadUrl(url string) ([]byte, error) {
	// Отправить GET-запрос по указанному URL
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []byte{}, nil
	}

	// Прочитать содержимое ответа в байтовый срез
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
