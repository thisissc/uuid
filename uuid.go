package uuid

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

func UUID4() string {
	out, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	return strings.ToLower(strings.Join(strings.Split(out.String(), "-"), ""))
}
