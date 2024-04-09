package subscription

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

type Subscription struct {
	State            RegistrationState `json:"state"`
	RegsitrationDate *string           `json:"registrationDate,omitempty"`
	Properties       *Properties       `json:"properties,omitempty"`
}

type Properties struct {
	TenantId             *string              `json:"tenantId,omitempty"`
	LocationPlacementId  *string              `json:"locationPlacementId,omitempty"`
	QuotaId              *string              `json:"quotaId,omitempty"`
	RegisteredFeatures   *[]Feature           `json:"registeredFeatures,omitempty"`
	AvailabilityZones    *AvailabilityZone    `json:"availabilityZones,omitempty"`
	SpendingLimit        *string              `json:"spendingLimit,omitempty"`
	AccountOwner         *map[string]string   `json:"accountOwner,omitempty"`
	ManagedByTenants     *[]map[string]string `json:"managedByTenants,omitempty"`
	AdditionalProperties *map[string]string   `json:"additionalProperties,omitempty"`
}

type Feature struct {
	Name  *string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
}

type AvailabilityZone struct {
	Location     *string        `json:"location,omitempty"`
	ZoneMappings *[]ZoneMapping `json:"zoneMppings,omitempty"`
}

type ZoneMapping struct {
	LogicalZone  *string `json:"logicalZone,omitempty"`
	PhysicalZone *string `json:"physicalZone,omitempty"`
}
type RegistrationState string

const (
	Registered   RegistrationState = "Registered"
	Unregistered RegistrationState = "Unregistered"
	Warned       RegistrationState = "Warned"
	Deleted      RegistrationState = "Deleted"
	Suspended    RegistrationState = "Suspended"
)
