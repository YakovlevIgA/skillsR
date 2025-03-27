package main

import (
	"sync"
	"testing"
)

func TestProcessNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint8
		expected []float64
	}{
		{
			name:     "Single number",
			input:    []uint8{2},
			expected: []float64{8},
		},
		{
			name:     "Multiple numbers",
			input:    []uint8{1, 3, 5},
			expected: []float64{1, 27, 125},
		},
		{
			name:     "Zero value",
			input:    []uint8{0},
			expected: []float64{0},
		},
		{
			name:     "Max uint8 value",
			input:    []uint8{255},
			expected: []float64{255 * 255 * 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(chan uint8)
			output := make(chan float64)
			var wg sync.WaitGroup
			wg.Add(1)

			// Запускаем обработчик
			go processNumbers(input, output, &wg)

			// Запускаем горутину для отправки данных
			go func() {
				for _, num := range tt.input {
					input <- num
				}
				close(input)
			}()

			// Собираем результаты
			var results []float64
			var resultsWg sync.WaitGroup
			resultsWg.Add(1)
			go func() {
				defer resultsWg.Done()
				for res := range output {
					results = append(results, res)
				}
			}()

			// Ожидаем завершения
			wg.Wait()
			resultsWg.Wait()

			// Проверяем результаты
			if len(results) != len(tt.expected) {
				t.Errorf("Expected %d results, got %d", len(tt.expected), len(results))
			}

			for i, expected := range tt.expected {
				if results[i] != expected {
					t.Errorf("At index %d: expected %f, got %f", i, expected, results[i])
				}
			}
		})
	}
}

func TestEmptyInput(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)

	go processNumbers(input, output, &wg)

	// Закрываем входной канал сразу
	close(input)

	var results []float64
	var resultsWg sync.WaitGroup
	resultsWg.Add(1)
	go func() {
		defer resultsWg.Done()
		for res := range output {
			results = append(results, res)
		}
	}()

	wg.Wait()
	resultsWg.Wait()

	if len(results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(results))
	}
}

func TestChannelClosing(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)

	go processNumbers(input, output, &wg)

	// Проверяем что output канал закрывается
	var outputClosed bool
	go func() {
		for {
			_, ok := <-output
			if !ok {
				outputClosed = true
				return
			}
		}
	}()

	close(input)
	wg.Wait()

	if !outputClosed {
		t.Error("Output channel was not closed")
	}
}
