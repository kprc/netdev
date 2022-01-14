package main

import (
	"encoding/base64"
	"fmt"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/kprc/netdev/tools/excel"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
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
	earLabel string
	timestamp int64
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

var testExcelDbCmd = &cobra.Command{
	Use: "test",
	Short: "test Data from db",
	Long: "test Data from db",
	Run: testExcel,
}

const (
	Version = "0.0.1"

)
func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "nd-tools -v")


	parseExcelCmd.Flags().StringVarP(&param.excelFile,"excelfile","e","","excel file name")
	parseExcelCmd.Flags().StringVarP(&param.sheet,"sheet","s","","excel sheet")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlPasswd,"passwd","c","","mysql password")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlHost,"host","d","","mysql host name")
	inserExcel2DbCmd.Flags().StringVarP(&param.excelFile,"excelfile","e","","excel file name")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlPort,"port","p","3306","mysql service port")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlUser,"user","u","","mysql user name")
	inserExcel2DbCmd.Flags().StringVarP(&param.sheet,"sheet","s","","excel sheet")
	inserExcel2DbCmd.Flags().StringVarP(&param.mysqlDbName,"database","r","","mysql database name")
	testExcelDbCmd.Flags().StringVarP(&param.earLabel,"earlabel","e","","test earlabel")
	testExcelDbCmd.Flags().Int64VarP(&param.timestamp,"timestamp","t",0,"timestamp")

	rootCmd.AddCommand(excelCmd)
	excelCmd.AddCommand(parseExcelCmd,inserExcel2DbCmd,testExcelDbCmd)

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
	defer f.Close()

	rows, err := f.GetRows(param.sheet)
	if err!=nil{
		fmt.Println(err)
		return
	}

	lastDay,lastMonth := 0, 8
	for idx, colCell := range rows[0] {
		if colCell != "" && base64.StdEncoding.EncodeToString([]byte(colCell)) != "IA=="{
			fmt.Println(idx," ")
			_,lastDay,lastMonth,_=excel.CellDate(colCell,lastMonth,lastDay)
			fmt.Println("")
		}

	}

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

func testExcel(_ *cobra.Command, _ []string) {

	if param.earLabel == ""{
		fmt.Println("need ear label ...")
		return
	}

	if param.timestamp == 0{
		fmt.Println("timestamp need...")
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

	//if err:=e.Open();err!=nil{
	//	fmt.Println("open excel error",err)
	//	return
	//}
	//defer e.Close()

	t:=time.Unix(param.timestamp,0)

	if err:=e.TestExcel(param.earLabel,t);err!=nil{
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

