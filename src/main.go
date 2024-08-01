package main

import "github.com/ribeirosaimon/aergia/internal/config"

func main() {
	config.NewAergiaServer(&config.AergiaConfig{
		ApiPort: ":8080",
	})
}
