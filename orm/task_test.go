// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package orm

import (
	"testing"
)

func init() {

	DB().AutoMigrate(&Task{})
}

func Test_CreateTask(t *testing.T) {
	var task Task
	task.NewTask()
	// user.GetUserFollowBooks()
	t.Fatal(task)
}

func Test_GetTask(t *testing.T) {
	var task Task
	DB().First(&task, 1)
	// user.GetUserFollowBooks()
	t.Fatal(task)
}
