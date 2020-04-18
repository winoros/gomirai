package message

// Message 消息
type Message struct {
	Type string `json:"type"`

	ID       int64     `json:"id,omitempty" message:"Source|Quote"`
	Text     string    `json:"text,omitempty" message:"Plain"`
	Time     int64     `json:"time,omitempty" message:"Source"`
	GroupID  int64     `json:"groupId,omitempty" message:"Quote"`
	SenderID int64     `json:"senderId,omitempty" message:"Quote"`
	Origin   []Message `json:"origin,omitempty" message:"Quote"`
	Target   int64     `json:"target,omitempty" message:"At"`
	Display  string    `json:"display,omitempty" message:"At"`
	FaceID   int64     `json:"faceId,omitempty" message:"Face"`
	Name     string    `json:"name,omitempty" message:"Face"`
	ImageID  string    `json:"imageId,omitempty" message:"Image"`
	URL      string    `json:"url,omitempty" message:"Image"`
	Path     string    `json:"path,omitempty" message:"Image"`
	XML      string    `json:"xml,omitempty" message:"Xml"`
	JSON     string    `json:"json,omitempty" message:"Json"`
	Content  string    `json:"content,omitempty" message:"App"`
}

func (m *Message) Detail() interface{} {
	switch m.Type {
	case "Plain":
		return Plain{Type: "Plain", Text: m.Text}
	case "At":
		return At{Type: "At", Target: m.Target, Display: m.Display}
	case "AtAll":
		return AtAll{Type: "AtAll"}
	case "Xml":
		return Xml{
			Type: "Xml",
			XML:  m.XML,
		}
	case "Json":
		return Json{
			Type: "Json",
			Json: m.JSON,
		}
	case "App":
		return App{
			Type:    "App",
			Content: m.Content,
		}
	case "Poke":
		return Poke{
			Type: "Poke",
			Name: m.Name,
		}
	case "Face":
		return Face{
			Type:   "Face",
			FaceID: m.FaceID,
			Name:   m.Name,
		}
	case "Image":
		return Image{
			Type:    "Image",
			ImageID: m.ImageID,
			URL:     m.URL,
			Path:    m.Path,
		}
	case "FlashImage":
		return FlashImage{
			Type:    "FlashImage",
			ImageID: m.ImageID,
			URL:     m.URL,
			Path:    m.Path,
		}
	case "Quote":
		return Quote{
			Type:     "Quote",
			Id:       m.ID,
			GroupID:  m.GroupID,
			SenderID: m.SenderID,
			TargetID: m.Target,
			Origin:   m.Origin,
		}
	case "Source":
		return Source{
			Type: "Source",
			Id:   m.ID,
			Time: m.Time,
		}
	default:
		return nil
	}
}
