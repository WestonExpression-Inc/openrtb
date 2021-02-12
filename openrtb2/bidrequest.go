package openrtb

// ORTBBidRequest compliant with OpenRTB 2.5 (See section 3.2.1 in OpenRTB 2.5 Spec)
type ORTBBidRequest struct {
	ID                  string                `json:"id"`                // REQUIRED: Unique ID of bid request
	Impressions         []ORTBImpression      `json:"imp"`               // REQUIRED Impressions offered. At least 1 is required. (See Section 3.2.4 in OpenRTB 2.5 spec)
	Site                *ORTBSite             `json:"site,omitempty"`    // Details about the publishers website. Only applicable and recommended for websites (See section 3.2.13 in OpenRTB 2.5 spec)
	Device              *ORTBDevice           `json:"device,omitempty"`  // Details about the user's device that the impression will be delivered to (See Section 3.2.18 in OpenRTB 2.5 spec)
	User                *ORTBUser             `json:"user,omitempty"`    // Details about the human user of the device; the advertising audience. (See Section 3.2.20 in OpenRTB 2.5 spec)
	TestMode            *int8                 `json:"test,omitempty"`    // (default 0) Flag indicates if auctions are in test mode (not billable) 0 = Live mode 1 = Test Mode
	MimeTypes           []string              `json:"mimes,omitempty"`   // (Default all allowed) Content MIME types supported. Popular MIME types may include “application/x-shockwave-flash”, “image/jpg”, and “image/gif”.
	App                 *ORTBApp              `json:"app,omitempty"`     // Details about publisher's the publisher’s app (i.e., non-browser applications). Only applicable and recommended for apps. (See Section 3.2.14 in OpenRTB 2.5 spec)
	AuctionType         int64                 `json:"at,omitempty"`      // (default 2) Auction type 1 = First price 2 = Second price plus. Exchange specific aution types defined using values greater than 500
	MaxBidTimeout       int64                 `json:"tmax,omitempty"`    // Maximum time (milliseconds) to submit bid to avoid timeout
	WhitelistSeats      []string              `json:"wseat,omitempty"`   // Whitelist of buyer seats allowd to bid on impression. Sead IDs must be communicated between and exchange a priori. Omission implies no seat restrictions
	BlacklistSeats      []string              `json:"bseat,omitempty"`   // Block list of buyer seats (e.g., advertisers, agencies) restricted from bidding on this impression. IDs of seats and knowledge of the buyer’s customers to which they refer must be coordinated between bidders and the exchange a priori. At most, only one of wseat and bseat should be used in the same request.
	AllImpressions      int8                  `json:"allimps,omitempty"` // (default 0) Flag indicates if exchange can verify that impressions offered represent all impressions available in context. (eg all on the web page, all video spots) to support road-blocking. 0 = no or unknown 1 = yes impressions offered represent all available
	Currencies          []string              `json:"cur,omitempty"`     // (ISO-4217 alpha codes) Allowed currencies for bids on this exchange. Recommended only if exchange accepts multiple currencies
	WhitelistLanguages  []string              `json:"wlang,omitempty"`   // White list of languages for creatives using ISO-639-1-alpha-2. Omission implies no specific restrictions, but buyers would be advised to consider language attribute in the device and/or content objects if available.
	BlockedCategories   []ORTBContentCategory `json:"bcat,omitempty"`    // Blocked advertiser categories using IAB content categories (See Table 5.1 in OpenRTB 2.5 Spec)
	BlockedAdvertisers  []string              `json:"badv,omitempty"`    // Blocked advertisers by their domains (e.g., "ford.com")
	BlockedApplications []string              `json:"bapp,omitempty"`    // Blocklist of applications by their platform-specific exchange independent application identifiers. On Android, these should be bundle or package names (e.g., com.foo.mygame). On iOS, these are numeric IDs.
	Source              *ORTBSource           `json:"source,omitempty"`  // Details about the inventory source and which entity makes the final decision. (See section 3.2.2 in OpenRTB spec )
	Regulations         *ORTBRegulations      `json:"regs,omitempty"`    // Any industrial, legal or governmental regulations in force for this request (See Section 3.2.3 in OpenRTB 2.5 spec)
	Extensions          interface{}           `json:"ext,omitempty"`     // Any exchange specific data to be included in the bid
}

// ORTBSource describes the nature and behavior of the entity that is the source of the bid request upstream from the
// exchange. The primary purpose of this object is to define post-auction or upstream decisioning when the exchange
// itself does not control the final decision. A common example of this is header bidding, but it can also apply to
// upstream server entities such as another RTB exchange, a mediation platform, or an ad server combines direct
// campaigns with 3rd party demand decisioning. (Section 3.2.2 in OpenRTB 2.5 spec)
type ORTBSource struct {
	FinalDecision *int8       `json:"df,omitempty"`     // Entity is responsible for the final impression sale decision, where 0 = exchange, 1 = upstream source.
	TransactionID string      `json:"tid,omitempty"`    // Transaction ID that must be common across all participants in this bid request (e.g., potentially multiple exchanges).
	PaymentChain  string      `json:"pchain,omitempty"` // Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol v1.0.
	Extensions    interface{} `json:"ext,omitempty"`    // Exchange specific data related to this source
}

// ORTBRegulations contains any legal, governmental, or industry regulations that apply to the request. The coppa
// flag signals whether or not the request falls under the United States Federal Trade Commission’s regulations
// for the United States Children’s Online Privacy Protection Act (“COPPA”). (Section 3.2.3 in OpenRTB 2.5 spec)
type ORTBRegulations struct {
	IsCOPPA    int8        `json:"coppa,omitempty"` // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	Extensions interface{} `json:"ext,omitempty"`   // Exchange specific data related to this regulation
}

