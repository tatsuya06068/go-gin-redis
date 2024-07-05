# go-gin-redis

package main

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func main() {
    // AWS Config の読み込み
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    // Cognito クライアントの作成
    svc := cognitoidentityprovider.NewFromConfig(cfg)

    // 新しいユーザーの作成
    userPoolID := "your-user-pool-id"
    username := "newuser"
    email := "newuser@example.com"
    password := "your-secure-password"

    input := &cognitoidentityprovider.AdminCreateUserInput{
        UserPoolId: &userPoolID,
        Username:   &username,
        UserAttributes: []cognitoidentityprovider.AttributeType{
            {
                Name:  aws.String("email"),
                Value: aws.String(email),
            },
            {
                Name:  aws.String("email_verified"),
                Value: aws.String("true"),
            },
        },
        TemporaryPassword: &password,
        MessageAction:     cognitoidentityprovider.MessageActionTypeSuppress,
    }

    result, err := svc.AdminCreateUser(context.TODO(), input)
    if err != nil {
        log.Fatalf("failed to create user: %v", err)
    }

    fmt.Printf("User created: %s\n", *result.User.Username)
}
