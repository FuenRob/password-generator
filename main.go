package main

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[]{}|;:,.<>?`~"
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func main() {
	var rootCmd = &cobra.Command{Use: "password", Short: "Generador de contraseñas seguras"}

	var length int

	var cmdGenerate = &cobra.Command{
		Use:   "generate",
		Short: "Genera una contraseña segura",
		Run: func(cmd *cobra.Command, args []string) {
			password := generatePassword(length)
			fmt.Println("Contraseña generada:", password)
		},
	}

	cmdGenerate.Flags().IntVarP(&length, "longitud", "l", 12, "Longitud de la contraseña")

	rootCmd.AddCommand(cmdGenerate)
	rootCmd.Execute()
}
