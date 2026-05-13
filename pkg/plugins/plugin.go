package plugin

import (
	sriovnetworkv1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
)

//go:generate ../../bin/mockgen -destination mock/mock_plugin.go -source plugin.go
type VendorPlugin interface {
	// Name returns the name of plugin
	Name() string
	// OnNodeStateChange is invoked when SriovNetworkNodeState CR is created or updated, return if need dain and/or reboot node
	OnNodeStateChange(*sriovnetworkv1.SriovNetworkNodeState) (bool, bool, error)
	// Apply config change
	Apply() error
	// CheckStatusChanges checks status changes on the SriovNetworkNodeState CR for configured VFs.
	CheckStatusChanges(*sriovnetworkv1.SriovNetworkNodeState) (bool, error)
	// NodeStateCheckLive checks if the min/max tx has been updated from previous state.
	NodeStateCheckLive(*sriovnetworkv1.SriovNetworkNodeState) bool
	// Apply live changes based on updated config for VFs. No reboot required.
	ApplyLiveChanges(*sriovnetworkv1.SriovNetworkNodeState) (bool, error)
}
