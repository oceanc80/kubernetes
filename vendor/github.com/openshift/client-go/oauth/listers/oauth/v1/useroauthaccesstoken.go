// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	oauthv1 "github.com/openshift/api/oauth/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// UserOAuthAccessTokenLister helps list UserOAuthAccessTokens.
// All objects returned here must be treated as read-only.
type UserOAuthAccessTokenLister interface {
	// List lists all UserOAuthAccessTokens in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*oauthv1.UserOAuthAccessToken, err error)
	// Get retrieves the UserOAuthAccessToken from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*oauthv1.UserOAuthAccessToken, error)
	UserOAuthAccessTokenListerExpansion
}

// userOAuthAccessTokenLister implements the UserOAuthAccessTokenLister interface.
type userOAuthAccessTokenLister struct {
	listers.ResourceIndexer[*oauthv1.UserOAuthAccessToken]
}

// NewUserOAuthAccessTokenLister returns a new UserOAuthAccessTokenLister.
func NewUserOAuthAccessTokenLister(indexer cache.Indexer) UserOAuthAccessTokenLister {
	return &userOAuthAccessTokenLister{listers.New[*oauthv1.UserOAuthAccessToken](indexer, oauthv1.Resource("useroauthaccesstoken"))}
}
