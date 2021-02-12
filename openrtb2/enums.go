package openrtb

// ORTBBannerAd indicates the types of ads that can be accepted by the exchange unless restricted by
// publisher site settings
// OpenRTB 2.5 Section 5.2
type ORTBBannerAd int8

const (
	// ORTBBannerAdXHTMLTextAd is an Ad Type of XHTML Text Ad (usually mobile)
	ORTBBannerAdXHTMLTextAd ORTBBannerAd = iota + 1
	// ORTBBannerAdXHTMLBannerAd is an Ad Type of XHTML Banner Ad. (usually mobile)
	ORTBBannerAdXHTMLBannerAd ORTBBannerAd = iota
	// ORTBBannerAdJavaScriptAd is an Ad Type of JavaScript Ad (must be valid XHTML i.e., Script Tags Included)
	ORTBBannerAdJavaScriptAd ORTBBannerAd = iota
	// ORTBBannerAdiframe is an Ad Type of iframe
	ORTBBannerAdiframe ORTBBannerAd = iota
)

// ORTBCreativeAttribute specifies a standard list of creative attributes that can describe an ad being served or
// serve as restrictions of thereof.
// OpenRTB 2.5 Section 5.3
type ORTBCreativeAttribute int8

const (
	// ORTBCreativeAttributeAudioAdAuto is Creative Attribute Audio Ad Auto (Auto-Play)
	ORTBCreativeAttributeAudioAdAuto ORTBCreativeAttribute = iota + 1
	// ORTBCreativeAttributeAudioAdInteractive is Creative Attribute Audio Ad Interactive (User Initiated)
	ORTBCreativeAttributeAudioAdInteractive ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeExpandableAuto is Creative Attribute Expandable Auto (Automatic)
	ORTBCreativeAttributeExpandableAuto ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeExpandableClick is Creative Attribute Expandable Click (User Initiated - Click)
	ORTBCreativeAttributeExpandableClick ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeExpandableRollover is Creative Attribute Expandable Rollover (User Initiated - Rollover)
	ORTBCreativeAttributeExpandableRollover ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeInBannerVideoAdAuto is Creative Attribute In-Banner Video Ad Auto (Auto-Play)
	ORTBCreativeAttributeInBannerVideoAdAuto ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeInBannerVideoAdInteractive is Creative Attribute In-Banner Video Ad Interactive (User Initiated)
	ORTBCreativeAttributeInBannerVideoAdInteractive ORTBCreativeAttribute = iota
	// ORTBCreativeAttributePop is Creative Attribute Pop (e.g., Over, Under, or Upon Exit)
	ORTBCreativeAttributePop ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeProvocative is Creative Attribute Provocative or Suggestive Imagery
	ORTBCreativeAttributeProvocative ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeShakyFlashing is Creative Attribute Shaky, Flashing, Flickering, Extreme Animation, Smileys
	ORTBCreativeAttributeShakyFlashing ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeSurveys Creative Attribute Surveys
	ORTBCreativeAttributeSurveys ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeTextOnly Creative Attribute Text Only
	ORTBCreativeAttributeTextOnly ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeUserInteractive Creative Attribute User Interactive (e.g., Embedded Games)
	ORTBCreativeAttributeUserInteractive ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeWindowsDialogorAlertStyle Creative Attribute Windows Dialog or Alert Style
	ORTBCreativeAttributeWindowsDialogorAlertStyle ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeHasAudioOnOffButton Creative Attribute Has Audio On/Off Button
	ORTBCreativeAttributeHasAudioOnOffButton ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeAdProvidesSkipButton Creative Attribute Ad Provides Skip Button (e.g. VPAID-rendered skip button on pre-roll video)
	ORTBCreativeAttributeAdProvidesSkipButton ORTBCreativeAttribute = iota
	// ORTBCreativeAttributeAdobeFlash Creative Attribute Adobe Flash
	ORTBCreativeAttributeAdobeFlash ORTBCreativeAttribute = iota
)

