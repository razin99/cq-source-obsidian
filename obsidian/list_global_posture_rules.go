package obsidian

import (
	"context"
	_ "embed"
	"encoding/json"
	"time"
)

func (o ObsidianClient) ListGlobalPostureRules(ctx context.Context, cursor string) (*ListGlobalPostureRulesResponse, error) {
	response, err := o.NewRequest(ctx, GqlOperation{
		OperationName: "listGlobalPostureRules",
		Query:         listGlobalPostureRulesQuery,
		Variables: struct {
			Cursor string `json:"cursor"`
		}{Cursor: cursor},
	})
	if err != nil {
		return nil, err
	}

	data := &ListGlobalPostureRulesResponse{}
	if err := json.NewDecoder(response.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

//go:embed gql/list_global_posture_rules.gql
var listGlobalPostureRulesQuery string

type PostureRule struct {
	RuleId         string   `json:"rule_id"`
	PlatformId     string   `json:"platform_id"`
	ProductIds     []string `json:"product_ids"`
	SecurityDomain struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"security_domain"`
	Name                    string   `json:"name"`
	RiskLevel               string   `json:"risk_level"`
	StandardIds             []string `json:"standard_ids"`
	ControlIds              []string `json:"control_ids"`
	ObsidianRule            bool     `json:"obsidian_rule"`
	ReleaseLabel            string   `json:"release_label"`
	DefaultRiskLevel        string   `json:"default_risk_level"`
	BenchmarkValueType      string   `json:"benchmark_value_type"`
	RuleValueType           string   `json:"rule_value_type"`
	Description             string   `json:"description"`
	RemediationInstructions string   `json:"remediation_instructions"`
	ExceptionsCount         struct {
		Active   uint64 `json:"active"`
		Inactive uint64 `json:"inactive"`
	} `json:"exceptions_count"`
	TenantStates []struct {
		TenantId   string `json:"tenant_id"`
		IsPassing  bool   `json:"is_passing"`
		Violations uint64 `json:"violations"`
		Tenant     struct {
			TenantId     string `json:"tenantId"`
			Name         string `json:"name"`
			ServiceId    string `json:"serviceId"`
			IsProduction bool   `json:"isProduction"`
			Sensitivity  struct {
				DisplayName string `json:"displayName"`
				Enum        string `json:"enum"`
			} `json:"sensitivity"`
			Platform string `json:"platform"`
		} `json:"tenant"`
		LastScanned     time.Time `json:"last_scanned"`
		RiskAccepted    bool      `json:"risk_accepted"`
		ExceptionsCount struct {
			Active   uint64 `json:"active"`
			Inactive uint64 `json:"inactive"`
		} `json:"exceptions_count"`
		CorrectionScoreChange float64 `json:"correction_score_change"`
	} `json:"tenant_states"`
	TotalViolations       uint64  `json:"total_violations"`
	CorrectionScoreChange float64 `json:"correction_score_change"`
}

type ListGlobalPostureRulesResponse struct {
	Data struct {
		ListGlobalPostureRules struct {
			HasMoreResults bool          `json:"has_more_results"`
			Cursor         string        `json:"cursor"`
			Rules          []PostureRule `json:"rules"`
			Total          uint64        `json:"total"`
		} `json:"listGlobalPostureRules"`
	} `json:"data"`
}
