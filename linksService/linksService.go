package linksService

import (
	"context"
	"main/crypter"
	"main/links"
	"main/repository"
)

// Implementation of Service
type LinksService struct {
	repository repository.IRepository
	crypter    crypter.ICrypter
}

// Creates new service by repository and crypter
func NewLinksService(rep repository.IRepository, crypt crypter.ICrypter) LinksService {
	return LinksService{repository: rep, crypter: crypt}
}

// Creates short link and pushes it to a repository
// Returns a struct with short link
func (linkServ LinksService) Create(ctx context.Context, req *links.CreateShortLinkRequest) (*links.CreateShortLinkResponse, error) {
	originalLink := req.OriginalLink.Url

	// Note: its possible to cut few symbols, but then we need to resolve collisions
	shortLink := string(linkServ.crypter.GetHashSumFromString(originalLink)[:])
	if _, err := linkServ.repository.GetByShortLink(string(shortLink)); err != nil {
		linkServ.repository.PushOriginalAndShort(originalLink, shortLink)
	}

	return &links.CreateShortLinkResponse{
		ShortLink: &links.Link{
			Url: shortLink,
		},
	}, nil
}

// Takes original link from a rep and returns it in a struct
func (linkServ LinksService) Retrive(ctx context.Context, req *links.RetriveOriginalLinkRequest) (*links.RetriveOriginalLinkResponse, error) {
	shortLink := req.ShortLink.Url
	originalLink, err := linkServ.repository.GetByShortLink(string(shortLink))
	return &links.RetriveOriginalLinkResponse{
		OriginalLink: &links.Link{
			Url: originalLink,
		},
	}, err
}
