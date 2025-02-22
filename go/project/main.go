/*
1. Необходимо поднять сервер. Порт для сервера необходимо подтянуть из файла env
2. Сервер может принимать GET и POST запросы
3. Необходимо написать 3 клиентов: Первый и второй клиент будут в горутинах слать по 100 POST запросов на сервер.
Третий клиент будет с каждые 5 секунд проверять состояние сервера. Доступен или нет.

4. Сделать имитацию ответов сервера: Сервер в рандомном порядке будет 1 и 2 клиенту в ответ отправлять статусы:
70 положительных ответов, 30 отрицательных
http.StatusOK (положительный статус)
http.StatusAccepted (положительный статус)
http.StatusBadRequest (отрицательный)
http.StatusInternalServerError (отрицательный)

5. Каждый клиент имеет по 2 воркера и каждый может отправлять по 5 запросов на сервер.
По завершению отправки, каждый клиент должен выводить статистику:

Отправлено запросов: #
Разбивка по статусам: 200 - #, 400 - #, 500 - # и тд

6. Ограничить пропускную способность сервера 5 запрос в секунду и ограничить отправку на стороне клиента 5 запрос в секунду
Если лимит превышен, то возвращать уникальное сообщение о том, что лимит превышен

7. После того, как будут отправлены все запросы клиентами на сервер и даны ответы, необходимо посчитать количество
положительных и отрицательных запросов
1. Всего для сервера
2. На каждого из клиентов
И сохранить это в JSON, который можно будет потом получить с сервера по GET запросу */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// ----------------------------
// Конфигурация и типы для сервера
// ----------------------------

// Options хранит настраиваемые параметры для сервера.
type Options struct {
	MaxConnections int
	Timeout        int
	Protocol       string
	Port           int
	Logging        bool
}

// ServerStats хранит статистику, специфичную для сервера.
type ServerStats struct {
	Positive int `json:"positive"` // колво положительных ответов
	Negative int `json:"negative"` // колво отрицательных ответов
}

// AllStats хранит статистику для сервера и для каждого клиента.
type AllStats struct {
	Server  ServerStats            `json:"server"`
	Clients map[string]ServerStats `json:"clients"`
}

// Server инкапсулирует конфигурацию сервера, статистику, ограничитель скорости и маршрутизатор.
type Server struct {
	Opts           Options
	Stats          *AllStats
	Limiter        chan struct{}  // ограничитель (лимитер) запросов
	Mux            *http.ServeMux // мультиплексор (маршрутизатор) HTTP-запросов
	StatsMu        sync.Mutex     // мьютекс для безопасного обновления статистики
	ClientCounters map[string]int // счетчики запросов для каждого клиента
}

// NewServer создаёт и инициализирует новый сервер.
func NewServer(opts Options) *Server {
	s := &Server{
		Opts: opts,
		Stats: &AllStats{
			Server:  ServerStats{},
			Clients: make(map[string]ServerStats),
		},
		// Создаем канал-лимитер с емкостью 5 (5 запросов в секунду)
		Limiter:        make(chan struct{}, 5),
		Mux:            http.NewServeMux(),
		ClientCounters: make(map[string]int),
	}

	// Запускаем тикер, который пополняет токены ограничителя (один токен каждые 200 мс)
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case s.Limiter <- struct{}{}:
			default:
				// Если канал заполнен, ничего не делаем
			}
		}
	}()

	return s
}

// rateLimitMiddleware ограничивает количество входящих запросов, используя ограничитель сервера.
func (s *Server) rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-s.Limiter:
			next(w, r)
		default:
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintln(w, "Превышен лимит запросов: слишком много запросов")
		}
	}
}

// mainHandler распределяет GET и POST запросы для эндпоинта "/".
func (s *Server) mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handleGet(w)
	} else if r.Method == http.MethodPost {
		s.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// handleGet – простой обработчик GET-запроса.
func (s *Server) handleGet(w http.ResponseWriter) {
	fmt.Fprintln(w, "Сервер работает")
}

// handlePost имитирует ответ сервера, обновляет статистику и возвращает код статуса.
func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bodyStr := string(body)
	fmt.Printf("Получен POST запрос: %s\n", bodyStr)

	// Имитация ответа:
	// 70% вероятность положительного ответа (200 или 202)
	// 30% вероятность отрицательного ответа (400 или 500)
	randomValue := rand.Float64() * 100
	var statusCode int
	if randomValue < 70 {
		if rand.Intn(2) == 0 {
			statusCode = http.StatusOK // 200
		} else {
			statusCode = http.StatusAccepted // 202
		}
	} else {
		if rand.Intn(2) == 0 {
			statusCode = http.StatusBadRequest // 400
		} else {
			statusCode = http.StatusInternalServerError // 500
		}
	}

	// Обновляем статистику (только для тех кодов ответа, которые нас интересуют)
	if statusCode == http.StatusOK || statusCode == http.StatusAccepted ||
		statusCode == http.StatusBadRequest || statusCode == http.StatusInternalServerError {
		s.StatsMu.Lock()
		// Обновляем глобальную статистику сервера.
		if statusCode == http.StatusOK || statusCode == http.StatusAccepted {
			s.Stats.Server.Positive++
		} else {
			s.Stats.Server.Negative++
		}

		// Определяем имя клиента по содержимому запроса.
		var clientName string
		if strings.Contains(bodyStr, "Клиент 1") {
			clientName = "Клиент 1"
		} else if strings.Contains(bodyStr, "Клиент 2") {
			clientName = "Клиент 2"
		}

		// Если имя клиента определено, обновляем статистику для этого клиента.
		if clientName != "" {
			cs := s.Stats.Clients[clientName]
			if statusCode == http.StatusOK || statusCode == http.StatusAccepted {
				cs.Positive++
			} else {
				cs.Negative++
			}
			s.Stats.Clients[clientName] = cs
		}
		s.StatsMu.Unlock()
	}

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Ответ сервера: %d\n", statusCode)
}

