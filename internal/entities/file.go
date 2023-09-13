package entities

type File struct {
	PostId  string `json:"postId"`
	FileUrl string `json:"fileUrl"`
	Kind    string `json:"kind"`
}

func NewFile(postId string, fileUrl string, kind string) *File {
	file := &File{
		PostId:  postId,
		FileUrl: fileUrl,
		Kind:    kind,
	}

	return file
}
