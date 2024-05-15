package main

import "fmt"

var tasks []TODO

func main() {
	var option int
	fmt.Println("TODO is running")
	for option != -1 {
		fmt.Println("\nPlease Select A Task TODO!")
		fmt.Println("1. Create a Task")
		fmt.Println("2. Read Tasks")
		fmt.Println("3. Mark Task Complete")
		fmt.Println("4. Delete a Task")
		fmt.Println("Type -1 to exit")
		fmt.Scan(&option)
		switch option {
		case 1:
			var task string
			count++
			fmt.Println("Please enter a task")
			fmt.Scan(&task)
			tasks = append(tasks, *createItem(task))
			fmt.Printf("Task %s with id %d is created!", task, count)
			break
		case 2:
			displayTasks()
		case 3:
			var id int
			fmt.Println("Please select one id to mark complete")
			displayTasks()
			fmt.Scan(&id)
			markComplete(id)
		case 4:
			if len(tasks) == 0 {
				fmt.Println("No Tasks Available to Delete.")
			} else {
				var id int
				fmt.Println("Please select one id to delete")
				displayTasks()
				fmt.Scan(&id)
				deleteItem(id)
			}
			break
		case -1:
			fmt.Println("Exiting TODO")
			break
		}
	}
}

var count int = 0

type TODO struct {
	id        int
	title     string
	completed bool
}

func createItem(task string) *TODO {
	var item = &TODO{id: count, title: task, completed: false}
	return item
}

func deleteItem(id int) {
	var i int = -1
	for index, task := range tasks {
		if task.id == id {
			var temp []TODO
			for _, element := range tasks {
				if element != task {
					temp = append(temp, element)
					i = index
				} else {
					fmt.Printf("Succesfully Deleted Task: %s", tasks[i].title)
				}
			}
			tasks = temp
		}
	}
	if i == -1 {
		fmt.Printf("No Task found with id %d", i)
	}
}

func displayTasks() {
	if len(tasks) == 0 {
		fmt.Println()
		fmt.Println("No Tasks Available ")
		fmt.Println()

	} else {
		for _, task := range tasks {
			fmt.Printf("\nid: %d, task: %s, completed: %v \n", task.id, task.title, task.completed)
		}
	}
}

func markComplete(id int) {
	i := -1
	for index, task := range tasks {
		if task.id == id {
			tasks[index].completed = true
			i = index
			fmt.Printf("\nUpdated task: %s to completed\n", tasks[index].title)
		}
	}
	if i == -1 {
		fmt.Printf("\nNo Task found with id %d\n", i)
	}
}
