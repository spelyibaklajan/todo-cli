package main

import "fmt"

func main() {
    todos := []string{}
    
    // добавляем задачи
    todos = append(todos, "Купить еду")
    todos = append(todos, "Учить Go")
    todos = append(todos, "Сделать проект")
    
    // показываем все задачи
    for i, todo := range todos {
        fmt.Printf("%d. %s\n", i+1, todo)
    }
	// удаляем задачу номер 2 (индекс 1)
	todos = append(todos[:1], todos[2:]...)
	fmt.Println("\nПосле удаления:")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
}
}