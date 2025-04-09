package testsdkgo

type DefaultPaginated struct {
	// MetaData contains metadata of the response, such as record count, pagination and other additional information.
	Meta DefaultPaginatedMeta `json:"_meta,omitempty"`
	// Errors specifies a list of errors that occurred, can be filled using error handlers.
	Errors DefaultPaginatedErrors `json:"errors,omitempty"`
}

type DefaultPaginatedErrors struct{}

type DefaultPaginatedMeta struct {
	Limit   int    `json:"limit,omitempty"`
	Total   int64  `json:"total,omitempty"`
	TraceID string `json:"traceId,omitempty"`
}

type DefaultResponse struct {
	// MetaData contains metadata of the response, such as record count, pagination and other additional information.
	Meta DefaultResponseMeta `json:"_meta,omitempty"`
	// Errors specifies a list of errors that occurred, can be filled using error handlers.
	Errors DefaultResponseErrors `json:"errors,omitempty"`
}

type DefaultResponseErrors struct{}

type DefaultResponseMeta struct {
	TraceID string `json:"traceId,omitempty"`
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type Pet struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
