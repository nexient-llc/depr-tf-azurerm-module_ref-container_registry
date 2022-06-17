package test

// Basic imports
import (
	"os"
	"path"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TerraTestSuite struct {
	suite.Suite
	TerraformOptions *terraform.Options
}

// setup to do before any test runs
func (suite *TerraTestSuite) SetupSuite() {
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(suite.T(), "../..", ".")
	_ = files.CopyFile(path.Join("..", "..", ".tool-versions"), path.Join(tempTestFolder, ".tool-versions"))
	pwd, _ := os.Getwd()
	suite.TerraformOptions = terraform.WithDefaultRetryableErrors(suite.T(), &terraform.Options{
		TerraformDir: tempTestFolder,
		VarFiles:     [](string){path.Join(pwd, "..", "test.tfvars")},
	})
	terraform.InitAndApplyAndIdempotent(suite.T(), suite.TerraformOptions)
}

// TearDownAllSuite has a TearDownSuite method, which will run after all the tests in the suite have been run.
func (suite *TerraTestSuite) TearDownSuite() {
	terraform.Destroy(suite.T(), suite.TerraformOptions)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TerraTestSuite))
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *TerraTestSuite) TestAcr() {
	expectedRegistryName := "democreusdev000acr000"
	expectedRgName := "democr-eus-dev-000-rg-000"
	// NOTE: "subscriptionID" is overridden by the environment variable "ARM_SUBSCRIPTION_ID". <>
	subscriptionID := ""

	azure.ContainerRegistryExists(suite.T(), expectedRegistryName, expectedRgName, subscriptionID)
	azure.ResourceGroupExists(suite.T(), expectedRgName, subscriptionID)
}
