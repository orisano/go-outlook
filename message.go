package outlook

type InferenceClassificationType string

const (
	InferenceClassFocused InferenceClassificationType = "Focused"
	InferenceClassOther   InferenceClassificationType = "Other"
)

type MultiValueLegacyExtendedProperty struct {
	Value      []string
	PropertyID string `json:"PropertyId"`
}

type SingleValueLegacyExtendedProperty struct {
	Value      string
	PropertyID string `json:"PropertyId"`
}

type Message struct {
	Attachments                   []Attachment
	BccRecipients                 []Recipient
	Body                          ItemBody
	BodyPreview                   string
	Categories                    []string
	CcRecipients                  []Recipient
	ChangeKey                     string
	ConversationID                string         `json:"ConversationId"`
	Created                       DateTimeOffset `json:"CreatedDateTime"`
	Extensions                    []Extension
	From                          Recipient
	HasAttachments                bool
	ID                            string `json:"Id"`
	Importance                    Importance
	InferenceClassification       InferenceClassificationType
	IsDeliveryReceiptRequested    bool
	IsDraft                       bool
	IsRead                        bool
	IsReadReceiptRequested        bool
	LastModified                  DateTimeOffset `json:"LastModifiedDateTime"`
	MultiValueExtendedProperties  []MultiValueLegacyExtendedProperty
	ParentFolderID                string         `json:"ParentFolderId"`
	Received                      DateTimeOffset `json:"ReceivedDateTime"`
	ReplyTo                       []Recipient
	Sender                        Recipient
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty
	Sent                          DateTimeOffset `json:"SentDateTime"`
	Subject                       string
	ToRecipients                  []Recipient
	UniqueBody                    ItemBody
	WebLink                       string
}