// ORTBAdPosition specifies the position of the ad as a relative measure of visibility or prominence. This
// OpenRTB table has values derived from the Inventory Quality Guidelines (IQG). Practitioners should
// keep in sync with updates to the IQG values as published on IAB.com. Values “4” - “7” apply to apps per
// the mobile addendum to IQG version 2.1.
// OpenRTB 2.5 Section 5.4
type ORTBAdPosition int8

const (
	// ORTBAdPositionUnknown is position Unknown
	ORTBAdPositionUnknown ORTBAdPosition = iota + 1
	// ORTBAdPositionAbovetheFold is position Above the Fold
	ORTBAdPositionAbovetheFold ORTBAdPosition = iota
	// ORTBAdPositionDEPRECATED is position DEPRECATED (May or may not be initially visible depending on screen size/resolution.)
	ORTBAdPositionDEPRECATED ORTBAdPosition = iota
	// ORTBAdPositionBelowtheFold is position Below the Fold
	ORTBAdPositionBelowtheFold ORTBAdPosition = iota
	// ORTBAdPositionHeader is position Header
	ORTBAdPositionHeader ORTBAdPosition = iota
	// ORTBAdPositionFooter is position Footer
	ORTBAdPositionFooter ORTBAdPosition = iota
	// ORTBAdPositionSidebar is position Sidebar
	ORTBAdPositionSidebar ORTBAdPosition = iota
	// ORTBAdPositionFullScreen is position Full Screen
	ORTBAdPositionFullScreen ORTBAdPosition = iota
)

// ORTBExpandableDirection is the direction in which an expandable ad may expand, given the positioning
// of the ad unit on the page and constraints imposed by the content.
// OpenRTB 2.5 Section 5.5
type ORTBExpandableDirection int8

const (
	// ORTBExpandableDirectionLeft is direction Left
	ORTBExpandableDirectionLeft ORTBExpandableDirection = iota + 1
	// ORTBExpandableDirectionRight is direction Right
	ORTBExpandableDirectionRight ORTBExpandableDirection = iota
	// ORTBExpandableDirectionUp is direction Up
	ORTBExpandableDirectionUp ORTBExpandableDirection = iota
	// ORTBExpandableDirectionDown is direction Down
	ORTBExpandableDirectionDown ORTBExpandableDirection = iota
	// ORTBExpandableDirectionFullScreen is direction Full Screen
	ORTBExpandableDirectionFullScreen ORTBExpandableDirection = iota
)

// ORTBAPIFramework is an API framework supported by the publisher.
// OpenRTB 2.5 Section 5.6
type ORTBAPIFramework int8

const (
	// ORTBAdPositionVPAID10 is API Framework VPAID 1.0
	ORTBAdPositionVPAID10 ORTBAdPosition = iota + 1
	// ORTBAdPositionVPAID20 is API Framework VPAID 2.0
	ORTBAdPositionVPAID20 ORTBAdPosition = iota
	// ORTBAdPositionMRAID1 is API Framework MRAID-1
	ORTBAdPositionMRAID1 ORTBAdPosition = iota
	// ORTBAdPositionORMMA is API Framework ORMMA
	ORTBAdPositionORMMA ORTBAdPosition = iota
	// ORTBAdPositionMRAID2 is API Framework MRAID-2
	ORTBAdPositionMRAID2 ORTBAdPosition = iota
	// ORTBAdPositionMRAID3 is API Framework MRAID-3
	ORTBAdPositionMRAID3 ORTBAdPosition = iota
)

// ORTBVideoLinearity indicates the options for video linearity. “In-stream” or “linear” video refers to
// preroll, post-roll, or mid-roll video ads where the user is forced to watch ad in order to see the video
// content. “Overlay” or “non-linear” refer to ads that are shown on top of the video content.
// This OpenRTB table has values derived from the Inventory Quality Guidelines (IQG). Practitioners should
// keep in sync with updates to the IQG values.
// OpenRTB 2.5 Section 5.7
type ORTBVideoLinearity int8

const (
	// ORTBVideoLinearityLinear is linearity of of Linear / In-Stream
	ORTBVideoLinearityLinear ORTBVideoLinearity = iota + 1
	// ORTBVideoLinearityNonLinear is linearity of of Non-Linear / Overlay
	ORTBVideoLinearityNonLinear ORTBVideoLinearity = iota
)

