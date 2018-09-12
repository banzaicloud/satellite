package providers

import "github.com/sirupsen/logrus"

type IdentifySlow struct {
	log logrus.FieldLogger
}

func (s *IdentifySlow) Identify() (string, error) {
	return "", nil
}
