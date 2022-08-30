package request

type File struct {
	Name      string `json:"name" binding:"required"`
	PrefixDir string `json:"prefix_dir" binding:"required"`
	Size      int    `json:"size" binding:"required"`
	SHA256    string `json:"sha256" binding:"required"`
	Provider  string `json:"provider" binding:"required"`
}

type PreFile struct {
	Name      string `json:"name" binding:"required"`
	PrefixDir string `json:"prefix_dir" binding:"required"`
	Size      int    `json:"size" binding:"required"`
	SHA256    string `json:"sha256" binding:"required"`
}

type Folder struct {
	Name      string `json:"name" binding:"required"`
	PrefixDir string `json:"prefix_dir" binding:"required"`
}

type ListFile struct {
	PrefixDir string `json:"prefix_dir" binding:"required"`
}
