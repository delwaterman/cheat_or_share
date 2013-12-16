package sims

import (
    "testing"
    "github.com/jacobsa/oglematchers"
    "github.com/jacobsa/ogletest"
    "github.com/jacobsa/oglemock"
    mock_strategies "github.com/delwaterman/cheat_or_share/mock_strategies"
)

// -------------------------------------------

type NewIndividualTest struct{}

func (t *NewIndividualTest) SetsMinAndMaxEnergy() {
    newIndividual := NewIndividual(1, 2)
    ogletest.ExpectThat((*newIndividual),
        oglematchers.DeepEquals(Individual{1, 2}))
}

func (t *NewIndividualTest) PanicsWhenMaxGreaterThanMin() {
    badIndividualFunc := func() {
        NewIndividual(2, 1)
    }
    expectedErrorMessage := "minEnergy(2) cannot be greater than maxEnergy(1)"
    ogletest.ExpectThat(badIndividualFunc,
                        oglematchers.Panics(
                            oglematchers.Equals(expectedErrorMessage)))
}

// -------------------------------------------

type SimpleIndividualBuilderTest struct{}

func (t *SimpleIndividualBuilderTest) ReturnsSimpleIndiviualFunction() {
    var newIndividualFunction IndivdualBuilder
    newIndividualFunction = SimpleIndividualBuilder(1, 2)
    ogletest.ExpectThat((*newIndividualFunction()),
        oglematchers.DeepEquals(Individual{1, 2}))
}

// -------------------------------------------

type NewRandIndividualTest struct {
    RandStrategyMock mock_strategies.MockRandStrategy
}

func (t *NewRandIndividualTest) SetUp(testInfo *ogletest.TestInfo) {
    controller := testInfo.MockController
    t.RandStrategyMock = mock_strategies.NewMockRandStrategy(controller, "MockRandomStrategy")
}

func (t *NewRandIndividualTest) ReturnsNewIndividual() {
    ogletest.ExpectCall(t.RandStrategyMock, "RandomInteger")(1, 10).
        WillOnce(oglemock.Return(1))
    ogletest.ExpectCall(t.RandStrategyMock, "RandomInteger")(2, 10).
        WillOnce(oglemock.Return(2))

    newIndividual := NewRandIndividual(t.RandStrategyMock)
    ogletest.ExpectThat((*newIndividual),
        oglematchers.DeepEquals(Individual{1, 2}))
}

func (t *NewRandIndividualTest) KeepsUpperBoundAtTen() {
    ogletest.ExpectCall(t.RandStrategyMock, "RandomInteger")(1, 10).
        WillOnce(oglemock.Return(10))
    ogletest.ExpectCall(t.RandStrategyMock, "RandomInteger")(10, 10).
        WillOnce(oglemock.Return(10))

    newIndividual := NewRandIndividual(t.RandStrategyMock)
    ogletest.ExpectThat((*newIndividual),
        oglematchers.DeepEquals(Individual{10, 10}))
}

// -------------------------------------------

type RandIndividualBuilderTest struct {
    RandStrategyMock mock_strategies.MockRandStrategy
}

func (t *RandIndividualBuilderTest) SetUp(testInfo *ogletest.TestInfo) {
    controller := testInfo.MockController
    t.RandStrategyMock = mock_strategies.NewMockRandStrategy(controller, "MockRandomStrategy")
}

func (t *RandIndividualBuilderTest) ReturnsRandBuilderFunction() {
    ogletest.ExpectCall(t.RandStrategyMock, "RandomInteger")(1, 10).
        WillOnce(oglemock.Return(1))
    ogletest.ExpectCall(t.RandStrategyMock, "RandomInteger")(2, 10).
        WillOnce(oglemock.Return(2))

    individualBuilderFunction := RandIndividualBuilder(t.RandStrategyMock)
    newIndividual := individualBuilderFunction()
    ogletest.ExpectThat((*newIndividual),
        oglematchers.DeepEquals(Individual{1, 2}))
}

// -----------------------------------------



func init() {
    ogletest.RegisterTestSuite(&NewIndividualTest{})
    ogletest.RegisterTestSuite(&NewRandIndividualTest{})
    ogletest.RegisterTestSuite(&SimpleIndividualBuilderTest{})
    ogletest.RegisterTestSuite(&RandIndividualBuilderTest{})
}

func TestIndividual(t *testing.T) {
    ogletest.RunTests(t)
}

// func TestIndividualExertEnergy(t *testing.T) {
//   individual := NewIndividual(1, 10)

//   // Test lower bound

//   if output := (&individual).ExertEnergy(float32(0.0)); output != 1 {
//     t.Errorf("ExertEnergy(individual) should return lower bound of %v; got %v.", 1, output)
//   }


//   if output := (&individual).ExertEnergy(float32(0.999999)); output != 10 {
//     t.Errorf("ExertEnergy(individual) should return upper bound of %v; got %v.", 10, output)
//   }
// }

