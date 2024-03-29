package obsidian

import (
	"context"
	_ "embed"
	"encoding/json"
	"time"
)

func (o ObsidianClient) GetAllEvents(ctx context.Context, cursor string) (*GetAllEventsResponse, error) {
	response, err := o.NewRequest(ctx, GqlOperation{
		OperationName: "getAllEvents",
		Query:         getAllEventsQuery,
		Variables: struct {
			Cursor string `json:"cursor"`
		}{Cursor: cursor},
	})
	if err != nil {
		return nil, err
	}

	data := &GetAllEventsResponse{}
	if err := json.NewDecoder(response.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

//go:embed gql/get_all_events.gql
var getAllEventsQuery string

type Event struct {
	Actors []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		User struct {
			Id string `json:"id"`
		} `json:"user"`
		IsAdmin bool `json:"isAdmin"`
		Email   struct {
			Raw string `json:"raw"`
		} `json:"email"`
		GeneralProperties struct {
			JobTitle string `json:"jobTitle"`
		} `json:"generalProperties"`
	} `json:"actors"`
	Targets []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		User struct {
			Id string `json:"id"`
		} `json:"user"`
		IsAdmin bool `json:"isAdmin"`
		Email   struct {
			Raw string `json:"raw"`
		} `json:"email"`
	} `json:"targets"`
	BrowserId                string    `json:"browserId"`
	Datetime                 time.Time `json:"datetime"`
	DeviceId                 string    `json:"deviceId"`
	EventFrequency           string    `json:"eventFrequency"`
	HumanReadableDescription struct {
		Plain  string `json:"plain"`
		Tagged string `json:"tagged"`
	} `json:"humanReadableDescription"`
	Id        string `json:"id"`
	IPAddress struct {
		Raw                 string `json:"raw"`
		Isp                 string `json:"isp"`
		Asn                 string `json:"asn"`
		IsMobile            bool   `json:"isMobile"`
		Hostname            string `json:"hostname"`
		VpnProxyDescription string `json:"vpnProxyDescription"`
		VpnProxyType        string `json:"vpnProxyType"`
	} `json:"ipAddress"`
	IsPrivilegedEvent bool `json:"isPrivilegedEvent"`
	Location          struct {
		City        string `json:"city"`
		Country     string `json:"country"`
		Subdivision string `json:"subdivision"`
	} `json:"location"`
	ObsidianEventName struct {
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"obsidianEventName"`
	Service struct {
		ServiceId string `json:"serviceId"`
		Name      string `json:"name"`
	} `json:"service"`
	SessionId    string `json:"sessionId"`
	TenantId     string `json:"tenantId"`
	TenantObject struct {
		Name     string `json:"name"`
		TenantId string `json:"tenantId"`
	} `json:"tenantObject"`
	SubService string `json:"subService"`
	Status     string `json:"status"`
	UserAgent  struct {
		Raw string `json:"raw"`
	} `json:"userAgent"`
}

type GetAllEventsResponse struct {
	Data struct {
		GetEvents struct {
			HasMoreResults bool    `json:"hasMoreResults"`
			Cursor         string  `json:"cursor"`
			Results        []Event `json:"results"`
		} `json:"getEvents"`
	} `json:"data"`
}
