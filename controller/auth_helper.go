package controller

import(
    "dublog/router"
)

type AuthHelper struct{}

var(
    _AuthHelper = &AuthHelper{}
)

func (c *AuthHelper) AuthToken(ctx *router.Context) error {
	return nil
}
