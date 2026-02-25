package wmtracker

type TopOrdersResponse struct {
	APIVersion string         `json:"apiVersion"`
	Data       TopOrderData   `json:"data"`
	Error      *ResponseError `json:"error"` // pointer so null is allowed
}

type OrdersResponse struct {
	APIVersion string          `json:"apiVersion"`
	Data       []OrderWithUser `json:"data"`
	Error      *ResponseError  `json:"error"` // pointer so null is allowed
}

type OrderWithUser struct {
	Order
	User UserShort `json:"user"` // Represents the user who created the order, with basic profile information.
}

type TopOrderData struct {
	Buy  []OrderWithUser `json:"buy"`
	Sell []OrderWithUser `json:"sell"`
}

type ResponseError struct {
	Request []string `json:"request"`
}

type UserShort struct {
	Id         string `json:"id"`
	IngameName string `json:"ingameName"`       // In-game name of the user.
	Avatar     string `json:"avatar,omitempty"` // Optional avatar image.
	Reputation int16  `json:"reputation"`       // Reputation score.
	Locale     string `json:"locale"`           // Preferred communication language (e.g., 'en', 'ko', 'es').
	Platform   string `json:"platform"`         // Gaming platform used by the user.
	Crossplay  bool   `json:"crossplay"`

	Status   string   `json:"status"`   // Current status of the user.
	Activity Activity `json:"activity"` // Addition to the status, current activity of the user.
	LastSeen string   `json:"lastSeen"` // Timestamp of the user's last online presence.
}

type Order struct {
	Id         string `json:"id"`                   // Is the unique identifier of the order.
	Type       string `json:"type"`                 // Specifies whether the order is a 'buy' or 'sell'.
	Platinum   int32  `json:"platinum"`             // Is the total platinum currency involved in the order.
	Quantity   int32  `json:"quantity"`             // Represents the number of items included in the order.
	PerTrade   int8   `json:"perTrade,omitempty"`   // (optional) indicates the items quantity per transaction.
	Rank       *int8  `json:"rank,omitempty"`       // (optional) specifies the rank or level of the item in the order.
	Charges    *int8  `json:"charges,omitempty"`    // (optional) specifies number of charges left (used in requiem mods).
	Subtype    string `json:"subtype,omitempty"`    // (optional) defines the specific subtype or category of the item.
	AmberStars *int8  `json:"amberStars,omitempty"` // (optional) denotes the count of amber stars in a sculpture order.
	CyanStars  *int8  `json:"cyanStars,omitempty"`  // (optional) denotes the count of cyan stars in a sculpture order.
	Visible    bool   `json:"visible"`              // (auth\mod) Indicates whether the order is publicly visible or not.
	CreatedAt  string `json:"createdAt"`            // Records the creation time of the order.
	UpdatedAt  string `json:"updatedAt"`            // Records the last modification time of the order.
	ItemId     string `json:"itemId"`               // Is the unique identifier of the item involved in the order.
	Group      string `json:"group"`                // User-defined group to which the order belongs
}

type Activity struct {
	Type      ActivityType `json:"type" `               // Name of the activity (e.g., 'on mission', 'dojo').
	Details   string       `json:"details,omitempty"`   // Optional specifics about the activity (e.g., mission name, solo/squad status).
	StartedAt string       `json:"startedAt,omitempty"` // Timestamp of the activity start.
}

type ActivityType string

const (
	UNKNOWN    ActivityType = "unknown"
	IDLE       ActivityType = "idle"
	ON_MISSION ActivityType = "on mission"
	IN_DOJO    ActivityType = "in dojo"
	IN_ORBITER ActivityType = "in orbiter"
	IN_RELAY   ActivityType = "in relay"
)

type RequestData struct {
	Slug string
	Rank int8
}
