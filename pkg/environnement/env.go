package environnement

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(envType string) string {
    err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		//create a empty .env file
		os.Create(".env")


	}

	switch envType {
		case "DB_USER":
			if val := os.Getenv("DB_USER"); val != "" {
            return val
        }
        return "user"
		case "DB_PASSWORD":
			if val := os.Getenv("DB_PASSWORD"); val != "" {
				return val
			}
			return "password"
		case "DB_HOST":
			if val := os.Getenv("DB_HOST"); val != "" {
				return val
			}
			return "127.0.0.1"
		case "DB_PORT":
			if val := os.Getenv("DB_PORT"); val != "" {
				return val
			}
			return "3306"
		case "DB_NAME":
			if val := os.Getenv("DB_NAME"); val != "" {
				return val
			}
			return "database"
		case "MAIL_SENDER":
			if val := os.Getenv("MAIL_SENDER"); val != "" {
				return val
			}
			return "noreply@grocery-store.online"
		case "MAIL_SERVER":
			if val := os.Getenv("MAIL_SERVER"); val != "" {
				return val
			}
			return "127.0.0.1"
		case "MAIL_PORT":
			if val := os.Getenv("MAIL_PORT"); val != "" {
				return val
			}
			return "1025"
		default:
			return ""
	}
}
