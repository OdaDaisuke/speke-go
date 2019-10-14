package main

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/OdaDaisuke/speke_go/src/copyProtection/drm"
)

const AES128_ID = "81376844-f976-481e-a84e-cc25d39b0b33"
const WIDEVINE_ID = "edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
const PLAYREADY_ID = "9a04f079-9840-4286-ab92-e65be0885f95"

type ResponseGenerator struct {
	wv *drm.Widevine
}

func NewResponseGenerator() *ResponseGenerator {
	return &ResponseGenerator{
		wv: drm.NewWidevine(),
	}
}

func (g *ResponseGenerator) Run(xmlBody string) (string, error) {
	fmt.Println(xmlBody)

	var cpix drm.CPIX
	xml.Unmarshal([]byte(xmlBody), &cpix)
	fmt.Println(cpix)
	systemID := cpix.DRMSystemList.DRMSystem.SystemId
	kid := cpix.ContentKeyList.ContentKey.Kid
	fmt.Println("log -------------", kid, systemID)
	fmt.Println(WIDEVINE_ID)
	fmt.Println(WIDEVINE_ID == systemID)
	switch systemID {
	case AES128_ID:
		fmt.Println("AES 128")
	case WIDEVINE_ID:
		psshData, err := g.wv.GetPSSHBody(kid, "content id here")
		fmt.Println("widevine ", psshData)
		if err != nil {
			return "", err
		}
		psshBox := drm.GeneratePSSH(psshData, systemID)
		cpix.DRMSystemList.DRMSystem.PSSH = psshBox
	case PLAYREADY_ID:
		fmt.Println("pr id")
	default:
		return "", errors.New("Invalid systemID.")
	}
	resCpix, err := xml.Marshal(cpix)
	if err != nil {
		return "", err
	}
	return string(resCpix), nil
}
