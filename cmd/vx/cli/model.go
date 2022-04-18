package cli

type FileStatus string

var (
	StatusCreated FileStatus = "Created"
	StatusUpdated FileStatus = "Updated"
)

type fileMetadata struct {
	Path             string
	Status           FileStatus
	ModificationTime string
	GoToStaging      bool
}

type fileMetadataArr []fileMetadata

func (arr fileMetadataArr) Len() int {
	return len(arr)
}
func (arr fileMetadataArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr fileMetadataArr) Less(i, j int) bool {
	return arr[i].Path > arr[j].Path
}

type commitHistory struct {
	Version string // computed
	Date    string
	Message string
}
