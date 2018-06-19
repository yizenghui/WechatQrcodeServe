package orm

import "time"

// Task 任务
type Task struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"size:255;sql:index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// NewTask 新增一个任务
func (task *Task) NewTask() {
	DB().Create(task)
}

// GetTaskByID 获取任务
func (task *Task) GetTaskByID(id int) {
	DB().First(task, id)
}

// Save 保存信息
func (task *Task) Save() {
	DB().Save(&task)
}