// ORTBImpression describes an ad placement or impression being auctioned. Each bid must conform to one of the offered
// Banner, Video, Audio, or Native types (Section 3.2.4 in OpenRTB 2.5 spec)
type ORTBImpression struct {
	ID                    string       `json:"id"`                          // REQUIRED: Unique ID of the impression within the context of the bid request (typically starts with 1 and increments)
	Metrics               []ORTBMetric `json:"metric,omitempty"`            // Array of metric objects (see section 3.2.5 in OpenRTB 2.5 spec)
	Banner                *ORTBBanner  `json:"banner,omitempty"`            // Banner object (See Section 3.2.6 in OpenRTB 2.5 spec)
	Video                 *ORTBVideo   `json:"video,omitempty"`             // A video object required if this impression is offered as a video ad opportunity. (See section 3.2.6 of OpenRTB 2.5 spec )
	PrivateMarketplace    *ORTBPMP     `json:"pmp,omitempty"`               // Describes any private marketplace deals in effect for this impression. (See section 3.2.11 in OpenRTB 2.5 spec)
	BidFloor              float64      `json:"bidfloor,omitempty"`          // (default 0.0) Minimum bid for this impression expressed in CPM
	Delay                 int64        `json:"exp,omitempty"`               // Advisory as to the number of seconds that may elapse between the auction and the actual impression.
	Audio                 *ORTBAudio   `json:"audio,omitempty"`             // An audio object required if this impression is offered as an audio ad opportunity (See section 3.2.8 of OpenRTB 2.5 Spec)
	Native                *ORTBNative  `json:"native,omitempty"`            // A native object required if this impression is offered as a native ad opportunity. (See section 3.2.9 of OpenRTB 2.5 spec)
	DisplayManager        string       `json:"displaymanager,omitempty"`    // Name of ad mediation partner, SDK technology, or player responsible for rendering ad. Sometimes used to customize ad code.
	DisplayManagerVersion string       `json:"displaymanagerver,omitempty"` // Version of display manager technology being used (see displaymanager)
	Interstitial          int8         `json:"instl,omitempty"`             // (Default 0) Flag if the add is interstitial. 1 = ad is interstitial or full screen, 0 = not interstitial
	TagID                 string       `json:"tagid,omitempty"`             // Identifier for specific ad placement or ad tag that was used to initiate auction. Useful for debugging or optimization by buyer
	BidFloorCurrency      string       `json:"bidfloorcur,omitempty"`       // (default "USD") Currency specified using ISO-4217 alpha codes. May be differet from currency returned by bidder
	ClickBrowser          int8         `json:"clickbrowser,omitempty"`      // Indicates the type of browser opened upon clicking the creative in an app, where 0 = embedded, 1 = native. Note that the Safari View Controller in iOS 9.x devices is considered a native browser for purposes of this attribute.
	Secure                *int8        `json:"secure,omitempty"`            // Flag to indicate if impression requires secure HTTPS URL creative assets and markup. 0 = non-secure 1 = secure. If omitted, secure state is unknown, but non-secure HTTP support can be assumed
	IFrameBuster          []string     `json:"iframebuster,omitempty"`      // Array of exchange specific names of supported iframe busters
	Extensions            interface{}  `json:"ext,omitempty"`               // Any exchange specific data to be included in the impression
}

// ORTBMetric metrics can offer insight into the impression to assist with decisioning such as average recent viewability,
// click-through rate, etc. (Section 3.2.5 in OpenRTB 2.5 spec)
type ORTBMetric struct {
	Type       string      `json:"type"`             // REQUIRED Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Value      float64     `json:"value"`            // REQUIRED Number representing the value of the metric. Probabilities must be in the range of 0.0 – 1.0.
	Vendor     string      `json:"vendor,omitempty"` // Source of the value using exchange curated string names which should be published to bidders a priori. If the exchange itself is the source versus a third party, “EXCHANGE” is recommended.
	Extensions interface{} `json:"ext,omitempty"`    // Any exchange specific data regarding this metric
}

// ORTBBanner most general type of impression. Although the term “banner” may have a very specific meaning in other contexts, here
// it can be many things including a simple static image, an expandable ad unit, or even in-banner video (refer to the Video object
// in Section 3.2.7 in OpenRTB 2.5 for the more generalized and full-featured video ad units). An array of Banner objects can also
// appear within the Video to describe optional companion ads defined in the VAST specification. (Section 3.2.6 in OpenRTB 2.5 spec)
type ORTBBanner struct {
	Width               *int64                    `json:"w,omitempty"`        // Width in device independent pixels (DIPS). If no format objects are specified, this is an exact width requirement. Otherwise, it is a preferred width.
	Height              *int64                    `json:"h,omitempty"`        // Height in device independent pixels (DIPS). If no format objects are specified, this is an exact height requirement. Otherwise, it is a preferred height.
	MimeTypes           []string                  `json:"mimes,omitempty"`    // (Default all allowed) Content MIME types supported. Popular MIME types may include “application/x-shockwave-flash”, “image/jpg”, and “image/gif”.
	Format              []ORTBFormat              `json:"format,omitempty"`   // Banner sizes permitted. If none are specified, then the use of the h and w attributes is highly recommended
	MaxWidth            int64                     `json:"wmax,omitempty"`     // DEPRECATED NOTE: Deprecated in favor of the format array. Maximum width in device independent pixels (DIPS).
	MaxHeight           int64                     `json:"hmax,omitempty"`     // DEPRECATED NOTE: Deprecated in favor of the format array. Maximum height in device independent pixels (DIPS).
	MinWidth            int64                     `json:"wmin,omitempty"`     // DEPRECATED NOTE: Deprecated in favor of the format array. Minimum width in device independent pixels (DIPS).
	MinHeight           int64                     `json:"hmin,omitempty"`     // DEPRECATED NOTE: Deprecated in favor of the format array. Minimum height in device independent pixels (DIPS).
	ID                  string                    `json:"id,omitempty"`       // Unique identifier for this banner object. Recommended when Banner objects are used with a video object to represent an array of companion ads. Values usually start at 1 and increase with each object; should be unique within an impression.
	BlockedTypes        []ORTBBannerAd            `json:"btype,omitempty"`    // (Default all allowed) Blocked banner ad types. (See Table 5.2 of OpenRTB 2.5 spec)
	BlockedAttributes   []ORTBCreativeAttribute   `json:"battr,omitempty"`    // (Default all allowed) Blocked creative attributes. (See Table 5.3 of OpenRTB 2.5 spec)
	Position            *ORTBAdPosition           `json:"pos,omitempty"`      // Ad position on the screen. (See Table 5.4 of OpenRTB 2.5 spec)
	TopFrame            int8                      `json:"topframe,omitempty"` // Indicates if the banner is in the top frame as opposed to an iframe, where 0 = no, 1 = yes
	ExpansionDirections []ORTBExpandableDirection `json:"expdir,omitempty"`   // (Default non expandable) Directions in which the banner may expand. (See Table 5.5 of OpenRTB 2.5 spec)
	SupportedAPIs       []ORTBAPIFramework        `json:"api,omitempty"`      // List of supported API frameworks for this banner. (See Table 5.6 in OpenRTB 2.5 spec).  If an API is not explicitly listed it is assumed not to be supported.
	VideoCompanion      int8                      `json:"vcm,omitempty"`      // Relevant only for banner objects used with a video object in an array of companion ads. Indicates the companion banner rendering mode relative elative to the associated video, where 0 = concurrent, 1 = end-card
	Extensions          interface{}               `json:"ext,omitempty"`      // Any exchange specific data related to the banner ad
}

