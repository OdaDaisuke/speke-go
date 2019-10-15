package drm

import (
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io"
	"log"

	config "github.com/OdaDaisuke/speke_go/src/copyProtection/config"
	"github.com/OdaDaisuke/speke_go/src/copyProtection/util"
)

type PlayReady struct {
	wrmHeader []byte
	pro       []byte
	config    *config.AppConfig
}

func NewPlayReady(c *config.AppConfig) *PlayReady {
	return &PlayReady{
		config: c,
	}
}

func (p *PlayReady) GenerateContentKey(kid string) string {
	bSeed := util.StrToHexBinStr(p.config.KeySeed)
	bkid := util.StrToHexBinStr(kid)
	fmt.Println(bSeed, bkid)
	s := sha256.New()
	io.WriteString(s, bSeed)
	io.WriteString(s, bkid)
	shaA := fmt.Sprintf("%x", s.Sum(nil))

	s = sha256.New()
	io.WriteString(s, bSeed)
	io.WriteString(s, bkid)
	io.WriteString(s, bSeed)
	shaB := fmt.Sprintf("%x", s.Sum(nil))

	s = sha256.New()
	io.WriteString(s, bSeed)
	io.WriteString(s, bkid)
	io.WriteString(s, bSeed)
	io.WriteString(s, bkid)
	shaC := fmt.Sprintf("%x", s.Sum(nil))

	var contentKey string
	for i := 0; i < 16; i++ {
		x := shaA[i] ^ shaA[i+16] ^ shaB[i] ^ shaB[i+16] ^ shaC[i] ^ shaC[i+16]
		contentKey += util.StrToHexBinStr(string(x))
	}
	return contentKey
}

func (p *PlayReady) GenerateWRMHeader(kid string, contentID string) {
	var err error
	wrm := &WRMHEADER{}
	p.wrmHeader, err = xml.Marshal(wrm)
	if err != nil {
		log.Print(err)
	}
}

func (p *PlayReady) ComputeChecksum() {
}

func (p *PlayReady) GeneratePlayReadyObject() {
	var b []byte
	b = p.wrmHeader
	p.pro = b
}

type WRMHEADER struct {
	XMLName xml.Name `xml:"WRMHEADER"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Version string   `xml:"version,attr"`
	DATA    struct {
		Text        string `xml:",chardata"`
		PROTECTINFO struct {
			Text   string `xml:",chardata"`
			ALGID  string `xml:"ALGID"`
			KEYLEN string `xml:"KEYLEN"`
		} `xml:"PROTECTINFO"`
		KID              string `xml:"KID"`
		CHECKSUM         string `xml:"CHECKSUM"`
		LAURL            string `xml:"LA_URL"`
		CUSTOMATTRIBUTES struct {
			Text          string `xml:",chardata"`
			IISDRMVERSION string `xml:"IIS_DRM_VERSION"`
		} `xml:"CUSTOMATTRIBUTES"`
	} `xml:"DATA"`
}
