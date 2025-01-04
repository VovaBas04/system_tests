package neuron_model

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"ginProject1/internal/tasks/config"
	"ginProject1/pkg/database/cache"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var Format = "Входные данные: переменная = значение - Выходные данные: результат. Входные данные разделяй запятой."

type GigaChatRequest struct {
	model     string
	maxTokens int
	prompt    string
}

func NewGigaChatRequest() *GigaChatRequest {
	return &GigaChatRequest{}
}

func (gigaChat *GigaChatRequest) GeneratePrompt(task string, constraints map[string]map[Constraint]string) error {
	var conditions string
	for variable, constraint := range constraints {
		for operator, value := range constraint {
			conditions += fmt.Sprintf("%s %s %s, ", variable, operator, value)
		}
	}
	count, err := config.GetTestCases()
	if err != nil {
		return err
	}

	gigaChat.prompt = fmt.Sprintf("Сгенерируй %d наборов тестовых данных для проверки задачи в формате: %s. Текст данной задачи - %s. Ограничения на входные переменные - %s", count, Format, task, conditions)
	return nil
}

func (gigaChat *GigaChatRequest) SetModel(model string) error {
	gigaChat.model = model

	return nil
}

func (gigaChat *GigaChatRequest) SetAdditionalFields(args ...interface{}) error {
	if _, ok := args[0].(int); !ok || len(args) > 1 {
		return errors.New("neuron gigachat requires an integer argument")
	}
	gigaChat.maxTokens = args[0].(int)

	return nil
}
func (gigaChat *GigaChatRequest) Do() (*http.Response, error) {
	const urlApi = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	body := map[string]interface{}{
		"model": gigaChat.model,
		"messages": []interface{}{
			map[string]interface{}{
				"role":    "user",
				"content": gigaChat.prompt,
			},
		},
		"max_tokens":         gigaChat.maxTokens,
		"n":                  1,
		"stream":             false,
		"repetition_penalty": 1,
		"update_interval":    0,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, urlApi, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	token, err := GetToken()
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client.Do(request)
}

func (gigaChat *GigaChatRequest) PreparedRequest(task string, constraints map[string]map[Constraint]string) error {
	err := gigaChat.SetModel("GigaChat")
	if err != nil {
		return err
	}
	err = gigaChat.SetAdditionalFields(500)
	if err != nil {
		return err
	}

	err = gigaChat.GeneratePrompt(task, constraints)
	if err != nil {
		return err
	}
	return nil
}

func GetToken() (string, error) {
	cacheKey := "token"
	redisDb, err := cache.Connect()
	if err != nil {
		return "", err
	}

	defer func(redisDb *cache.RedisDb) {
		err := redisDb.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(redisDb)

	ctx := context.Background()
	cacheToken, err := redisDb.Get(ctx, cacheKey).Result()
	if err == nil {
		return cacheToken, nil
	}
	if !errors.Is(err, redis.Nil) {
		return "", err
	}

	const urlToken = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	urlValues := url.Values{}
	urlValues.Add("scope", "GIGACHAT_API_PERS")
	body := urlValues.Encode()

	request, err := http.NewRequest(http.MethodPost, urlToken, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")

	uuid4, _ := uuid.NewUUID()
	request.Header.Set("RqUID", uuid4.String())

	token, err := config.GetBasicToken()
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return "", errors.New(string(data))
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(data, &tokenResponse)
	if err != nil {
		return "", err
	}

	redisDb.Set(ctx, cacheKey, tokenResponse.AccessToken, 30*time.Minute)
	return tokenResponse.AccessToken, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expire_at"`
}
