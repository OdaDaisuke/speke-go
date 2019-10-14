package drm

import (
	"encoding/base64"

	"github.com/OdaDaisuke/speke_go/src/copyProtection/util"
)

// GeneratePSSHにヘッダをくっつけPSSH BOXを生成する。
// ISO 23001-7参照。
// mp4 full boxの形式で返す。
func GeneratePSSH(psshData []byte, systemID string) string {
	strLen := len(psshData)
	dataSize := util.ToHex(strLen, false)
	payloadBinary := util.StrToHexBinStr(systemID) + util.StrToHexBinStr(dataSize) + string(psshData)
	psshType := "pssh"
	version := util.ToHex(0, true)
	flag := util.ToHexWithDigit(0, 3, true)
	boxSize := 4 + len(psshType) + len(version) + len(flag) + len(payloadBinary)
	headerBinary := util.ToHex(boxSize, true) + psshType + version + flag
	var psshBox []byte
	psshBox = append(psshBox, []byte(headerBinary)...)
	psshBox = append(psshBox, []byte(payloadBinary)...)
	return base64.StdEncoding.EncodeToString(psshBox)
}
