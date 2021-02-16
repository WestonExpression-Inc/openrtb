package openrtb

// ORTBBidResponse is the top-level bid response object (i.e., the unnamed outer JSON object). The id attribute
// is a reflection of the bid request ID for logging purposes. Similarly, bidid is an optional response
// tracking ID for bidders. If specified, it can be included in the subsequent win notice call if the bidder
// wins. At least one seatbid object is required, which contains at least one bid for an impression. Other
// attributes are optional.
// To express a “no-bid”, the options are to return an empty response with HTTP 204. Alternately if the
// bidder wishes to convey to the exchange a reason
// OpenRTB 2.5 Section 4.2.1
type ORTBBidResponse struct {
	ID          string           `json:"id"`                   // REQUIRED ID of the bid request to which this is a response.
	SeatBid     []ORTBSeatBid    `json:"seatbid,omitempty"`    // Array of seatbid objects; 1+ required if a bid is to be made (See OpenRTB 2.5 Section 4.2.2)
	BidID       string           `json:"bidid,omitempty"`      // Bidder generated response ID to assist with logging/tracking.
	Currency    string           `json:"cur,omitempty"`        // (Default "USD") Bid currency using ISO-4217 alpha codes.
	CustomData  string           `json:"customdata,omitempty"` // Optional feature to allow a bidder to set data in the exchange’s cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include “escaped” quotation marks.
	NoBidReason *ORTBNoBidReason `json:"nbr,omitempty"`        // Reason for not bidding. (See Table 5.24 in OpenRTB 2.5 spec)
	Extensions  interface{}      `json:"ext,omitempty"`        // Additional bidder specific data for the bid
}

// ORTBSeatBid A bid response can contain multiple SeatBid objects, each on behalf of a different bidder seat and
// each containing one or more individual bids. If multiple impressions are presented in the request, the group
// attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if
// it is only interested in winning any if it can win them all as a group.
// OpenRTB 2.5 Section 4.2.2
type ORTBSeatBid struct {
	Bid        []ORTBBid   `json:"bid"`             // REQUIRED Array of 1+ Bid objects (See OpenRTB 2.5 Section 4.2.3) each related to an impression. Multiple bids can relate to the same impression
	Seat       string      `json:"seat,omitempty"`  // ID of the buyer seat (e.g., advertiser, agency) on whose behalf this bid is made.
	Group      int8        `json:"group,omitempty"` // (Default 0) 0 = impressions can be won individually; 1 = impressions must be won or lost as a group.
	Extensions interface{} `json:"ext,omitempty"`   // Additional bidder specific data for the seat bid
}

// ORTBBid A SeatBid object contains one or more Bid objects, each of which relates to a specific impression
// in the bid request via the impid attribute and constitutes an offer to buy that impression for a given price.
// OpenRTB 2.5 Section 4.2.2
type ORTBBid struct {
	ID                string                  `json:"id"`                       // REQUIRED Bidder generated bid ID to assist with logging/tracking.
	ImpressionID      string                  `json:"impid"`                    // REQUIRED ID of the Imp object in the related bid request.
	Price             float64                 `json:"price"`                    // REQUIRED Bid price expressed as CPM although the actual transaction is for a unit impression only. Note that while the type indicates float, integer math is highly recommended when handling currencies.
	WinURL            string                  `json:"nurl,omitempty"`           // Win notice URL called by the exchange if the bid wins (not necessarily indicative of a delivered, viewed, or billable ad); optional means of serving ad markup. Substitution macros (OpenRTB 2.5 Section 4.4) may be included in both the URL and optionally returned markup.
	BillingURL        string                  `json:"burl,omitempty"`           // Billing notice URL called by the exchange when a winning bid becomes billable based on exchange-specific business policy (e.g., typically delivered, viewed, etc.). Substitution macros (OpenRTB 2.5 Section 4.4 may be included.
	LossURL           string                  `json:"lurl,omitempty"`           // Loss notice URL called by the exchange when a bid is known to have been lost. Substitution macros (Section 4.4) may be included. Exchange-specific policy may preclude support for loss notices or the disclosure of winning clearing prices resulting in ${AUCTION_PRICE} macros being removed (i.e., replaced with a zero-length string).
	ADMarkup          string                  `json:"adm,omitempty"`            // Optional means of conveying ad markup in case the bid wins; supersedes the win notice if markup is included in both. Substitution macros (Section 4.4) may be included.
	AdID              string                  `json:"adid,omitempty"`           // ID of a preloaded ad to be served if the bid wins.
	AdvertiserDomain  []string                `json:"adomain,omitempty"`        // Advertiser domain for block list checking (e.g., “ford.com”). This can be an array of for the case of rotating creatives. Exchanges can mandate that only one domain is allowed.
	Bundle            string                  `json:"bundle,omitempty"`         // A platform-specific application identifier intended to be unique to the app and independent of the exchange. On Android, this should be a bundle or package name (e.g., com.foo.mygame). On iOS, it is a numeric ID.
	ImageURL          string                  `json:"iurl,omitempty"`           // URL without cache-busting to an image that is representative of the content of the campaign for ad quality/safety checking.
	CampaignID        string                  `json:"cid,omitempty"`            // Campaign ID to assist with ad quality checking; the collection of creatives for which iurl should be representative.
	CreativeID        string                  `json:"crid,omitempty"`           // Creative ID to assist with ad quality checking.
	TacticID          string                  `json:"tactic,omitempty"`         // Tactic ID to enable buyers to label bids for reporting to the exchange the tactic through which their bid was submitted. The specific usage and meaning of the tactic ID should be communicated between buyer and exchanges a priori.
	ContentCategories []ORTBContentCategory   `json:"cat,omitempty"`            // IAB content categories of the creative. (OpenRTB 2.5 Table 5.1)
	CreativeAttribues []ORTBCreativeAttribute `json:"attr,omitempty"`           // Set of attributes describing the creative. (OpenRTB 2.5 Table 5.3)
	MarkupAPI         *ORTBAPIFramework       `json:"api,omitempty"`            // API required by the markup if applicable. (OpenRTB 2.5 Table 5.6)
	VideoProtocol     *ORTBProtocol           `json:"protocol,omitempty"`       // Video response protocol of the markup if applicable. (OpenRTB 2.5 Table 5.8)
	QAGRating         *ORTBIQGMediaRating     `json:"qagmediarating,omitempty"` // Creative media rating per IQG guidelines. (OpenRTB 2.5 Table 5.19)
	Language          string                  `json:"language,omitempty"`       // Language of the creative using ISO-639-1-alpha-2. The nonstandard code “xx” may also be used if the creative has no linguistic content (e.g., a banner with just a company logo).
	DealID            string                  `json:"dealid,omitempty"`         // Reference to the deal.id from the bid request if this bid pertains to a private marketplace direct deal.
	Width             int64                   `json:"w,omitempty"`              // Width of the creative in device independent pixels (DIPS).
	Height            int64                   `json:"h,omitempty"`              // Height of the creative in device independent pixels (DIPS).
	WidthRatio        int64                   `json:"wratio,omitempty"`         // Relative width of the creative when expressing size as a ratio. Required for Flex Ads.
	HeightRatio       int64                   `json:"hratio,omitempty"`         // Relative height of the creative when expressing size as a ratio. Required for Flex Ads.
	Expirey           int64                   `json:"exp,omitempty"`            // Number of seconds the bidder is willing to wait between the auction and the actual impression.
	Extensions        interface{}             `json:"ext,omitempty"`            // Additional bidder specific data for the bid
}
