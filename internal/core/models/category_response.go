package models

type CategoryResponse struct {
	ID                       string           `json:"id"`
	Name                     string           `json:"name"`
	Picture                  interface{}      `json:"picture"`
	Permalink                interface{}      `json:"permalink"`
	TotalItemsInThisCategory int              `json:"total_items_in_this_category"`
	PathFromRoot             []CategoryPath   `json:"path_from_root"`
	ChildrenCategories       []interface{}    `json:"children_categories"`
	AttributeTypes           string           `json:"attribute_types"`
	Settings                 CategorySettings `json:"settings"`
	ChannelsSettings         []ChannelSetting `json:"channels_settings"`
	MetaCategID              interface{}      `json:"meta_categ_id"`
	Attributable             bool             `json:"attributable"`
	DateCreated              string           `json:"date_created"`
}

type CategoryPath struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategorySettings struct {
	AdultContent            bool          `json:"adult_content"`
	BuyingAllowed           bool          `json:"buying_allowed"`
	BuyingModes             []string      `json:"buying_modes"`
	CatalogDomain           string        `json:"catalog_domain"`
	CoverageAreas           string        `json:"coverage_areas"`
	Currencies              []string      `json:"currencies"`
	Fragile                 bool          `json:"fragile"`
	ImmediatePayment        string        `json:"immediate_payment"`
	ItemConditions          []string      `json:"item_conditions"`
	ItemsReviewsAllowed     bool          `json:"items_reviews_allowed"`
	ListingAllowed          bool          `json:"listing_allowed"`
	MaxDescriptionLength    int           `json:"max_description_length"`
	MaxPicturesPerItem      int           `json:"max_pictures_per_item"`
	MaxPicturesPerItemVar   int           `json:"max_pictures_per_item_var"`
	MaxSubTitleLength       int           `json:"max_sub_title_length"`
	MaxTitleLength          int           `json:"max_title_length"`
	MaxVariationsAllowed    int           `json:"max_variations_allowed"`
	MaximumPrice            interface{}   `json:"maximum_price"`
	MaximumPriceCurrency    string        `json:"maximum_price_currency"`
	MinimumPrice            int           `json:"minimum_price"`
	MinimumPriceCurrency    string        `json:"minimum_price_currency"`
	MirrorCategory          interface{}   `json:"mirror_category"`
	MirrorMasterCategory    interface{}   `json:"mirror_master_category"`
	MirrorSlaveCategories   []interface{} `json:"mirror_slave_categories"`
	Price                   string        `json:"price"`
	ReservationAllowed      string        `json:"reservation_allowed"`
	Restrictions            []interface{} `json:"restrictions"`
	RoundedAddress          bool          `json:"rounded_address"`
	SellerContact           string        `json:"seller_contact"`
	ShippingOptions         []string      `json:"shipping_options"`
	ShippingProfile         string        `json:"shipping_profile"`
	ShowContactInformation  bool          `json:"show_contact_information"`
	SimpleShipping          string        `json:"simple_shipping"`
	Stock                   string        `json:"stock"`
	SubVertical             interface{}   `json:"sub_vertical"`
	Subscribable            bool          `json:"subscribable"`
	Tags                    []interface{} `json:"tags"`
	Vertical                string        `json:"vertical"`
	VipSubdomain            string        `json:"vip_subdomain"`
	BuyerProtectionPrograms []interface{} `json:"buyer_protection_programs"`
	Status                  string        `json:"status"`
}

type ChannelSetting struct {
	Channel  string `json:"channel"`
	Settings struct {
		MinimumPrice int    `json:"minimum_price"`
		Status       string `json:"status"`
	} `json:"settings"`
}
