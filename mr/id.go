package mr

import (
	"strings"

	"github.com/google/uuid"
)

// NewId 生成一个唯一的id
func NewId() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
