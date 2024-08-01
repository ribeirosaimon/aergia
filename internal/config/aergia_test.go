package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	go func() {
		NewAergiaServer(&AergiaConfig{
			ApiPort: ":8081",
		})
	}()

	t.Run("up server", func(t *testing.T) {
		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8081/health", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Erro ao ler o corpo da resposta: %v", err)
		}
		defer resp.Body.Close()

		var dtoHealth dto.Health

		if err = json.Unmarshal(body, &dtoHealth); err != nil {
			t.Fatalf("Erro ao decodificar JSON: %v", err)
		}

		assert.NotEmpty(t, dtoHealth.Status)
		assert.Equal(t, "up", dtoHealth.Status)
	})
}
