package deferMany

type deferred struct {
	tasks []func()
}

func New() Defer {
	return &deferred{}
}

func (d *deferred) Add(task func()) {
	d.tasks = append(d.tasks, task)
}

func (d *deferred) Defer() {
	for _, task := range d.tasks {
		task()
	}
}

func (d *deferred) Return() func() {
	returnedTasks := d.tasks
	d.tasks = nil
	return func() {
		for _, task := range returnedTasks {
			task()
		}
	}
}
