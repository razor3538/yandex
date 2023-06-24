package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// saveBytesCmd represents the saveBytesCmd command
var saveBytesCmd = &cobra.Command{
	Use:   "saveBytes",
	Short: "Сохранение текста",
	Long:  `Сохранение текста по имени`,
	Run: func(cmd *cobra.Command, args []string) {
		var bytesService = services.NewByteService()

		bytes, _ := cmd.Flags().GetString("bytes")
		meta, _ := cmd.Flags().GetString("meta")
		name, _ := cmd.Flags().GetString("name")

		statusCode, err := bytesService.Save(bytes, name, meta)
		if err != nil {
			println(err.Error())
		}

		if statusCode == 201 {
			fmt.Println("Сохранение текста прошло успешно")
		} else if statusCode == 400 {
			fmt.Println("Что то пошло не так. Проверьте правильность ввода данных")
		} else if statusCode == 401 {
			fmt.Println("Ошибка авторизации. Попробуйте войти еще раз")
		}
	},
}

func init() {
	saveBytesCmd.Flags().StringP("name", "n", "", "Название для бинарных данных")
	saveBytesCmd.Flags().StringP("meta", "m", "", "Мета информация")
	saveBytesCmd.Flags().StringP("bytes", "b", "", "Байты для сохранения")

	rootCmd.AddCommand(saveBytesCmd)
}
