package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	// Создаем временный HTTP-сервер для тестирования
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}))
	defer ts.Close()

	// Создаем временную директорию для сохранения загруженных файлов
	tempDir, err := os.MkdirTemp("", "wget")
	if err != nil {
		t.Fatalf("ошибка создания временной директории: %v", err)
	}
	defer os.RemoveAll(tempDir) // Удаляем временную директорию после теста

	err = DownloadPage(ts.URL, tempDir)
	if err != nil {
		t.Errorf("ошибка загрузки страницы: %v", err)
	}

	// Проверяем, что файл был создан
	parsedURL, _ := url.Parse(ts.URL)
	filePath := filepath.Join(tempDir, parsedURL.Host, "index.html")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("файл не был создан: %s", filePath)
	}
}

func TestExtractLinks(t *testing.T) {
	htmlContent := `
	<html>
		<body>
			<a href="http://example.com/page1">Page 1</a>
			<a href="/page2">Page 2</a>
		</body>
	</html>`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(htmlContent))
	}))
	defer ts.Close()

	links, err := ExtractLinks(ts.URL)
	if err != nil {
		t.Errorf("ошибка извлечения ссылок: %v", err)
	}

	expectedLinks := []string{
		"http://example.com/page1",
		ts.URL + "/page2",
	}

	if len(links) != len(expectedLinks) {
		t.Errorf("ожидалось %d ссылок, получено %d", len(expectedLinks), len(links))
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("ожидалось %s, получено %s", expectedLinks[i], link)
		}
	}
}

func TestDownloadSite(t *testing.T) {
	htmlContent := `
	<html>
		<body>
			<a href="/page1">Page 1</a>
		</body>
	</html>`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/page1" {
			w.Write([]byte("Page 1 content"))
		} else {
			w.Write([]byte(htmlContent))
		}
	}))
	defer ts.Close()

	tempDir, err := os.MkdirTemp("", "wget")
	if err != nil {
		t.Fatalf("ошибка создания временной директории: %v", err)
	}
	defer os.RemoveAll(tempDir)

	err = DownloadSite(ts.URL, tempDir)
	if err != nil {
		t.Errorf("ошибка загрузки сайта: %v", err)
	}

	parsedURL, _ := url.Parse(ts.URL)
	indexPath := filepath.Join(tempDir, parsedURL.Host, "index.html")
	page1Path := filepath.Join(tempDir, parsedURL.Host, "page1.html")

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("файл index.html не был создан: %s", indexPath)
	}

	if _, err := os.Stat(page1Path); os.IsNotExist(err) {
		t.Errorf("файл page1.html не был создан: %s", page1Path)
	}
}
