package close

type Prompt struct {
	CustomerCount Customers `label:"客数" suffix:"(\",\"または\"、\"区切りで複数入力)"`
	Income        Income    `label:"収支" suffix:" ※収入はプラス、支出はマイナスで表記してください。"`
	Cash          Cash      `label:"レジ(各枚数)"`
}

type Customers struct {
	Count string `label:"客数" mdblk-type:"list" mdblk-list-separate-with:",、"`
}

type Income struct {
	Likaman                   string `label:"リカーマウンテン(オープン報告分)" mdblk-type:"list,omitempty"`
	LikamanPayBack            string `label:"リカーマウンテン(瓶回収分)" mdblk-type:"list,omitempty"`
	ConvenienceStore          string `label:"コンビニ(氷・お菓子類)" mdblk-type:"list,omitempty"`
	DifferenceFromMobileOrder string `label:"モバイルオーダー差分" mdblk-type:"list" mdblk-default:"0"`
}

type Cash struct {
	TenThousandYenBill  string `label:"1万円札" mdblk-format:"- ${label}   x${value}" mdblk-default:"0" mdblk-total-rate:"10000"`
	FiveThousandYenBill string `label:"5千円札" mdblk-format:"- ${label}   x${value}" mdblk-default:"0" mdblk-total-rate:"5000"`
	ThousandYenBill     string `label:"1千円札" mdblk-format:"- ${label}   x${value}" mdblk-default:"0" mdblk-total-rate:"1000"`
	FiveHundredCoin     string `label:"500円硬貨" mdblk-format:"- ${label} x${value}" mdblk-default:"0" mdblk-total-rate:"500"`
	HundredCoin         string `label:"100円硬貨" mdblk-format:"- ${label} x${value}" mdblk-default:"0" mdblk-total-rate:"100"`
	Total               string `label:"-" mdblk-format:"\n計 ${total}円"`
}
