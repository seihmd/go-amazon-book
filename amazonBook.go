package amazonbook

import "net/http"

type AmazonProductAPI struct {
	ID           string
	Secret       string
	AssociateTag string
	Host         string
	Client       *http.Client
}

// New returns AmazonProductAPI
func New(id, secret, assoc, host string) AmazonProductAPI {
	return AmazonProductAPI{
		ID:           id,
		Secret:       secret,
		AssociateTag: assoc,
		Host:         host,
	}
}

// ItemLookupByASIN takes a ASIN and returns the result
func (api AmazonProductAPI) ItemLookupByASIN(asin string) (*Response, error) {
	params := map[string]string{
		"ItemId":        asin,
		"ResponseGroup": "Large",
	}
	return api.genSignAndFetch("ItemLookup", params)
}

// ItemLookupByISBN takes a ISBN and returns the result
func (api AmazonProductAPI) ItemLookupByISBN(isbn string) (*Response, error) {
	params := map[string]string{
		"ItemId":        isbn,
		"IdType":        "ISBN",
		"ResponseGroup": "Large",
		"SearchIndex":   "Books",
	}
	return api.genSignAndFetch("ItemLookup", params)
}

// GetBrowseNodeNewReleases request BrowseNodeLookup with ResponseGroup NewReleases
func (api AmazonProductAPI) GetBrowseNodeNewReleases(nodeID string) (*Response, error) {
	params := map[string]string{
		"BrowseNodeId":  nodeID,
		"ResponseGroup": "NewReleases",
	}
	return api.genSignAndFetch("BrowseNodeLookup", params)
}

// GetBrowseNodeTopSellers request BrowseNodeLookup with ResponseGroup TopSellers
func (api AmazonProductAPI) GetBrowseNodeTopSellers(nodeID string) (*Response, error) {
	params := map[string]string{
		"BrowseNodeId":  nodeID,
		"ResponseGroup": "TopSellers",
	}
	return api.genSignAndFetch("BrowseNodeLookup", params)
}
