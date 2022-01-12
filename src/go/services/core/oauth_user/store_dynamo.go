package user

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
// 	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
// 	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
// 	"github.com/google/uuid"
// 	oauth_provider "github.com/koblas/grpc-todo/services/core/oauth_user/provider"
// 	"golang.org/x/net/context"
// )

// type oauthDynamoStore struct {
// 	client *dynamodb.Client
// 	table  *string
// }

// func NewOauthDynamoStore() OAuthStore {
// 	cfg, err := config.LoadDefaultConfig(context.TODO())

// 	if err != nil {
// 		panic(err)
// 	}

// 	client := dynamodb.NewFromConfig(cfg)

// 	return &oauthDynamoStore{
// 		client: client,
// 		table:  aws.String("app-user"),
// 	}
// }

// func (store *oauthDynamoStore) ListByUserId(userId string) ([]OauthUser, error) {
// 	// TODO - Doesn't handle paging, but should typically be a small enough result set (~5 items) that it doesn't matter
// 	out, err := store.client.Query(context.TODO(), &dynamodb.QueryInput{
// 		TableName:              store.table,
// 		KeyConditionExpression: aws.String("user_id = :user_id"),
// 		ExpressionAttributeValues: map[string]types.AttributeValue{
// 			":user_id": &types.AttributeValueMemberS{Value: userId},
// 		},
// 	})

// 	if err != nil || out == nil || len(out.Items) == 0 {
// 		return nil, err
// 	}

// 	result := []OauthUser{}
// 	for _, item := range out.Items {
// 		oauth := OauthUser{}
// 		if err := attributevalue.UnmarshalMap(item, &oauth); err != nil {
// 			return nil, err
// 		}

// 		result = append(result, oauth)
// 	}

// 	return result, nil
// }

// func (store *oauthDynamoStore) Associate(userId string, provider string, providerId string, token oauth_provider.TokenResult) error {
// 	if err := store.Remove(userId, provider); err != nil {
// 		return err
// 	}

// 	record := OauthUser{
// 		ID:             uuid.NewString(),
// 		UserId:         userId,
// 		Provider:       provider,
// 		ProviderId:     providerId,
// 		ProviderMerged: fmt.Sprintf("provider#%s#%s", provider, providerId),
// 		AccessToken:    token.AccessToken,
// 		RefreshToken:   token.RefreshToken,
// 		ExpiresAt:      *token.Expires,
// 	}

// 	av, err := attributevalue.MarshalMap(record)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal Record, %w", err)
// 	}

// 	_, err = store.client.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
// 		TransactItems: []types.TransactWriteItem{
// 			{
// 				Put: &types.Put{
// 					TableName: store.table,
// 					Item:      av,
// 				},
// 			},
// 			{
// 				Put: &types.Put{
// 					TableName: store.table,
// 					Item: map[string]types.AttributeValue{
// 						"id":          &types.AttributeValueMemberS{Value: record.ProviderMerged},
// 						"user_id":     &types.AttributeValueMemberS{Value: "user_id#" + record.UserId},
// 						"provider":    &types.AttributeValueMemberS{Value: record.Provider},
// 						"provider_id": &types.AttributeValueMemberS{Value: record.ProviderId},
// 					},
// 				},
// 			},
// 		},
// 	})

// 	_, err = store.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
// 		TableName: store.table,
// 		Item:      av,
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (store *oauthDynamoStore) Remove(userId string, provider string) error {
// 	items, found := store.database[userId]
// 	if !found {
// 		return nil
// 	}

// 	for idx, item := range items {
// 		if item.UserId == userId && item.Provider == provider {
// 			store.database[userId] = append(items[:idx], items[idx+1:]...)
// 			return nil
// 		}
// 	}

// 	return nil
// }

// func (store *oauthDynamoStore) FindByProviderId(provider string, providerId string) (*OauthUser, error) {
// 	out, err := store.client.Query(context.TODO(), &dynamodb.QueryInput{
// 		TableName:              store.table,
// 		IndexName:              aws.String("app-oauth-by-provider"),
// 		KeyConditionExpression: aws.String("provider = :provider AND provider_id = :provider_id"),
// 		ExpressionAttributeValues: map[string]types.AttributeValue{
// 			":provider":    &types.AttributeValueMemberS{Value: provider},
// 			":provider_id": &types.AttributeValueMemberS{Value: providerId},
// 		},
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(out.Items) == 0 {
// 		return nil, nil
// 	}

// 	oauth := OauthUser{}
// 	if err := attributevalue.UnmarshalMap(out.Items[0], &oauth); err != nil {
// 		return nil, err
// 	}

// 	return &oauth, nil
// }
