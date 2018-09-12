package providers

import "github.com/sirupsen/logrus"

type IdentifySlow struct {
	Log logrus.FieldLogger
}

func (s *IdentifySlow) Identify() (string, error) {

	return "", nil
}
