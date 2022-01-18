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

var insertElectricityCmd = &cobra.Command{
	Use: "insertElectricity",

	Short: "insert excel content of electricity to mysql databases",

	Long: `insert excel content of electricity to mysql databases`,

	Run: insertElectricity,
}

var insertFodderCmd = &cobra.Command{
	Use: "insertFodder",

	Short: "insert fodder usage to mysql databases",

	Long: `insert fodder usage to mysql databases`,

	Run: insertFodder,
}

var insertWaterCmd = &cobra.Command{
	Use: "insertWater",

	Short: "insert water usage to mysql databases",

	Long: `insert water usage to mysql databases`,

	Run: insertWater,
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


	insertElectricityCmd.Flags().StringVarP(&param.mysqlPasswd,"passwd","c","","mysql password")
	insertElectricityCmd.Flags().StringVarP(&param.mysqlHost,"host","d","","mysql host name")
	insertElectricityCmd.Flags().StringVarP(&param.excelFile,"excelfile","e","","excel file name")
	insertElectricityCmd.Flags().StringVarP(&param.mysqlPort,"port","p","3306","mysql service port")
	insertElectricityCmd.Flags().StringVarP(&param.mysqlUser,"user","u","","mysql user name")
	insertElectricityCmd.Flags().StringVarP(&param.sheet,"sheet","s","","excel sheet")
	insertElectricityCmd.Flags().StringVarP(&param.mysqlDbName,"database","r","","mysql database name")

	insertFodderCmd.Flags().StringVarP(&param.mysqlPasswd,"passwd","c","","mysql password")
	insertFodderCmd.Flags().StringVarP(&param.mysqlHost,"host","d","","mysql host name")
	insertFodderCmd.Flags().StringVarP(&param.excelFile,"excelfile","e","","excel file name")
	insertFodderCmd.Flags().StringVarP(&param.mysqlPort,"port","p","3306","mysql service port")
	insertFodderCmd.Flags().StringVarP(&param.mysqlUser,"user","u","","mysql user name")
	insertFodderCmd.Flags().StringVarP(&param.sheet,"sheet","s","","excel sheet")
	insertFodderCmd.Flags().StringVarP(&param.mysqlDbName,"database","r","","mysql database name")


	insertWaterCmd.Flags().StringVarP(&param.mysqlPasswd,"passwd","c","","mysql password")
	insertWaterCmd.Flags().StringVarP(&param.mysqlHost,"host","d","","mysql host name")
	insertWaterCmd.Flags().StringVarP(&param.excelFile,"excelfile","e","","excel file name")
	insertWaterCmd.Flags().StringVarP(&param.mysqlPort,"port","p","3306","mysql service port")
	insertWaterCmd.Flags().StringVarP(&param.mysqlUser,"user","u","","mysql user name")
	insertWaterCmd.Flags().StringVarP(&param.sheet,"sheet","s","","excel sheet")
	insertWaterCmd.Flags().StringVarP(&param.mysqlDbName,"database","r","","mysql database name")


	testExcelDbCmd.Flags().StringVarP(&param.earLabel,"earlabel","e","","test earlabel")
	testExcelDbCmd.Flags().Int64VarP(&param.timestamp,"timestamp","t",0,"timestamp")

	rootCmd.AddCommand(excelCmd)
	excelCmd.AddCommand(parseExcelCmd,insertElectricityCmd,testExcelDbCmd,insertFodderCmd,insertWaterCmd)

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

func insertElectricity(_ *cobra.Command, _ []string)  {
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


	if err:=e.ReadAndInsertElectricity();err!=nil{
		fmt.Println(err)
		return
	}

}


func insertFodder(_ *cobra.Command, _ []string)  {
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


	if err:=e.ReadAndInsertFodder();err!=nil{
		fmt.Println(err)
		return
	}

}

func insertWater(_ *cobra.Command, _ []string)  {
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


	if err:=e.ReadAndInsertWater();err!=nil{
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

