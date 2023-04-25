// Code generated by kun; DO NOT EDIT.
// github.com/RussellLuo/kun

package main

import (
	"context"
	"net/http"

	"github.com/RussellLuo/kun/pkg/httpcodec"
	"github.com/RussellLuo/kun/pkg/httpoption"
	"github.com/RussellLuo/kun/pkg/oas2"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPRouter(svc Service, codecs httpcodec.Codecs, opts ...httpoption.Option) chi.Router {
	r := chi.NewRouter()
	options := httpoption.NewOptions(opts...)

	r.Method("GET", "/api", oas2.Handler(OASv2APIDoc, options.ResponseSchema()))

	var codec httpcodec.Codec
	var validator httpoption.Validator
	var kitOptions []kithttp.ServerOption

	codec = codecs.EncodeDecoder("Chat")
	validator = options.RequestValidator("Chat")
	r.Method(
		"POST", "/chat",
		kithttp.NewServer(
			MakeEndpointOfChat(svc),
			decodeChatRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("CreateDocuments")
	validator = options.RequestValidator("CreateDocuments")
	r.Method(
		"POST", "/upsert",
		kithttp.NewServer(
			MakeEndpointOfCreateDocuments(svc),
			decodeCreateDocumentsRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("DebugSplitDocument")
	validator = options.RequestValidator("DebugSplitDocument")
	r.Method(
		"POST", "/debug/split",
		kithttp.NewServer(
			MakeEndpointOfDebugSplitDocument(svc),
			decodeDebugSplitDocumentRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("DeleteDocuments")
	validator = options.RequestValidator("DeleteDocuments")
	r.Method(
		"POST", "/delete",
		kithttp.NewServer(
			MakeEndpointOfDeleteDocuments(svc),
			decodeDeleteDocumentsRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("UploadFile")
	validator = options.RequestValidator("UploadFile")
	r.Method(
		"POST", "/upload",
		kithttp.NewServer(
			MakeEndpointOfUploadFile(svc),
			decodeUploadFileRequest(codec, validator),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(kitOptions,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	return r
}

func decodeChatRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req ChatRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeCreateDocumentsRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req CreateDocumentsRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeDebugSplitDocumentRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req DebugSplitDocumentRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeDeleteDocumentsRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req DeleteDocumentsRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeUploadFileRequest(codec httpcodec.Codec, validator httpoption.Validator) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req UploadFileRequest

		if err := codec.DecodeRequestBody(r, &_req); err != nil {
			return nil, err
		}

		if err := validator.Validate(&_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}
