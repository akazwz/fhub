package request

// UploadFile  上传文件
type UploadFile struct {
	File      bool   `json:"file" form:"file" binding:"required"`
	Filename  string `json:"filename" form:"filename" binding:"required"`
	PrefixDir string `json:"prefix_dir" form:"prefix_dir"`
	Size      int64  `json:"size" form:"size" binding:"required"`
	Sha256    string `json:"sha256" form:"sha256" binding:"required"`
	QKey      string `json:"qkey" form:"qkey"`
	CID       string `json:"cid" form:"cid"`
}
