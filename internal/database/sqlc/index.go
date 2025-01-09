package sqlc

var Q *Queries

func GetDBTX() DBTX {
	return Q.db
}
