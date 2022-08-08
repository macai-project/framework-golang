package helpers

import "github.com/gofrs/uuid"

func MidecFromEAN(ean string) string {
	namespace, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	return uuid.NewV5(namespace, ean).String()
}
