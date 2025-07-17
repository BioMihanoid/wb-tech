package main

import "fmt"

/*
	Реализовать паттерн проектирования «Адаптер» на любом примере.

	Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого,
	который ожидает клиент.

	Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура) и другой,
	несовместимый по интерфейсу потребитель — напишите адаптер, который реализует нужный интерфейс
	и делегирует вызовы к встроенному объекту.

	Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.
*/

type Reader interface {
	Read() string
}

type AltReader struct{}

func (r *AltReader) AltRead() string {
	return "alt read"
}

type ReaderAdapter struct {
	AltReader *AltReader
}

func (r *ReaderAdapter) Read() string {
	return r.AltReader.AltRead()
}

func main() {
	alt := &AltReader{}

	adapter := &ReaderAdapter{AltReader: alt}

	fmt.Println(adapter.Read())
}
