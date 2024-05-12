package open

import (
	promptgroup "hackbar-report/internal/usecase/prompt-group"
	"io"
)

type Prompt struct {
	Stocking                 Stocking                 `label:"仕入れ" suffix:"(\",\"区切りで複数入力)"`
	FixtureRestockingRequest FixtureRestockingRequest `label:"その他備品補充依頼(to 井出くん)"`
	Cash                     Cash                     `label:"レジ(各枚数)"`
}

type Stocking struct {
	Likaman          string `label:"リカーマウンテン(酒・割り材・氷)"`
	ConvenienceStore string `label:"コンビニ(氷・お菓子類)"`
}

type FixtureRestockingRequest struct {
	WetTowel   string `label:"おしぼり"`
	PaperTowel string `label:"ペーパータオル"`
	Canning    string `label:"缶詰" suffix:"(スタンダード・プレミアム)"`
	Other      string `label:"その他" suffix:"(あれば入力)"`
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
