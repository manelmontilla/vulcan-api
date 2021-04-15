// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": Application User Types
//
// Command:
// $ main

package client

import (
	"github.com/goadesign/goa"
	"time"
)

// assetGroupPayload user type.
type assetGroupPayload struct {
	// Asset ID
	AssetID *string `form:"asset_id,omitempty" json:"asset_id,omitempty" yaml:"asset_id,omitempty" xml:"asset_id,omitempty"`
}

// Validate validates the assetGroupPayload type instance.
func (ut *assetGroupPayload) Validate() (err error) {
	if ut.AssetID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "asset_id"))
	}
	return
}

// Publicize creates AssetGroupPayload from assetGroupPayload
func (ut *assetGroupPayload) Publicize() *AssetGroupPayload {
	var pub AssetGroupPayload
	if ut.AssetID != nil {
		pub.AssetID = *ut.AssetID
	}
	return &pub
}

// AssetGroupPayload user type.
type AssetGroupPayload struct {
	// Asset ID
	AssetID string `form:"asset_id" json:"asset_id" yaml:"asset_id" xml:"asset_id"`
}

// Validate validates the AssetGroupPayload type instance.
func (ut *AssetGroupPayload) Validate() (err error) {
	if ut.AssetID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "asset_id"))
	}
	return
}

// assetPayload user type.
type assetPayload struct {
	// The alias of the asset in Vulcan
	Alias *string `form:"alias,omitempty" json:"alias,omitempty" yaml:"alias,omitempty" xml:"alias,omitempty"`
	// Environmental CVSS
	EnvironmentalCvss *string `form:"environmental_cvss,omitempty" json:"environmental_cvss,omitempty" yaml:"environmental_cvss,omitempty" xml:"environmental_cvss,omitempty"`
	// Identifier
	Identifier *string `form:"identifier,omitempty" json:"identifier,omitempty" yaml:"identifier,omitempty" xml:"identifier,omitempty"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Rolfp plus scope vector
	Rolfp *string `form:"rolfp,omitempty" json:"rolfp,omitempty" yaml:"rolfp,omitempty" xml:"rolfp,omitempty"`
	// Scannable
	Scannable *bool `form:"scannable,omitempty" json:"scannable,omitempty" yaml:"scannable,omitempty" xml:"scannable,omitempty"`
	// Type
	Type *string `form:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
}

// Validate validates the assetPayload type instance.
func (ut *assetPayload) Validate() (err error) {
	if ut.Identifier == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "identifier"))
	}
	return
}

// Publicize creates AssetPayload from assetPayload
func (ut *assetPayload) Publicize() *AssetPayload {
	var pub AssetPayload
	if ut.Alias != nil {
		pub.Alias = ut.Alias
	}
	if ut.EnvironmentalCvss != nil {
		pub.EnvironmentalCvss = ut.EnvironmentalCvss
	}
	if ut.Identifier != nil {
		pub.Identifier = *ut.Identifier
	}
	if ut.Options != nil {
		pub.Options = ut.Options
	}
	if ut.Rolfp != nil {
		pub.Rolfp = ut.Rolfp
	}
	if ut.Scannable != nil {
		pub.Scannable = ut.Scannable
	}
	if ut.Type != nil {
		pub.Type = ut.Type
	}
	return &pub
}

// AssetPayload user type.
type AssetPayload struct {
	// The alias of the asset in Vulcan
	Alias *string `form:"alias,omitempty" json:"alias,omitempty" yaml:"alias,omitempty" xml:"alias,omitempty"`
	// Environmental CVSS
	EnvironmentalCvss *string `form:"environmental_cvss,omitempty" json:"environmental_cvss,omitempty" yaml:"environmental_cvss,omitempty" xml:"environmental_cvss,omitempty"`
	// Identifier
	Identifier string `form:"identifier" json:"identifier" yaml:"identifier" xml:"identifier"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Rolfp plus scope vector
	Rolfp *string `form:"rolfp,omitempty" json:"rolfp,omitempty" yaml:"rolfp,omitempty" xml:"rolfp,omitempty"`
	// Scannable
	Scannable *bool `form:"scannable,omitempty" json:"scannable,omitempty" yaml:"scannable,omitempty" xml:"scannable,omitempty"`
	// Type
	Type *string `form:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
}

