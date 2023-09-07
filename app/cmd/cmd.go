package cmd

import (
	"api-doc/app/http"
	"api-doc/app/lib/config"
	"api-doc/app/model"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

var Main = cobra.Command{
	Use: "main",
	Run: func(cmd *cobra.Command, args []string) {
		http.StartWebServer() //开启web server服务
	},
}

var tableName string
var GormDto = cobra.Command{
	Use: "gorm:dto",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GormDto command run begin")
		if tableName == "" {
			panic("要生成的表名称必须输入")
		}
		fmt.Printf("gen table [%s] struct\r\n", tableName)
		db := model.Instance()
		g := gen.NewGenerator(gen.Config{
			OutPath:      "./app/model/query",
			ModelPkgPath: "./dto",
			Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		})
		dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
			"bigint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
			"int":    func(detailType gorm.ColumnType) (dataType string) { return "int" },
			"tinyint": func(detailType gorm.ColumnType) (dataType string) {
				return "int"
			},
		}

		g.WithDataTypeMap(dataMap)
		g.UseDB(db)
		wt := g.GenerateModel(tableName)
		g.ApplyBasic(wt)
		g.Execute()
		path, _ := os.Getwd()
		oldFile := fmt.Sprintf("%s/app/model/dto/%s.gen.go", path, tableName)
		newFile := fmt.Sprintf("%s/app/model/dto/%s.go", path, tableName)
		err := os.Rename(oldFile, newFile)
		if err != nil {
			fmt.Println("rename err ", err)
		}
	},
}

// init cmd包初始化函数
func init() {
	//初始化配置
	config.Init()

	//配置GormDto命令的选项参数
	GormDto.Flags().StringVarP(&tableName, "table", "t", "", "要生成dto的表名称")
}
