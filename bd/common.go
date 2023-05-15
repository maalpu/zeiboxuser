package bd

import (
	"os"

	"github.com/maalpu/zeiboxuser/models"
	"github.com/maalpu/zeiboxuser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	var err error
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
