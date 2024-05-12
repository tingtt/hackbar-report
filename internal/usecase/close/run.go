package close

import (
	promptgroup "hackbar-report/internal/usecase/prompt-group"
	"io"
)

type Prompt struct {
	Buy  Buy  `label:"収支"`
	Cash Cash `label:"レジ(各枚数)"`
}

type Buy struct {
	Likaman                   string `label:"リカーマウンテン(オープン報告分)"`
	ConvenienceStore          string `label:"コンビニ(氷・お菓子類)"`
	DifferenceFromMobileOrder string `label:"モバイルオーダー差分"`
}

type Cash struct {
	TenThousandYenBill  string `label:"1万円札"`
	FiveThousandYenBill string `label:"5千円札"`
	ThousandYenBill     string `label:"1千円札"`
	FiveHundredCoin     string `label:"500円硬貨"`
	HundredCoin         string `label:"100円硬貨"`
}

func Run(out io.Writer, in io.Reader) (*Prompt, error) {
	p := &Prompt{}

	err := promptgroup.Run(out, in, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
