package robot

const (
	robotApiKey = "dda101c5f2d44f88beb22f2e18f93a6a"
	robotAprUrl = "http://www.tuling123.com/openapi/api"
)

type LocalReq struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	Loc    string `json:"loc"`
	UserId string `json:"userid"`
}

type TextTypeJson struct {
	Code int64  `json:"code"`
	Text string `json:"text"`
}

type LinkTypeJson struct {
	Code int64  `json:"code"`
	Text string `json:"text"`
	Url  string `json:"url"`
}

type PieceOfNews struct {
	Article   string `json:"article"`
	Source    string `json:"source"`
	Icon      string `json:"icon"`
	DetailUrl string `json:"detailurl"`
}

type NewsTypeJson struct {
	Code int64         `json:"code"`
	Text string        `json:"text"`
	List []PieceOfNews `json:"list"`
}

type PieceOfMenu struct {
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Info      string `json:"info"`
	DetailUrl string `json:"detailurl"`
}

type MenuTypeJson struct {
	Code int64         `json:"code"`
	Text string        `json:"text"`
	List []PieceOfMenu `json:"list"`
}