// Validate validates the AssetPayload type instance.
func (ut *AssetPayload) Validate() (err error) {
	if ut.Identifier == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "identifier"))
	}
	return
}

// assetUpdatePayload user type.
type assetUpdatePayload struct {
	// Alias
	Alias *string `form:"alias,omitempty" json:"alias,omitempty" yaml:"alias,omitempty" xml:"alias,omitempty"`
	// Environmental CVSS
	EnvironmentalCvss *string `form:"environmental_cvss,omitempty" json:"environmental_cvss,omitempty" yaml:"environmental_cvss,omitempty" xml:"environmental_cvss,omitempty"`
	// Identifier
	Identifier *string `form:"identifier,omitempty" json:"identifier,omitempty" yaml:"identifier,omitempty" xml:"identifier,omitempty"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Rolfp plus scope vector
	Rolfp *string `form:"rolfp,omitempty" json:"rolfp,omitempty" yaml:"rolfp,omitempty" xml:"rolfp,omitempty"`
	// Scannable
	Scannable *bool `form:"scannable,omitempty" json:"scannable,omitempty" yaml:"scannable,omitempty" xml:"scannable,omitempty"`
	// Type
	Type *string `form:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
}

// Publicize creates AssetUpdatePayload from assetUpdatePayload
func (ut *assetUpdatePayload) Publicize() *AssetUpdatePayload {
	var pub AssetUpdatePayload
	if ut.Alias != nil {
		pub.Alias = ut.Alias
	}
	if ut.EnvironmentalCvss != nil {
		pub.EnvironmentalCvss = ut.EnvironmentalCvss
	}
	if ut.Identifier != nil {
		pub.Identifier = ut.Identifier
	}
	if ut.Options != nil {
		pub.Options = ut.Options
	}
	if ut.Rolfp != nil {
		pub.Rolfp = ut.Rolfp
	}
	if ut.Scannable != nil {
		pub.Scannable = ut.Scannable
	}
	if ut.Type != nil {
		pub.Type = ut.Type
	}
	return &pub
}

// AssetUpdatePayload user type.
type AssetUpdatePayload struct {
	// Alias
	Alias *string `form:"alias,omitempty" json:"alias,omitempty" yaml:"alias,omitempty" xml:"alias,omitempty"`
	// Environmental CVSS
	EnvironmentalCvss *string `form:"environmental_cvss,omitempty" json:"environmental_cvss,omitempty" yaml:"environmental_cvss,omitempty" xml:"environmental_cvss,omitempty"`
	// Identifier
	Identifier *string `form:"identifier,omitempty" json:"identifier,omitempty" yaml:"identifier,omitempty" xml:"identifier,omitempty"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Rolfp plus scope vector
	Rolfp *string `form:"rolfp,omitempty" json:"rolfp,omitempty" yaml:"rolfp,omitempty" xml:"rolfp,omitempty"`
	// Scannable
	Scannable *bool `form:"scannable,omitempty" json:"scannable,omitempty" yaml:"scannable,omitempty" xml:"scannable,omitempty"`
	// Type
	Type *string `form:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
}

// createAssetPayload user type.
type createAssetPayload struct {
	Assets []*assetPayload `form:"assets,omitempty" json:"assets,omitempty" yaml:"assets,omitempty" xml:"assets,omitempty"`
	Groups []string        `form:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty" xml:"groups,omitempty"`
}

