package close

import "hackbar-report/internal/usecase/open"

type Prompt struct {
	Buy  Buy       `label:"収支"`
	Cash open.Cash `label:"レジ(各枚数)"`
}

type Buy struct {
	Likaman                   string `label:"リカーマウンテン(オープン報告分)" mdblk-type:"list,omitempty"`
	ConvenienceStore          string `label:"コンビニ(氷・お菓子類)" mdblk-type:"list,omitempty"`
	DifferenceFromMobileOrder string `label:"モバイルオーダー差分" mdblk-type:"list,omitempty"`
}
