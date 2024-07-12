package model

type HappinessIndexMap map[string]float64

/* HappinessByFacetId */

type HappinessByFacetIdResponse struct {
	Data HappinessByFacetIdData `json:"data"`
}

func NewHappinessByFacetIdResponse(data HappinessByFacetIdData) *HappinessByFacetIdResponse {
	return &HappinessByFacetIdResponse{Data: data}
}

type HappinessByFacetIdData struct {
	Facet   string         `json:"facet"`
	Indices []IndexMapping `json:"indices"`
}

func NewHappinessByFacetIdData(facet string) *HappinessByFacetIdData {
	return &HappinessByFacetIdData{
		Facet:   facet,
		Indices: []IndexMapping{},
	}
}

/* HappinessByFacetIds */

type HappinessByFacetIdsRequest struct {
	Counties []string `json:"counties" binding:"required"`
	Metrics  []string `json:"metrics"`
}

type HappinessByFacetIdsResponse struct {
	Data HappinessByFacetIdsData `json:"data"`
}

func NewHappinessByFacetIdsResponse(data HappinessByFacetIdsData) *HappinessByFacetIdsResponse {
	return &HappinessByFacetIdsResponse{Data: data}
}

type HappinessByFacetIdsData struct {
	Facet   string          `json:"facet"`
	Indices []IndexMapping  `json:"indices"`
	Metrics []MetricMapping `json:"metrics"`
}

func NewHappinessByFacetIdsData(facet string) *HappinessByFacetIdsData {
	return &HappinessByFacetIdsData{
		Facet:   facet,
		Indices: []IndexMapping{},
		Metrics: []MetricMapping{},
	}
}

type IndexMapping struct {
	Id    string  `json:"id"`
	Value float64 `json:"value"`
}

type MetricMapping struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
