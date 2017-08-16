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
	ConversationID                string   `json:"ConversationId"`
	Created                       DateTime `json:"CreatedDateTime"`
	From                          Recipient
	HasAttachments                bool
	ID                            string `json:"Id"`
	Importance                    Importance
	InferenceClassification       InferenceClassificationType
	IsDeliveryReceiptRequested    bool
	IsDraft                       bool
	IsRead                        bool
	IsReadReceiptRequested        bool
	LastModified                  DateTime `json:"LastModifiedDateTime"`
	MultiValueExtendedProperties  []MultiValueLegacyExtendedProperty
	ParentFolderID                string   `json:"ParentFolderId"`
	Received                      DateTime `json:"ReceivedDateTime"`
	ReplyTo                       []Recipient
	Sender                        Recipient
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty
	Sent                          DateTime `json:"SentDateTime"`
	Subject                       string
	ToRecipients                  []Recipient
	UniqueBody                    ItemBody
	WebLink                       string
}