// ORTBProtocol represents the options for the various bid response protocols that could be
// supported by an exchange.
// OpenRTB 2.5 Section 5.8
type ORTBProtocol int8

const (
	// ORTBProtocolVAST10 is VAST 1.0 protocol
	ORTBProtocolVAST10 ORTBProtocol = iota + 1
	// ORTBProtocolVAST20 is VAST 2.0 protocol
	ORTBProtocolVAST20 ORTBProtocol = iota
	// ORTBProtocolVAST30 is VAST 3.0 protocol
	ORTBProtocolVAST30 ORTBProtocol = iota
	// ORTBProtocolVAST10Wrapper is VAST 1.0 Wrapper protocol
	ORTBProtocolVAST10Wrapper ORTBProtocol = iota
	// ORTBProtocolVAST20Wrapper is VAST 2.0 Wrapper protocol
	ORTBProtocolVAST20Wrapper ORTBProtocol = iota
	// ORTBProtocolVAST30Wrapper is VAST 3.0 Wrapper protocol
	ORTBProtocolVAST30Wrapper ORTBProtocol = iota
	// ORTBProtocolVAST40 is VAST 4.0 protocol
	ORTBProtocolVAST40 ORTBProtocol = iota
	// ORTBProtocolVAST40Wrapper is VAST 4.0 Wrapper protocol
	ORTBProtocolVAST40Wrapper ORTBProtocol = iota
	// ORTBProtocolDAAST10 is DAAST 1.0 protocol
	ORTBProtocolDAAST10 ORTBProtocol = iota
	// ORTBProtocolDAAST10Wrapper is DAAST 1.0 Wrapper protocol
	ORTBProtocolDAAST10Wrapper ORTBProtocol = iota
)

// ORTBVideoPlacement represents the various types of video placements derived largely from the IAB Digital
// Video Guidelines
// OpenRTB 2.5 Section 5.9
type ORTBVideoPlacement int8

const (
	// ORTBVideoPlacementInStream In-Stream Played before, during or after the streaming video content that the
	// consumer has requested (e.g., Pre-roll, Mid-roll, Post-roll).
	// OpenRTB API Specification Version 2.5 IAB Technology Lab
	// www.iab.com/openrtb Page 48
	ORTBVideoPlacementInStream ORTBVideoPlacement = iota + 1
	// ORTBVideoPlacementInBanner Exists within a web banner that leverages the banner space to deliver a video
	// experience as opposed to another static or rich media format. The format relies on the existence of display
	// ad inventory on the page for its delivery.
	ORTBVideoPlacementInBanner ORTBVideoPlacement = iota
	// ORTBVideoPlacementInArticle Loads and plays dynamically between paragraphs of editorial content; existing as
	//  a standalone branded message.
	ORTBVideoPlacementInArticle ORTBVideoPlacement = iota
	// ORTBVideoPlacementInFeed Found in content, social, or product feeds.
	ORTBVideoPlacementInFeed ORTBVideoPlacement = iota
	// ORTBVideoPlacementInterstitial Covers the entire or a portion of screen area, but is always on screen
	// while displayed (i.e. cannot be scrolled out of view). Note that a full-screen interstitial (e.g., in
	// mobile) can be distinguished from a floating/slider unit by the imp.instl field.
	ORTBVideoPlacementInterstitial ORTBVideoPlacement = iota
)

// ORTBPlaybackMethod represents the various playback methods
// OpenRTB 2.5 Section 5.10
type ORTBPlaybackMethod int8

