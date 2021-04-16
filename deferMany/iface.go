package deferMany

type Deferrer interface {
	// Defer should be passed to a defer block in the method that is calling Add to add items.
	// this ensures that, should an error occur, any items put into Add are cleaned up before returning.
	Defer()

	// Add appends a new task to complete when Defer or the value of Return is called
	Add(task func())

	// Return gives back a function that will execute all of the tasks that were Add'ed
	Return() (allTasks func())
}