// Validate validates the createAssetPayload type instance.
func (ut *createAssetPayload) Validate() (err error) {
	if ut.Assets == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "assets"))
	}
	for _, e := range ut.Assets {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Publicize creates CreateAssetPayload from createAssetPayload
func (ut *createAssetPayload) Publicize() *CreateAssetPayload {
	var pub CreateAssetPayload
	if ut.Assets != nil {
		pub.Assets = make([]*AssetPayload, len(ut.Assets))
		for i2, elem2 := range ut.Assets {
			pub.Assets[i2] = elem2.Publicize()
		}
	}
	if ut.Groups != nil {
		pub.Groups = ut.Groups
	}
	return &pub
}

// CreateAssetPayload user type.
type CreateAssetPayload struct {
	Assets []*AssetPayload `form:"assets" json:"assets" yaml:"assets" xml:"assets"`
	Groups []string        `form:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty" xml:"groups,omitempty"`
}

// Validate validates the CreateAssetPayload type instance.
func (ut *CreateAssetPayload) Validate() (err error) {
	if ut.Assets == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "assets"))
	}
	for _, e := range ut.Assets {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// digestPayload user type.
type digestPayload struct {
	// End Date
	EndDate *string `form:"end_date,omitempty" json:"end_date,omitempty" yaml:"end_date,omitempty" xml:"end_date,omitempty"`
	// Start Date
	StartDate *string `form:"start_date,omitempty" json:"start_date,omitempty" yaml:"start_date,omitempty" xml:"start_date,omitempty"`
}

// Publicize creates DigestPayload from digestPayload
func (ut *digestPayload) Publicize() *DigestPayload {
	var pub DigestPayload
	if ut.EndDate != nil {
		pub.EndDate = ut.EndDate
	}
	if ut.StartDate != nil {
		pub.StartDate = ut.StartDate
	}
	return &pub
}

// DigestPayload user type.
type DigestPayload struct {
	// End Date
	EndDate *string `form:"end_date,omitempty" json:"end_date,omitempty" yaml:"end_date,omitempty" xml:"end_date,omitempty"`
	// Start Date
	StartDate *string `form:"start_date,omitempty" json:"start_date,omitempty" yaml:"start_date,omitempty" xml:"start_date,omitempty"`
}

// findingOverwritePayload user type.
type findingOverwritePayload struct {
	// Notes
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" yaml:"notes,omitempty" xml:"notes,omitempty"`
	// Status
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
}

// Validate validates the findingOverwritePayload type instance.
func (ut *findingOverwritePayload) Validate() (err error) {
	if ut.Status == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "status"))
	}
	if ut.Notes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "notes"))
	}
	return
}

// Publicize creates FindingOverwritePayload from findingOverwritePayload
func (ut *findingOverwritePayload) Publicize() *FindingOverwritePayload {
	var pub FindingOverwritePayload
	if ut.Notes != nil {
		pub.Notes = *ut.Notes
	}
	if ut.Status != nil {
		pub.Status = *ut.Status
	}
	return &pub
}

// FindingOverwritePayload user type.
type FindingOverwritePayload struct {
	// Notes
	Notes string `form:"notes" json:"notes" yaml:"notes" xml:"notes"`
	// Status
	Status string `form:"status" json:"status" yaml:"status" xml:"status"`
}

// Validate validates the FindingOverwritePayload type instance.
func (ut *FindingOverwritePayload) Validate() (err error) {
	if ut.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "status"))
	}
	if ut.Notes == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "notes"))
	}
	return
}

// groupPayload user type.
type groupPayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
}

// Validate validates the groupPayload type instance.
func (ut *groupPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	return
}

// Publicize creates GroupPayload from groupPayload
func (ut *groupPayload) Publicize() *GroupPayload {
	var pub GroupPayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Options != nil {
		pub.Options = ut.Options
	}
	return &pub
}

// GroupPayload user type.
type GroupPayload struct {
	// name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
}

// Validate validates the GroupPayload type instance.
func (ut *GroupPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	return
}

// policyPayload user type.
type policyPayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the policyPayload type instance.
func (ut *policyPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	return
}

