package mockup

import "github.com/kprc/nbsnetwork/tools/httputil"

func NewIndexSource() error {
	hp := httputil.NewHttpPost(nil,true,2,2)
	hp.ProtectPost("url","data")

	return nil
}


