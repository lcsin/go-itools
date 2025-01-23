package ideepseek

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	url := "https://api.deepseek.com"
	apikey := os.Getenv("DeepSeek_APIKEY")
	client := NewClient(url, apikey).WithModel(ModelR1)

	t.Log(client)
}