const (
	// ORTBPlaybackMethodLoadWithSound Initiates on Page Load with Sound On
	ORTBPlaybackMethodLoadWithSound ORTBPlaybackMethod = iota + 1
	// ORTBPlaybackMethodLoadNoSound Initiates on Page Load with Sound Off by Default
	ORTBPlaybackMethodLoadNoSound ORTBPlaybackMethod = iota
	// ORTBPlaybackMethodClickWithSound Initiates on Click with Sound On
	ORTBPlaybackMethodClickWithSound ORTBPlaybackMethod = iota
	// ORTBPlaybackMethodMouseoverWithSound Initiates on Mouse-Over with Sound On
	ORTBPlaybackMethodMouseoverWithSound ORTBPlaybackMethod = iota
	// ORTBPlaybackMethodOnEnterWithSound Initiates on Entering Viewport with Sound On
	ORTBPlaybackMethodOnEnterWithSound ORTBPlaybackMethod = iota
	// ORTBPlaybackMethodOnEnterNoSound Initiates on Entering Viewport with Sound Off by Default
	ORTBPlaybackMethodOnEnterNoSound ORTBPlaybackMethod = iota
)

// ORTBPlaybackCessationMode represents the various modes for when playback terminates.
// OpenRTB 2.5 Section 5.11
type ORTBPlaybackCessationMode int8

const (
	// ORTBPlaybackCessationModeOnCompleation On Video Completion or when Terminated by User
	ORTBPlaybackCessationModeOnCompleation ORTBPlaybackCessationMode = iota + 1
	// ORTBPlaybackCessationModeOnLeaveViewport On Leaving Viewport or when Terminated by User
	ORTBPlaybackCessationModeOnLeaveViewport ORTBPlaybackCessationMode = iota
	// ORTBPlaybackCessationModeUntilFinished On Leaving Viewport Continues as a Floating/Slider
	// Unit until Video Completion or when Terminated by User
	ORTBPlaybackCessationModeUntilFinished ORTBPlaybackCessationMode = iota
)

// ORTBStartDelay represents the various options for the video or audio start delay. If the start
// delay value is greater than 0, then the position is mid-roll and the value indicates the start
// delay.
// OpenRTB 2.5 Section 5.12
type ORTBStartDelay int64

// Note: Any StartDelay > 0 represents start delay in seconds
const (
	// ORTBStartDelayPreRoll is the preroll flag
	ORTBStartDelayPreRoll ORTBStartDelay = 0
	// ORTBStartDelayGenericMidRoll
	ORTBStartDelayGenericMidRoll ORTBStartDelay = -1
	// ORTBStartDelayGenericPostRoll
	ORTBStartDelayGenericPostRoll ORTBStartDelay = -2
)

// ORTBProductionQuality represents the options for content quality. These values are defined by the
// IAB; refer to www.iab.com/wp-content/uploads/2015/03/long-form-video-final.pdf for more information.
// OpenRTB 2.5 Section 5.13
type ORTBProductionQuality int8

const (
	// ORTBProductionQualityUnknown is Unknown quality
	ORTBProductionQualityUnknown ORTBProductionQuality = iota
	// ORTBProductionQualityProfessionallyProduced is Professionally Produced quality
	ORTBProductionQualityProfessionallyProduced ORTBProductionQuality = iota
	// ORTBProductionQualityProsumer is Prosumer quality
	ORTBProductionQualityProsumer ORTBProductionQuality = iota
	// ORTBProductionQualityUserGenerated is User Generated (UGC) quality
	ORTBProductionQualityUserGenerated ORTBProductionQuality = iota
)

// ORTBCompanion represents the options to indicate markup types allowed for companion ads that apply to
// video and audio ads. This table is derived from VAST 2.0+ and DAAST 1.0 specifications. Refer to
// www.iab.com/guidelines/digital-video-suite for more information.
// OpenRTB 2.5 Section 5.14
type ORTBCompanion int8

const (
	// ORTBCompanionStatic is resource of type Static Resource
	ORTBCompanionStatic ORTBCompanion = iota + 1
	// ORTBCompanionHTML is resource of type HTML Resource
	ORTBCompanionHTML ORTBCompanion = iota
	// ORTBCompanioniframe is resource of type iframe Resource
	ORTBCompanioniframe ORTBCompanion = iota
)

// ORTBContentDeliveryMethod represents the various options for the delivery of video or audio content.
// OpenRTB 2.5 Section 5.15
type ORTBContentDeliveryMethod int8

