package main

import (
  "github.com/aws/aws-sdk-go/service/secretsmanager"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "encoding/json"
  "fmt"
  "log"
)

const REGION = "ap-northeast-1"

func main() {
  secretName := "GO_AWS_SECRET_MANAGER"

  value, err := getSecret(secretName)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(value)
}

func getSecret(secretName string) (string, error){
  svc := secretsmanager.New(session.New(),
    aws.NewConfig().WithRegion(REGION))

  input := &secretsmanager.GetSecretValueInput{
    SecretId:     aws.String(secretName),
    VersionStage: aws.String("AWSCURRENT"),
  }

  result, err := svc.GetSecretValue(input)
  if err != nil {
    return "", err
  }

  secretString := aws.StringValue(result.SecretString)
  res := make(map[string]interface{})
  if err := json.Unmarshal([]byte(secretString), &res); err != nil {
    return "", err
  }
  return res["AWS_GO_SECRET_KEY"].(string), nil
}
