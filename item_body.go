package outlook

type BodyType int

const (
	BodyText BodyType = iota
	BodyHTML
)

type ItemBody struct {
	ContentType BodyType
	Content     string
}
