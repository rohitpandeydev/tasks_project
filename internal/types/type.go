package types

import "time"

type Task struct {
	ID          int       `csv:"id"`
	Description string    `csv:"description"`
	CreatedAt   time.Time `csv:"createdAt"`
	IsComplete  bool      `csv:"isComplete"`
}
