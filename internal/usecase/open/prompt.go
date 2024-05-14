package open

type Prompt struct {
	Stocking                 Stocking                 `label:"仕入れ" suffix:"(\",\"区切りで複数入力)"`
	FixtureRestockingRequest FixtureRestockingRequest `label:"その他備品補充依頼(to 井出くん)"`
	Cash                     Cash                     `label:"レジ(各枚数)"`
}

type Stocking struct {
	Likaman          string `label:"リカーマウンテン(酒・割り材・氷)" mdblk-type:"list,omitempty" mdblk-list-separate-with:","`
	ConvenienceStore string `label:"コンビニ(氷・お菓子類)" mdblk-type:"list,omitempty" mdblk-list-separate-with:","`
}

type FixtureRestockingRequest struct {
	WetTowel   string `label:"おしぼり" mdblk-type:"list,omitempty"`
	PaperTowel string `label:"ペーパータオル" mdblk-type:"list,omitempty"`
	Sponge     string `label:"スポンジ" mdblk-type:"list,omitempty"`
	Detergent  string `label:"洗剤" mdblk-type:"list,omitempty"`
	Canning    string `label:"缶詰" suffix:"(スタンダード・プレミアム)" mdblk-type:"list,omitempty"`
	Other      string `label:"その他" suffix:"(あれば入力)" mdblk-type:"list,omitempty"`
}

type Cash struct {
	TenThousandYenBill  string `label:"1万円札" mdblk-format:"- ${label} x${value}" mdblk-default:"0"`
	FiveThousandYenBill string `label:"5千円札" mdblk-format:"- ${label} x${value}" mdblk-default:"0"`
	ThousandYenBill     string `label:"1千円札" mdblk-format:"- ${label} x${value}" mdblk-default:"0"`
	FiveHundredCoin     string `label:"500円硬貨" mdblk-format:"- ${label} x${value}" mdblk-default:"0"`
	HundredCoin         string `label:"100円硬貨" mdblk-format:"- ${label} x${value}" mdblk-default:"0"`
}
