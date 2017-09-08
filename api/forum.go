package api

type ForumController interface {
	Create(*ForumCreateRequest) (*ForumCreateResponse, error)
	Update(*ForumUpdateRequest) (*ForumUpdateResponse, error)
	List()
	Delete()
	Get()
}
