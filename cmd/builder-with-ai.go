package main

import (
	"fmt"
	cfg "github.com/iamvkosarev/builder-with-ai/internal/config"
)

func main() {
	config := cfg.MustLoad()
	fmt.Println(config.CloneRepositoryDir)
}
