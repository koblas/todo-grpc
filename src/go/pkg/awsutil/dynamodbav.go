package awsutil

// Copyright 2023 Ryan Collingham
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// “Software”), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR
// ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
// CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package dynamodbav provides Marshal/Unmarshal utilities for DynamoDB,
// intended to complement the AWS Go SDK V2 [github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue].

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

// MarshalItem marshals a generic type into its DB representation. This
// function is just a simple wrapper around attributevalue.MarshalMap() but is
// included for completeness.
func MarshalItem(v interface{}) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMap(v)
	if err != nil {
		return nil, errors.Wrap(err, "attributevalue.MarshalMap")
	}

	return av, nil
}

// MarshalList marshals a list of values into a list of their DB representations.
func MarshalList[T any](vals []T) ([]map[string]types.AttributeValue, error) {
	results := make([]map[string]types.AttributeValue, len(vals))

	for i := range vals {
		item, err := MarshalItem(vals[i])
		if err != nil {
			return nil, err
		}

		results[i] = item
	}

	return results, nil
}

// UnmarshalItem unmarshals a value from its DB represenations. The type to
// return must be specified as a generic parameter, for example:
//
//	val, err := dynamodbav.UnmarshalItem[MyType](out.Item)
func UnmarshalItem[T any](item map[string]types.AttributeValue) (*T, error) {
	result := new(T)

	if err := attributevalue.UnmarshalMap(item, result); err != nil {
		return nil, errors.Wrap(err, "attributevalue.UnmarshalMap")
	}

	return result, nil
}

// UnmarshalList unmarshals a list of values from their DB representations.
// The type to return must be specified as a generic parameter, for example:
//
//	vals, err := dynamodbav.UnmarshalList[MyType](out.Items)
func UnmarshalList[T any](items []map[string]types.AttributeValue) ([]T, error) {
	var results []T
	if err := attributevalue.UnmarshalListOfMaps(items, &results); err != nil {
		return nil, errors.Wrap(err, "attributevalue.UnmarshalListOfMaps")
	}

	return results, nil
}
