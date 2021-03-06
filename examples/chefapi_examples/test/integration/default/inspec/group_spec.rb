# Inspec tests for the group chef api go module
#

describe command('/go/src/chefapi_test/bin/group') do
  its('stderr') { should match(%r{^Issue recreating group1. POST https://localhost/organizations/test/groups: 409}) }
  its('stderr') { should_not match(/error|no such file|cannot find|not used|undefined/) }
  its('stdout') { should match(%r{^List initial groups map\[(?=.*admins:https://localhost/organizations/test/groups/admins)(?=.*billing-admins:https://localhost/organizations/test/groups/billing-admins)(?=.*clients:https://localhost/organizations/test/groups/clients)(?=.*users:https://localhost/organizations/test/groups/users)(?=.*public_key_read_access:https://localhost/organizations/test/groups/public_key_read_access).*\]EndInitialList}) }
  its('stdout') { should match(%r{^Added group1 \&\{https://localhost/organizations/test/groups/group1\}}) }
  its('stdout') { should match(%r{^List groups after adding group1 map\[(?=.*group1:https://localhost/organizations/test/groups/group1)(?=.*admins:https://localhost/organizations/test/groups/admins)(?=.*billing-admins:https://localhost/organizations/test/groups/billing-admins)(?=.*clients:https://localhost/organizations/test/groups/clients)(?=.*users:https://localhost/organizations/test/groups/users)(?=.*public_key_read_access:https://localhost/organizations/test/groups/public_key_read_access).*\]EndAddList}) }
  its('stdout') { should match(/^Get group1 \{Name:group1 GroupName:group1 OrgName:test Actors:\[\] Clients:\[\] Groups:\[\] Users:\[\]\}/) }
  its('stdout') { should match(/^Update group1 \{Name:group1 GroupName:group1-new OrgName: Actors:\[\] Clients:\[\] Groups:\[\] Users:\[pivotal\]\}/) }
  its('stdout') { should match(/^Delete group1 <nil>/) }
  its('stdout') { should match(%r{^List groups after cleanup map\[(?=.*admins:https://localhost/organizations/test/groups/admins)(?=.*billing-admins:https://localhost/organizations/test/groups/billing-admins)(?=.*clients:https://localhost/organizations/test/groups/clients)(?=.*users:https://localhost/organizations/test/groups/users)(?=.*public_key_read_access:https://localhost/organizations/test/groups/public_key_read_access).*\]EndFinalList}) }
end
