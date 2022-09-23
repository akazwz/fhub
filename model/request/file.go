package request

type File struct {
	Name        string `json:"name" binding:"required"`
	ParentID    string `json:"parent_id" binding:"required"`
	Size        uint64 `json:"size" binding:"required"`
	ContentHash string `json:"content_hash" binding:"required"`
	Provider    string `json:"provider" binding:"required"`
	URI         string `json:"uri" binding:"required"`
}

type CreateFile struct {
	Name         string     `json:"name"`
	Size         uint64     `json:"size"`
	ContentHash  string     `json:"content_hash"`
	ParentID     string     `json:"parent_id"`
	PartInfoList []PartInfo `json:"part_info_list"`
}

type CompleteMultipartUpload struct {
	Name        string `json:"name"`
	Size        uint64 `json:"size"`
	Key         string `json:"key"`
	UploadId    string `json:"upload_id"`
	ContentHash string `json:"content_hash"`
	ParentId    string `json:"parent_id"`
	PartCount   int32  `json:"part_count"`
}

type PartInfo struct {
	PartNumber int32 `json:"part_number" form:"part_number"`
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
