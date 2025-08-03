package openrouter

import "os"

func GetOpenRouterHeaders() map[string]string {
	headers := make(map[string]string)
	headers["HTTP-Referer"] = os.Getenv("OPENROUTER_HTTP_REFERER")
	headers["X-Title"] = os.Getenv("OPENROUTER_X_TITLE")
	return headers
}
