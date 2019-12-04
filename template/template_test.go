package dayTEMPLATE

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type dayTEMPLATESuite struct {
	suite.Suite
}

func (suite *dayTEMPLATESuite) SetupTest() {
}

func (suite *dayTEMPLATESuite) TestStuff() {

}

func TestDayTEMPLATESuite(t *testing.T) {
	suite.Run(t, new(dayTEMPLATESuite))
}
