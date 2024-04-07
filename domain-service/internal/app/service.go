package app

// @title Client service
// @version @0.9
// @description swagger для api к клиентскому сервису

// @basePath /api

// Implementation структура для реализации различных ручек
type Implementation struct {
}

// NewImplementation конструктор для Implementation
func NewImplementation() (*Implementation, error) {
	return &Implementation{}, nil
}
