package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}