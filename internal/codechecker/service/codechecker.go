package service

import (
	"bytes"
	"fmt"
	"ginProject1/internal/codechecker/repository"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"sync"
)

type IService interface {
	CheckCode(code string, tests []map[string]string) (bool, error)
}

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{repository: repository}
}

func (service *Service) CheckCode(code string, tests []map[string]string) (bool, error) {
	fileName := fmt.Sprintf("main_%s.py", uuid.New().String())
	file, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	defer func() {
		file.Close()
		os.Remove(fileName)
	}()
	_, err = file.WriteString(code)
	if err != nil {
		return false, err
	}

	ch := make(chan bool)
	defer close(ch)
	wg := &sync.WaitGroup{}
	wg.Add(len(tests))

	for i := 0; i < len(tests); i++ {
		go func() {
			defer wg.Done()
			cmd := exec.Command("python3", fileName)
			var stdin bytes.Buffer

			for key, value := range tests[i] {
				if key != "answer" {
					stdin.WriteString(value)
				}
			}
			cmd.Stdin = &stdin
			output, err := cmd.CombinedOutput()
			if err != nil {
				ch <- false

				return
			}
			ch <- string(output) == tests[i]["answer"]
		}()
	}

	ok := true
	for i := 0; i < len(tests); i++ {
		if <-ch == false {
			fmt.Println("Тут")
			ok = false
		}
	}
	wg.Wait()

	return ok, nil
}
