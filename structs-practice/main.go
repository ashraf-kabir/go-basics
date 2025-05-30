package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

//type displayer interface {
//	Display()
//}

// embedded interface
type outputtable interface {
	saver
	Display()
}

//type outputtable interface {
//	Save() error
//	Display()
//}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving data failed.")
		return err
	}

	fmt.Println("Date saved successfully.")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// any value allowed, switch statement
func printSomething(value any) {
	//switch value.(type) {
	//case int:
	//	fmt.Println("Integer ", value)
	//case string:
	//	fmt.Println("String ", value)
	//case float64:
	//	fmt.Println("Float64 ", value)
	//default:
	//	fmt.Println("Unknown type ", value)
	//}

	intVal, ok := value.(int)
	if ok {
		//intVal + 1
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatVal)
	}
}
