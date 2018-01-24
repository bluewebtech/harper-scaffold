package bootstrap

import (
	"app/config"
	"app/routes"
	"fmt"
	"github.com/bluewebtech/harper"
)

func Run() *harper.Application {
	fmt.Print(config.Configuration)

	app := harper.New()
	app.View.SetTemplateLoader("template", "resources/views/")

	return routes.Routes(app)

	/*
		for key, value := range config {
			fmt.Println("Key:", key, "Value:", value)
		}
	*/

	/*
		dir, _ := os.Getwd()
		path := filepath.FromSlash(dir + "/config")
		fileList := []string{}

		filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			file := filepath.Base(path)

			if strings.Contains(file, ".") && file[0:1] != "." {
				fileList = append(fileList, strings.Title(strings.Replace(file, ".go", "", -1)))
			}

			return nil
		})

		for _, file := range fileList {
			fmt.Print(file)
		}
	*/
}
