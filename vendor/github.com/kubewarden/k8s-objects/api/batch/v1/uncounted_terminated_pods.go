// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// UncountedTerminatedPods UncountedTerminatedPods holds UIDs of Pods that have terminated but haven't been accounted in Job status counters.
//
// swagger:model UncountedTerminatedPods
type UncountedTerminatedPods struct {

	// Failed holds UIDs of failed Pods.
	Failed []string `json:"failed,omitempty"`

	// Succeeded holds UIDs of succeeded Pods.
	Succeeded []string `json:"succeeded,omitempty"`
}