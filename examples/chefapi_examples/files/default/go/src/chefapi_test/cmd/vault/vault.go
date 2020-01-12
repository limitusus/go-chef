//
// Test the go-chef/chef chef vault support against a live server
//
package main

import (
	"fmt"
        "chefapi_test/testapi"
        "github.com/go-chef/chef"
	"os"
)


// main Exercise the chef server vault support
func main() {
	client := testapi.Client()

	// Add users to the test organizations
        invite := chef.AddNow { User: "usrv", }
        invite2 := chef.AddNow { User: "usrv2", }
	// TODO: Add a node

	// List vaults before an item is created
	vaultList, err := client.Vaults.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue listing vaults %+v\n",err)
	}
	fmt.Printf("List vaults before creation %+v\n", vaultList)

	// Create a vault item
	item, err := client.Vaults.Create('testv', 'secrets')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue creating testv secrets vault item %+v\n", err)
	}
	fmt.Printf("Created testv secrets vault item %+v\n", item)

	// Add content to the vault item
	item, err = client.Vaults.Update(item, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue updating testv secrets vault item %+v\n", err)
	}
	fmt.Printf("Updated testv secrets vault item %+v\n", item)

	// List vaults after an item is created
	vaultList, err = client.Vaults.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue listing vaults %+v\n",err)
	}
	fmt.Printf("List vaults after creation %+v\n", vaultList)

	// TODO:  List items in a vault

	// Get vault item
	vaultItem, err = client.Vaults.Get("testv", "secrets")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue getting vault item testv secrets %+v\n",err)
	}
	fmt.Printf("Delete testv secrets vault item%+v\n", vaultList)

	// Delete vault contents
	err = client.Vaults.Delete("testv", "secrets")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue deleting vault item testv secrets %+v\n",err)
	}
	fmt.Println("Delete testv secrets vault item")

	// List vaults after all items are deleted
	vaultList, err = client.Vaults.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue listing vaults %+v\n",err)
	}
	fmt.Printf("List vaults after creation %+v\n", vaultList)

}

// add user and node to the admin and client list
Vaults.GetItem(vaultName, itemName)  (*VaultItem, error)
// Add user2
// Change a value
// Do things using usrv id - admin
// Do things using usrv2 id
Vaults.DeleteItem(vaultName, vaultItem)  (error)
Vaults.List() *VaultListResult, err error
