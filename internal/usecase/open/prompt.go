package open

type Prompt struct {
	Stocking                 Stocking                 `label:"仕入れ" suffix:"(\",\"または\"、\"区切りで複数入力)"`
	FixtureRestockingRequest FixtureRestockingRequest `label:"その他備品補充依頼(to 井出くん)" suffix:" ※十分在庫がある場合はスキップしてください。"`
	Cash                     Cash                     `label:"レジ(各枚数)"`
}

type Stocking struct {
	Likaman          string `label:"リカーマウンテン(酒・割り材・氷)" mdblk-type:"list,omitempty" mdblk-list-separate-with:",、"`
	ConvenienceStore string `label:"コンビニ(氷・お菓子類)" mdblk-type:"list,omitempty" mdblk-list-separate-with:",、"`
}

type FixtureRestockingRequest struct {
	WetTowel   string `label:"おしぼり" suffix:"[y/n/任意メッセージ]" mdblk-type:"list,omitempty"`
	PaperTowel string `label:"ペーパータオル" suffix:"[y/n/任意メッセージ]" mdblk-type:"list,omitempty"`
	Sponge     string `label:"スポンジ" suffix:"[y/n/任意メッセージ]" mdblk-type:"list,omitempty"`
	Detergent  string `label:"洗剤" suffix:"[y/n/任意メッセージ]" mdblk-type:"list,omitempty"`
	Canning    string `label:"缶詰" suffix:"(スタンダード・プレミアム)[y/n/任意メッセージ]" mdblk-type:"list,omitempty"`
	Other      string `label:"その他" suffix:"(あれば入力)" mdblk-type:"list,omitempty"`
}

type Cash struct {
	TenThousandYenBill  string `label:"1万円札" mdblk-format:"- ${label}   x${value}" mdblk-default:"0" mdblk-total-rate:"10000"`
	FiveThousandYenBill string `label:"5千円札" mdblk-format:"- ${label}   x${value}" mdblk-default:"0" mdblk-total-rate:"5000"`
	ThousandYenBill     string `label:"1千円札" mdblk-format:"- ${label}   x${value}" mdblk-default:"0" mdblk-total-rate:"1000"`
	FiveHundredCoin     string `label:"500円硬貨" mdblk-format:"- ${label} x${value}" mdblk-default:"0" mdblk-total-rate:"500"`
	HundredCoin         string `label:"100円硬貨" mdblk-format:"- ${label} x${value}" mdblk-default:"0" mdblk-total-rate:"100"`
	Total               string `label:"-" mdblk-format:"\n計 ${total}円"`
	Diff                string `label:"先日差分" mdblk-format:"**差分 ${value}**"`
}
