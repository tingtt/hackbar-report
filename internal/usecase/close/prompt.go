package close

import "hackbar-report/internal/usecase/open"

type Prompt struct {
	Income Income    `label:"収支" suffix:" ※収入はプラス、支出はマイナスで表記してください。"`
	Cash   open.Cash `label:"レジ(各枚数)"`
}

type Income struct {
	Likaman                   string `label:"リカーマウンテン(オープン報告分)" mdblk-type:"list,omitempty"`
	LikamanPayBack            string `label:"リカーマウンテン(瓶回収分)" mdblk-type:"list,omitempty"`
	ConvenienceStore          string `label:"コンビニ(氷・お菓子類)" mdblk-type:"list,omitempty"`
	DifferenceFromMobileOrder string `label:"モバイルオーダー差分" mdblk-type:"list,omitempty"`
}
