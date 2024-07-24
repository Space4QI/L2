package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

// DownloadPage загружает страницу и сохраняет её на диск
func DownloadPage(urlStr string, baseDir string) error {
	resp, err := http.Get(urlStr) // Отправляем HTTP GET запрос
	if err != nil {
		return fmt.Errorf("ошибка запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("получен статус: %v", resp.Status)
	}

	parsedURL, err := url.Parse(urlStr) // Парсим URL
	if err != nil {
		return fmt.Errorf("ошибка парсинга URL: %v", err)
	}

	// Формируем путь к файлу для сохранения страницы
	filePath := filepath.Join(baseDir, parsedURL.Host, parsedURL.Path)
	if strings.HasSuffix(urlStr, "/") {
		filePath = filepath.Join(filePath, "index.html")
	} else if !strings.Contains(filepath.Base(filePath), ".") {
		filePath += ".html"
	}

	// Создаем необходимые директории
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("ошибка создания директорий: %v", err)
	}

	// Создаем файл для записи данных
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %v", err)
	}
	defer file.Close()

	// Копируем данные из ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	fmt.Printf("Скачано: %s\n", urlStr)
	return nil
}

// ExtractLinks извлекает все ссылки с данной HTML-страницы
func ExtractLinks(urlStr string) ([]string, error) {
	resp, err := http.Get(urlStr) // Отправляем HTTP GET запрос
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("получен статус: %v", resp.Status)
	}

	z := html.NewTokenizer(resp.Body) // Создаем токенайзер для HTML-документа
	var links []string

	for {
		tt := z.Next() // Получаем следующий токен
		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "a" { // Если токен является тегом <a>, то извлекаем значение атрибута href
				for _, a := range t.Attr {
					if a.Key == "href" {
						u, err := url.Parse(a.Val)
						if err != nil {
							continue
						}
						base, err := url.Parse(urlStr)
						if err != nil {
							continue
						}
						absURL := base.ResolveReference(u) // Преобразуем относительный URL в абсолютный
						links = append(links, absURL.String())
					}
				}
			}
		}
	}
}

// DownloadSite загружает сайт целиком
func DownloadSite(urlStr string, baseDir string) error {
	queue := []string{urlStr}        // Очередь URL для обработки
	visited := make(map[string]bool) // Карта посещённых URL

	for len(queue) > 0 {
		currURL := queue[0]
		queue = queue[1:]

		if visited[currURL] {
			continue
		}
		visited[currURL] = true

		err := DownloadPage(currURL, baseDir) // Загружаем страницу
		if err != nil {
			fmt.Printf("Ошибка загрузки страницы %s: %v\n", currURL, err)
			continue
		}

		links, err := ExtractLinks(currURL) // Извлекаем ссылки со страницы
		if err != nil {
			fmt.Printf("Ошибка извлечения ссылок %s: %v\n", currURL, err)
			continue
		}
		queue = append(queue, links...) // Добавляем ссылки в очередь
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Использование: wget URL DIRECTORY")
		return
	}

	urlStr := os.Args[1]
	baseDir := os.Args[2]

	err := DownloadSite(urlStr, baseDir)
	if err != nil {
		fmt.Printf("Ошибка загрузки сайта: %v\n", err)
	}
}
