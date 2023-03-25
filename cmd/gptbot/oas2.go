// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package main

import (
	"reflect"

	"github.com/RussellLuo/kun/pkg/httpcodec"
	"github.com/RussellLuo/kun/pkg/oas2"
	"github.com/go-aie/gptbot"
)

var (
	base = `swagger: "2.0"
info:
  title: "GPTBot-API"
  version: "1.0.0"
  description: "Hi here! This is the API documentation for GPTBot.\n//"
  license:
    name: "MIT"
host: "example.com"
basePath: "/"
schemes:
  - "https"
consumes:
  - "application/json"
produces:
  - "application/json"
`

	paths = `
paths:
  /chat:
    post:
      description: "Chat sends question to the bot for an answer."
      operationId: "Chat"
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/ChatRequestBody"
      %s
  /upsert:
    post:
      description: "CreateDocuments feeds documents into the vector store."
      operationId: "CreateDocuments"
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/CreateDocumentsRequestBody"
      %s
  /delete:
    post:
      description: "DeleteDocuments deletes the specified documents from the vector store."
      operationId: "DeleteDocuments"
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/DeleteDocumentsRequestBody"
      %s
  /upload:
    post:
      description: "UploadFile uploads a file and then feeds the text into the vector store."
      operationId: "UploadFile"
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/UploadFileRequestBody"
      %s
`
)

func getResponses(schema oas2.Schema) []oas2.OASResponses {
	return []oas2.OASResponses{
		oas2.GetOASResponses(schema, "Chat", 200, &ChatResponse{}),
		oas2.GetOASResponses(schema, "CreateDocuments", 200, &CreateDocumentsResponse{}),
		oas2.GetOASResponses(schema, "DeleteDocuments", 200, &DeleteDocumentsResponse{}),
		oas2.GetOASResponses(schema, "UploadFile", 200, &UploadFileResponse{}),
	}
}

func getDefinitions(schema oas2.Schema) map[string]oas2.Definition {
	defs := make(map[string]oas2.Definition)

	oas2.AddDefinition(defs, "ChatRequestBody", reflect.ValueOf(&struct {
		Question string `json:"question"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "Chat", 200, (&ChatResponse{}).Body())

	oas2.AddDefinition(defs, "CreateDocumentsRequestBody", reflect.ValueOf(&struct {
		Documents []*gptbot.Document `json:"documents"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "CreateDocuments", 200, (&CreateDocumentsResponse{}).Body())

	oas2.AddDefinition(defs, "DeleteDocumentsRequestBody", reflect.ValueOf(&struct {
		DocumentIds []string `json:"document_ids"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "DeleteDocuments", 200, (&DeleteDocumentsResponse{}).Body())

	oas2.AddDefinition(defs, "UploadFileRequestBody", reflect.ValueOf(&struct {
		File *httpcodec.FormFile `json:"file"`
	}{}))
	oas2.AddResponseDefinitions(defs, schema, "UploadFile", 200, (&UploadFileResponse{}).Body())

	return defs
}

func OASv2APIDoc(schema oas2.Schema) string {
	resps := getResponses(schema)
	paths := oas2.GenPaths(resps, paths)

	defs := getDefinitions(schema)
	definitions := oas2.GenDefinitions(defs)

	return base + paths + definitions
}