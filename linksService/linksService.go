package linksService

import (
	"context"
	"errors"
	"log"
	"main/crypter"
	"main/links"
	"main/repository"
	"net/url"
)

// Implementation of Service
type LinksService struct {
	repository repository.IRepository
	// TODO: other fields can be read from file
	crypter crypter.ICrypter
	// Note: the service is gonna try to generate url with this length, but its not guarenteed
	maximumUrlLength int
	// Count of attempts that the service is gonna do before it increments a length
	maximumAttempts int
}

// Creates new service by repository and crypter
func NewLinksService(rep repository.IRepository, crypt crypter.ICrypter) LinksService {
	return LinksService{
		repository:       rep,
		crypter:          crypt,
		maximumUrlLength: 7,
		maximumAttempts:  10, // Todo: think about a better value
	}
}

func (linkServ LinksService) generateShortLink(originalLink string) string {
	strToHash := originalLink
	attempts := 0
	shortUrlLength := linkServ.maximumUrlLength
	shortLink := ""
	for {
		// Todo: think about a case when maximumUrlLength could be greater than possible hash length
		shortLink = string(linkServ.crypter.GetHashSumFromString(strToHash))[:shortUrlLength]
		// Note: its to expensive to check database everytime
		if value, err := linkServ.repository.GetByShortLink(string(shortLink)); err == nil {
			if value != strToHash {
				attempts++
				if attempts > linkServ.maximumAttempts {
					shortUrlLength++
					strToHash = originalLink
					attempts = 0
					continue
				}
				strToHash += strToHash
				continue
			}
			log.Println("Get existed from rep")
			return shortLink
		}
		break
	}

	linkServ.repository.PushOriginalAndShort(originalLink, shortLink)
	return shortLink
}

// Creates short link and pushes it to a repository
// Returns a struct with short link
func (linkServ LinksService) Create(ctx context.Context, req *links.CreateShortLinkRequest) (*links.CreateShortLinkResponse, error) {
	originalLink := req.OriginalLink.Url
	if err := isValidUrl(originalLink); err != nil {
		return nil, err
	}
	// Note: its possible to cut few symbols, but then we need to resolve collisions
	// I've tried to do that in generateShortLink method
	// shortLink := string(linkServ.crypter.GetHashSumFromString(originalLink)[:])
	// if _, err := linkServ.repository.GetByShortLink(string(shortLink)); err != nil {
	// 	linkServ.repository.PushOriginalAndShort(originalLink, shortLink)
	// }

	shortLink := linkServ.generateShortLink(originalLink)

	return &links.CreateShortLinkResponse{
		ShortLink: &links.Link{
			Url: shortLink,
		},
	}, nil
}

// Takes original link from a rep and returns it in a struct
func (linkServ LinksService) Retrive(ctx context.Context, req *links.RetriveOriginalLinkRequest) (*links.RetriveOriginalLinkResponse, error) {
	shortLink := req.ShortLink.Url
	if shortLink == "" {
		return nil, errors.New("[ERROR]: invalid request (empty short string)")
	}
	originalLink, err := linkServ.repository.GetByShortLink(string(shortLink))
	return &links.RetriveOriginalLinkResponse{
		OriginalLink: &links.Link{
			Url: originalLink,
		},
	}, err
}

func isValidUrl(urlToValidate string) error {
	// Todo: move validation to the service
	parsedUrl, err := url.Parse(urlToValidate)
	if err != nil {
		return errors.New("[ERROR]:" + err.Error())
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return errors.New("[ERROR]: invalid url scheme:" + parsedUrl.Scheme)
	}
	if parsedUrl.Host == "" {
		return errors.New("[ERROR]: invalid url host")
	}
	return nil
}
