package service

import (
	"sort"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
)

type FolderService struct{}

func (s *FolderService) CreateFolder(folder model.Folder) error {
	err := folder.Create(global.GDB)
	return err
}

func (s *FolderService) FindFilesByParentID(uid, parentID string) ([]model.File, error) {
	files, err := model.FindFilesByParentID(global.GDB, uid, parentID)
	if err != nil {
		return nil, err
	}
	return files, err
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
