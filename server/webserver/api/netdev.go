package api

import (
	"encoding/json"
	"fmt"
	"github.com/kprc/netdev/config"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"github.com/kprc/netdev/server/webserver/msg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

const (
	maxUploadSize = 1 << 24 //16M
)

type WebApi struct {
	db *mysqlconn.NetDevDbConn
}

func NewWebApi(db *mysqlconn.NetDevDbConn) *WebApi {
	return &WebApi{
		db: db,
	}
}

func (wa *WebApi) FoodTower(writer http.ResponseWriter, request *http.Request) {
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

		ft := &msg.MsgFoodTower{}
		if err = json.Unmarshal(contents, ft); err != nil {
			result := PackResult(Failure, "json unmarshal failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		if err = sql.InsertFoodTower(wa.db, ft); err != nil {
			result := PackResult(Failure, "Insert into db failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}
		result := PackResult(Success, "success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}

}

func (wa *WebApi) Water(writer http.ResponseWriter, request *http.Request) {
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
		water := &msg.MsgWater{}
		if err = json.Unmarshal(contents, water); err != nil {
			result := PackResult(Failure, "json unmarshal failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		if err = sql.InsertWater(wa.db, water); err != nil {
			result := PackResult(Failure, "Insert into db failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		result := PackResult(Success, "success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func (wa *WebApi) Weigh(writer http.ResponseWriter, request *http.Request) {
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
		weigh := &msg.MsgWeigh{}
		if err = json.Unmarshal(contents, weigh); err != nil {
			result := PackResult(Failure, "json unmarshal failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		if err = sql.InsertWeigh(wa.db, weigh); err != nil {
			result := PackResult(Failure, "Insert into db failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		result := PackResult(Success, "success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func (wa *WebApi) UniPhase(writer http.ResponseWriter, request *http.Request) {
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

		uni := &msg.MsgUniphase{}
		if err = json.Unmarshal(contents, uni); err != nil {
			result := PackResult(Failure, "json unmarshal failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		if err = sql.InsertUniphase(wa.db, uni); err != nil {
			result := PackResult(Failure, "Insert into db failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		result := PackResult(Success, "success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func (wa *WebApi) Triphase(writer http.ResponseWriter, request *http.Request) {
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
		tri := &msg.MsgTriphase{}
		if err = json.Unmarshal(contents, tri); err != nil {
			result := PackResult(Failure, "json unmarshal failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		if err = sql.InsertTriphase(wa.db, tri); err != nil {
			result := PackResult(Failure, "Insert into db failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		result := PackResult(Success, "success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func (wa *WebApi) IndexSource(writer http.ResponseWriter, request *http.Request) {
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
		tri := &msg.MsgIndexSource{}
		if err = json.Unmarshal(contents, tri); err != nil {
			result := PackResult(Failure, "json unmarshal failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		if err = sql.InsertIndexSource(wa.db, tri); err != nil {
			result := PackResult(Failure, "Insert into db failed")
			writer.WriteHeader(200)
			writer.Write(result.Bytes())
			return
		}

		result := PackResult(Success, "success")
		writer.WriteHeader(200)
		writer.Write(result.Bytes())
	}
}

func (wa *WebApi) UploadFile(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(maxUploadSize)
	defer request.MultipartForm.RemoveAll()
	file, h, err := request.FormFile("filename")
	if err != nil {
		return
	}
	defer file.Close()
	//fmt.Fprintf(w, "%v", h.Header)
	f, err := os.OpenFile(path.Join(config.UploadDir(), h.Filename), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	writer.Write([]byte("success"))
	fmt.Println("save file", h.Filename, "success")
}
