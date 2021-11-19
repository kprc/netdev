package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FoodTower(writer http.ResponseWriter, request *http.Request){
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {
		fmt.Println(string(contents))
		//todo...


		result:=PackResult(Success,"success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}

}

func Water(writer http.ResponseWriter, request *http.Request)  {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {
		fmt.Println(string(contents))
		//todo...


		result:=PackResult(Success,"success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func Weigh(writer http.ResponseWriter, request *http.Request)  {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {
		fmt.Println(string(contents))
		//todo...


		result:=PackResult(Success,"success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func UniPhase(writer http.ResponseWriter, request *http.Request)  {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {
		fmt.Println(string(contents))
		//todo...


		result:=PackResult(Success,"success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func Triphase(writer http.ResponseWriter, request *http.Request)  {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {
		fmt.Println(string(contents))
		//todo...


		result:=PackResult(Success,"success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}