package comment

type CommentRepository interface {
	FindByUUID(string) (*Comment, error)
	Store(*Comment) (string, error)
	Update(*Comment) error
	DeleteByUUID(string) error
}

type commentRepositoryMongo struct {
}

func (rr *commentRepositoryMongo) FindByUUID(uuid string) (*Comment, error) {
	return nil, nil
}

func (rr *commentRepositoryMongo) Store(r *Comment) (string, error) {
	return "", nil
}

func (rr *commentRepositoryMongo) Update(r *Comment) error {
	return nil
}

func (rr *commentRepositoryMongo) DeleteByUUID(uuid string) error {
	return nil
}

type CommentRepositoryMongoConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (c *CommentRepositoryMongoConfig) Connect() (CommentRepository, func() error, error) {
	return nil, nil, nil
}
