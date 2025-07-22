package zabbix

const (
	// SelectExtendedOutput may be given as a SelectQuery in search parameters
	// to return all available feilds for all objects in the search results.
	SelectExtendedOutput = "extend"

	// SelectCount may be given as a SelectQuery for supported search parameters
	// to return only the number of available search results, instead of the
	// search result details.
	SelectCount = "count"
)

// For `GetParameters` field: `SortOrder`
const (
	GetParametersSortOrderASC  = "ASC"
	GetParametersSortOrderDESC = "DESC"
)

// GetParameters struct is used as embedded struct for some other structs within package
//
// see for details: https://www.zabbix.com/documentation/5.0/manual/api/reference_commentary#common_get_method_parameters
type GetParameters struct {
	CountOutput            bool                   `json:"countOutput,omitempty"`
	Editable               bool                   `json:"editable,omitempty"`
	ExcludeSearch          bool                   `json:"excludeSearch,omitempty"`
	Filter                 map[string]interface{} `json:"filter,omitempty"`
	Limit                  int                    `json:"limit,omitempty"`
	Output                 SelectQuery            `json:"output,omitempty"`
	PreserveKeys           bool                   `json:"preservekeys,omitempty"`
	Search                 map[string]string      `json:"search,omitempty"`
	SearchByAny            bool                   `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool                   `json:"searchWildcardsEnabled,omitempty"`
	SortField              []string               `json:"sortfield,omitempty"`
	SortOrder              []string               `json:"sortorder,omitempty"` // has defined consts, see above
	StartSearch            bool                   `json:"startSearch,omitempty"`
}

// SelectQuery represents the query data type for a Zabbix API call.
// Wherever a SelectQuery is required, one of SelectFields, SelectExtendedOutput
// or SelectCount should be given.
type SelectQuery interface{}

// SelectFields may be given as a SelectQuery in search parameters where each
// member string is the name of a JSON field which should be returned for each
// search result.
//
// For example, for a Host search query:
//
//	query := SelectFields{ "hostid", "host", "name" }
type SelectFields []string
