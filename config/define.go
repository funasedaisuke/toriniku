package config

const (
	// ServerPort サーバーポート
	ServerPort = ":8000"
)

// ShopReq 店舗名、URL一括取得APIリクエストボディ
type ShopReq struct {
	URL  string `json:"url"`
	Shop string `json:"shop"`
}

// Shops seleniumレスポンス
type Shops struct {
	ContentType string `json:"Content-Type"`
	ShopList    []Shop `json:"shop_list"`
}

// Shop Shops要素
type Shop struct {
	Prefecture string `json:"prefecture"`
	ShopName   string `json:"shop_name"`
	URL        string `json:"url"`
}
