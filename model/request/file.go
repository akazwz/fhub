package request

type File struct {
	Name     string `json:"name" binding:"required"`
	ParentID string `json:"parent_id" binding:"required"`
	Size     uint64 `json:"size" binding:"required"`
	SHA256   string `json:"sha256" binding:"required"`
	Provider string `json:"provider" binding:"required"`
}

type PreFile struct {
	Name     string `json:"name" binding:"required"`
	ParentID string `json:"parent_id" binding:"required"`
	Size     uint64 `json:"size" binding:"required"`
	SHA256   string `json:"sha256" binding:"required"`
}

type Folder struct {
	Name     string `json:"name" binding:"required"`
	ParentID string `json:"parent_id" binding:"required"`
}

type ListFile struct {
	ParentID string `json:"parent_id" binding:"required"`
}
