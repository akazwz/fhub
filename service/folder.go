package service

import (
	"fmt"
	"github.com/akazwz/fhub/utils"
	"os"
	"sort"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
)

type FolderService struct{}

func (s *FolderService) CreateFolder(folder model.Folder) error {
	err := folder.Create(global.GDB)
	return err
}

func (s *FolderService) DeleteFolder(folder model.Folder) error {
	err := folder.Delete(global.GDB)
	return err
}

func (s *FolderService) RenameFolder(folder model.Folder) error {
	err := folder.Rename(global.GDB)
	return err
}

func (s *FolderService) FindFilesByParentID(uid, parentID string) ([]model.File, error) {
	files, err := model.FindFilesByParentID(global.GDB, uid, parentID)
	filesWithURL := make([]model.File, 0)
	for _, file := range files {
		provider := model.Provider{ContentHash: file.ContentHash}
		err := provider.FindProviderByContentHash(global.GDB)
		if err != nil {
			return nil, err
		}

		client := global.WasabiClient
		bucket := os.Getenv("WASABI_BUCKET_NAME")

		contentDisposition := fmt.Sprintf("attachment; filename=%s", file.Name)
		object, err := utils.S3Storage.GetPresignGetObjectURL(client, bucket, provider.Key, contentDisposition)
		file.URL = object.URL
		filesWithURL = append(filesWithURL, file)
	}

	if err != nil {
		return nil, err
	}
	return filesWithURL, err
}

func (s *FolderService) FindFoldersByParentID(uid, parentID string) ([]model.Folder, error) {
	folder := model.Folder{
		Model: model.Model{ID: parentID},
		UID:   uid,
	}
	folders, err := folder.GetChildFolders(global.GDB)
	if err != nil {
		return nil, err
	}
	return folders, err
}

func (s *FolderService) FindPath(uid, parentID string) ([]model.FolderPath, error) {
	folder := &model.Folder{
		ParentID: parentID,
		UID:      uid,
	}

	err := folder.GetPath(global.GDB)
	if err != nil {
		return nil, err
	}
	// 反转数组
	sort.SliceStable(folder.Path, func(i, j int) bool {
		return true
	})
	return folder.Path, err
}
