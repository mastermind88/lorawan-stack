// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package store

import "github.com/TheThingsNetwork/ttn/pkg/ttnpb"

// Gateway is the interface of all things that can be a gateway.
type Gateway interface {
	// GetGateway returns the ttnpb.Gateway that represents this gateway.
	GetGateway() *ttnpb.Gateway

	// SetAttributes sets the free-form attributes.
	SetAttributes(map[string]string)

	// SetAntennas sets the antennas.
	SetAntennas([]ttnpb.GatewayAntenna)

	// SetRadios sets the radios.
	SetRadios([]ttnpb.GatewayRadio)
}

// GatewaySpecializer returns a new Gateway with the given base ttnpb.Gateway.
type GatewaySpecializer func(ttnpb.Gateway) Gateway

// GatewayStore is a store that holds Gateways.
// nolint: dupl
type GatewayStore interface {
	// Create creates a new gateway.
	Create(Gateway) error

	// GetByID finds a gateway by ID and retrieves it.
	GetByID(ttnpb.GatewayIdentifiers, GatewaySpecializer) (Gateway, error)

	// ListByOrganizationOrUser returns all the gateways to which an organization
	// or user is collaborator of.
	ListByOrganizationOrUser(ttnpb.OrganizationOrUserIdentifiers, GatewaySpecializer) ([]Gateway, error)

	// Update updates the gateway.
	Update(Gateway) error

	// TODO(gomezjdaniel#274): use sql 'ON DELETE CASCADE' when CockroachDB implements it.
	// Delete deletes a gateway.
	Delete(ttnpb.GatewayIdentifiers) error

	// SaveAPIKey stores an API Key attached to a gateway.
	SaveAPIKey(ttnpb.GatewayIdentifiers, ttnpb.APIKey) error

	// GetAPIKey retrieves an API key by value and the gateway identifiers.
	GetAPIKey(string) (ttnpb.GatewayIdentifiers, ttnpb.APIKey, error)

	// GetAPIKeyByName retrieves an API key from a gateway.
	GetAPIKeyByName(ttnpb.GatewayIdentifiers, string) (ttnpb.APIKey, error)

	// UpdateAPIKeyRights updates the right of an API key.
	UpdateAPIKeyRights(ttnpb.GatewayIdentifiers, string, []ttnpb.Right) error

	// ListAPIKeys list all the API keys that a gateway has.
	ListAPIKeys(ttnpb.GatewayIdentifiers) ([]ttnpb.APIKey, error)

	// DeleteAPIKey deletes a given API key from a gateway.
	DeleteAPIKey(ttnpb.GatewayIdentifiers, string) error

	// SetCollaborator inserts or updates a collaborator within a gateway.
	// If the list of rights is empty the collaborator will be unset.
	SetCollaborator(ttnpb.GatewayCollaborator) error

	// HasCollaboratorRights checks whether a collaborator has a given set of rights
	// to a gateway. It returns false if the collaborationship does not exist.
	HasCollaboratorRights(ttnpb.GatewayIdentifiers, ttnpb.OrganizationOrUserIdentifiers, ...ttnpb.Right) (bool, error)

	// ListCollaborators retrieves all the gateway collaborators.
	// Optionally a list of rights can be passed to filter them.
	ListCollaborators(ttnpb.GatewayIdentifiers, ...ttnpb.Right) ([]ttnpb.GatewayCollaborator, error)

	// ListCollaboratorRights returns the rights a given collaborator has for a
	// Gateway. Returns empty list if the collaborationship does not exist.
	ListCollaboratorRights(ttnpb.GatewayIdentifiers, ttnpb.OrganizationOrUserIdentifiers) ([]ttnpb.Right, error)

	// LoadAttributes loads extra attributes into the gateway if it's an Attributer.
	LoadAttributes(ttnpb.GatewayIdentifiers, Gateway) error

	// StoreAttributes writes the extra attributes on the gatewat if it's an
	// Attributer to the store.
	StoreAttributes(ttnpb.GatewayIdentifiers, Gateway) error
}