// ORTBVideo in-stream video impression. Many of the fields are non-essential for minimally viable transactions
// but are included to offer fine control when needed. Video in OpenRTB generally assumes compliance with the
// VAST standard. As such, the notion of companion ads is supported by optionally including an array of Banner
// objects (refer to the Banner object in Section 3.2.6) that define these companion ads. (Section 3.2.7 in OpenRTB 2.5 spec)
type ORTBVideo struct {
	MimeTypes         []string                    `json:"mimes"`                    // REQUIRED Content MIME types supported. Popular MIME types may include “application/x-shockwave-flash”, “image/jpg”, and “image/gif”.
	MinDuration       int64                       `json:"minduration,omitempty"`    // REQUIRED Minimum video ad duration in seconds.
	MaxDuration       int64                       `json:"maxduration,omitempty"`    // REQUIRED Maximum video ad duration in seconds.
	AllowedProtocols  []ORTBProtocol              `json:"protocols,omitempty"`      // Supported video protocols. At least one supported protocol must be specified in either the protocol or protocols attribute. (See Table 5.8 in OpenRTB 2.5 spec)
	Width             int64                       `json:"w,omitempty"`              // Width of the video player in device independent pixels (DIPS).
	Height            int64                       `json:"h,omitempty"`              // Height of the video player in device independent pixels (DIPS).
	Protocol          ORTBProtocol                `json:"protocol,omitempty"`       // DEPRICATED NOTE: Deprecated in favor of protocols. Supported video protocol. At least one supported protocol must be specified in either the protocol or protocols attribute. (See Table 5.8 in OpenRTB 2.5 spec)
	StartDelay        *ORTBStartDelay             `json:"startdelay,omitempty"`     // Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll ad placements. (See Table 5.12 in OpenRTB 2.5 spec)
	PlacementType     ORTBVideoPlacement          `json:"placement,omitempty"`      // Placement type for the impression (See Table 5.9 in OpenRTB 2.5 spec)
	Linearity         ORTBVideoLinearity          `json:"linearity,omitempty"`      // Indicates if the impression must be linear, nonlinear, etc. If none specified, assume all are allowed. (See Table 5.7 in OpenRTB 2.5 spec)
	AllowSkip         *int8                       `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes. If a bidder sends markup/creative that is itself skippable, the Bid object should include the attr array with an element of 16 indicating skippable video.
	MinimumSkip       int64                       `json:"skipmin,omitempty"`        // Videos of total duration greater than this number of seconds can be skippable; only applicable if the ad is skippable.
	SkipDisablePeriod int64                       `json:"skipafter,omitempty"`      // Number of seconds a video must play before skipping is enabled; only applicable if the ad is skippable.
	Sequence          int8                        `json:"sequence,omitempty"`       // If multiple ad impressions are offered in the same bid request, the sequence number will allow for the coordinated delivery of multiple creatives.
	BlockedAttributes []ORTBCreativeAttribute     `json:"battr,omitempty"`          // (Default all allowed) Blocked creative attributes. (See Table 5.3 of OpenRTB 2.5 spec)
	MaximumExtended   int64                       `json:"maxextended,omitempty"`    // (Default not allowed) Maximum extended ad duration if an extension is allowed. blank or 0 = extension not allowed. -1 = extension is allowed, and there is no time limit imposed. >0 =  value represents the number of seconds of extended play supported beyond the max duration value.
	MinimumBitrate    int64                       `json:"minbitrate,omitempty"`     // (Default 250) Minimum bit rate in Kbps.
	MaximumBitrate    int64                       `json:"maxbitrate,omitempty"`     // (Default 4000) Maximum bit rate in Kbps.
	BoxingAllowed     int8                        `json:"boxingallowed,omitempty"`  // (Default 1) Indicates if letter-boxing of 4:3 content into a 16:9 window is allowed, where 0 = no, 1 = yes.
	PlaybackMethod    []ORTBPlaybackMethod        `json:"playbackmethod,omitempty"` // (Default All) Playback methods that may be in use. If none are specified, any method may be used. (See Table 5.10 in OpenRTB 2.5 spec) Only one method is typically used in practice. As a result, this array may be converted to an integer in a future version of the specification.
	PlaybackEnd       ORTBPlaybackCessationMode   `json:"playbackend,omitempty"`    // The event that causes playback to end. (See Table 5.11 in OpenRTB 2.5 spec)
	DeliveryMethods   []ORTBContentDeliveryMethod `json:"delivery,omitempty"`       // (Default All) Supported delivery methods (e.g., streaming, progressive). If none specified, assume all are supported. (See Table 5.15 of OpenRTB 2.5 spec)
	Position          *ORTBAdPosition             `json:"pos,omitempty"`            // Ad position on the screen. (see Table 5.4 of OpenRTB 2.5 spec)
	CompanionAds      []ORTBBanner                `json:"companionad,omitempty"`    //	The array of Banner objects if companion ads are available. (See Section 3.2.6 in OpenRTB 2.5 spec)
	SupportedAPIs     []ORTBAPIFramework          `json:"api,omitempty"`            // List of supported API frameworks for this banner. (See Table 5.6 in OpenRTB 2.5 spec).  If an API is not explicitly listed it is assumed not to be supported.
	CompanionTypes    []ORTBCompanion             `json:"companiontype,omitempty"`  // Supported VAST companion ad types. (See Table 5.14 in OpenRTB 2.5 spec) Recommended if companion Banner objects are included via the companionad array.
	Extensions        interface{}                 `json:"ext,omitempty"`            // Exchange specific data related to this Video item
}

// ORTBAudio is an audio type impression. Many of the fields are non-essential for minimally viable transactions
// but are included to offer fine control when needed. Audio in OpenRTB generally assumes compliance with the D
// AAST standard. As such, the notion of companion ads is supported by optionally including an array of Banner
// objects (refer to the Banner object in Section 3.2.6) that define these companion ads. (Section 3.2.8 in OpenRTB 2.5 spec)
type ORTBAudio struct {
	MimeTypes           []string                    `json:"mimes"`                   // Content MIME types supported (e.g., “audio/mp4”).
	MinDuration         int64                       `json:"minduration,omitempty"`   // Minimum audio ad duration in seconds.
	MaxDuration         int64                       `json:"maxduration,omitempty"`   // Maximum audio ad duration in seconds.
	AllowedProtools     []ORTBProtocol              `json:"protocols,omitempty"`     // The array of supported audio protocols. (See Table 5.8 in OpenRTB 2.5 spec)
	StartDelay          *ORTBStartDelay             `json:"startdelay,omitempty"`    // Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll ad placements. (See Table 5.12 in OpenRTB 2.5 spec)
	Sequence            int64                       `json:"sequence,omitempty"`      // If multiple ad impressions are offered in the same bid request, the sequence number will allow for the coordinated delivery of multiple creatives.
	BlockedAttributes   []ORTBCreativeAttribute     `json:"battr,omitempty"`         // (Default all allowed) Blocked creative attributes. (See Table 5.3 of OpenRTB 2.5 spec)
	MaximumExtended     int64                       `json:"maxextended,omitempty"`   // (Default not allowed) Maximum extended ad duration if an extension is allowed. blank or 0 = extension not allowed. -1 = extension is allowed, and there is no time limit imposed. >0 =  value represents the number of seconds of extended play supported beyond the max duration value.
	MinimumBitrate      int64                       `json:"minbitrate,omitempty"`    // (Default 250) Minimum bit rate in Kbps.
	MaximumBitrate      int64                       `json:"maxbitrate,omitempty"`    // (Default 4000) Maximum bit rate in Kbps.
	DeliveryMethods     []ORTBContentDeliveryMethod `json:"delivery,omitempty"`      // (Default All) Supported delivery methods (e.g., streaming, progressive). If none specified, assume all are supported. (See Table 5.15 of OpenRTB 2.5 spec)
	CompanionTypes      []ORTBCompanion             `json:"companiontype,omitempty"` // Supported VAST companion ad types. (See Table 5.14 in OpenRTB 2.5 spec) Recommended if companion Banner objects are included via the companionad array.
	MaximumSequence     int64                       `json:"maxseq,omitempty"`        // Maximum number of ads that can be played in an ad pod.
	Feed                ORTBFeedType                `json:"feed,omitempty"`          // Type of audio feed. (See Table 5.16 in OpenRTB 2.5 spec)
	IsStitched          int8                        `json:"stitched,omitempty"`      // Indicates if the ad is stitched with audio content or delivered independently, where 0 = no, 1 = yes.
	VolumeNormalization *ORTBVolumeNormalization    `json:"nvol,omitempty"`          // Volume normalization mode. (See Table 5.17 in OpenRTB 2.5 spec)
}

// ORTBNative is a native type impression. Native ad units are intended to blend seamlessly into the surrounding
// content. As such, the response must be well-structured to afford the publisher fine-grained control over
// rendering. The Native Subcommittee has developed a companion specification to OpenRTB called the Dynamic
// Native Ads API. It defines the request parameters and response markup structure of native ad units. This
// object provides the means of transporting request parameters as an opaque string so that the specific
// parameters can evolve separately under the auspices of the Dynamic Native Ads API. Similarly, the ad markup
// served will be structured according to that specification. (Section 2.3.9 in OpenRTB 2.5 spec)
type ORTBNative struct {
	Request           string                  `json:"request"`         // REQUIRED Request payload complying with the Native Ad Specification. See natice Ad Spec
	Version           string                  `json:"ver,omitempty"`   // (Default 1.1) The version of the dynamic native ads API to which request complies; highly recommended for efficient parsing.
	SupportedAPIs     []ORTBAPIFramework      `json:"api,omitempty"`   // (Default None) List of supported API frameworks for this banner. (See Table 5.6 in OpenRTB 2.5 spec).  If an API is not explicitly listed it is assumed not to be supported.
	BlockedAttributes []ORTBCreativeAttribute `json:"battr,omitempty"` // (Default all allowed) Blocked creative attributes. (See Table 5.3 of OpenRTB 2.5 spec)
}

// ORTBFormat represents size format (Section 3.2.10 in OpenRTB 2.5 spec)
type ORTBFormat struct {
	Width       int64       `json:"w,omitempty"`      // Width in device independent pixels (DIPS).
	Height      int64       `json:"h,omitempty"`      // Height in device independent pixels (DIPS).
	WidthRatio  int64       `json:"wratio,omitempty"` // Relative width when expressing size as a ratio
	HeightRatio int64       `json:"hratio,omitempty"` // Relative height when expressing size as a ratio.
	MinWidth    int64       `json:"wmin,omitempty"`   // The minimum width in device independent pixels (DIPS) at which the ad will be displayed the size is expressed as a ratio.
	Extensions  interface{} `json:"ext,omitempty"`    // Any exchange specific data related to the format
}

// ORTBPMP is a private marketplace container for direct deals between buyers and sellers that may pertain to
// this impression. The actual deals are represented as a collection of Deal objects. Refer to Section 7.3
// in OpenRTB 2.5 spec for more details. (Section 3.2.11 in OpenRTB 2.5 spec)
type ORTBPMP struct {
	IsPrivate  int8        `json:"private_auction,omitempty"` // (Default 0) Indicator of auction eligibility to seats named in a “direct deal” object. 0 = all bids accepted 1 = bids are restricted to the deals specified and the terms thereof
	Deals      []ORTBDeal  `json:"deals,omitempty"`           // Deals that convey the specific deals applicable to this impression. (See section 3.2.12 of OpenRTB 2.5 Spec)
	Extensions interface{} `json:"ext,omitempty"`             // Exchange specific data related to this Private Markerplace
}

// ORTBDeal constitutes a specific deal that was struck a priori between a buyer and a seller. Its presence
// with the Pmp collection indicates that this impression is available under the terms of that deal.
// Refer to Section 7.3 in OpenRTB 2.5 spec for more details. (Section 3.2.12 in OpenRTB spec)
type ORTBDeal struct {
	ID                   string   `json:"id"`                    // REQUIRED A unique identifier for the direct deal.
	BidFloor             float64  `json:"bidfloor,omitempty"`    // (default 0.0) Minimum bid for this impression expressed in CPM
	BidFloorCurrency     string   `json:"bidfloorcur,omitempty"` // (default "USD") Currency specified using ISO-4217 alpha codes. May be differet from currency returned by bidder
	AuctionType          int64    `json:"at,omitempty"`          // Auction type. If “1”, then first price auction. If “2”, then second price auction. If “3”, the passed bidfloor indicates the a priori agreed upon deal price. Additional auction types can be defined as per the exchange’s business rules.
	WhitelistSeats       []string `json:"wseat,omitempty"`       // Whitelist of buyer seats allowd to bid on impression. Sead IDs must be communicated between and exchange a priori. Omission implies no seat restrictions
	WhitelistAdvertisers []string `json:"wadomain,omitempty"`    // The array of advertiser domains allowed to bid on this Direct Deal. For example, (“advertiser1.com”,”advertiser2.com”) indicates that only the listed advertisers are allowed to bid on this direct deal.
}

// ORTBSite meant to be included if the ad supported content is a website as opposed to a non-browser application.
// A bid request must not contain both a Site and an App object. At a minimum, it is useful to provide a site ID
// or page URL, but this is not strictly required. (Section 3.2.13 in OpenRTB 2.5 Spec)
type ORTBSite struct {
	ID                string                `json:"id,omitempty"`            // Exchange-specific site ID.
	Name              string                `json:"name,omitempty"`          // Site name (may be masked at publisher’s request).
	Publisher         *ORTBPublisher        `json:"publisher,omitempty"`     // Details about the publisher of the site. (See section 3.2.15 of OpenRTB 2.5 spec)
	Domain            string                `json:"domain,omitempty"`        // Domain of the site, used for advertiser side blocking. For example, “foo.com”.
	ContentCategories []ORTBContentCategory `json:"cat,omitempty"`           // Array of IAB content categories of the site. (See Table 5.1 of OpenRTB 2.5 spec)
	SectionCategories []ORTBContentCategory `json:"sectioncat,omitempty"`    // Array of IAB content categories that describe the current section of the site. (See Table 5.1 of OpenRTB 2.5 spec)
	PageCategories    []ORTBContentCategory `json:"pagecat,omitempty"`       // Array of IAB content categories that describe the current page or view of the site. (See Table 5.1 of OpenRTB 2.5 spec)
	Page              string                `json:"page,omitempty"`          // URL of the page where the impression will be shown.
	Referrer          string                `json:"ref,omitempty"`           // Referrer URL that caused navigation to the current page.
	SearchString      string                `json:"search,omitempty"`        // Search string that caused navigation to the current page.
	IsMobile          int8                  `json:"mobile,omitempty"`        // Indicates if the site has been programmed to optimize layout when viewed on mobile devices, where 0 = no, 1 = yes
	PrivacyPolicy     int8                  `json:"privacypolicy,omitempty"` // Specifies whether the site has a privacy policy. “1” means there is a policy. “0” means there is not.
	Content           ORTBContent           `json:"content,omitempty"`       // Details about the content within the site. (See section 3.2.16 of OpenRTB spec)
	Keywords          string                `json:"keywords,omitempty"`      // Comma-separated list of keywords about the site
	Extensions        interface{}           `json:"ext,omitempty"`           // Exchange specific data related to this site
}

// ORTBApp should be included if the ad supported content is a non-browser application (typically in mobile) as opposed
// to a website. A bid request must not contain both an App and a Site object. At a minimum, it is useful to provide an
// App ID or bundle, but this is not strictly require (Section 3.2.14 of OpenRTB spec)
type ORTBApp struct {
	ID                string                `json:"id,omitempty"`            // Application ID on the exchange.
	Name              string                `json:"name,omitempty"`          // Application name (may be masked at publisher’s request).
	Bundle            string                `json:"bundle,omitempty"`        // A platform-specific application identifier intended to be unique to the app and independent of the exchange. On Android, this should be a bundle or package name (e.g., com.foo.mygame). On iOS, it is a numeric ID.
	Domain            string                `json:"domain,omitempty"`        // The domain of the application (e.g., “mygame.foo.com”).
	StoreURL          string                `json:"storeurl,omitempty"`      // For QAG 1.5 compliance, an app store URL for an installed app should be passed in the bid request.
	ContentCategories []ORTBContentCategory `json:"cat,omitempty"`           // Array of IAB content categories of the overall application. (See Table 5.1 of OpenRTB 2.5 spec)
	SectionCategories []ORTBContentCategory `json:"sectioncat,omitempty"`    // Array of IAB content categories that describe the current subsection of the app. (See Table 5.1 of OpenRTB 2.5 spec)
	PageCategories    []ORTBContentCategory `json:"pagecat,omitempty"`       // Array of IAB content categories that describe the current page or view of the app. (See Table 5.1 of OpenRTB 2.5 spec)
	Version           string                `json:"ver,omitempty"`           // Application version
	PrivacyPolicy     int8                  `json:"privacypolicy,omitempty"` // Flag if the app has a privacy policy, where 0 = no, 1 = yes.
	IsPaid            int8                  `json:"paid,omitempty"`          // Flag if the app is a paid version “1” if the application is a paid version; else “0” (i.e., free).
	Publisher         *ORTBPublisher        `json:"publisher,omitempty"`     // Details about publisher
	Content           *ORTBContent          `json:"content,omitempty"`       // Details about app content
	Keywords          string                `json:"keywords,omitempty"`      // Comma-separated list of keywords about the app
	Extensions        interface{}           `json:"ext,omitempty"`           // Exchange specific data related to the app
}

// ORTBPublisher describes the publisher of the media in which the ad will be displayed. The publisher is
// typically the seller in an OpenRTB transaction. (Section 3.2.15 in OpenRTB 2.5 spec)
type ORTBPublisher struct {
	ID                string                `json:"id,omitempty"`     // Publisher ID on the exchange.
	Name              string                `json:"name,omitempty"`   // Publisher name (may be masked at publisher’s request).
	ContentCategories []ORTBContentCategory `json:"cat,omitempty"`    // Array of IAB content categories of the site. (See Table 5.1 of OpenRTB 2.5 spec)
	Domain            string                `json:"domain,omitempty"` // Publisher’s highest level domain name, for example “foopub.com”
	Extensions        interface{}           `json:"ext,omitempty"`    // Exchange specific data related to the publisher
}

// ORTBContent describes the content in which the impression will appear, which may be syndicated or non-syndicated
// content. This object may be useful when syndicated content contains impressions and does not necessarily match
// the publisher’s general content. The exchange might or might not have knowledge of the page where the content is
// running, as a result of the syndication method. For example, there might be a video impression embedded in an
// iframe on an unknown web property or device. (Section 3.2.16 of OpenRTB 2.5 spec)
type ORTBContent struct {
	ID                 string                 `json:"id,omitempty"`                 // The unique ID used to identify the content
	Episode            int64                  `json:"episode,omitempty"`            // Content episode number (typically applies to video content).
	Title              string                 `json:"title,omitempty"`              // Content title. Video examples: “Search Committee” (television) or “A New Hope” (movie) or “Endgame” (made for web)
	Series             string                 `json:"series,omitempty"`             // Content series. Video examples: “The Office” (television) or “Star Wars” (movie) or “Arby ‘N’ The Chief” (made for web)
	Season             string                 `json:"season,omitempty"`             // Content season. e.g., “Season 3” (typically applies to video content).
	Artist             string                 `json:"artist,omitempty"`             // Artist credited with the content.
	Genre              string                 `json:"genre,omitempty"`              // Genre that best describes the content (e.g., rock, pop, etc).
	Album              string                 `json:"album,omitempty"`              // Album to which the content belongs; typically for audio.
	ISRC               string                 `json:"isrc,omitempty"`               // International Standard Recording Code conforming to ISO3901.
	Producer           *ORTBProducer          `json:"producer,omitempty"`           // Details about the content producer.
	URL                string                 `json:"url,omitempty"`                // Original URL of the content, for buy-­‐side contextualization or review.
	ContentCategories  []ORTBContentCategory  `json:"cat,omitempty"`                // Array of IAB content categories of the site. (See Table 5.1 of OpenRTB 2.5 spec)
	ProductionQuality  *ORTBProductionQuality `json:"prodq,omitempty"`              // Production quality. (See section 5.13 of OpenRTB 2.5 spec)
	VideoQuality       *ORTBProductionQuality `json:"videoquality,omitempty"`       // Video quality per the IAB’s classification. (See Table 5.13 of OpenRTB 2.5 spec)
	Context            ORTBContentContext     `json:"context,omitempty"`            // Specifies the type of content (game, video, text, etc.). (See Table 5.18 of OpenRTB 2.5 spec)
	ContentRating      string                 `json:"contentrating,omitempty"`      // Content rating (e.g., MPAA)
	UserRating         string                 `json:"userrating,omitempty"`         // User rating of the content (e.g., number of stars, likes, etc.).
	MediaRating        ORTBIQGMediaRating     `json:"qagmediarating,omitempty"`     // Media rating of the content, per QAG guidelines. (See Table 5.19 of OpenRTB 2.5 spec)
	Keywords           string                 `json:"keywords,omitempty"`           // Comma-separated list of keywords about the content
	IsLiveStream       int8                   `json:"livestream,omitempty"`         // Is content live? E.g., live video stream, live blog.“1” means content is live. “0” means it is not live.
	SourceRelationship int8                   `json:"sourcerelationship,omitempty"` // 1 for “direct”; 0 for “indirect”
	Lenth              int64                  `json:"len,omitempty"`                // Length of content (appropriate for video or audio) in seconds.
	Language           string                 `json:"language,omitempty"`           // Language of the content. Use alpha-­‐2/ISO 639-­‐1 codes
	IsEmbeddable       int8                   `json:"embeddable,omitempty"`         // From QAG Video Addendum. If content can be embedded (such as an embeddable video player) this value should be set to “1”.  If content cannot be embedded, then this should be set to “0”
	AdditionalData     []ORTBData             `json:"data,omitempty"`               // Additional content data. Each data object represents a different data source. (See section 3.2.21 of OpenRTB 2.5 spec)
	Extensions         interface{}            `json:"ext,omitempty"`                // Exchange specific data related to the content
}

// ORTBProducer defines the producer of the content in which the ad will be shown. This is particularly useful when
// the content is syndicated and may be distributed through different publishers and thus when the producer and
// publisher are not necessarily the same entity. (Section 3.2.17 of OpenRTB 2.5 spec)
type ORTBProducer struct {
	ID                string                `json:"id,omitempty"`     // Content producer or originator ID. Useful if the content is syndicated, and may be posted on a site using embed tags.
	Name              string                `json:"name,omitempty"`   // Content producer or originator name (e.g., “Warner Bros”).
	ContentCategories []ORTBContentCategory `json:"cat,omitempty"`    // Array of IAB content categories of the site. (See Table 5.1 of OpenRTB 2.5 spec)
	Domain            string                `json:"domain,omitempty"` // URL of the content producer
	Extensions        interface{}           `json:"ext,omitempty"`    // Exchange specific data related to the producer
}

// ORTBDevice provides information pertaining to the device through which the user is interacting. Device information
// includes its hardware, platform, location, and carrier data. The device can refer to a mobile handset, a desktop
// computer, set-top box, or another digital device. (Section 3.2.18 of OpenRTB 2.5 spec)
type ORTBDevice struct {
	UserAgent            string              `json:"ua,omitempty"`             // Browser user agent string.
	GeoLocation          *ORTBGeolocation    `json:"geo,omitempty"`            // Geography as derived from the device’s location services (e.g., cell tower triangulation, GPS) or IP address. (See section 3.2.19 of OpenRTB 2.5 spec)
	IPv4                 string              `json:"ip,omitempty"`             // IPv4 address is closest to the device.
	Height               int64               `json:"h,omitempty"`              // Physical height of the screen in pixels.
	Width                int64               `json:"w,omitempty"`              // Physical width of the screen in pixels.
	AdvertisingIFA       string              `json:"ifa,omitempty"`            // Native identifier for advertisers; an opaque ID assigned by the device or browser for use as an advertising identifier. (e.g. Apple’s IFA, Android’s Advertising ID, etc)
	DoNotTrack           *int8               `json:"dnt,omitempty"`            // Standard "Do Not Track" flag as set in the header by the browser, where 0 = tracking is unrestricted, 1 = do not track.
	LimtTracking         *int8               `json:"lmt,omitempty"`            // “Limit Ad Tracking” signal commercially endorsed (e.g., iOS, Android), where 0 = tracking is unrestricted, 1 = tracking must be limited per commercial guidelines.
	IPv6                 string              `json:"ipv6,omitempty"`           // IP address closest to device as IPv6.
	DeviceType           ORTBDeviceType      `json:"devicetype,omitempty"`     // Device type being used. (See table 5.21 in OpenRTB 2.5 spec)
	Make                 string              `json:"make,omitempty"`           // Device make (e.g., “Apple”).
	Model                string              `json:"model,omitempty"`          // Device model (e.g., “iPhone 5s”)
	OS                   string              `json:"os,omitempty"`             // Device operating system (e.g., “iOS”).
	OSVersion            string              `json:"osv,omitempty"`            // Device operating system version (e.g., “3.1.2”).
	HardwareVersion      string              `json:"hwv,omitempty"`            // Hardware Version of Device in the Device Object (e.g. iPhone Xs would have the Xs in the hwv field.)
	PPI                  int64               `json:"ppi,omitempty"`            // Screen size as pixels per linear inch.
	PixelRatio           float64             `json:"pxratio,omitempty"`        // Ratio of physical pixels to device-independent pixels.
	SupportsJS           int8                `json:"js,omitempty"`             // Flag if device supports javascript. “1” if the device supports JavaScript; else “0”.
	GeoAPIJS             int8                `json:"geofetch,omitempty"`       // Flag if  geolocation API will be available to JavaScript code running in the banner, where 0 = no, 1 = yes.
	FlashVersion         string              `json:"flashver,omitempty"`       // Version of Flash supported by the browser.
	Language             string              `json:"language,omitempty"`       // Browser language; use alpha-­‐2/ISO 639-­‐1 codes.
	Carrier              string              `json:"carrier,omitempty"`        // Carrier or ISP (e.g., “VERIZON”). “WIFI” is often used in mobile to indicate high bandwidth (e.g., video friendly vs. cellular).
	MCCMNC               string              `json:"mccmnc,omitempty"`         // Mobile carrier as the concatenated MCC-MNC code (e.g.,“310-005” identifies Verizon Wireless CDMA in the USA).
	ConnectionType       *ORTBConnectionType `json:"connectiontype,omitempty"` // Detected data connection type for the device. (See table 5.22 in OpenRTB 2.5 Spec)
	DeviceIDSHA1         string              `json:"didsha1,omitempty"`        // SHA1 hashed device ID; IMEI when available, else MEID or ESN. OpenRTB’s preferred method for device ID hashing is SHA1.
	DeviceIDMD5          string              `json:"didmd5,omitempty"`         // MD5 hashed device ID; IMEI when available, else MEID or ESN. Should be interpreted as case insensitive.
	DevicePlatformIDSHA1 string              `json:"dpidsha1,omitempty"`       // SHA1 hashed platform-­‐specific ID (e.g., Android ID or UDID for iOS). OpenRTB’s preferred method for device ID hash is SHA1.
	DevicePlatformIDMD5  string              `json:"dpidmd5,omitempty"`        // MD5 hashed platform-­‐specific ID (e.g., Android ID or UDID for iOS). Should be interpreted as case insensitive.
	MACSHA1              string              `json:"macsha1,omitempty"`        // MAC address of the device; hashed via SHA1.
	MACMD5               string              `json:"macmd5,omitempty"`         // MAC address of the device; hashed via MD5.
	Extensions           interface{}         `json:"ext,omitempty"`            // Exchange specific data related to the device
}

// ORTBGeolocation encapsulates various methods for specifying a geographic location. When subordinate to a Device
// object, it indicates the location of the device which can also be interpreted as the user’s current location. When
// subordinate to a User object, it indicates the location of the user’s home base (i.e., not necessarily their current
// location). The lat/lon attributes should only be passed if they conform to the accuracy depicted in the type attribute.
// For example, the centroid of a geographic region such as postal code should not be passed.
// (Section 3.2.19 of OpenRTB 2.5 spec)
type ORTBGeolocation struct {
	Latitude   float64               `json:"lat,omitempty"`           // Latitude from -90.0 to +90.0, where negative is south.
	Longitude  float64               `json:"lon,omitempty"`           // Longitude from -180.0 to +180.0, where negative is west.
	Type       ORTBLocationType      `json:"type,omitempty"`          // Source of location data; recommended when passing lat/lon. (See Table 5.20 in OpenRTB 2.5 spec)
	Country    string                `json:"country,omitempty"`       // Country using ISO-­‐3166-­‐1 Alpha-­‐3.
	Region     string                `json:"region,omitempty"`        // Region using ISO 3166-­‐2.
	City       string                `json:"city,omitempty"`          // City code using the United Nations Code for Trade and Transport Locations unece.org/cefact/locode/
	Accuracy   int64                 `json:"accuracy,omitempty"`      // Estimated location accuracy in meters. Recommended when lat/lon are specified and derived from a device’s location services (i.e., type = 1). Note that this is the accuracy as reported from the device. Consult OS-specific documentation (e.g., Android, iOS) for exact interpretation.
	LastFix    int64                 `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established. Note that devices may cache location data across multiple fetches. Ideally, this value should be from the time the actual fix was taken.
	IPService  ORTBIPLocationService `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable (i.e., type = 2). (See Table 5.23 in OpenRTB 2.5 spec)
	RegionFIPS string                `json:"regionfips104,omitempty"` // Region of a country using FIPS 10-­‐4 notation (alternative to ISO 3166-­‐2).
	MetroCode  string                `json:"metro,omitempty"`         // Google Metro codes are similar to, but not the same as, Nielsen DMAs.
	ZIPCode    string                `json:"zip,omitempty"`           // Zip/postal code.
	UTCOffset  int64                 `json:"utcoffset,omitempty"`     // Local time as the number +/- of minutes from UTC.
	Extensions interface{}           `json:"ext,omitempty"`           // Exchange specific data related to the geolocation
}

// ORTBUser contains information known or derived about the human user of the device (i.e., the audience for advertising).
// The user id is an exchange artifact and may be subject to rotation or other privacy policies. However, this user ID must
// be stable long enough to serve reasonably as the basis for frequency capping and retargeting.
// (Section 3.2.20 in OpenRTB 2.5 spec)
type ORTBUser struct {
	ID          string           `json:"id,omitempty"`         // Unique consumer ID of this user on the exchange.
	BuyerUserID string           `json:"buyeruid,omitempty"`   // Buyer’s user ID for this user as mapped by exchange for the buyer.
	YearOfBirth int              `json:"yob,omitempty"`        // Year of birth as a 4-­‐digit integer.
	Gender      string           `json:"gender,omitempty"`     // Gender as “M” male, “F” female, “O” Other. (Null indicates unknown).
	Keywords    string           `json:"keywords,omitempty"`   // Comma-separated list of keywords consumer interests or intent
	CustomData  string           `json:"customdata,omitempty"` // If supported by the exchange, this is custom data that the bidder had stored in the exchange’s cookie. The string may be in base85 cookie safe characters and be in any format. This may useful for storing user features. Note: Proper JSON encoding must be used to include “escaped” quotation marks.
	GeoLocation *ORTBGeolocation `json:"geo,omitempty"`        // Home geo for the user (e.g., based off of registration data); this is different from the current location of the access device (that is defined by the geo object embedded in the Device Object); see Device Object.
	Data        []ORTBData       `json:"data,omitempty"`       // Additional user data. Each Data object represents a different data source.
	Extensions  interface{}      `json:"ext,omitempty"`        // Exchange specific data related to the user
}

// ORTBData and ORTBSegment objects together allow additional data about the related object (e.g., user, content)
// to be specified. This data may be from multiple sources whether from the exchange itself or third parties as
// specified by the id field. A bid request can mix data objects from multiple providers. The specific data providers
// in use should be published by the exchange a priori to its bidders. (Section 3.2.21 of OpenRTB 2.5 spec)
type ORTBData struct {
	ID         string        `json:"id,omitempty"`      // The unique ID used to identify the data
	Name       string        `json:"name,omitempty"`    // Data provider name
	Segments   []ORTBSegment `json:"segment,omitempty"` // Segment objects. (See section 3.2.22 of OpenRTB 2.5 spec)
	Extensions interface{}   `json:"ext,omitempty"`     // Exchange specific data related to the data
}

// ORTBSegment essentially key-value pairs that convey specific units of data. The parent Data object is a collection
// of such values from a given data provider. The specific segment names and value options must be published by the
// exchange a priori to its bidders. (Section 3.2.22 of OpenRTB 2.5 spec)
type ORTBSegment struct {
	ID         string      `json:"id,omitempty"`    // ID of a data provider’s segment applicable to the user.
	Name       string      `json:"name,omitempty"`  // Name of a data provider’s segment applicable to the user.
	Value      string      `json:"value,omitempty"` // String representation of the data segment value.
	Extensions interface{} `json:"ext,omitempty"`   // Exchange specific data related to the segment
}
