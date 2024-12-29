package crawler

const (
	TitleSelector       = "html head title"
	DescriptionSelector = "meta[name=\"description\" i][content]"
	IconSelector        = "link[rel=\"icon\" i], link[rel=\"shortcut icon\" i]"
	KeywordsSelector    = "meta[name=\"keywords\" i][content]"
	RobotsSelector      = "meta[name=\"robots\" i][content]"
	RatingSelector      = "meta[name=\"rating\" i][content]"
	CanonicalSelector   = "link[rel=\"canonical\" i][href]"
	AlternateSelector   = "link[rel=\"alternate\" i]"
)

const (
	OpenGraphSelector            = "meta[property*=\"og:\" i][content]" // Selector to get every OpenGraph element
	OpenGraphURLSelector         = "meta[property=\"og:url\" i][content]"
	OpenGraphSiteNameSelector    = "meta[property=\"og:site_name\" i][content]"
	OpenGraphTitleSelector       = "meta[property=\"og:title\" i][content]"
	OpenGraphDescriptionSelector = "meta[property=\"og:description\" i][content]"
	OpenGraphLocaleSelector      = "meta[property=\"og:locale\" i][content]"
	OpenGraphTypeSelector        = "meta[property=\"og:type\" i][content]"
	OpenGraphImageSelector       = "meta[property=\"og:image\" i][content]"
)
