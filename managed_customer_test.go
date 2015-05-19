package gads

import "testing"

func testNewManagedCustomerService(t *testing.T) (service *ManagedCustomerService) {
	auth := testAuthSetup(t)
	return NewManagedCustomerService(&auth)
}

func TestNewManagedCustomer(t *testing.T) {
	//	service := testNewManagedCustomerService(t)
}
