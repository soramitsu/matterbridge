// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AccessPackageAssignment undocumented
type AccessPackageAssignment struct {
	// Entity is the base model of AccessPackageAssignment
	Entity
	// CatalogID undocumented
	CatalogID *string `json:"catalogId,omitempty"`
	// AccessPackageID undocumented
	AccessPackageID *string `json:"accessPackageId,omitempty"`
	// AssignmentPolicyID undocumented
	AssignmentPolicyID *string `json:"assignmentPolicyId,omitempty"`
	// TargetID undocumented
	TargetID *string `json:"targetId,omitempty"`
	// AssignmentStatus undocumented
	AssignmentStatus *string `json:"assignmentStatus,omitempty"`
	// AssignmentState undocumented
	AssignmentState *string `json:"assignmentState,omitempty"`
	// IsExtended undocumented
	IsExtended *bool `json:"isExtended,omitempty"`
	// ExpiredDateTime undocumented
	ExpiredDateTime *time.Time `json:"expiredDateTime,omitempty"`
	// AccessPackage undocumented
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`
	// AccessPackageAssignmentPolicy undocumented
	AccessPackageAssignmentPolicy *AccessPackageAssignmentPolicy `json:"accessPackageAssignmentPolicy,omitempty"`
	// Target undocumented
	Target *AccessPackageSubject `json:"target,omitempty"`
	// AccessPackageAssignmentRequests undocumented
	AccessPackageAssignmentRequests []AccessPackageAssignmentRequestObject `json:"accessPackageAssignmentRequests,omitempty"`
	// AccessPackageAssignmentResourceRoles undocumented
	AccessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRole `json:"accessPackageAssignmentResourceRoles,omitempty"`
}