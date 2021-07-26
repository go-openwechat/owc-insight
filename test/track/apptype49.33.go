package PPPPP

import "encoding/xml"

type Msg struct {
	XMLName xml.Name `xml:"msg"`
	Text    string   `xml:",chardata"`
	Appmsg  struct {
		Text        string `xml:",chardata"`
		Appid       string `xml:"appid,attr"`
		Sdkver      string `xml:"sdkver,attr"`
		Title       string `xml:"title"`
		Des         string `xml:"des"`
		Username    string `xml:"username"`
		Action      string `xml:"action"`
		Type        string `xml:"type"`
		Showtype    string `xml:"showtype"`
		Content     string `xml:"content"`
		URL         string `xml:"url"`
		Lowurl      string `xml:"lowurl"`
		Forwardflag string `xml:"forwardflag"`
		Dataurl     string `xml:"dataurl"`
		Lowdataurl  string `xml:"lowdataurl"`
		Contentattr string `xml:"contentattr"`
		Streamvideo struct {
			Text                 string `xml:",chardata"`
			Streamvideourl       string `xml:"streamvideourl"`
			Streamvideototaltime string `xml:"streamvideototaltime"`
			Streamvideotitle     string `xml:"streamvideotitle"`
			Streamvideowording   string `xml:"streamvideowording"`
			Streamvideoweburl    string `xml:"streamvideoweburl"`
			Streamvideothumburl  string `xml:"streamvideothumburl"`
			Streamvideoaduxinfo  string `xml:"streamvideoaduxinfo"`
			Streamvideopublishid string `xml:"streamvideopublishid"`
		} `xml:"streamvideo"`
		CanvasPageItem struct {
			Text          string `xml:",chardata"`
			CanvasPageXml string `xml:"canvasPageXml"`
		} `xml:"canvasPageItem"`
		Appattach struct {
			Text           string `xml:",chardata"`
			Attachid       string `xml:"attachid"`
			Cdnthumburl    string `xml:"cdnthumburl"`
			Cdnthumbmd5    string `xml:"cdnthumbmd5"`
			Cdnthumblength string `xml:"cdnthumblength"`
			Cdnthumbheight string `xml:"cdnthumbheight"`
			Cdnthumbwidth  string `xml:"cdnthumbwidth"`
			Cdnthumbaeskey string `xml:"cdnthumbaeskey"`
			Aeskey         string `xml:"aeskey"`
			Encryver       string `xml:"encryver"`
			Fileext        string `xml:"fileext"`
			Islargefilemsg string `xml:"islargefilemsg"`
		} `xml:"appattach"`
		Extinfo           string `xml:"extinfo"`
		Androidsource     string `xml:"androidsource"`
		Sourceusername    string `xml:"sourceusername"`
		Sourcedisplayname string `xml:"sourcedisplayname"`
		Commenturl        string `xml:"commenturl"`
		Thumburl          string `xml:"thumburl"`
		Mediatagname      string `xml:"mediatagname"`
		Messageaction     string `xml:"messageaction"`
		Messageext        string `xml:"messageext"`
		Emoticongift      struct {
			Text        string `xml:",chardata"`
			Packageflag string `xml:"packageflag"`
			Packageid   string `xml:"packageid"`
		} `xml:"emoticongift"`
		Emoticonshared struct {
			Text        string `xml:",chardata"`
			Packageflag string `xml:"packageflag"`
			Packageid   string `xml:"packageid"`
		} `xml:"emoticonshared"`
		Designershared struct {
			Text                 string `xml:",chardata"`
			Designeruin          string `xml:"designeruin"`
			Designername         string `xml:"designername"`
			Designerrediretcturl string `xml:"designerrediretcturl"`
		} `xml:"designershared"`
		Emotionpageshared struct {
			Text      string `xml:",chardata"`
			Tid       string `xml:"tid"`
			Title     string `xml:"title"`
			Desc      string `xml:"desc"`
			IconUrl   string `xml:"iconUrl"`
			SecondUrl string `xml:"secondUrl"`
			PageType  string `xml:"pageType"`
		} `xml:"emotionpageshared"`
		Webviewshared struct {
			Text             string `xml:",chardata"`
			ShareUrlOriginal string `xml:"shareUrlOriginal"`
			ShareUrlOpen     string `xml:"shareUrlOpen"`
			JsAppId          string `xml:"jsAppId"`
			PublisherId      string `xml:"publisherId"`
		} `xml:"webviewshared"`
		TemplateID string `xml:"template_id"`
		Md5        string `xml:"md5"`
		Weappinfo  struct {
			Text                     string `xml:",chardata"`
			Pagepath                 string `xml:"pagepath"`
			Username                 string `xml:"username"`
			Appid                    string `xml:"appid"`
			Version                  string `xml:"version"`
			Type                     string `xml:"type"`
			Weappiconurl             string `xml:"weappiconurl"`
			Weapppagethumbrawurl     string `xml:"weapppagethumbrawurl"`
			ShareId                  string `xml:"shareId"`
			Appservicetype           string `xml:"appservicetype"`
			Secflagforsinglepagemode string `xml:"secflagforsinglepagemode"`
			Videopageinfo            struct {
				Text        string `xml:",chardata"`
				Thumbwidth  string `xml:"thumbwidth"`
				Thumbheight string `xml:"thumbheight"`
				Fromopensdk string `xml:"fromopensdk"`
			} `xml:"videopageinfo"`
		} `xml:"weappinfo"`
		Statextstr     string `xml:"statextstr"`
		MusicShareItem struct {
			Text          string `xml:",chardata"`
			MusicDuration string `xml:"musicDuration"`
		} `xml:"musicShareItem"`
		Findernamecard struct {
			Text        string `xml:",chardata"`
			Username    string `xml:"username"`
			Avatar      string `xml:"avatar"`
			Nickname    string `xml:"nickname"`
			AuthJob     string `xml:"auth_job"`
			AuthIcon    string `xml:"auth_icon"`
			AuthIconURL string `xml:"auth_icon_url"`
		} `xml:"findernamecard"`
		FinderGuarantee struct {
			Text  string `xml:",chardata"`
			Scene string `xml:"scene"`
		} `xml:"finderGuarantee"`
		Directshare string `xml:"directshare"`
		Gamecenter  struct {
			Text     string `xml:",chardata"`
			Namecard struct {
				Text    string `xml:",chardata"`
				IconUrl string `xml:"iconUrl"`
				Name    string `xml:"name"`
				Desc    string `xml:"desc"`
				Tail    string `xml:"tail"`
				JumpUrl string `xml:"jumpUrl"`
			} `xml:"namecard"`
		} `xml:"gamecenter"`
		PatMsg struct {
			Text     string `xml:",chardata"`
			ChatUser string `xml:"chatUser"`
			Records  struct {
				Text      string `xml:",chardata"`
				RecordNum string `xml:"recordNum"`
			} `xml:"records"`
		} `xml:"patMsg"`
		Secretmsg struct {
			Text        string `xml:",chardata"`
			Issecretmsg string `xml:"issecretmsg"`
		} `xml:"secretmsg"`
		Websearch struct {
			Text        string `xml:",chardata"`
			RecCategory string `xml:"rec_category"`
			ChannelId   string `xml:"channelId"`
		} `xml:"websearch"`
	} `xml:"appmsg"`
	Fromusername string `xml:"fromusername"`
	Scene        string `xml:"scene"`
	Appinfo      struct {
		Text    string `xml:",chardata"`
		Version string `xml:"version"`
		Appname string `xml:"appname"`
	} `xml:"appinfo"`
	Commenturl string `xml:"commenturl"`
} 

