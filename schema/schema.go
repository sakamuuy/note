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

type Action int

const (
	Write Action = iota + 1
	Edit
	Delete
)

func (a Action) string() string {
	switch a {
	case Write:
		return "write"
	case Edit:
		return "edit"
	case Delete:
		return "delete"
	default:
		return "unknown"
	}
}