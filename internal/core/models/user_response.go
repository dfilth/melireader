package models

type UserResponse struct {
	ID               int              `json:"id"`
	Nickname         string           `json:"nickname"`
	CountryID        string           `json:"country_id"`
	Address          UserAddress      `json:"address"`
	UserType         string           `json:"user_type"`
	SiteID           string           `json:"site_id"`
	Permalink        string           `json:"permalink"`
	SellerReputation SellerReputation `json:"seller_reputation"`
	Status           UserStatus       `json:"status"`
}

type UserAddress struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type SellerReputation struct {
	LevelID           string             `json:"level_id"`
	PowerSellerStatus interface{}        `json:"power_seller_status"`
	Transactions      SellerTransactions `json:"transactions"`
}

type SellerTransactions struct {
	Period string `json:"period"`
	Total  int    `json:"total"`
}

type UserStatus struct {
	SiteStatus string `json:"site_status"`
}
