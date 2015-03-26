package db

import(
    "database/sql"
    _"github.com/go-sql-driver/mysql"
)

type Context struct{
    *sql.DB
}
