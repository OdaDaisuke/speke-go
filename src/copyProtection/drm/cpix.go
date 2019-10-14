package drm

import "encoding/xml"

type CPIX struct {
	XMLName        xml.Name `xml:"CPIX"`
	Text           string   `xml:",chardata"`
	ID             string   `xml:"id,attr"`
	Cpix           string   `xml:"cpix,attr"`
	Pskc           string   `xml:"pskc,attr"`
	Speke          string   `xml:"speke,attr"`
	Ds             string   `xml:"ds,attr"`
	Enc            string   `xml:"enc,attr"`
	ContentKeyList struct {
		Text       string `xml:",chardata"`
		ContentKey struct {
			Text string `xml:",chardata"`
			Kid  string `xml:"kid,attr"`
		} `xml:"ContentKey"`
	} `xml:"ContentKeyList"`
	DRMSystemList struct {
		Text      string `xml:",chardata"`
		DRMSystem struct {
			Text                  string `xml:",chardata"`
			Kid                   string `xml:"kid,attr"`
			SystemId              string `xml:"systemId,attr"`
			PSSH                  string `xml:"PSSH"`
			ContentProtectionData string `xml:"ContentProtectionData"`
			URIExtXKey            string `xml:"URIExtXKey"`
			KeyFormat             string `xml:"KeyFormat"`
			KeyFormatVersions     string `xml:"KeyFormatVersions"`
			ProtectionHeader      string `xml:"ProtectionHeader"`
		} `xml:"DRMSystem"`
	} `xml:"DRMSystemList"`
}
