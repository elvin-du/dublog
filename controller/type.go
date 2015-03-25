package controller

import(
    "dublog/router"
)

type (
    Action func(*router.Context)
    PreHandler func(*router.Context) error
)