// Publicize creates PolicyPayload from policyPayload
func (ut *policyPayload) Publicize() *PolicyPayload {
	var pub PolicyPayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// PolicyPayload user type.
type PolicyPayload struct {
	// name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the PolicyPayload type instance.
func (ut *PolicyPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	return
}

// policySettingPayload user type.
type policySettingPayload struct {
	// Check Type Name
	ChecktypeName *string `form:"checktype_name,omitempty" json:"checktype_name,omitempty" yaml:"checktype_name,omitempty" xml:"checktype_name,omitempty"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
}

// Validate validates the policySettingPayload type instance.
func (ut *policySettingPayload) Validate() (err error) {
	if ut.ChecktypeName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "checktype_name"))
	}
	return
}

// Publicize creates PolicySettingPayload from policySettingPayload
func (ut *policySettingPayload) Publicize() *PolicySettingPayload {
	var pub PolicySettingPayload
	if ut.ChecktypeName != nil {
		pub.ChecktypeName = *ut.ChecktypeName
	}
	if ut.Options != nil {
		pub.Options = ut.Options
	}
	return &pub
}

// PolicySettingPayload user type.
type PolicySettingPayload struct {
	// Check Type Name
	ChecktypeName string `form:"checktype_name" json:"checktype_name" yaml:"checktype_name" xml:"checktype_name"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
}

// Validate validates the PolicySettingPayload type instance.
func (ut *PolicySettingPayload) Validate() (err error) {
	if ut.ChecktypeName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "checktype_name"))
	}
	return
}

// policySettingUploadPayload user type.
type policySettingUploadPayload struct {
	// Check Type Name
	ChecktypeName *string `form:"checktype_name,omitempty" json:"checktype_name,omitempty" yaml:"checktype_name,omitempty" xml:"checktype_name,omitempty"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
}

// Publicize creates PolicySettingUploadPayload from policySettingUploadPayload
func (ut *policySettingUploadPayload) Publicize() *PolicySettingUploadPayload {
	var pub PolicySettingUploadPayload
	if ut.ChecktypeName != nil {
		pub.ChecktypeName = ut.ChecktypeName
	}
	if ut.Options != nil {
		pub.Options = ut.Options
	}
	return &pub
}

// PolicySettingUploadPayload user type.
type PolicySettingUploadPayload struct {
	// Check Type Name
	ChecktypeName *string `form:"checktype_name,omitempty" json:"checktype_name,omitempty" yaml:"checktype_name,omitempty" xml:"checktype_name,omitempty"`
	// Options
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
}

// policyUpdatePayload user type.
type policyUpdatePayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Publicize creates PolicyUpdatePayload from policyUpdatePayload
func (ut *policyUpdatePayload) Publicize() *PolicyUpdatePayload {
	var pub PolicyUpdatePayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// PolicyUpdatePayload user type.
type PolicyUpdatePayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// programPayload user type.
type programPayload struct {
	// Autosend
	Autosend *bool `form:"autosend,omitempty" json:"autosend,omitempty" yaml:"autosend,omitempty" xml:"autosend,omitempty"`
	// Disabled
	Disabled *bool `form:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" xml:"disabled,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// PolicyGroups
	PolicyGroups []*programPolicyGroupPayload `form:"policy_groups,omitempty" json:"policy_groups,omitempty" yaml:"policy_groups,omitempty" xml:"policy_groups,omitempty"`
}

