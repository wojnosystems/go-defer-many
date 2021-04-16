package deferMany

type deferer struct {
	tasks []func()
}

func New() Deferrer {
	return &deferer{}
}

func (d *deferer) Add(task func()) {
	d.tasks = append(d.tasks, task)
}

func (d *deferer) Defer() {
	for _, task := range d.tasks {
		task()
	}
}

func (d *deferer) Return() func() {
	returnedTasks := d.tasks
	d.tasks = nil
	return func() {
		for _, task := range returnedTasks {
			task()
		}
	}
}