const (
	// ORTBContentDeliveryMethodStreaming is Streaming delivery method
	ORTBContentDeliveryMethodStreaming ORTBContentDeliveryMethod = iota
	// ORTBContentDeliveryMethodProgressive is Progressive delivery method
	ORTBContentDeliveryMethodProgressive ORTBContentDeliveryMethod = iota
	// ORTBContentDeliveryMethodDownload is Download delivery method
	ORTBContentDeliveryMethodDownload ORTBContentDeliveryMethod = iota
)

// ORTBFeedType represents the types of feeds, typically for audio
// OpenRTB 2.5 Section 5.16
type ORTBFeedType int8

const (
	// ORTBFeedTypeMusicService is Music Service feed
	ORTBFeedTypeMusicService ORTBFeedType = iota + 1
	// ORTBFeedTypeFMAMBroadcast is FM/AM Broadcast feed
	ORTBFeedTypeFMAMBroadcast ORTBFeedType = iota
	// ORTBFeedTypePodcast is Podcast feed
	ORTBFeedTypePodcast ORTBFeedType = iota
)

// ORTBVolumeNormalization represents the the types of volume normalization modes, typically for audio.
// OpenRTB 2.5 Section 5.17
type ORTBVolumeNormalization int8

const (
	// ORTBVolumeNormalizationNone normalizes by None
	ORTBVolumeNormalizationNone ORTBVolumeNormalization = iota
	// ORTBVolumeNormalizationAdVolumeAverageNormalizedtoContent normalizes by Ad Volume Average Normalized to Content
	ORTBVolumeNormalizationAdVolumeAverageNormalizedtoContent ORTBVolumeNormalization = iota
	// ORTBVolumeNormalizationAdVolumePeakNormalizedtoContent normalizes by Ad Volume Peak Normalized to Content
	ORTBVolumeNormalizationAdVolumePeakNormalizedtoContent ORTBVolumeNormalization = iota
	// ORTBVolumeNormalizationAdLoudnessNormalizedtoContent normalizes by Ad Loudness Normalized to Content
	ORTBVolumeNormalizationAdLoudnessNormalizedtoContent ORTBVolumeNormalization = iota
	// ORTBVolumeNormalizationCustomVolumeNormalization normalizes by Custom Volume Normalization
	ORTBVolumeNormalizationCustomVolumeNormalization ORTBVolumeNormalization = iota
)

// ORTBContentContext represents the various options for indicating the type of content being used or consumed
// by the user in which the impression will appear. This OpenRTB table has values derived from the
// Inventory Quality Guidelines (IQG). Practitioners should keep in sync with updates to the IQG values.
// OpenRTB 2.5 Section 5.18
type ORTBContentContext int8

const (
	// ORTBContentContextVideo is a Video (i.e., video file or stream such as Internet TV broadcasts)
	ORTBContentContextVideo ORTBContentContext = iota + 1
	// ORTBContentContextGame is a Game (i.e., an interactive software game)
	ORTBContentContextGame ORTBContentContext = iota
	// ORTBContentContextMusic is a Music (i.e., audio file or stream such as Internet radio broadcasts)
	ORTBContentContextMusic ORTBContentContext = iota
	// ORTBContentContextApplication is a Application (i.e., an interactive software application)
	ORTBContentContextApplication ORTBContentContext = iota
	// ORTBContentContextText is a Text (i.e., primarily textual document such as a web page, eBook, or news article)
	ORTBContentContextText ORTBContentContext = iota
	// ORTBContentContextOther is a Other (i.e., none of the other categories applies)
	ORTBContentContextOther ORTBContentContext = iota
	// ORTBContentContextUnknown is a Unknown
	ORTBContentContextUnknown ORTBContentContext = iota
)

// ORTBIQGMediaRating represents the media ratings used in describing content based on the IQG 2.1
// categorization. Refer to www.iab.com/guidelines/digital-video-suite for more information.
// OpenRTB 2.5 Section 5.19
type ORTBIQGMediaRating int8

