package model

import (
	"fyne.io/fyne/v2/data/binding"
)

// TODO: wrap datalist
type Folders struct {
	binding.DataList
}

func NewFoldersFromSlice(folders []Folder) *Folders {
	fd := &Folders{
		DataList: binding.NewStringList(),
	}

	for _, v := range folders {
		fd.DataList.(binding.StringList).Append(v.Name)
	}

	return fd
}

func NewFolders() *Folders {
	return NewFoldersFromSlice([]Folder{})
}

func (fd *Folders) StringSlice() []string {
	count := fd.DataList.Length()
	r := make([]string, count)
	for i := 0; i < count; i++ {
		item, _ := fd.DataList.GetItem(i)
		value, _ := item.(binding.String).Get()
		r[i] = value
	}
	return r
}
