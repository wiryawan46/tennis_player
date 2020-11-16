package tennis_player

import (
	"fmt"
	"os"
	configEnv "github.com/joho/godotenv"
)

func main() {
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}
}