// Publicize creates ProgramPayload from programPayload
func (ut *programPayload) Publicize() *ProgramPayload {
	var pub ProgramPayload
	if ut.Autosend != nil {
		pub.Autosend = ut.Autosend
	}
	if ut.Disabled != nil {
		pub.Disabled = ut.Disabled
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.PolicyGroups != nil {
		pub.PolicyGroups = make([]*ProgramPolicyGroupPayload, len(ut.PolicyGroups))
		for i2, elem2 := range ut.PolicyGroups {
			pub.PolicyGroups[i2] = elem2.Publicize()
		}
	}
	return &pub
}

// ProgramPayload user type.
type ProgramPayload struct {
	// Autosend
	Autosend *bool `form:"autosend,omitempty" json:"autosend,omitempty" yaml:"autosend,omitempty" xml:"autosend,omitempty"`
	// Disabled
	Disabled *bool `form:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" xml:"disabled,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// PolicyGroups
	PolicyGroups []*ProgramPolicyGroupPayload `form:"policy_groups,omitempty" json:"policy_groups,omitempty" yaml:"policy_groups,omitempty" xml:"policy_groups,omitempty"`
}

// programPolicyGroupPayload user type.
type programPolicyGroupPayload struct {
	// group
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" yaml:"group_id,omitempty" xml:"group_id,omitempty"`
	// policy
	PolicyID *string `form:"policy_id,omitempty" json:"policy_id,omitempty" yaml:"policy_id,omitempty" xml:"policy_id,omitempty"`
}

// Publicize creates ProgramPolicyGroupPayload from programPolicyGroupPayload
func (ut *programPolicyGroupPayload) Publicize() *ProgramPolicyGroupPayload {
	var pub ProgramPolicyGroupPayload
	if ut.GroupID != nil {
		pub.GroupID = ut.GroupID
	}
	if ut.PolicyID != nil {
		pub.PolicyID = ut.PolicyID
	}
	return &pub
}

// ProgramPolicyGroupPayload user type.
type ProgramPolicyGroupPayload struct {
	// group
	GroupID *string `form:"group_id,omitempty" json:"group_id,omitempty" yaml:"group_id,omitempty" xml:"group_id,omitempty"`
	// policy
	PolicyID *string `form:"policy_id,omitempty" json:"policy_id,omitempty" yaml:"policy_id,omitempty" xml:"policy_id,omitempty"`
}

// programUpdatePayload user type.
type programUpdatePayload struct {
	// Autosend
	Autosend *bool `form:"autosend,omitempty" json:"autosend,omitempty" yaml:"autosend,omitempty" xml:"autosend,omitempty"`
	// Disabled
	Disabled *bool `form:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" xml:"disabled,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// PolicyGroups
	PolicyGroups []*programPolicyGroupPayload `form:"policy_groups,omitempty" json:"policy_groups,omitempty" yaml:"policy_groups,omitempty" xml:"policy_groups,omitempty"`
}

// Publicize creates ProgramUpdatePayload from programUpdatePayload
func (ut *programUpdatePayload) Publicize() *ProgramUpdatePayload {
	var pub ProgramUpdatePayload
	if ut.Autosend != nil {
		pub.Autosend = ut.Autosend
	}
	if ut.Disabled != nil {
		pub.Disabled = ut.Disabled
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.PolicyGroups != nil {
		pub.PolicyGroups = make([]*ProgramPolicyGroupPayload, len(ut.PolicyGroups))
		for i2, elem2 := range ut.PolicyGroups {
			pub.PolicyGroups[i2] = elem2.Publicize()
		}
	}
	return &pub
}

// ProgramUpdatePayload user type.
type ProgramUpdatePayload struct {
	// Autosend
	Autosend *bool `form:"autosend,omitempty" json:"autosend,omitempty" yaml:"autosend,omitempty" xml:"autosend,omitempty"`
	// Disabled
	Disabled *bool `form:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" xml:"disabled,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// PolicyGroups
	PolicyGroups []*ProgramPolicyGroupPayload `form:"policy_groups,omitempty" json:"policy_groups,omitempty" yaml:"policy_groups,omitempty" xml:"policy_groups,omitempty"`
}

// recipientsPayload user type.
type recipientsPayload struct {
	// Emails
	Emails []string `form:"emails,omitempty" json:"emails,omitempty" yaml:"emails,omitempty" xml:"emails,omitempty"`
}

// Validate validates the recipientsPayload type instance.
func (ut *recipientsPayload) Validate() (err error) {
	if ut.Emails == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "emails"))
	}
	return
}

// Publicize creates RecipientsPayload from recipientsPayload
func (ut *recipientsPayload) Publicize() *RecipientsPayload {
	var pub RecipientsPayload
	if ut.Emails != nil {
		pub.Emails = ut.Emails
	}
	return &pub
}

// RecipientsPayload user type.
type RecipientsPayload struct {
	// Emails
	Emails []string `form:"emails" json:"emails" yaml:"emails" xml:"emails"`
}

// Validate validates the RecipientsPayload type instance.
func (ut *RecipientsPayload) Validate() (err error) {
	if ut.Emails == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "emails"))
	}
	return
}

// scanPayload user type.
type scanPayload struct {
	// Program ID
	ProgramID *string `form:"program_id,omitempty" json:"program_id,omitempty" yaml:"program_id,omitempty" xml:"program_id,omitempty"`
	// Group ID
	ScheduledTime *time.Time `form:"scheduled_time,omitempty" json:"scheduled_time,omitempty" yaml:"scheduled_time,omitempty" xml:"scheduled_time,omitempty"`
}

// Validate validates the scanPayload type instance.
func (ut *scanPayload) Validate() (err error) {
	if ut.ProgramID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "program_id"))
	}
	return
}

