package schema

type Schema int

const (
	Folder Schema = iota + 1
	File
	Tag
)

func (s Schema) String() string {
	switch s {
	case Folder:
		return "folder"
	case File:
		return "file"
	case Tag:
		return "tag"
	default:
		return "Unknown"
	}
}
