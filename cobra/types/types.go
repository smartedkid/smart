package types

import "time"

type Goods struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	ClickNum        int       `json:"click_num"`
	SoldNum         int       `json:"sold_num"`
	FavNum          int       `json:"fav_num"`
	MarketPrice     float64   `json:"market_price"`
	ShopPrice       float64   `json:"shop_price"`
	GoodsBrief      string    `json:"goods_brief"`
	ShipFree        int       `json:"ship_free"`
	Images          string    `json:"images"`
	DescImages      string    `json:"desc_images"`
	GoodsFrontImage string    `json:"goods_front_image"`
	IsNew           int       `json:"is_new"`
	IsHot           int       `json:"is_hot"`
	UpdateTime      time.Time `json:"update_time"`
	CategoryId      int       `json:"category_id"`
	BrandsId        int       `json:"brands_id"`
	OnSale          int       `json:"on_sale"`
	GoodsSn         string    `json:"goods_sn"`
	AddTime         time.Time `json:"add_time"`
	IsDeleted       int       `json:"is_deleted"`
}
