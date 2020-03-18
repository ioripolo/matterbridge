// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PrivilegedRoleAssignmentRequestObject undocumented
type PrivilegedRoleAssignmentRequestObject struct {
	// Entity is the base model of PrivilegedRoleAssignmentRequestObject
	Entity
	// Schedule undocumented
	Schedule *GovernanceSchedule `json:"schedule,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// RoleID undocumented
	RoleID *string `json:"roleId,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// AssignmentState undocumented
	AssignmentState *string `json:"assignmentState,omitempty"`
	// RequestedDateTime undocumented
	RequestedDateTime *time.Time `json:"requestedDateTime,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Duration undocumented
	Duration *string `json:"duration,omitempty"`
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// TicketNumber undocumented
	TicketNumber *string `json:"ticketNumber,omitempty"`
	// TicketSystem undocumented
	TicketSystem *string `json:"ticketSystem,omitempty"`
	// RoleInfo undocumented
	RoleInfo *PrivilegedRole `json:"roleInfo,omitempty"`
}