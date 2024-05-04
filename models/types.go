package models

import (
	"assignment_week17/db"
)

type ApiError struct {
	Error string `json:"error"`
}

type MongoResponse struct {
	Code       int             `json:"code"`
	Message    string          `json:"msg"`
	DbResponse []db.DbResponse `json:"records,omitempty"`
}

func NewMongoResponse(code int, msg string, dbResponse []db.DbResponse) *MongoResponse {
	return &MongoResponse{
		Code:       code,
		Message:    msg,
		DbResponse: dbResponse,
	}
}

type InMemory struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewInmemoryResponse(key string, value string) *InMemory {
	return &InMemory{
		Key:   key,
		Value: value,
	}
}
