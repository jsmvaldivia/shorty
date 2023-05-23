package app

type ShortenerService struct {
	Shortener  ShortenerStrategy
	Repository Repository
}

type ShortenerStrategy interface {
	shorten(url string) (string, error)
}

type Repository interface {
	save(bigUrl string, shortUrl string) error
}

func (s *ShortenerService) CreateShortUrl(url string) (string, error) {
	shortUrl, err := s.Shortener.shorten(url)
	if err != nil {
		return "", err
	}

	err = s.Repository.save(url, shortUrl)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func newShortenerService() {
	return
}

type InMemoryRepository struct {
	urlMap map[string]string
}

func (i *InMemoryRepository) save(bigUrl string, shortUrl string) error {
	i.urlMap[shortUrl] = bigUrl
	return nil
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{urlMap: make(map[string]string)}
}
