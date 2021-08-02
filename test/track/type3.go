package PPPPP

import "encoding/xml"

type Msg struct {
	XMLName xml.Name `xml:"msg"`
	Text    string   `xml:",chardata"`
	Img     struct {
		Text           string `xml:",chardata"`
		Aeskey         string `xml:"aeskey,attr"`
		Encryver       string `xml:"encryver,attr"`
		Cdnthumbaeskey string `xml:"cdnthumbaeskey,attr"`
		Cdnthumburl    string `xml:"cdnthumburl,attr"`
		Cdnthumblength string `xml:"cdnthumblength,attr"`
		Cdnthumbheight string `xml:"cdnthumbheight,attr"`
		Cdnthumbwidth  string `xml:"cdnthumbwidth,attr"`
		Cdnmidheight   string `xml:"cdnmidheight,attr"`
		Cdnmidwidth    string `xml:"cdnmidwidth,attr"`
		Cdnhdheight    string `xml:"cdnhdheight,attr"`
		Cdnhdwidth     string `xml:"cdnhdwidth,attr"`
		Cdnmidimgurl   string `xml:"cdnmidimgurl,attr"`
		Length         string `xml:"length,attr"`
		Cdnbigimgurl   string `xml:"cdnbigimgurl,attr"`
		Hdlength       string `xml:"hdlength,attr"`
		Md5            string `xml:"md5,attr"`
	} `xml:"img"`
} 

