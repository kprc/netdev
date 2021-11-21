package config

import (
	"bytes"
	"encoding/json"
	"github.com/kprc/nbsnetwork/tools"
	"path"
	"sync"
)

const (
	netdevHome = ".netdev"
	confFile   = "conf.json"
	cmdSock    = ".cmd-sock"
)

var (
	netconf       *NetDevConf
	netconfOnceDo sync.Once
)

func defaultConf() *NetDevConf {
	return &NetDevConf{
		WConf: &WebSeverConf{
			ListenServer: ":60010",
		},
		UConf: &UdpServerConf{
			ListenServer: ":60012",
		},
		Db: &DatabaseConf{
			Driver: "mysql",
			DbName: "netdev",
			User:   "root",
			Passwd: "Rickey@123",
		},
	}
}

func GetNetDevConf() *NetDevConf {
	if netconf == nil {
		netconfOnceDo.Do(func() {
			netconf = defaultConf()
			netconf.load()
			netconf.compareAndSave()
		})
	}

	return netconf
}

func InitConfig() {
	nc := defaultConf()
	nc.save()
}

func (nc *NetDevConf) compareAndSave() {
	data, err := json.Marshal(*nc)
	if err != nil {
		return
	}

	ncdefault := defaultConf()

	var dataDefault []byte
	dataDefault, err = json.Marshal(ncdefault)
	if err != nil {
		return
	}

	if cp := bytes.Compare(data, dataDefault); cp != 0 {
		nc.save()
		return
	}
	var dataload []byte

	ncload := &NetDevConf{}
	ncload.load()

	dataload, err = json.Marshal(ncload)
	if err != nil {
		return
	}

	if cp := bytes.Compare(dataload, dataDefault); cp != 0 {
		nc.save()
		return
	}

}

func NetDevHome() string {
	h, _ := tools.Home()

	return path.Join(h, netdevHome)
}

func ConfFile() string {
	return path.Join(NetDevHome(), confFile)
}

func CmdSockFile() string {
	return path.Join(NetDevHome(), cmdSock)
}

type WebSeverConf struct {
	ListenServer string `json:"listen_server"`
}

type UdpServerConf struct {
	ListenServer string `json:"listen_server"`
}

type DatabaseConf struct {
	Driver string `json:"driver"`
	DbName string `json:"db_name"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

type NetDevConf struct {
	WConf *WebSeverConf  `json:"w_conf"`
	UConf *UdpServerConf `json:"u_conf"`
	Db    *DatabaseConf  `json:"db"`
}

func (nc *NetDevConf) load() error {
	cfile := ConfFile()

	fdata, err := tools.OpenAndReadAll(cfile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(fdata, nc); err != nil {
		return err
	}

	return nil
}

func (nc *NetDevConf) save() error {
	cfile := ConfFile()

	if j, err := json.MarshalIndent(*nc, "\t", " "); err != nil {
		return err
	} else {
		return tools.Save2File(j, cfile)
	}
}