// Publicize creates ScanPayload from scanPayload
func (ut *scanPayload) Publicize() *ScanPayload {
	var pub ScanPayload
	if ut.ProgramID != nil {
		pub.ProgramID = *ut.ProgramID
	}
	if ut.ScheduledTime != nil {
		pub.ScheduledTime = ut.ScheduledTime
	}
	return &pub
}

// ScanPayload user type.
type ScanPayload struct {
	// Program ID
	ProgramID string `form:"program_id" json:"program_id" yaml:"program_id" xml:"program_id"`
	// Group ID
	ScheduledTime *time.Time `form:"scheduled_time,omitempty" json:"scheduled_time,omitempty" yaml:"scheduled_time,omitempty" xml:"scheduled_time,omitempty"`
}

// Validate validates the ScanPayload type instance.
func (ut *ScanPayload) Validate() (err error) {
	if ut.ProgramID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "program_id"))
	}
	return
}

// schedulePayload user type.
type schedulePayload struct {
	// Cron Expression
	Cron *string `form:"cron,omitempty" json:"cron,omitempty" yaml:"cron,omitempty" xml:"cron,omitempty"`
}

// Publicize creates SchedulePayload from schedulePayload
func (ut *schedulePayload) Publicize() *SchedulePayload {
	var pub SchedulePayload
	if ut.Cron != nil {
		pub.Cron = ut.Cron
	}
	return &pub
}

// SchedulePayload user type.
type SchedulePayload struct {
	// Cron Expression
	Cron *string `form:"cron,omitempty" json:"cron,omitempty" yaml:"cron,omitempty" xml:"cron,omitempty"`
}

// scheduleUpdatePayload user type.
type scheduleUpdatePayload struct {
	// Cron Expression
	Cron *string `form:"cron,omitempty" json:"cron,omitempty" yaml:"cron,omitempty" xml:"cron,omitempty"`
}

// Publicize creates ScheduleUpdatePayload from scheduleUpdatePayload
func (ut *scheduleUpdatePayload) Publicize() *ScheduleUpdatePayload {
	var pub ScheduleUpdatePayload
	if ut.Cron != nil {
		pub.Cron = ut.Cron
	}
	return &pub
}

// ScheduleUpdatePayload user type.
type ScheduleUpdatePayload struct {
	// Cron Expression
	Cron *string `form:"cron,omitempty" json:"cron,omitempty" yaml:"cron,omitempty" xml:"cron,omitempty"`
}

// teamMemberPayload user type.
type teamMemberPayload struct {
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// Member role. Valid values are: owner, member
	Role *string `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
	// User ID
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" yaml:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Publicize creates TeamMemberPayload from teamMemberPayload
func (ut *teamMemberPayload) Publicize() *TeamMemberPayload {
	var pub TeamMemberPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Role != nil {
		pub.Role = ut.Role
	}
	if ut.UserID != nil {
		pub.UserID = ut.UserID
	}
	return &pub
}

// TeamMemberPayload user type.
type TeamMemberPayload struct {
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// Member role. Valid values are: owner, member
	Role *string `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
	// User ID
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" yaml:"user_id,omitempty" xml:"user_id,omitempty"`
}

// teamMemberUpdatePayload user type.
type teamMemberUpdatePayload struct {
	// Member role. Valid values are: owner, member
	Role *string `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
}

// Publicize creates TeamMemberUpdatePayload from teamMemberUpdatePayload
func (ut *teamMemberUpdatePayload) Publicize() *TeamMemberUpdatePayload {
	var pub TeamMemberUpdatePayload
	if ut.Role != nil {
		pub.Role = ut.Role
	}
	return &pub
}

// TeamMemberUpdatePayload user type.
type TeamMemberUpdatePayload struct {
	// Member role. Valid values are: owner, member
	Role *string `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
}

