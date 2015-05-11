package gads

import "testing"

func testNewManagedCustomerService(t *testing.T) (service *LocationCriterionService) {
	return &NewManagedCustomerService{Auth: testAuthSetup(t)}
}

func TestNewManagedCustomer(t *testing.T) {
	//	service := testNewManagedCustomerService(t)
}