// statsHandler возвращает собранную статистику в формате JSON.
func (s *Server) statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	s.StatsMu.Lock()
	defer s.StatsMu.Unlock()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(s.Stats)
}

// ServeHTTP позволяет серверу реализовать интерфейс http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

// ----------------------------
// Клиентский код
// ----------------------------

// worker отправляет указанное количество запросов, ожидая токен от клиентского ограничителя.
func worker(clientName string, workerID, requestsPerWorker int, serverURL string, clientLimiter <-chan struct{}, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= requestsPerWorker; i++ {
		<-clientLimiter
		data := fmt.Sprintf("%s - Worker %d: POST запрос номер %d", clientName, workerID, i)
		resp, err := http.Post(serverURL, "text/plain", bytes.NewBufferString(data))
		if err != nil {
			fmt.Printf("%s - Worker %d: ошибка отправки POST: %v\n", clientName, workerID, err)
			results <- -1
		} else {
			results <- resp.StatusCode
			resp.Body.Close()
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// runClient запускает указанное количество воркеров для клиента и выводит локальную статистику.
func runClient(clientName, serverURL string, workerCount, requestsPerWorker int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Клиентский ограничитель: разрешает 5 запросов в секунду.
	clientLimiter := make(chan struct{}, 5)
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case clientLimiter <- struct{}{}:
			default:
			}
		}
	}()

	results := make(chan int, workerCount*requestsPerWorker)
	var workerWG sync.WaitGroup
	for w := 1; w <= workerCount; w++ {
		workerWG.Add(1)
		go worker(clientName, w, requestsPerWorker, serverURL, clientLimiter, results, &workerWG)
	}
	workerWG.Wait()
	close(results)

	totalRequests := 0
	statusCounts := make(map[int]int)
	for status := range results {
		totalRequests++
		statusCounts[status]++
	}

	fmt.Printf("\nСтатистика для %s:\n", clientName)
	fmt.Printf("Всего отправлено запросов: %d\n", totalRequests)
	for code, count := range statusCounts {
		if code == -1 {
			fmt.Printf("Ошибки отправки: %d\n", count)
		} else {
			fmt.Printf("Статус %d: %d\n", code, count)
		}
	}
}

// checkServerStatusWithTimeout проверяет доступность сервера каждые 5 секунд в течение заданного времени.
func checkServerStatusWithTimeout(serverURL string, duration time.Duration, done chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	timeout := time.NewTimer(duration)
	defer timeout.Stop()

	for {
		select {
		case <-ticker.C:
			resp, err := http.Get(serverURL)
			if err != nil {
				fmt.Printf("Ошибка проверки сервера: %v\n", err)
			} else {
				fmt.Printf("Проверка сервера: код статуса %d\n", resp.StatusCode)
				resp.Body.Close()
			}
		case <-timeout.C:
			fmt.Println("Время проверки сервера истекло.")
			done <- struct{}{}
			return
		}
	}
}

// ----------------------------
//        ФУНКЦИЯ MAIN
// ----------------------------

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла, используются настройки по умолчанию")
	}
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	var port int
	fmt.Sscanf(portStr, "%d", &port)

	// Создаем настройки сервера и инициализируем сервер.
	opts := Options{
		MaxConnections: 100,
		Timeout:        30,
		Protocol:       "http",
		Port:           port,
		Logging:        true,
	}
	server := NewServer(opts)

	// Регистрируем эндпоинты на mux сервера.
	server.Mux.HandleFunc("/", server.rateLimitMiddleware(server.mainHandler))
	server.Mux.HandleFunc("/stats", server.statsHandler)

	// Запускаем сервер в отдельной горутине.
	go func() {
		addr := fmt.Sprintf(":%d", server.Opts.Port)
		fmt.Printf("Запуск сервера на %s\n", addr)
		if err := http.ListenAndServe(addr, server); err != nil {
			log.Fatalf("Ошибка сервера: %v", err)
		}
	}()

	serverURL := fmt.Sprintf("http://localhost:%d", server.Opts.Port)

	// Запускаем клиентов.
	var clientsWG sync.WaitGroup
	clientsWG.Add(2)
	// Каждый клиент использует 2 воркера, которые отправляют по 50 запросов (всего 100 запросов на клиента).
	go runClient("Клиент 1", serverURL, 2, 50, &clientsWG)
	go runClient("Клиент 2", serverURL, 2, 50, &clientsWG)

	// Запускаем проверку доступности сервера с таймаутом 30 секунд.
	done := make(chan struct{})
	go checkServerStatusWithTimeout(serverURL, 30*time.Second, done)

	clientsWG.Wait()
	fmt.Println("\nВсе клиенты завершили отправку запросов.")
	fmt.Println("Статистику можно получить по адресу:", serverURL+"/stats")

	<-done
	fmt.Println("Завершение работы приложения.")
}
