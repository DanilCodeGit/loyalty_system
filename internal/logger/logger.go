package logger

import (
	"bytes"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

var sugar zap.SugaredLogger

func WithLogging(h http.Handler) http.Handler {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar = *logger.Sugar()

	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		var requestBodyBuffer bytes.Buffer
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			requestBody, err := io.ReadAll(r.Body)
			if err != nil {

			} else {
				requestBodyBuffer.Write(requestBody)
				r.Body = io.NopCloser(bytes.NewBuffer(requestBody))
			}
		}

		h.ServeHTTP(&lw, r) // внедряем реализацию http.ResponseWriter

		duration := time.Since(start)

		// отправляем сведения о запросе в zap
		sugar.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"status", responseData.status, // получаем перехваченный код статуса ответа
			"duration", duration,
			"size", responseData.size,
			"request_body", requestBodyBuffer.String(), // получаем перехваченный размер ответа
		)

	}
	// возвращаем функционально расширенный хендлер
	return http.HandlerFunc(logFn)
}

type (
	// Берём структуру для хранения сведений об ответе
	responseData struct {
		status int
		size   int
	}

	// добавляем реализацию http.ResponseWriter
	loggingResponseWriter struct {
		http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size // захватываем размер
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}
