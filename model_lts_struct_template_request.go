package entity

type StructTemplateRequest struct {
	Content      string           `json:"content,omitempty"`
	LogGroupId   string           `json:"log_group_id"`
	ParseType    string           `json:"parse_type,omitempty"`
	TemplateId   string           `json:"template_id"`
	TemplateType string           `json:"template_type"`
	TemplateName string           `json:"template_name"`
	LogStreamId  string           `json:"log_stream_id"`
	ProjectId    string           `json:"project_id"`
	RegexRules   *string          `json:"regex_rules,omitempty"`
	Layers       *int             `json:"layers,omitempty"`
	Tokenizer    string           `json:"tokenizer,omitempty"`
	LogFormat    *string          `json:"log_format,omitempty"`
	DemoFields   []DemoFieldsInfo `json:"demo_fields"`
	TagFields    []TagFieldsInfo  `json:"tag_fields"`
}

type DemoFieldsInfo struct {
	IsAnalysis      bool   `json:"is_analysis"`
	Content         string `json:"content,omitempty"`
	FieldName       string `json:"field_name,omitempty"`
	Type            string `json:"type"`
	UserDefinedName string `json:"userDefinedName,omitempty"`
	Index           int    `json:"index,omitempty"`
}

type TagFieldsInfo struct {
	FieldName  string  `json:"fieldName"`
	Type       string  `json:"type"`
	Content    *string `json:"content,omitempty"`
	IsAnalysis *bool   `json:"isAnalysis,omitempty"`
}

fun (s *StructTemplateRequest) ToDemoFieldsInfo() {
	s.DemoFields = []DemoFieldsInfo{
		{
			Type:      "string",
			FieldName: "remote_ip",
			Index:     0,
		},
		{
			Type:      "string",
			FieldName: "local_ip",
			Index:     1,
		},
		{
			Type:      "string",
			FieldName: "local_part",
			Index:     2,
		},
		{
			Type:      "string",
			FieldName: "t",
			Index:     3,
		},
		{
			Type:      "string",
			FieldName: "tt",
			Index:     4,
		},
		{
			Type:      "string",
			FieldName: "method",
			Index:     5,
		},
		{
			Type:      "string",
			FieldName: "uri",
			Index:     6,
		},
		{
			Type:      "string",
			FieldName: "protocal",
			Index:     7,
		},
		{
			Type:      "string",
			FieldName: "code",
			Index:     8,
		},
		{
			Type:      "string",
			FieldName: "send_bytes",
			Index:     9,
		},
		{
			Type:      "string",
			FieldName: "cost",
			Index:     10,
		},
	}
}