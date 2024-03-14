package models

import "time"

type ItemResponse struct {
	ID                           string        `json:"id"`
	StartTime                    time.Time     `json:"start_time,omitempty"`
	SiteID                       string        `json:"site_id"`
	Title                        string        `json:"title"`
	SellerID                     int           `json:"seller_id"`
	CategoryID                   string        `json:"category_id"`
	OfficialStoreID              interface{}   `json:"official_store_id"`
	Price                        float64       `json:"price"`
	BasePrice                    float64       `json:"base_price"`
	OriginalPrice                interface{}   `json:"original_price"`
	CurrencyID                   string        `json:"currency_id"`
	InitialQuantity              int           `json:"initial_quantity"`
	BuyingMode                   string        `json:"buying_mode"`
	ListingTypeID                string        `json:"listing_type_id"`
	Condition                    string        `json:"condition"`
	Permalink                    string        `json:"permalink"`
	ThumbnailID                  string        `json:"thumbnail_id"`
	Thumbnail                    string        `json:"thumbnail"`
	VideoID                      interface{}   `json:"video_id"`
	Descriptions                 []interface{} `json:"descriptions"`
	AcceptsMercadoPago           bool          `json:"accepts_mercadopago"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	Shipping                     struct {
		Mode         string        `json:"mode"`
		Methods      []interface{} `json:"methods"`
		Tags         []interface{} `json:"tags"`
		Dimensions   string        `json:"dimensions"`
		LocalPickUp  bool          `json:"local_pick_up"`
		FreeShipping bool          `json:"free_shipping"`
		LogisticType string        `json:"logistic_type"`
		StorePickUp  bool          `json:"store_pick_up"`
	} `json:"shipping"`
	InternationalDeliveryMode string `json:"international_delivery_mode"`
	SellerAddress             struct {
		City struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"state"`
		Country struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		SearchLocation struct {
			City struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
		} `json:"search_location"`
		ID int `json:"id"`
	} `json:"seller_address"`
	SellerContact interface{}   `json:"seller_contact"`
	Location      struct{}      `json:"location"`
	CoverageAreas []interface{} `json:"coverage_areas"`
	Attributes    []struct {
		ID        string        `json:"id"`
		Name      string        `json:"name"`
		ValueID   interface{}   `json:"value_id"`
		ValueName string        `json:"value_name"`
		Values    []interface{} `json:"values"`
		ValueType string        `json:"value_type"`
	} `json:"attributes"`
	ListingSource    string        `json:"listing_source"`
	Variations       []interface{} `json:"variations"`
	Status           string        `json:"status"`
	SubStatus        []interface{} `json:"sub_status"`
	Tags             []string      `json:"tags"`
	Warranty         interface{}   `json:"warranty"`
	CatalogProductID string        `json:"catalog_product_id"`
	DomainID         string        `json:"domain_id"`
	ParentItemID     interface{}   `json:"parent_item_id"`
	DealIDs          []interface{} `json:"deal_ids"`
	AutomaticRelist  bool          `json:"automatic_relist"`
	DateCreated      string        `json:"date_created"`
	LastUpdated      string        `json:"last_updated"`
	Health           float64       `json:"health"`
	CatalogListing   bool          `json:"catalog_listing"`
}