const (
	// ORTBIQGMediaRatingAllAudiences rated for All Audiences
	ORTBIQGMediaRatingAllAudiences ORTBIQGMediaRating = iota + 1
	// ORTBIQGMediaRatingEveryoneOver12 rated for Everyone Over 12
	ORTBIQGMediaRatingEveryoneOver12 ORTBIQGMediaRating = iota
	// ORTBIQGMediaRatingMatureAudiences rated for Mature Audiences
	ORTBIQGMediaRatingMatureAudiences ORTBIQGMediaRating = iota
)

// ORTBLocationType represents the options to indicate how the geographic information was determined.
// OpenRTB 2.5 Section 5.20
type ORTBLocationType int8

const (
	// ORTBLocationTypeGPSLocationServices location derived from GPS/Location Services
	ORTBLocationTypeGPSLocationServices ORTBLocationType = iota + 1
	// ORTBLocationTypeIPAddress location derived from IP Address
	ORTBLocationTypeIPAddress ORTBLocationType = iota
	// ORTBLocationTypeUserprovided location derived from User provided (e.g., registration data)
	ORTBLocationTypeUserprovided ORTBLocationType = iota
)

// ORTBDeviceType represets the type of device from which the impression originated.
// OpenRTB version 2.2 of the specification added distinct values for Mobile and Tablet. It is
// recommended that any bidder adding support for 2.2 treat a value of 1 as an acceptable alias of 4 & 5.
// This OpenRTB table has values derived from the Inventory Quality Guidelines (IQG). Practitioners should
// keep in sync with updates to the IQG values.
// OpenRTB 2.5 Section 5.21
type ORTBDeviceType int8

const (
	// ORTBDeviceTypeMobileTablet device of type Mobile/Tablet
	ORTBDeviceTypeMobileTablet ORTBDeviceType = iota + 1
	// ORTBDeviceTypePersonalComputer device of type Personal Computer
	ORTBDeviceTypePersonalComputer ORTBDeviceType = iota
	// ORTBDeviceTypeConnectedTV device of type Connected TV
	ORTBDeviceTypeConnectedTV ORTBDeviceType = iota
	// ORTBDeviceTypePhone device of type Phone
	ORTBDeviceTypePhone ORTBDeviceType = iota
	// ORTBDeviceTypeTablet device of type Tablet
	ORTBDeviceTypeTablet ORTBDeviceType = iota
	// ORTBDeviceTypeConnectedDevice device of type Connected Device
	ORTBDeviceTypeConnectedDevice ORTBDeviceType = iota
	// ORTBDeviceTypeSetTopBox device of type Set Top Box
	ORTBDeviceTypeSetTopBox ORTBDeviceType = iota
)

// ORTBConnectionType represents the various options for the type of device connectivity.
// OpenRTB 2.5 Section 5.22
type ORTBConnectionType int8

const (
	// ORTBConnectionTypeUnknown is a connection of type Unknown
	ORTBConnectionTypeUnknown ORTBConnectionType = iota
	// ORTBConnectionTypeEthernet is a connection of type Ethernet
	ORTBConnectionTypeEthernet ORTBConnectionType = iota
	// ORTBConnectionTypeWIFI is a connection of type WIFI
	ORTBConnectionTypeWIFI ORTBConnectionType = iota
	// ORTBConnectionTypeCellularNetworkUnknown is a connection of type Cellular Network - Unknown Generation
	ORTBConnectionTypeCellularNetworkUnknown ORTBConnectionType = iota
	// ORTBConnectionType2G is a connection of type Cellular Network - 2G
	ORTBConnectionType2G ORTBConnectionType = iota
	// ORTBConnectionType3G is a connection of type Cellular Network - 3G
	ORTBConnectionType3G ORTBConnectionType = iota
	// ORTBConnectionType4G is a connection of type Cellular Network - 4G
	ORTBConnectionType4G ORTBConnectionType = iota
)

// ORTBIPLocationService represents the services and/or vendors used for resolving IP addresses to geolocations.
// OpenRTB 2.5 Section 5.23
type ORTBIPLocationService int8

