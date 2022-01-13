package main

import (
	"fmt"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/kprc/netdev/tools/excel"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"strconv"
)

var param = struct {
	version bool
	excelFile  string
	sheet      string
	mysqlHost string
	mysqlUser string
	mysqlPasswd  string
	mysqlPort string
	mysqlDbName string
}{}

var rootCmd = &cobra.Command{
	Use: "nd-tools",

	Short: "netdev node for service node",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

var excelCmd = &cobra.Command{
	Use: "excel",

	Short: "excel sub command",

	Long: `excel sub command`,

	Args: cobra.NoArgs,
}
var parseExcelCmd = &cobra.Command{
	Use: "parse",

	Short: "test to parse excel",

	Long: `test to parse excel`,

	Run: parseExcel,
}

var inserExcel2DbCmd = &cobra.Command{
	Use: "insert2db",

	Short: "insert excel content to mysql databases",

	Long: `insert excel content to mysql databases`,

	Run: insertExcel,
}

const (
	Version = "0.0.1"

)
func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "nd-tools -v")


	parseExcelCmd.Flags().StringVarP(&param.excelFile,"excel","e","","excel file name")
	parseExcelCmd.Flags().StringVarP(&param.sheet,"sheet","s","","excel sheet")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlPasswd,"passwd","c","","mysql password")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlHost,"host","d","","mysql host name")
	inserExcel2DbCmd.Flags().StringVarP(&param.excelFile,"excel","e","","excel file name")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlPort,"port","p","3306","mysql service port")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlUser,"user","u","","mysql user name")
	inserExcel2DbCmd.Flags().StringVarP(&param.sheet,"sheet","s","","excel sheet")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlDbName,"database","r","","mysql database name")

	rootCmd.AddCommand(excelCmd)
	excelCmd.AddCommand(parseExcelCmd,inserExcel2DbCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func parseExcel(_ *cobra.Command, _ []string)  {
	if param.excelFile == ""{
		fmt.Println("please input excel file")
		return
	}

	if b:=tools.FileExists(param.excelFile);!b{
		fmt.Println("excel file not exists")
		return
	}

	f, err := excelize.OpenFile(param.excelFile)
	if err != nil {
		println(err.Error())
		return
	}

	cell, err := f.GetCellValue(param.sheet, "C1")
	if err != nil {
		println(err.Error())
		return
	}
	println(cell)

	y,m,d,ts:=excel.CellDate(cell,8,0)

	fmt.Println(y,m,d,ts)
}

func insertExcel(_ *cobra.Command, _ []string)  {
	if param.excelFile == ""{
		fmt.Println("please input excel file")
		return
	}

	if b:=tools.FileExists(param.excelFile);!b{
		fmt.Println("excel file not exists")
		return
	}

	port := 0

	if param.mysqlPort != ""{
		var err error
		if port,err = strconv.Atoi(param.mysqlPort);err!=nil{
			fmt.Println(err)
			return
		}
	}

	e:=excel.NewExcel(param.excelFile,param.sheet,param.mysqlUser,param.mysqlPasswd,param.mysqlHost,param.mysqlDbName,port)

	if err:=e.Open();err!=nil{
		fmt.Println("open excel error",err)
		return
	}
	defer e.Close()


	if err:=e.ReadAndInsert();err!=nil{
		fmt.Println(err)
		return
	}



}

func mainRun(_ *cobra.Command, _ []string) {

	if param.version {
		fmt.Println(Version)
		return
	}

}

