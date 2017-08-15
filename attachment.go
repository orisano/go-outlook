package outlook

type ReferenceAttachmentProviders string

const (
	ReferenceAttachmentProvidersDropbox          ReferenceAttachmentProviders = "Dropbox"
	ReferenceAttachmentProvidersOneDriveBusiness ReferenceAttachmentProviders = "OneDriveBusiness"
	ReferenceAttachmentProvidersOneDriveConsumer ReferenceAttachmentProviders = "OneDriveConsumer"
	ReferenceAttachmentProvidersOther            ReferenceAttachmentProviders = "Other"
)

type ReferenceAttachmentPermissions string

const (
	ReferenceAttachmentPermissionOther            ReferenceAttachmentPermissions = "Other"
	ReferenceAttachmentPermissionView             ReferenceAttachmentPermissions = "View"
	ReferenceAttachmentPermissionEdit             ReferenceAttachmentPermissions = "Edit"
	ReferenceAttachmentPermissionAnonymousView    ReferenceAttachmentPermissions = "AnonymousView"
	ReferenceAttachmentPermissionAnonymousEdit    ReferenceAttachmentPermissions = "AnonymousEdit"
	ReferenceAttachmentPermissionOrganizationView ReferenceAttachmentPermissions = "OrganizationView"
	ReferenceAttachmentPermissionOrganizationEdit ReferenceAttachmentPermissions = "OrganizationEdit"
)

type Attachment struct {
	ContentType  string
	IsInline     bool
	LastModified DateTimeOffset `json:"LastModifiedDateTime"`
	Name         string
	Size         int
	ID           string `json:"Id"`

	// for ItemAttachment
	Item Item

	// for FileAttachment
	ContentBytes    []byte
	ContentID       string `json:"ContentId"`
	ContentLocation string
	IsContactPhoto  bool

	// for ReferenceAttachment
	IsFolder     bool
	Permission   ReferenceAttachmentPermissions
	PreviewURL   string `json:"PreviewUrl"`
	ProviderType ReferenceAttachmentProviders
	SourceURL    string `json:"SourceUrl"`
	ThumbnailURL string `json:"ThumbnailUrl"`
}