const (
	// ORTBIPLocationServiceip2location is ip2location Location Service
	ORTBIPLocationServiceip2location ORTBIPLocationService = iota + 1
	// ORTBIPLocationServiceNeustar is Neustar (Quova) Location Service
	ORTBIPLocationServiceNeustar ORTBIPLocationService = iota
	// ORTBIPLocationServiceMaxMind is MaxMind Location Service
	ORTBIPLocationServiceMaxMind ORTBIPLocationService = iota
	// ORTBIPLocationServiceNetAcuity is NetAcuity (Digital Element) Location Service
	ORTBIPLocationServiceNetAcuity ORTBIPLocationService = iota
)

// ORTBNoBidReason represents the options for a bidder to signal the exchange as to why it did not offer a bid
// for the impression.
// OpenRTB 2.5 Section 5.24
type ORTBNoBidReason int8

const (
	// ORTBNoBidReasonUnknownError exchange said Unknown Error for no bid reason
	ORTBNoBidReasonUnknownError ORTBNoBidReason = iota
	// ORTBNoBidReasonTechnicalError exchange said Technical Error for no bid reason
	ORTBNoBidReasonTechnicalError ORTBNoBidReason = iota
	// ORTBNoBidReasonInvalidRequest exchange said Invalid Request for no bid reason
	ORTBNoBidReasonInvalidRequest ORTBNoBidReason = iota
	// ORTBNoBidReasonKnownWebSpider exchange said Known Web Spider for no bid reason
	ORTBNoBidReasonKnownWebSpider ORTBNoBidReason = iota
	// ORTBNoBidReasonSuspectedNonHuman exchange said Suspected Non-Human Traffic for no bid reason
	ORTBNoBidReasonSuspectedNonHuman ORTBNoBidReason = iota
	// ORTBNoBidReasonProxyIP exchange said Cloud, Data center, or Proxy IP for no bid reason
	ORTBNoBidReasonProxyIP ORTBNoBidReason = iota
	// ORTBNoBidReasonUnsupportedDevice exchange said Unsupported Device for no bid reason
	ORTBNoBidReasonUnsupportedDevice ORTBNoBidReason = iota
	// ORTBNoBidReasonBlockedPublisherRrSite exchange said Blocked Publisher or Site for no bid reason
	ORTBNoBidReasonBlockedPublisherRrSite ORTBNoBidReason = iota
	// ORTBNoBidReasonUnmatchedUser exchange said Unmatched User for no bid reason
	ORTBNoBidReasonUnmatchedUser ORTBNoBidReason = iota
	// ORTBNoBidReasonDailyReaderCapMet exchange said Daily Reader Cap Met for no bid reason
	ORTBNoBidReasonDailyReaderCapMet ORTBNoBidReason = iota
	// ORTBNoBidReasonDailyDomainCapMet exchange said Daily Domain Cap Met for no bid reason
	ORTBNoBidReasonDailyDomainCapMet ORTBNoBidReason = iota
)

// ORTBLossReason represents options for an exchange to inform a bidder as to the reason why they did
// not win an impression.
// OpenRTB 2.5 Section 5.25
type ORTBLossReason int64