// teamPayload user type.
type teamPayload struct {
	// description
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// tag
	Tag *string `form:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
}

// Validate validates the teamPayload type instance.
func (ut *teamPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Tag == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "tag"))
	}
	return
}

// Publicize creates TeamPayload from teamPayload
func (ut *teamPayload) Publicize() *TeamPayload {
	var pub TeamPayload
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Tag != nil {
		pub.Tag = *ut.Tag
	}
	return &pub
}

// TeamPayload user type.
type TeamPayload struct {
	// description
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// tag
	Tag string `form:"tag" json:"tag" yaml:"tag" xml:"tag"`
}

// Validate validates the TeamPayload type instance.
func (ut *TeamPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Tag == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "tag"))
	}
	return
}

// teamUpdatePayload user type.
type teamUpdatePayload struct {
	// description
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// tag
	Tag *string `form:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
}

// Publicize creates TeamUpdatePayload from teamUpdatePayload
func (ut *teamUpdatePayload) Publicize() *TeamUpdatePayload {
	var pub TeamUpdatePayload
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Tag != nil {
		pub.Tag = ut.Tag
	}
	return &pub
}

// TeamUpdatePayload user type.
type TeamUpdatePayload struct {
	// description
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// tag
	Tag *string `form:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
}

// userPayload user type.
type userPayload struct {
	// Active (Default: true)
	Active *bool `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	// Admin
	Admin *bool `form:"admin,omitempty" json:"admin,omitempty" yaml:"admin,omitempty" xml:"admin,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// First Name
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty" xml:"firstname,omitempty"`
	// Last Name
	Lastname *string `form:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty" xml:"lastname,omitempty"`
	// Observer
	Observer *bool `form:"observer,omitempty" json:"observer,omitempty" yaml:"observer,omitempty" xml:"observer,omitempty"`
}

// Finalize sets the default values for userPayload type instance.
func (ut *userPayload) Finalize() {
	var defaultAdmin = false
	if ut.Admin == nil {
		ut.Admin = &defaultAdmin
	}
	var defaultObserver = false
	if ut.Observer == nil {
		ut.Observer = &defaultObserver
	}
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Active != nil {
		pub.Active = ut.Active
	}
	if ut.Admin != nil {
		pub.Admin = *ut.Admin
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Firstname != nil {
		pub.Firstname = ut.Firstname
	}
	if ut.Lastname != nil {
		pub.Lastname = ut.Lastname
	}
	if ut.Observer != nil {
		pub.Observer = *ut.Observer
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	// Active (Default: true)
	Active *bool `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	// Admin
	Admin bool `form:"admin" json:"admin" yaml:"admin" xml:"admin"`
	// Email
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// First Name
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty" xml:"firstname,omitempty"`
	// Last Name
	Lastname *string `form:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty" xml:"lastname,omitempty"`
	// Observer
	Observer bool `form:"observer" json:"observer" yaml:"observer" xml:"observer"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	return
}

// userUpdatePayload user type.
type userUpdatePayload struct {
	// Active (Default: true)
	Active *bool `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	// Admin
	Admin *bool `form:"admin,omitempty" json:"admin,omitempty" yaml:"admin,omitempty" xml:"admin,omitempty"`
	// First Name
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty" xml:"firstname,omitempty"`
	// Last Name
	Lastname *string `form:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty" xml:"lastname,omitempty"`
}

// Finalize sets the default values for userUpdatePayload type instance.
func (ut *userUpdatePayload) Finalize() {
	var defaultAdmin = false
	if ut.Admin == nil {
		ut.Admin = &defaultAdmin
	}
}

// Publicize creates UserUpdatePayload from userUpdatePayload
func (ut *userUpdatePayload) Publicize() *UserUpdatePayload {
	var pub UserUpdatePayload
	if ut.Active != nil {
		pub.Active = ut.Active
	}
	if ut.Admin != nil {
		pub.Admin = *ut.Admin
	}
	if ut.Firstname != nil {
		pub.Firstname = ut.Firstname
	}
	if ut.Lastname != nil {
		pub.Lastname = ut.Lastname
	}
	return &pub
}

// UserUpdatePayload user type.
type UserUpdatePayload struct {
	// Active (Default: true)
	Active *bool `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	// Admin
	Admin bool `form:"admin" json:"admin" yaml:"admin" xml:"admin"`
	// First Name
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty" xml:"firstname,omitempty"`
	// Last Name
	Lastname *string `form:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty" xml:"lastname,omitempty"`
}