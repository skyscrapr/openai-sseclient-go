# Open AI Server-sent events (SSE) client in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/skyscrapr/openai-sseclient-go.svg)](https://pkg.go.dev/github.com/skyscrapr/openai-sseclient-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/skyscrapr/openai-sseclient-go)](https://goreportcard.com/report/github.com/skyscrapr/openai-sseclient-go)
[![codecov](https://codecov.io/gh/skyscrapr/openai-sseclient-go/branch/main/graph/badge.svg?token=KG8ZTLDHS3)](https://codecov.io/gh/skyscrapr/openai-sseclient-go)
![Github Actions Workflow](https://github.com/skyscrapr/openai-sseclient-go/actions/workflows/go.yml/badge.svg)
![License](https://img.shields.io/dub/l/vibe-d.svg)


I needed streaming support for my terraform provider for OpenAI. 
I couldn't use any other availble clients due to some unusual implementation details from OpenAI.
Therefore I created this client based on the work by [@cryptix](https://github.com/cryptix/goSSEClient)
