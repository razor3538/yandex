package cmd

import (
	"client/internal/services"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteBytesCmd represents the deleteBytesCmd command
var deleteBytesCmd = &cobra.Command{
	Use:   "deleteBytesCmd",
	Short: "Удаление текста",
	Long:  `Удаление байтов по переданному имени и по пользователю`,
	Run: func(cmd *cobra.Command, args []string) {
		var bytesService = services.NewByteService()

		name, _ := cmd.Flags().GetString("name")

		result, statusCode, err := bytesService.Delete(name)
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
	deleteBytesCmd.Flags().StringP("name", "n", "", "Название текста для его удаления")

	rootCmd.AddCommand(deleteBytesCmd)
}
