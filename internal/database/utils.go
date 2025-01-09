package database

import (
	"Adeeb_Go/internal/database/sqlc"
	"Adeeb_Go/utils"
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

func UUIDToString(uuid pgtype.UUID) string {
	uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:16])
	return uuidStr
}

func StringToUUID(str string) (uuid pgtype.UUID, err error) {

	switch len(str) {
	case 36:
		str = str[0:8] + str[9:13] + str[14:18] + str[19:23] + str[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return uuid, fmt.Errorf("cannot parse UUID %v", str)
	}

	buf, err := hex.DecodeString(str)
	if err != nil {
		return uuid, err
	}

	copy(uuid.Bytes[:], buf)

	uuid.Valid = true

	return uuid, err
}

// Used to get a formated SQL string to set a value
//
// Note: when using enums, decouple it's validation and pass the value in built-in types,
// So that it'll be formated in the right way.
func GetFormatedSetString(field string, value any) string {
	if utils.IsString(value) {
		return fmt.Sprintf("%v = '%v'", field, value)
	}
	return fmt.Sprintf("%v = %v", field, value)
}

type Field struct {
	Name  string
	Value any
}

func UpdateQuery(table string, id pgtype.UUID, fields []Field) error {
	var setStrings []string
	var fieldsValues []any
	for i := 0; i < len(fields); i++ {
		// if the field is nil or instantiated using the zero value for the type
		// we don't include it in the update query
		if utils.IsZeroValue(fields[i].Value) || fields[i].Value == nil {
			continue
		}
		setStrings = append(setStrings, GetFormatedSetString(fields[i].Name, fields[i].Value))
		fieldsValues = append(fieldsValues, fields[i].Value)
	}

	setStrings = append(setStrings, "updated_at = CURRENT_TIMESTAMP")

	query := fmt.Sprintf(
		`UPDATE %v SET %v WHERE id = $1`,
		table,
		strings.Join(setStrings, ","),
	)

	fmt.Println(query)

	_, err := sqlc.GetDBTX().Exec(context.TODO(), query, id)
	if err != nil {
		return fmt.Errorf("Upate Failed, error: %v", err)
	}
	return nil
}
