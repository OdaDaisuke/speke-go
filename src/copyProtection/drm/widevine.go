package drm

import (
	pb "github.com/OdaDaisuke/speke_go/src/copyProtection/pb"
	"github.com/golang/protobuf/proto"
)

type Widevine struct {
}

func NewWidevine() *Widevine {
	return &Widevine{}
}

func (w *Widevine) GetPSSHBody(kid string, contentInfo string) ([]byte, error) {
	bkid := []byte(kid)
	bbkid := [][]byte{bkid}
	p := &pb.WidevinePsshData{
		KeyId:     bbkid,
		ContentId: []byte(contentInfo),
	}
	res, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return res, nil
}
