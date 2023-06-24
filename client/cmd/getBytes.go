package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// getBytesCmd represents the getBytesCmd command
var getBytesCmd = &cobra.Command{
	Use:   "getBytes",
	Short: "Показать текст",
	Long:  `Возвращает байты по переданному имени`,
	Run: func(cmd *cobra.Command, args []string) {
		var bytesService = services.NewByteService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := bytesService.Get(name)
		if err != nil {
			println(err.Error())
		}

		if statusCode == 200 {
			println(result)
		} else if statusCode == 400 {
			fmt.Println("Что то пошло не так. Проверьте правильность ввода данных")
		} else if statusCode == 401 {
			fmt.Println("Ошибка авторизации. Попробуйте войти еще раз")
		}
	},
}

func init() {
	getBytesCmd.Flags().StringP("name", "n", "", "Название текста для его получения")

	rootCmd.AddCommand(getBytesCmd)
}