// NOTE: Reason code >= 1000 is Exchange specific reason (should be communicated to bidders a priori)
const (
	// ORTBLossReasonBidWon lost bid due to Bid Won
	ORTBLossReasonBidWon ORTBLossReason = iota
	// ORTBLossReasonInternalError lost bid due to Internal Error
	ORTBLossReasonInternalError ORTBLossReason = iota
	// ORTBLossReasonImpressionOpportunityExpired lost bid due to Impression Opportunity Expired
	ORTBLossReasonImpressionOpportunityExpired ORTBLossReason = iota
	// ORTBLossReasonInvalidBidResponse lost bid due to Invalid Bid Response
	ORTBLossReasonInvalidBidResponse ORTBLossReason = iota
	// ORTBLossReasonInvalidDealID lost bid due to Invalid Deal ID
	ORTBLossReasonInvalidDealID ORTBLossReason = iota
	// ORTBLossReasonInvalidAuctionID lost bid due to Invalid Auction ID
	ORTBLossReasonInvalidAuctionID ORTBLossReason = iota
	// ORTBLossReasonInvalidAdvertiserDomain lost bid due to Invalid (i.e., malformed) Advertiser Domain
	ORTBLossReasonInvalidAdvertiserDomain ORTBLossReason = iota
	// ORTBLossReasonMissingMarkup lost bid due to Missing Markup
	ORTBLossReasonMissingMarkup ORTBLossReason = iota
	// ORTBLossReasonMissingCreativeID lost bid due to Missing Creative ID
	ORTBLossReasonMissingCreativeID ORTBLossReason = iota
	// ORTBLossReasonMissingBidPrice lost bid due to Missing Bid Price
	ORTBLossReasonMissingBidPrice ORTBLossReason = iota
	// ORTBLossReasonMissingMinimumCreativeApprovalData lost bid due to Missing Minimum Creative Approval Data
	ORTBLossReasonMissingMinimumCreativeApprovalData ORTBLossReason = iota
	// ORTBLossReasonBidwasBelowAuctionFloor lost bid due to Bid was Below Auction Floor
	ORTBLossReasonBidwasBelowAuctionFloor ORTBLossReason = iota
	// ORTBLossReasonBidwasBelowDealFloor lost bid due to Bid was Below Deal Floor
	ORTBLossReasonBidwasBelowDealFloor ORTBLossReason = iota
	// ORTBLossReasonLosttoHigherBid lost bid due to Lost to Higher Bid
	ORTBLossReasonLosttoHigherBid ORTBLossReason = iota
	// ORTBLossReasonLosttoaBidforaPMPDeal lost bid due to Lost to a Bid for a PMP Deal
	ORTBLossReasonLosttoaBidforaPMPDeal ORTBLossReason = iota
	// ORTBLossReasonBuyerSeatBlocked lost bid due to Buyer Seat Blocked
	ORTBLossReasonBuyerSeatBlocked ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredGeneralreasonunknown lost bid due to Creative Filtered - General; reason unknown.
	ORTBLossReasonCreativeFilteredGeneralreasonunknown ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredPendingprocessingbyExchange lost bid due to Creative Filtered - Pending processing by Exchange (e.g., approval, transcoding, etc.)
	ORTBLossReasonCreativeFilteredPendingprocessingbyExchange ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredDisapprovedbyExchange lost bid due to Creative Filtered - Disapproved by Exchange
	ORTBLossReasonCreativeFilteredDisapprovedbyExchange ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredSizeNotAllowed lost bid due to Creative Filtered - Size Not Allowed
	ORTBLossReasonCreativeFilteredSizeNotAllowed ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredIncorrectCreativeFormat lost bid due to Creative Filtered - Incorrect Creative Format
	ORTBLossReasonCreativeFilteredIncorrectCreativeFormat ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredAdvertiserExclusions lost bid due to Creative Filtered - Advertiser Exclusions
	ORTBLossReasonCreativeFilteredAdvertiserExclusions ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredAppBundleExclusions lost bid due to Creative Filtered - App Bundle Exclusions
	ORTBLossReasonCreativeFilteredAppBundleExclusions ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredNotSecure lost bid due to Creative Filtered - Not Secure
	ORTBLossReasonCreativeFilteredNotSecure ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredLanguageExclusions lost bid due to Creative Filtered - Language Exclusions
	ORTBLossReasonCreativeFilteredLanguageExclusions ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredCategoryExclusions lost bid due to Creative Filtered - Category Exclusions
	ORTBLossReasonCreativeFilteredCategoryExclusions ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredCreativeAttributeExclusions lost bid due to Creative Filtered - Creative Attribute Exclusions
	ORTBLossReasonCreativeFilteredCreativeAttributeExclusions ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredAdTypeExclusions lost bid due to Creative Filtered - Ad Type Exclusions
	ORTBLossReasonCreativeFilteredAdTypeExclusions ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredAnimationTooLong lost bid due to Creative Filtered - Animation Too Long
	ORTBLossReasonCreativeFilteredAnimationTooLong ORTBLossReason = iota
	// ORTBLossReasonCreativeFilteredNotAllowedinPMPDeal lost bid due to Creative Filtered - Not Allowed in PMP Deal
	ORTBLossReasonCreativeFilteredNotAllowedinPMPDeal ORTBLossReason = iota
)
