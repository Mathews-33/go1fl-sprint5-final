package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает методы для разбора и вывода информации об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает слайс строк с данными, используя переданный DataParser.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
