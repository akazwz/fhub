package utils

func GetCategoryByName(name string) string {
	image := []string{"png", "jpg", "jpeg", "bmp", "gif", "webp", "psd", "svg", "tiff"}
	video := []string{"mp4", "mov", "avi", "flv", "wmv", "mpg", "mkv", "f4v", "rmvb", "rm", "3gb"}
	audio := []string{"mp3", "aac", "wav", "cda", "wma", "mid", "aif", "aiff", "mid", "ra", "vqf", "ape"}
	ext := GetExtByName(name)
	switch true {
	case ISContainString(image, ext):
		return "image"
	case ISContainString(video, ext):
		return "video"
	case ISContainString(audio, ext):
		return "audio"
	default:
		return ""
	}
}

func GetExtByName(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return name[i+1:]
		}
	}
	return ""
}

func GetPureNameByName(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return name[:i]
		}
	}
	return ""
}

func ISContainString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
