package model

type Folders []Folder

func (f Folders) StringSlice() []string {
	x := make([]string, len(f))
	for k, v := range f {
		x[k] = v.Name
	}
	return x
}
