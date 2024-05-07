package domain

type GetCompanionInfoRequestSwag struct {
	Destination string `json:"destination" in:"query=destination"`
}

type GetRouteInfoRequestSwag struct {
	OneOfPath string `json:"one_of_path" in:"query=one_of_path"`
}

type CreateRouteRequestPayloadSwag struct {
	Path []string `json:"path"`
}
type CreateRouteRequestSwag struct {
	Payload *CreateRouteRequestPayloadSwag `json:"payload" in:"body=json"`
}

type CreateCompanionRequestPayloadSwag struct {
	Destination string `json:"destination"`
}
type CreateCompanionRequestSwag struct {
	Payload *CreateCompanionRequestPayloadSwag `json:"payload" in:"body=json"`
}
