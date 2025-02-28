//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azsecrets

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets/internal/generated"
	shared "github.com/Azure/azure-sdk-for-go/sdk/keyvault/internal"
)

// DeletedSecret consists of the previous ID, attributes, tags, and information on when it will be purged.
type DeletedSecret struct {
	// The secret management attributes.
	Properties *Properties

	// The secret id.
	ID *string

	// Name of the secret
	Name *string

	// The url of the recovery object, used to identify and recover the deleted secret.
	RecoveryID *string

	// READ-ONLY; The time when the secret was deleted, in UTC
	DeletedOn *time.Time

	// READ-ONLY; The time when the secret is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time
}

// Secret - A secret consisting of a value, id and its attributes.
type Secret struct {
	// The secret management attributes.
	Properties *Properties

	// The secret id.
	ID *string

	// The name of the secret
	Name *string

	// The secret value.
	Value *string
}

// Properties - The secret management properties.
type Properties struct {
	// The content type of the secret.
	ContentType *string

	// READ-ONLY; Creation time in UTC.
	CreatedOn *time.Time

	// Determines whether the object is enabled.
	Enabled *bool

	// Expiry date in UTC.
	ExpiresOn *time.Time

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a secret backing a certificate, then managed
	// will be true.
	Managed *bool

	// READ-ONLY; If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV
	// certificate.
	KeyID *string

	// NotBefore is the secret's not before date in UTC.
	NotBefore *time.Time

	// READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.
	RecoverableDays *int32

	// READ-ONLY; Reflects the deletion recovery level currently in effect for secrets in the current vault. If it contains 'Purgeable', the secret can be permanently
	// deleted by a privileged user; otherwise, only the
	// system can purge the secret, at the end of the retention interval.
	RecoveryLevel *string

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string

	// READ-ONLY; Last updated time in UTC.
	UpdatedOn *time.Time

	// VaultURL is the vault url the secret came from
	VaultURL *string

	// Version is the version of the secret
	Version *string

	// Name is the name of the secret
	Name *string
}

func (s *Properties) toGenerated() *generated.SecretAttributes {
	if s == nil {
		return nil
	}
	return &generated.SecretAttributes{
		RecoverableDays: s.RecoverableDays,
		RecoveryLevel:   (*generated.DeletionRecoveryLevel)(s.RecoveryLevel),
		Enabled:         s.Enabled,
		Expires:         s.ExpiresOn,
		NotBefore:       s.NotBefore,
		Created:         s.CreatedOn,
		Updated:         s.UpdatedOn,
	}
}

// create a SecretAttributes object from an generated.SecretAttributes object
func secretPropertiesFromGenerated(i *generated.SecretAttributes, ID, contentType, keyID *string, managed *bool, tags map[string]*string) *Properties {
	if i == nil {
		return nil
	}
	vaultURL, name, version := shared.ParseID(ID)
	return &Properties{
		ContentType:     contentType,
		CreatedOn:       i.Created,
		Enabled:         i.Enabled,
		ExpiresOn:       i.Expires,
		KeyID:           keyID,
		Managed:         managed,
		Name:            name,
		NotBefore:       i.NotBefore,
		RecoverableDays: i.RecoverableDays,
		RecoveryLevel:   (*string)(i.RecoveryLevel),
		Tags:            tags,
		UpdatedOn:       i.Updated,
		VaultURL:        vaultURL,
		Version:         version,
	}
}

// SecretItem contains secret metadata.
type SecretItem struct {
	// The secret management attributes.
	Properties *Properties

	// Secret identifier.
	ID *string

	// Name of the secret
	Name *string
}

// create a SecretItem from the generated.SecretItem model
func secretItemFromGenerated(i *generated.SecretItem) *SecretItem {
	if i == nil {
		return nil
	}

	_, name, _ := shared.ParseID(i.ID)
	return &SecretItem{
		Properties: secretPropertiesFromGenerated(i.Attributes, i.ID, i.ContentType, nil, i.Managed, i.Tags),
		ID:         i.ID,
		Name:       name,
	}
}

// DeletedSecretItem - The deleted secret item containing metadata about the deleted secret.
type DeletedSecretItem struct {
	// The secret management attributes.
	Properties *Properties

	// Secret identifier.
	ID *string

	// The name of the deleted secret
	Name *string

	// The url of the recovery object, used to identify and recover the deleted secret.
	RecoveryID *string

	// READ-ONLY; The time when the secret was deleted, in UTC
	DeletedOn *time.Time

	// READ-ONLY; The time when the secret is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time
}

func deletedSecretItemFromGenerated(i *generated.DeletedSecretItem) *DeletedSecretItem {
	if i == nil {
		return nil
	}

	_, name, _ := shared.ParseID(i.ID)
	return &DeletedSecretItem{
		Properties:         secretPropertiesFromGenerated(i.Attributes, i.ID, i.ContentType, nil, i.Managed, i.Tags),
		Name:               name,
		ID:                 i.ID,
		RecoveryID:         i.RecoveryID,
		DeletedOn:          i.DeletedDate,
		ScheduledPurgeDate: i.ScheduledPurgeDate,
	}
}
