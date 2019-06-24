package configuration

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://ws.coverity.com/v9"

// NewConfigurationService creates an initializes a ConfigurationService.
func NewConfigurationService(cli *soap.Client) ConfigurationService {
	return &configurationService{cli}
}

// ConfigurationService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ConfigurationService interface {
	// CopyStream was auto-generated from WSDL.
	CopyStream(CopyStream *CopyStream) (*CopyStreamResponse, error)

	// CreateAttribute was auto-generated from WSDL.
	CreateAttribute(CreateAttribute *CreateAttribute) (*CreateAttributeResponse, error)

	// CreateComponentMap was auto-generated from WSDL.
	CreateComponentMap(CreateComponentMap *CreateComponentMap) (*CreateComponentMapResponse, error)

	// CreateGroup was auto-generated from WSDL.
	CreateGroup(CreateGroup *CreateGroup) (*CreateGroupResponse, error)

	// CreateLdapConfiguration was auto-generated from WSDL.
	CreateLdapConfiguration(CreateLdapConfiguration *CreateLdapConfiguration) (*CreateLdapConfigurationResponse, error)

	// CreateProject was auto-generated from WSDL.
	CreateProject(CreateProject *CreateProject) (*CreateProjectResponse, error)

	// CreateRole was auto-generated from WSDL.
	CreateRole(CreateRole *CreateRole) (*CreateRoleResponse, error)

	// CreateStream was auto-generated from WSDL.
	CreateStream(CreateStream *CreateStream) (*CreateStreamResponse, error)

	// CreateStreamInProject was auto-generated from WSDL.
	CreateStreamInProject(CreateStreamInProject *CreateStreamInProject) (*CreateStreamInProjectResponse, error)

	// CreateTriageStore was auto-generated from WSDL.
	CreateTriageStore(CreateTriageStore *CreateTriageStore) (*CreateTriageStoreResponse, error)

	// CreateUser was auto-generated from WSDL.
	CreateUser(CreateUser *CreateUser) (*CreateUserResponse, error)

	// DeleteAttribute was auto-generated from WSDL.
	DeleteAttribute(DeleteAttribute *DeleteAttribute) (*DeleteAttributeResponse, error)

	// DeleteComponentMap was auto-generated from WSDL.
	DeleteComponentMap(DeleteComponentMap *DeleteComponentMap) (*DeleteComponentMapResponse, error)

	// DeleteGroup was auto-generated from WSDL.
	DeleteGroup(DeleteGroup *DeleteGroup) (*DeleteGroupResponse, error)

	// DeleteLdapConfiguration was auto-generated from WSDL.
	DeleteLdapConfiguration(DeleteLdapConfiguration *DeleteLdapConfiguration) (*DeleteLdapConfigurationResponse, error)

	// DeleteProject was auto-generated from WSDL.
	DeleteProject(DeleteProject *DeleteProject) (*DeleteProjectResponse, error)

	// DeleteRole was auto-generated from WSDL.
	DeleteRole(DeleteRole *DeleteRole) (*DeleteRoleResponse, error)

	// DeleteSnapshot was auto-generated from WSDL.
	DeleteSnapshot(DeleteSnapshot *DeleteSnapshot) (*DeleteSnapshotResponse, error)

	// DeleteStream was auto-generated from WSDL.
	DeleteStream(DeleteStream *DeleteStream) (*DeleteStreamResponse, error)

	// DeleteTriageStore was auto-generated from WSDL.
	DeleteTriageStore(DeleteTriageStore *DeleteTriageStore) (*DeleteTriageStoreResponse, error)

	// DeleteUser was auto-generated from WSDL.
	DeleteUser(DeleteUser *DeleteUser) (*DeleteUserResponse, error)

	// ExecuteNotification was auto-generated from WSDL.
	ExecuteNotification(ExecuteNotification *ExecuteNotification) (*ExecuteNotificationResponse, error)

	// GetAllLdapConfigurations was auto-generated from WSDL.
	GetAllLdapConfigurations(GetAllLdapConfigurations *GetAllLdapConfigurations) (*GetAllLdapConfigurationsResponse, error)

	// GetAllPermissions was auto-generated from WSDL.
	GetAllPermissions(GetAllPermissions *GetAllPermissions) (*GetAllPermissionsResponse, error)

	// GetAllRoles was auto-generated from WSDL.
	GetAllRoles(GetAllRoles *GetAllRoles) (*GetAllRolesResponse, error)

	// GetArchitectureAnalysisConfiguration was auto-generated from
	// WSDL.
	GetArchitectureAnalysisConfiguration(GetArchitectureAnalysisConfiguration *GetArchitectureAnalysisConfiguration) (*GetArchitectureAnalysisConfigurationResponse, error)

	// GetAttribute was auto-generated from WSDL.
	GetAttribute(GetAttribute *GetAttribute) (*GetAttributeResponse, error)

	// GetAttributes was auto-generated from WSDL.
	GetAttributes(GetAttributes *GetAttributes) (*GetAttributesResponse, error)

	// GetBackupConfiguration was auto-generated from WSDL.
	GetBackupConfiguration(GetBackupConfiguration *GetBackupConfiguration) (*GetBackupConfigurationResponse, error)

	// GetCategoryNames was auto-generated from WSDL.
	GetCategoryNames(GetCategoryNames *GetCategoryNames) (*GetCategoryNamesResponse, error)

	// GetCheckerNames was auto-generated from WSDL.
	GetCheckerNames(GetCheckerNames *GetCheckerNames) (*GetCheckerNamesResponse, error)

	// GetCommitState was auto-generated from WSDL.
	GetCommitState(GetCommitState *GetCommitState) (*GetCommitStateResponse, error)

	// GetComponent was auto-generated from WSDL.
	GetComponent(GetComponent *GetComponent) (*GetComponentResponse, error)

	// GetComponentMaps was auto-generated from WSDL.
	GetComponentMaps(GetComponentMaps *GetComponentMaps) (*GetComponentMapsResponse, error)

	// GetDefectStatuses was auto-generated from WSDL.
	GetDefectStatuses(GetDefectStatuses *GetDefectStatuses) (*GetDefectStatusesResponse, error)

	// GetDeleteSnapshotJobInfo was auto-generated from WSDL.
	GetDeleteSnapshotJobInfo(GetDeleteSnapshotJobInfo *GetDeleteSnapshotJobInfo) (*GetDeleteSnapshotJobInfoResponse, error)

	// GetDeveloperStreamsProjects was auto-generated from WSDL.
	GetDeveloperStreamsProjects(GetDeveloperStreamsProjects *GetDeveloperStreamsProjects) (*GetDeveloperStreamsProjectsResponse, error)

	// GetGroup was auto-generated from WSDL.
	GetGroup(GetGroup *GetGroup) (*GetGroupResponse, error)

	// GetGroups was auto-generated from WSDL.
	GetGroups(GetGroups *GetGroups) (*GetGroupsResponse, error)

	// GetLastUpdateTimes was auto-generated from WSDL.
	GetLastUpdateTimes(GetLastUpdateTimes *GetLastUpdateTimes) (*GetLastUpdateTimesResponse, error)

	// GetLdapServerDomains was auto-generated from WSDL.
	GetLdapServerDomains(GetLdapServerDomains *GetLdapServerDomains) (*GetLdapServerDomainsResponse, error)

	// GetLicenseConfiguration was auto-generated from WSDL.
	GetLicenseConfiguration(GetLicenseConfiguration *GetLicenseConfiguration) (*GetLicenseConfigurationResponse, error)

	// GetLicenseState was auto-generated from WSDL.
	GetLicenseState(GetLicenseState *GetLicenseState) (*GetLicenseStateResponse, error)

	// GetLoggingConfiguration was auto-generated from WSDL.
	GetLoggingConfiguration(GetLoggingConfiguration *GetLoggingConfiguration) (*GetLoggingConfigurationResponse, error)

	// GetMessageOfTheDay was auto-generated from WSDL.
	GetMessageOfTheDay(GetMessageOfTheDay *GetMessageOfTheDay) (*GetMessageOfTheDayResponse, error)

	// GetOutputFileForSnapshot was auto-generated from WSDL.
	GetOutputFileForSnapshot(GetOutputFileForSnapshot *GetOutputFileForSnapshot) (*GetOutputFileForSnapshotResponse, error)

	// GetProjects was auto-generated from WSDL.
	GetProjects(GetProjects *GetProjects) (*GetProjectsResponse, error)

	// GetRole was auto-generated from WSDL.
	GetRole(GetRole *GetRole) (*GetRoleResponse, error)

	// GetServerTime was auto-generated from WSDL.
	GetServerTime(GetServerTime *GetServerTime) (*GetServerTimeResponse, error)

	// GetSignInConfiguration was auto-generated from WSDL.
	GetSignInConfiguration(GetSignInConfiguration *GetSignInConfiguration) (*GetSignInConfigurationResponse, error)

	// GetSkeletonizationConfiguration was auto-generated from WSDL.
	GetSkeletonizationConfiguration(GetSkeletonizationConfiguration *GetSkeletonizationConfiguration) (*GetSkeletonizationConfigurationResponse, error)

	// GetSnapshotInformation was auto-generated from WSDL.
	GetSnapshotInformation(GetSnapshotInformation *GetSnapshotInformation) (*GetSnapshotInformationResponse, error)

	// GetSnapshotPurgeDetails was auto-generated from WSDL.
	GetSnapshotPurgeDetails(GetSnapshotPurgeDetails *GetSnapshotPurgeDetails) (*GetSnapshotPurgeDetailsResponse, error)

	// GetSnapshotsForStream was auto-generated from WSDL.
	GetSnapshotsForStream(GetSnapshotsForStream *GetSnapshotsForStream) (*GetSnapshotsForStreamResponse, error)

	// GetStreams was auto-generated from WSDL.
	GetStreams(GetStreams *GetStreams) (*GetStreamsResponse, error)

	// GetSystemConfig was auto-generated from WSDL.
	GetSystemConfig(GetSystemConfig *GetSystemConfig) (*GetSystemConfigResponse, error)

	// GetTriageStores was auto-generated from WSDL.
	GetTriageStores(GetTriageStores *GetTriageStores) (*GetTriageStoresResponse, error)

	// GetTypeNames was auto-generated from WSDL.
	GetTypeNames(GetTypeNames *GetTypeNames) (*GetTypeNamesResponse, error)

	// GetUser was auto-generated from WSDL.
	GetUser(GetUser *GetUser) (*GetUserResponse, error)

	// GetUsers was auto-generated from WSDL.
	GetUsers(GetUsers *GetUsers) (*GetUsersResponse, error)

	// GetVersion was auto-generated from WSDL.
	GetVersion(GetVersion *GetVersion) (*GetVersionResponse, error)

	// ImportLicense was auto-generated from WSDL.
	ImportLicense(ImportLicense *ImportLicense) (*ImportLicenseResponse, error)

	// MergeTriageStores was auto-generated from WSDL.
	MergeTriageStores(MergeTriageStores *MergeTriageStores) (*MergeTriageStoresResponse, error)

	// Notify was auto-generated from WSDL.
	Notify(Notify *Notify) (*NotifyResponse, error)

	// SetAcceptingNewCommits was auto-generated from WSDL.
	SetAcceptingNewCommits(SetAcceptingNewCommits *SetAcceptingNewCommits) (*SetAcceptingNewCommitsResponse, error)

	// SetArchitectureAnalysisConfiguration was auto-generated from
	// WSDL.
	SetArchitectureAnalysisConfiguration(SetArchitectureAnalysisConfiguration *SetArchitectureAnalysisConfiguration) (*SetArchitectureAnalysisConfigurationResponse, error)

	// SetBackupConfiguration was auto-generated from WSDL.
	SetBackupConfiguration(SetBackupConfiguration *SetBackupConfiguration) (*SetBackupConfigurationResponse, error)

	// SetLoggingConfiguration was auto-generated from WSDL.
	SetLoggingConfiguration(SetLoggingConfiguration *SetLoggingConfiguration) (*SetLoggingConfigurationResponse, error)

	// SetMessageOfTheDay was auto-generated from WSDL.
	SetMessageOfTheDay(SetMessageOfTheDay *SetMessageOfTheDay) (*SetMessageOfTheDayResponse, error)

	// SetSkeletonizationConfiguration was auto-generated from WSDL.
	SetSkeletonizationConfiguration(SetSkeletonizationConfiguration *SetSkeletonizationConfiguration) (*SetSkeletonizationConfigurationResponse, error)

	// SetSnapshotPurgeDetails was auto-generated from WSDL.
	SetSnapshotPurgeDetails(SetSnapshotPurgeDetails *SetSnapshotPurgeDetails) (*SetSnapshotPurgeDetailsResponse, error)

	// UpdateAttribute was auto-generated from WSDL.
	UpdateAttribute(UpdateAttribute *UpdateAttribute) (*UpdateAttributeResponse, error)

	// UpdateComponentMap was auto-generated from WSDL.
	UpdateComponentMap(UpdateComponentMap *UpdateComponentMap) (*UpdateComponentMapResponse, error)

	// UpdateGroup was auto-generated from WSDL.
	UpdateGroup(UpdateGroup *UpdateGroup) (*UpdateGroupResponse, error)

	// UpdateLdapConfiguration was auto-generated from WSDL.
	UpdateLdapConfiguration(UpdateLdapConfiguration *UpdateLdapConfiguration) (*UpdateLdapConfigurationResponse, error)

	// UpdateProject was auto-generated from WSDL.
	UpdateProject(UpdateProject *UpdateProject) (*UpdateProjectResponse, error)

	// UpdateRole was auto-generated from WSDL.
	UpdateRole(UpdateRole *UpdateRole) (*UpdateRoleResponse, error)

	// UpdateSignInConfiguration was auto-generated from WSDL.
	UpdateSignInConfiguration(UpdateSignInConfiguration *UpdateSignInConfiguration) (*UpdateSignInConfigurationResponse, error)

	// UpdateSnapshotInfo was auto-generated from WSDL.
	UpdateSnapshotInfo(UpdateSnapshotInfo *UpdateSnapshotInfo) (*UpdateSnapshotInfoResponse, error)

	// UpdateStream was auto-generated from WSDL.
	UpdateStream(UpdateStream *UpdateStream) (*UpdateStreamResponse, error)

	// UpdateTriageStore was auto-generated from WSDL.
	UpdateTriageStore(UpdateTriageStore *UpdateTriageStore) (*UpdateTriageStoreResponse, error)

	// UpdateUser was auto-generated from WSDL.
	UpdateUser(UpdateUser *UpdateUser) (*UpdateUserResponse, error)
}

// DateTime in WSDL format.
type DateTime string

// DeleteSnapshotJobStatus was auto-generated from WSDL.
type DeleteSnapshotJobStatus string

// Validate validates DeleteSnapshotJobStatus.
func (v DeleteSnapshotJobStatus) Validate() bool {
	for _, vv := range []string{
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CovRemoteServiceException was auto-generated from WSDL.
type CovRemoteServiceException struct {
	ErrorCode *int    `xml:"errorCode,omitempty" json:"errorCode,omitempty" yaml:"errorCode,omitempty"`
	Message   *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// AttributeDefinitionDataObj was auto-generated from WSDL.
type AttributeDefinitionDataObj struct {
	AttributeDefinitionId *AttributeDefinitionIdDataObj `xml:"attributeDefinitionId,omitempty" json:"attributeDefinitionId,omitempty" yaml:"attributeDefinitionId,omitempty"`
	AttributeType         *string                       `xml:"attributeType,omitempty" json:"attributeType,omitempty" yaml:"attributeType,omitempty"`
	BuiltIn               *bool                         `xml:"builtIn,omitempty" json:"builtIn,omitempty" yaml:"builtIn,omitempty"`
	ConfigurableValues    []*AttributeValueDataObj      `xml:"configurableValues,omitempty" json:"configurableValues,omitempty" yaml:"configurableValues,omitempty"`
	DefaultValue          *string                       `xml:"defaultValue,omitempty" json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
	Description           *string                       `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	DisplayDescription    *string                       `xml:"displayDescription,omitempty" json:"displayDescription,omitempty" yaml:"displayDescription,omitempty"`
	DisplayName           *string                       `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	ShowInTriage          *bool                         `xml:"showInTriage,omitempty" json:"showInTriage,omitempty" yaml:"showInTriage,omitempty"`
}

// AttributeDefinitionIdDataObj was auto-generated from WSDL.
type AttributeDefinitionIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// AttributeDefinitionSpecDataObj was auto-generated from WSDL.
type AttributeDefinitionSpecDataObj struct {
	AttributeName            *string                          `xml:"attributeName,omitempty" json:"attributeName,omitempty" yaml:"attributeName,omitempty"`
	AttributeType            *string                          `xml:"attributeType,omitempty" json:"attributeType,omitempty" yaml:"attributeType,omitempty"`
	AttributeValueChangeSpec *AttributeValueChangeSpecDataObj `xml:"attributeValueChangeSpec,omitempty" json:"attributeValueChangeSpec,omitempty" yaml:"attributeValueChangeSpec,omitempty"`
	DefaultValue             *string                          `xml:"defaultValue,omitempty" json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
	Description              *string                          `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	ShowInTriage             *bool                            `xml:"showInTriage,omitempty" json:"showInTriage,omitempty" yaml:"showInTriage,omitempty"`
}

// AttributeValueChangeSpecDataObj was auto-generated from WSDL.
type AttributeValueChangeSpecDataObj struct {
	AttributeValueIds []*AttributeValueIdDataObj   `xml:"attributeValueIds,omitempty" json:"attributeValueIds,omitempty" yaml:"attributeValueIds,omitempty"`
	AttributeValues   []*AttributeValueSpecDataObj `xml:"attributeValues,omitempty" json:"attributeValues,omitempty" yaml:"attributeValues,omitempty"`
}

// AttributeValueDataObj was auto-generated from WSDL.
type AttributeValueDataObj struct {
	AttributeValueId *AttributeValueIdDataObj `xml:"attributeValueId,omitempty" json:"attributeValueId,omitempty" yaml:"attributeValueId,omitempty"`
	Deprecated       *bool                    `xml:"deprecated,omitempty" json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	DisplayName      *string                  `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	IssueKindList    []*string                `xml:"issueKindList,omitempty" json:"issueKindList,omitempty" yaml:"issueKindList,omitempty"`
}

// AttributeValueIdDataObj was auto-generated from WSDL.
type AttributeValueIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// AttributeValueSpecDataObj was auto-generated from WSDL.
type AttributeValueSpecDataObj struct {
	Deprecated *bool   `xml:"deprecated,omitempty" json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Name       *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// BackupConfigurationDataObj was auto-generated from WSDL.
type BackupConfigurationDataObj struct {
	BackupLocation   *string `xml:"backupLocation,omitempty" json:"backupLocation,omitempty" yaml:"backupLocation,omitempty"`
	BackupTime       *string `xml:"backupTime,omitempty" json:"backupTime,omitempty" yaml:"backupTime,omitempty"`
	FridayEnabled    *bool   `xml:"fridayEnabled,omitempty" json:"fridayEnabled,omitempty" yaml:"fridayEnabled,omitempty"`
	MondayEnabled    *bool   `xml:"mondayEnabled,omitempty" json:"mondayEnabled,omitempty" yaml:"mondayEnabled,omitempty"`
	SaturdayEnabled  *bool   `xml:"saturdayEnabled,omitempty" json:"saturdayEnabled,omitempty" yaml:"saturdayEnabled,omitempty"`
	SundayEnabled    *bool   `xml:"sundayEnabled,omitempty" json:"sundayEnabled,omitempty" yaml:"sundayEnabled,omitempty"`
	ThursdayEnabled  *bool   `xml:"thursdayEnabled,omitempty" json:"thursdayEnabled,omitempty" yaml:"thursdayEnabled,omitempty"`
	TuesdayEnabled   *bool   `xml:"tuesdayEnabled,omitempty" json:"tuesdayEnabled,omitempty" yaml:"tuesdayEnabled,omitempty"`
	WednesdayEnabled *bool   `xml:"wednesdayEnabled,omitempty" json:"wednesdayEnabled,omitempty" yaml:"wednesdayEnabled,omitempty"`
}

// CommitStateDataObj was auto-generated from WSDL.
type CommitStateDataObj struct {
	CurrentCommitCount    *int  `xml:"currentCommitCount,omitempty" json:"currentCommitCount,omitempty" yaml:"currentCommitCount,omitempty"`
	IsAcceptingNewCommits *bool `xml:"isAcceptingNewCommits,omitempty" json:"isAcceptingNewCommits,omitempty" yaml:"isAcceptingNewCommits,omitempty"`
}

// ComponentDataObj was auto-generated from WSDL.
type ComponentDataObj struct {
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	ComponentId     *ComponentIdDataObj      `xml:"componentId,omitempty" json:"componentId,omitempty" yaml:"componentId,omitempty"`
	Subscribers     []*string                `xml:"subscribers,omitempty" json:"subscribers,omitempty" yaml:"subscribers,omitempty"`
}

// ComponentDefectRuleDataObj was auto-generated from WSDL.
type ComponentDefectRuleDataObj struct {
	ComponentId  *ComponentIdDataObj `xml:"componentId,omitempty" json:"componentId,omitempty" yaml:"componentId,omitempty"`
	DefaultOwner *string             `xml:"defaultOwner,omitempty" json:"defaultOwner,omitempty" yaml:"defaultOwner,omitempty"`
}

// ComponentIdDataObj was auto-generated from WSDL.
type ComponentIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// ComponentMapDataObj was auto-generated from WSDL.
type ComponentMapDataObj struct {
	ComponentMapId     *ComponentMapIdDataObj        `xml:"componentMapId,omitempty" json:"componentMapId,omitempty" yaml:"componentMapId,omitempty"`
	ComponentPathRules []*ComponentPathRuleDataObj   `xml:"componentPathRules,omitempty" json:"componentPathRules,omitempty" yaml:"componentPathRules,omitempty"`
	Components         []*ComponentDataObj           `xml:"components,omitempty" json:"components,omitempty" yaml:"components,omitempty"`
	DefectRules        []*ComponentDefectRuleDataObj `xml:"defectRules,omitempty" json:"defectRules,omitempty" yaml:"defectRules,omitempty"`
	Description        *string                       `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// ComponentMapFilterSpecDataObj was auto-generated from WSDL.
type ComponentMapFilterSpecDataObj struct {
	NamePattern *string `xml:"namePattern,omitempty" json:"namePattern,omitempty" yaml:"namePattern,omitempty"`
}

// ComponentMapIdDataObj was auto-generated from WSDL.
type ComponentMapIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// ComponentMapSpecDataObj was auto-generated from WSDL.
type ComponentMapSpecDataObj struct {
	ComponentMapName      *string                       `xml:"componentMapName,omitempty" json:"componentMapName,omitempty" yaml:"componentMapName,omitempty"`
	ComponentPathRules    []*ComponentPathRuleDataObj   `xml:"componentPathRules,omitempty" json:"componentPathRules,omitempty" yaml:"componentPathRules,omitempty"`
	Components            []*ComponentDataObj           `xml:"components,omitempty" json:"components,omitempty" yaml:"components,omitempty"`
	DefectRules           []*ComponentDefectRuleDataObj `xml:"defectRules,omitempty" json:"defectRules,omitempty" yaml:"defectRules,omitempty"`
	Description           *string                       `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	ForceDeleteComponents *bool                         `xml:"forceDeleteComponents,omitempty" json:"forceDeleteComponents,omitempty" yaml:"forceDeleteComponents,omitempty"`
}

// ComponentPathRuleDataObj was auto-generated from WSDL.
type ComponentPathRuleDataObj struct {
	ComponentId *ComponentIdDataObj `xml:"componentId,omitempty" json:"componentId,omitempty" yaml:"componentId,omitempty"`
	PathPattern *string             `xml:"pathPattern,omitempty" json:"pathPattern,omitempty" yaml:"pathPattern,omitempty"`
}

// ConfigurationDataObj was auto-generated from WSDL.
type ConfigurationDataObj struct {
	CommitPort     *int64  `xml:"commitPort,omitempty" json:"commitPort,omitempty" yaml:"commitPort,omitempty"`
	DbDialect      *string `xml:"dbDialect,omitempty" json:"dbDialect,omitempty" yaml:"dbDialect,omitempty"`
	DbDriver       *string `xml:"dbDriver,omitempty" json:"dbDriver,omitempty" yaml:"dbDriver,omitempty"`
	IssueExportUrl *string `xml:"issueExportUrl,omitempty" json:"issueExportUrl,omitempty" yaml:"issueExportUrl,omitempty"`
	MaindbName     *string `xml:"maindbName,omitempty" json:"maindbName,omitempty" yaml:"maindbName,omitempty"`
	MaindbUrl      *string `xml:"maindbUrl,omitempty" json:"maindbUrl,omitempty" yaml:"maindbUrl,omitempty"`
	MaindbUser     *string `xml:"maindbUser,omitempty" json:"maindbUser,omitempty" yaml:"maindbUser,omitempty"`
}

// CopyStream was auto-generated from WSDL.
type CopyStream struct {
	ProjectId      *ProjectIdDataObj `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	SourceStreamId *StreamIdDataObj  `xml:"sourceStreamId,omitempty" json:"sourceStreamId,omitempty" yaml:"sourceStreamId,omitempty"`
}

// CopyStreamResponse was auto-generated from WSDL.
type CopyStreamResponse struct {
	Return *StreamDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// CreateAttribute was auto-generated from WSDL.
type CreateAttribute struct {
	AttributeDefinitionSpec *AttributeDefinitionSpecDataObj `xml:"attributeDefinitionSpec,omitempty" json:"attributeDefinitionSpec,omitempty" yaml:"attributeDefinitionSpec,omitempty"`
}

// CreateAttributeResponse was auto-generated from WSDL.
type CreateAttributeResponse struct {
}

// CreateComponentMap was auto-generated from WSDL.
type CreateComponentMap struct {
	ComponentMapSpec *ComponentMapSpecDataObj `xml:"componentMapSpec,omitempty" json:"componentMapSpec,omitempty" yaml:"componentMapSpec,omitempty"`
}

// CreateComponentMapResponse was auto-generated from WSDL.
type CreateComponentMapResponse struct {
}

// CreateGroup was auto-generated from WSDL.
type CreateGroup struct {
	GroupSpec *GroupSpecDataObj `xml:"groupSpec,omitempty" json:"groupSpec,omitempty" yaml:"groupSpec,omitempty"`
}

// CreateGroupResponse was auto-generated from WSDL.
type CreateGroupResponse struct {
}

// CreateLdapConfiguration was auto-generated from WSDL.
type CreateLdapConfiguration struct {
	LdapConfigurationSpec *LdapConfigurationSpecDataObj `xml:"ldapConfigurationSpec,omitempty" json:"ldapConfigurationSpec,omitempty" yaml:"ldapConfigurationSpec,omitempty"`
}

// CreateLdapConfigurationResponse was auto-generated from WSDL.
type CreateLdapConfigurationResponse struct {
}

// CreateProject was auto-generated from WSDL.
type CreateProject struct {
	ProjectSpec *ProjectSpecDataObj `xml:"projectSpec,omitempty" json:"projectSpec,omitempty" yaml:"projectSpec,omitempty"`
}

// CreateProjectResponse was auto-generated from WSDL.
type CreateProjectResponse struct {
}

// CreateRole was auto-generated from WSDL.
type CreateRole struct {
	RoleSpec *RoleSpecDataObj `xml:"roleSpec,omitempty" json:"roleSpec,omitempty" yaml:"roleSpec,omitempty"`
}

// CreateRoleResponse was auto-generated from WSDL.
type CreateRoleResponse struct {
}

// CreateStream was auto-generated from WSDL.
type CreateStream struct {
	StreamSpec *StreamSpecDataObj `xml:"streamSpec,omitempty" json:"streamSpec,omitempty" yaml:"streamSpec,omitempty"`
}

// CreateStreamInProject was auto-generated from WSDL.
type CreateStreamInProject struct {
	ProjectId  *ProjectIdDataObj  `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	StreamSpec *StreamSpecDataObj `xml:"streamSpec,omitempty" json:"streamSpec,omitempty" yaml:"streamSpec,omitempty"`
}

// CreateStreamInProjectResponse was auto-generated from WSDL.
type CreateStreamInProjectResponse struct {
}

// CreateStreamResponse was auto-generated from WSDL.
type CreateStreamResponse struct {
}

// CreateTriageStore was auto-generated from WSDL.
type CreateTriageStore struct {
	TriageStoreSpec *TriageStoreSpecDataObj `xml:"triageStoreSpec,omitempty" json:"triageStoreSpec,omitempty" yaml:"triageStoreSpec,omitempty"`
}

// CreateTriageStoreResponse was auto-generated from WSDL.
type CreateTriageStoreResponse struct {
}

// CreateUser was auto-generated from WSDL.
type CreateUser struct {
	UserSpec *UserSpecDataObj `xml:"userSpec,omitempty" json:"userSpec,omitempty" yaml:"userSpec,omitempty"`
}

// CreateUserResponse was auto-generated from WSDL.
type CreateUserResponse struct {
}

// DeleteAttribute was auto-generated from WSDL.
type DeleteAttribute struct {
	AttributeDefinitionId *AttributeDefinitionIdDataObj `xml:"attributeDefinitionId,omitempty" json:"attributeDefinitionId,omitempty" yaml:"attributeDefinitionId,omitempty"`
}

// DeleteAttributeResponse was auto-generated from WSDL.
type DeleteAttributeResponse struct {
}

// DeleteComponentMap was auto-generated from WSDL.
type DeleteComponentMap struct {
	ComponentMapId *ComponentMapIdDataObj `xml:"componentMapId,omitempty" json:"componentMapId,omitempty" yaml:"componentMapId,omitempty"`
}

// DeleteComponentMapResponse was auto-generated from WSDL.
type DeleteComponentMapResponse struct {
}

// DeleteGroup was auto-generated from WSDL.
type DeleteGroup struct {
	GroupId *GroupIdDataObj `xml:"groupId,omitempty" json:"groupId,omitempty" yaml:"groupId,omitempty"`
}

// DeleteGroupResponse was auto-generated from WSDL.
type DeleteGroupResponse struct {
}

// DeleteLdapConfiguration was auto-generated from WSDL.
type DeleteLdapConfiguration struct {
	ServerDomainIdDataObj *ServerDomainIdDataObj `xml:"serverDomainIdDataObj,omitempty" json:"serverDomainIdDataObj,omitempty" yaml:"serverDomainIdDataObj,omitempty"`
}

// DeleteLdapConfigurationResponse was auto-generated from WSDL.
type DeleteLdapConfigurationResponse struct {
}

// DeleteProject was auto-generated from WSDL.
type DeleteProject struct {
	ProjectId *ProjectIdDataObj `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
}

// DeleteProjectResponse was auto-generated from WSDL.
type DeleteProjectResponse struct {
}

// DeleteRole was auto-generated from WSDL.
type DeleteRole struct {
	RoleId *RoleIdDataObj `xml:"roleId,omitempty" json:"roleId,omitempty" yaml:"roleId,omitempty"`
}

// DeleteRoleResponse was auto-generated from WSDL.
type DeleteRoleResponse struct {
}

// DeleteSnapshot was auto-generated from WSDL.
type DeleteSnapshot struct {
	SnapshotId []*SnapshotIdDataObj `xml:"snapshotId,omitempty" json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
}

// DeleteSnapshotJobInfoDataObj was auto-generated from WSDL.
type DeleteSnapshotJobInfoDataObj struct {
	SnapshotId *int64                   `xml:"snapshotId,omitempty" json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
	Status     *DeleteSnapshotJobStatus `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
}

// DeleteSnapshotResponse was auto-generated from WSDL.
type DeleteSnapshotResponse struct {
}

// DeleteStream was auto-generated from WSDL.
type DeleteStream struct {
	StreamId    *StreamIdDataObj `xml:"streamId,omitempty" json:"streamId,omitempty" yaml:"streamId,omitempty"`
	OnlyIfEmpty *bool            `xml:"onlyIfEmpty,omitempty" json:"onlyIfEmpty,omitempty" yaml:"onlyIfEmpty,omitempty"`
}

// DeleteStreamResponse was auto-generated from WSDL.
type DeleteStreamResponse struct {
}

// DeleteTriageStore was auto-generated from WSDL.
type DeleteTriageStore struct {
	TriageStoreId *TriageStoreIdDataObj `xml:"triageStoreId,omitempty" json:"triageStoreId,omitempty" yaml:"triageStoreId,omitempty"`
}

// DeleteTriageStoreResponse was auto-generated from WSDL.
type DeleteTriageStoreResponse struct {
}

// DeleteUser was auto-generated from WSDL.
type DeleteUser struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
}

// DeleteUserResponse was auto-generated from WSDL.
type DeleteUserResponse struct {
}

// ExecuteNotification was auto-generated from WSDL.
type ExecuteNotification struct {
	Viewname *string `xml:"viewname,omitempty" json:"viewname,omitempty" yaml:"viewname,omitempty"`
}

// ExecuteNotificationResponse was auto-generated from WSDL.
type ExecuteNotificationResponse struct {
}

// FeatureUpdateTimeDataObj was auto-generated from WSDL.
type FeatureUpdateTimeDataObj struct {
	FeatureName    *string   `xml:"featureName,omitempty" json:"featureName,omitempty" yaml:"featureName,omitempty"`
	LastUpdateDate *DateTime `xml:"lastUpdateDate,omitempty" json:"lastUpdateDate,omitempty" yaml:"lastUpdateDate,omitempty"`
}

// GetAllLdapConfigurations was auto-generated from WSDL.
type GetAllLdapConfigurations struct {
}

// GetAllLdapConfigurationsResponse was auto-generated from WSDL.
type GetAllLdapConfigurationsResponse struct {
	Return []*LdapConfigurationDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAllPermissions was auto-generated from WSDL.
type GetAllPermissions struct {
}

// GetAllPermissionsResponse was auto-generated from WSDL.
type GetAllPermissionsResponse struct {
	Return []*PermissionDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAllRoles was auto-generated from WSDL.
type GetAllRoles struct {
}

// GetAllRolesResponse was auto-generated from WSDL.
type GetAllRolesResponse struct {
	Return []*RoleDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetArchitectureAnalysisConfiguration was auto-generated from
// WSDL.
type GetArchitectureAnalysisConfiguration struct {
}

// GetArchitectureAnalysisConfigurationResponse was auto-generated
// from WSDL.
type GetArchitectureAnalysisConfigurationResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAttribute was auto-generated from WSDL.
type GetAttribute struct {
	AttributeDefinitionId *AttributeDefinitionIdDataObj `xml:"attributeDefinitionId,omitempty" json:"attributeDefinitionId,omitempty" yaml:"attributeDefinitionId,omitempty"`
}

// GetAttributeResponse was auto-generated from WSDL.
type GetAttributeResponse struct {
	Return *AttributeDefinitionDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAttributes was auto-generated from WSDL.
type GetAttributes struct {
}

// GetAttributesResponse was auto-generated from WSDL.
type GetAttributesResponse struct {
	Return []*AttributeDefinitionDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetBackupConfiguration was auto-generated from WSDL.
type GetBackupConfiguration struct {
}

// GetBackupConfigurationResponse was auto-generated from WSDL.
type GetBackupConfigurationResponse struct {
	Return *BackupConfigurationDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetCategoryNames was auto-generated from WSDL.
type GetCategoryNames struct {
}

// GetCategoryNamesResponse was auto-generated from WSDL.
type GetCategoryNamesResponse struct {
	Return []*LocalizedValueDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetCheckerNames was auto-generated from WSDL.
type GetCheckerNames struct {
}

// GetCheckerNamesResponse was auto-generated from WSDL.
type GetCheckerNamesResponse struct {
	Return []*string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetCommitState was auto-generated from WSDL.
type GetCommitState struct {
}

// GetCommitStateResponse was auto-generated from WSDL.
type GetCommitStateResponse struct {
	Return *CommitStateDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetComponent was auto-generated from WSDL.
type GetComponent struct {
	ComponentId *ComponentIdDataObj `xml:"componentId,omitempty" json:"componentId,omitempty" yaml:"componentId,omitempty"`
}

// GetComponentMaps was auto-generated from WSDL.
type GetComponentMaps struct {
	FilterSpec *ComponentMapFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetComponentMapsResponse was auto-generated from WSDL.
type GetComponentMapsResponse struct {
	Return []*ComponentMapDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetComponentResponse was auto-generated from WSDL.
type GetComponentResponse struct {
	Return *ComponentDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDefectStatuses was auto-generated from WSDL.
type GetDefectStatuses struct {
}

// GetDefectStatusesResponse was auto-generated from WSDL.
type GetDefectStatusesResponse struct {
	Return []*string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDeleteSnapshotJobInfo was auto-generated from WSDL.
type GetDeleteSnapshotJobInfo struct {
	SnapshotId []*SnapshotIdDataObj `xml:"snapshotId,omitempty" json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
}

// GetDeleteSnapshotJobInfoResponse was auto-generated from WSDL.
type GetDeleteSnapshotJobInfoResponse struct {
	Return []*DeleteSnapshotJobInfoDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDeveloperStreamsProjects was auto-generated from WSDL.
type GetDeveloperStreamsProjects struct {
	FilterSpec *ProjectFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetDeveloperStreamsProjectsResponse was auto-generated from
// WSDL.
type GetDeveloperStreamsProjectsResponse struct {
	Return []*ProjectDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetGroup was auto-generated from WSDL.
type GetGroup struct {
	GroupId *GroupIdDataObj `xml:"groupId,omitempty" json:"groupId,omitempty" yaml:"groupId,omitempty"`
}

// GetGroupResponse was auto-generated from WSDL.
type GetGroupResponse struct {
	Return *GroupDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetGroups was auto-generated from WSDL.
type GetGroups struct {
	FilterSpec *GroupFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
	PageSpec   *PageSpecDataObj        `xml:"pageSpec,omitempty" json:"pageSpec,omitempty" yaml:"pageSpec,omitempty"`
}

// GetGroupsResponse was auto-generated from WSDL.
type GetGroupsResponse struct {
	Return *GroupsPageDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetLastUpdateTimes was auto-generated from WSDL.
type GetLastUpdateTimes struct {
}

// GetLastUpdateTimesResponse was auto-generated from WSDL.
type GetLastUpdateTimesResponse struct {
	Return []*FeatureUpdateTimeDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetLdapServerDomains was auto-generated from WSDL.
type GetLdapServerDomains struct {
}

// GetLdapServerDomainsResponse was auto-generated from WSDL.
type GetLdapServerDomainsResponse struct {
	Return []*ServerDomainIdDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetLicenseConfiguration was auto-generated from WSDL.
type GetLicenseConfiguration struct {
}

// GetLicenseConfigurationResponse was auto-generated from WSDL.
type GetLicenseConfigurationResponse struct {
	Return *LicenseDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetLicenseState was auto-generated from WSDL.
type GetLicenseState struct {
}

// GetLicenseStateResponse was auto-generated from WSDL.
type GetLicenseStateResponse struct {
	Return *LicenseStateDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetLoggingConfiguration was auto-generated from WSDL.
type GetLoggingConfiguration struct {
}

// GetLoggingConfigurationResponse was auto-generated from WSDL.
type GetLoggingConfigurationResponse struct {
	Return *LoggingConfigurationDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMessageOfTheDay was auto-generated from WSDL.
type GetMessageOfTheDay struct {
}

// GetMessageOfTheDayResponse was auto-generated from WSDL.
type GetMessageOfTheDayResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetOutputFileForSnapshot was auto-generated from WSDL.
type GetOutputFileForSnapshot struct {
	SnapshotId *SnapshotIdDataObj `xml:"snapshotId,omitempty" json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
	FileName   *string            `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
}

// GetOutputFileForSnapshotResponse was auto-generated from WSDL.
type GetOutputFileForSnapshotResponse struct {
	Return *OutputFileDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetProjects was auto-generated from WSDL.
type GetProjects struct {
	FilterSpec *ProjectFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetProjectsResponse was auto-generated from WSDL.
type GetProjectsResponse struct {
	Return []*ProjectDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRole was auto-generated from WSDL.
type GetRole struct {
	RoleId *RoleIdDataObj `xml:"roleId,omitempty" json:"roleId,omitempty" yaml:"roleId,omitempty"`
}

// GetRoleResponse was auto-generated from WSDL.
type GetRoleResponse struct {
	Return *RoleDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetServerTime was auto-generated from WSDL.
type GetServerTime struct {
}

// GetServerTimeResponse was auto-generated from WSDL.
type GetServerTimeResponse struct {
	Return *DateTime `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetSignInConfiguration was auto-generated from WSDL.
type GetSignInConfiguration struct {
}

// GetSignInConfigurationResponse was auto-generated from WSDL.
type GetSignInConfigurationResponse struct {
	Return *SignInSettingsDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetSkeletonizationConfiguration was auto-generated from WSDL.
type GetSkeletonizationConfiguration struct {
}

// GetSkeletonizationConfigurationResponse was auto-generated from
// WSDL.
type GetSkeletonizationConfigurationResponse struct {
	Return *SkeletonizationConfigurationDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetSnapshotInformation was auto-generated from WSDL.
type GetSnapshotInformation struct {
	SnapshotIds []*SnapshotIdDataObj `xml:"snapshotIds,omitempty" json:"snapshotIds,omitempty" yaml:"snapshotIds,omitempty"`
}

// GetSnapshotInformationResponse was auto-generated from WSDL.
type GetSnapshotInformationResponse struct {
	Return []*SnapshotInfoDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetSnapshotPurgeDetails was auto-generated from WSDL.
type GetSnapshotPurgeDetails struct {
}

// GetSnapshotPurgeDetailsResponse was auto-generated from WSDL.
type GetSnapshotPurgeDetailsResponse struct {
	Return *SnapshotPurgeDetailsObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetSnapshotsForStream was auto-generated from WSDL.
type GetSnapshotsForStream struct {
	StreamId   *StreamIdDataObj           `xml:"streamId,omitempty" json:"streamId,omitempty" yaml:"streamId,omitempty"`
	FilterSpec *SnapshotFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetSnapshotsForStreamResponse was auto-generated from WSDL.
type GetSnapshotsForStreamResponse struct {
	Return []*SnapshotIdDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetStreams was auto-generated from WSDL.
type GetStreams struct {
	FilterSpec *StreamFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetStreamsResponse was auto-generated from WSDL.
type GetStreamsResponse struct {
	Return []*StreamDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetSystemConfig was auto-generated from WSDL.
type GetSystemConfig struct {
}

// GetSystemConfigResponse was auto-generated from WSDL.
type GetSystemConfigResponse struct {
	Return *ConfigurationDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetTriageStores was auto-generated from WSDL.
type GetTriageStores struct {
	FilterSpec *TriageStoreFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetTriageStoresResponse was auto-generated from WSDL.
type GetTriageStoresResponse struct {
	Return []*TriageStoreDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetTypeNames was auto-generated from WSDL.
type GetTypeNames struct {
}

// GetTypeNamesResponse was auto-generated from WSDL.
type GetTypeNamesResponse struct {
	Return []*LocalizedValueDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetUser was auto-generated from WSDL.
type GetUser struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUserResponse was auto-generated from WSDL.
type GetUserResponse struct {
	Return *UserDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetUsers was auto-generated from WSDL.
type GetUsers struct {
	FilterSpec *UserFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
	PageSpec   *PageSpecDataObj       `xml:"pageSpec,omitempty" json:"pageSpec,omitempty" yaml:"pageSpec,omitempty"`
}

// GetUsersResponse was auto-generated from WSDL.
type GetUsersResponse struct {
	Return *UsersPageDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetVersion was auto-generated from WSDL.
type GetVersion struct {
}

// GetVersionResponse was auto-generated from WSDL.
type GetVersionResponse struct {
	Return *VersionDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GroupDataObj was auto-generated from WSDL.
type GroupDataObj struct {
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	Local           *bool                    `xml:"local,omitempty" json:"local,omitempty" yaml:"local,omitempty"`
	Name            *GroupIdDataObj          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	SyncEnabled     *bool                    `xml:"syncEnabled,omitempty" json:"syncEnabled,omitempty" yaml:"syncEnabled,omitempty"`
}

// GroupFilterSpecDataObj was auto-generated from WSDL.
type GroupFilterSpecDataObj struct {
	Ldap             *bool             `xml:"ldap,omitempty" json:"ldap,omitempty" yaml:"ldap,omitempty"`
	NamePattern      *string           `xml:"namePattern,omitempty" json:"namePattern,omitempty" yaml:"namePattern,omitempty"`
	ProjectIdDataObj *ProjectIdDataObj `xml:"projectIdDataObj,omitempty" json:"projectIdDataObj,omitempty" yaml:"projectIdDataObj,omitempty"`
	UserList         []*string         `xml:"userList,omitempty" json:"userList,omitempty" yaml:"userList,omitempty"`
}

// GroupIdDataObj was auto-generated from WSDL.
type GroupIdDataObj struct {
	DisplayName *string                `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Domain      *ServerDomainIdDataObj `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	Name        *string                `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// GroupSpecDataObj was auto-generated from WSDL.
type GroupSpecDataObj struct {
	ClearRoles      *bool                    `xml:"clearRoles,omitempty" json:"clearRoles,omitempty" yaml:"clearRoles,omitempty"`
	Domain          *ServerDomainIdDataObj   `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	Local           *bool                    `xml:"local,omitempty" json:"local,omitempty" yaml:"local,omitempty"`
	Name            *string                  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	SyncEnabled     *bool                    `xml:"syncEnabled,omitempty" json:"syncEnabled,omitempty" yaml:"syncEnabled,omitempty"`
}

// GroupsPageDataObj was auto-generated from WSDL.
type GroupsPageDataObj struct {
	Groups               []*GroupDataObj `xml:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	TotalNumberOfRecords *int            `xml:"totalNumberOfRecords,omitempty" json:"totalNumberOfRecords,omitempty" yaml:"totalNumberOfRecords,omitempty"`
}

// ImportLicense was auto-generated from WSDL.
type ImportLicense struct {
	LicenseSpecDataObj *LicenseSpecDataObj `xml:"licenseSpecDataObj,omitempty" json:"licenseSpecDataObj,omitempty" yaml:"licenseSpecDataObj,omitempty"`
}

// ImportLicenseResponse was auto-generated from WSDL.
type ImportLicenseResponse struct {
}

// LdapConfigurationDataObj was auto-generated from WSDL.
type LdapConfigurationDataObj struct {
	AnonymousBind         *bool                  `xml:"anonymousBind,omitempty" json:"anonymousBind,omitempty" yaml:"anonymousBind,omitempty"`
	BaseDN                *string                `xml:"baseDN,omitempty" json:"baseDN,omitempty" yaml:"baseDN,omitempty"`
	BindName              *string                `xml:"bindName,omitempty" json:"bindName,omitempty" yaml:"bindName,omitempty"`
	BindPassword          *string                `xml:"bindPassword,omitempty" json:"bindPassword,omitempty" yaml:"bindPassword,omitempty"`
	GroupFilter           *string                `xml:"groupFilter,omitempty" json:"groupFilter,omitempty" yaml:"groupFilter,omitempty"`
	GroupFullName         *bool                  `xml:"groupFullName,omitempty" json:"groupFullName,omitempty" yaml:"groupFullName,omitempty"`
	GroupMember           *string                `xml:"groupMember,omitempty" json:"groupMember,omitempty" yaml:"groupMember,omitempty"`
	GroupName             *string                `xml:"groupName,omitempty" json:"groupName,omitempty" yaml:"groupName,omitempty"`
	GroupObjectClass      *string                `xml:"groupObjectClass,omitempty" json:"groupObjectClass,omitempty" yaml:"groupObjectClass,omitempty"`
	GroupSearchBase       *string                `xml:"groupSearchBase,omitempty" json:"groupSearchBase,omitempty" yaml:"groupSearchBase,omitempty"`
	Primary               *bool                  `xml:"primary,omitempty" json:"primary,omitempty" yaml:"primary,omitempty"`
	SecureConnection      *bool                  `xml:"secureConnection,omitempty" json:"secureConnection,omitempty" yaml:"secureConnection,omitempty"`
	ServerDomain          *string                `xml:"serverDomain,omitempty" json:"serverDomain,omitempty" yaml:"serverDomain,omitempty"`
	ServerDomainIdDataObj *ServerDomainIdDataObj `xml:"serverDomainIdDataObj,omitempty" json:"serverDomainIdDataObj,omitempty" yaml:"serverDomainIdDataObj,omitempty"`
	ServerPort            *int64                 `xml:"serverPort,omitempty" json:"serverPort,omitempty" yaml:"serverPort,omitempty"`
	TlsEnabled            *bool                  `xml:"tlsEnabled,omitempty" json:"tlsEnabled,omitempty" yaml:"tlsEnabled,omitempty"`
	UserEmail             *string                `xml:"userEmail,omitempty" json:"userEmail,omitempty" yaml:"userEmail,omitempty"`
	UserFirstName         *string                `xml:"userFirstName,omitempty" json:"userFirstName,omitempty" yaml:"userFirstName,omitempty"`
	UserLastName          *string                `xml:"userLastName,omitempty" json:"userLastName,omitempty" yaml:"userLastName,omitempty"`
	UserName              *string                `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	UserObjectClass       *string                `xml:"userObjectClass,omitempty" json:"userObjectClass,omitempty" yaml:"userObjectClass,omitempty"`
	UserSearchBase        *string                `xml:"userSearchBase,omitempty" json:"userSearchBase,omitempty" yaml:"userSearchBase,omitempty"`
}

// LdapConfigurationSpecDataObj was auto-generated from WSDL.
type LdapConfigurationSpecDataObj struct {
	AnonymousBind    *bool   `xml:"anonymousBind,omitempty" json:"anonymousBind,omitempty" yaml:"anonymousBind,omitempty"`
	BaseDN           *string `xml:"baseDN,omitempty" json:"baseDN,omitempty" yaml:"baseDN,omitempty"`
	BindName         *string `xml:"bindName,omitempty" json:"bindName,omitempty" yaml:"bindName,omitempty"`
	BindPassword     *string `xml:"bindPassword,omitempty" json:"bindPassword,omitempty" yaml:"bindPassword,omitempty"`
	DisplayName      *string `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	GroupFilter      *string `xml:"groupFilter,omitempty" json:"groupFilter,omitempty" yaml:"groupFilter,omitempty"`
	GroupFullName    *bool   `xml:"groupFullName,omitempty" json:"groupFullName,omitempty" yaml:"groupFullName,omitempty"`
	GroupMember      *string `xml:"groupMember,omitempty" json:"groupMember,omitempty" yaml:"groupMember,omitempty"`
	GroupName        *string `xml:"groupName,omitempty" json:"groupName,omitempty" yaml:"groupName,omitempty"`
	GroupObjectClass *string `xml:"groupObjectClass,omitempty" json:"groupObjectClass,omitempty" yaml:"groupObjectClass,omitempty"`
	GroupSearchBase  *string `xml:"groupSearchBase,omitempty" json:"groupSearchBase,omitempty" yaml:"groupSearchBase,omitempty"`
	Primary          *bool   `xml:"primary,omitempty" json:"primary,omitempty" yaml:"primary,omitempty"`
	SecureConnection *bool   `xml:"secureConnection,omitempty" json:"secureConnection,omitempty" yaml:"secureConnection,omitempty"`
	ServerDomain     *string `xml:"serverDomain,omitempty" json:"serverDomain,omitempty" yaml:"serverDomain,omitempty"`
	ServerPort       *int64  `xml:"serverPort,omitempty" json:"serverPort,omitempty" yaml:"serverPort,omitempty"`
	TlsEnabled       *bool   `xml:"tlsEnabled,omitempty" json:"tlsEnabled,omitempty" yaml:"tlsEnabled,omitempty"`
	UserEmail        *string `xml:"userEmail,omitempty" json:"userEmail,omitempty" yaml:"userEmail,omitempty"`
	UserFirstName    *string `xml:"userFirstName,omitempty" json:"userFirstName,omitempty" yaml:"userFirstName,omitempty"`
	UserLastName     *string `xml:"userLastName,omitempty" json:"userLastName,omitempty" yaml:"userLastName,omitempty"`
	UserName         *string `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	UserObjectClass  *string `xml:"userObjectClass,omitempty" json:"userObjectClass,omitempty" yaml:"userObjectClass,omitempty"`
	UserSearchBase   *string `xml:"userSearchBase,omitempty" json:"userSearchBase,omitempty" yaml:"userSearchBase,omitempty"`
}

// LicenseDataObj was auto-generated from WSDL.
type LicenseDataObj struct {
	Customer           *string   `xml:"customer,omitempty" json:"customer,omitempty" yaml:"customer,omitempty"`
	ExpirationDate     *DateTime `xml:"expirationDate,omitempty" json:"expirationDate,omitempty" yaml:"expirationDate,omitempty"`
	LicenseEditionName *string   `xml:"licenseEditionName,omitempty" json:"licenseEditionName,omitempty" yaml:"licenseEditionName,omitempty"`
	Loc                *int64    `xml:"loc,omitempty" json:"loc,omitempty" yaml:"loc,omitempty"`
	LocLimit           *int64    `xml:"locLimit,omitempty" json:"locLimit,omitempty" yaml:"locLimit,omitempty"`
	UserCount          *int      `xml:"userCount,omitempty" json:"userCount,omitempty" yaml:"userCount,omitempty"`
	UserLimit          *string   `xml:"userLimit,omitempty" json:"userLimit,omitempty" yaml:"userLimit,omitempty"`
}

// LicenseSpecDataObj was auto-generated from WSDL.
type LicenseSpecDataObj struct {
	LicenseDataFile *string `xml:"licenseDataFile,omitempty" json:"licenseDataFile,omitempty" yaml:"licenseDataFile,omitempty"`
}

// LicenseStateDataObj was auto-generated from WSDL.
type LicenseStateDataObj struct {
	DesktopAnalysisEnabled *bool `xml:"desktopAnalysisEnabled,omitempty" json:"desktopAnalysisEnabled,omitempty" yaml:"desktopAnalysisEnabled,omitempty"`
}

// LocalizedValueDataObj was auto-generated from WSDL.
type LocalizedValueDataObj struct {
	DisplayName *string `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Name        *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// LoggingConfigurationDataObj was auto-generated from WSDL.
type LoggingConfigurationDataObj struct {
	AccessControlLogging      *bool `xml:"accessControlLogging,omitempty" json:"accessControlLogging,omitempty" yaml:"accessControlLogging,omitempty"`
	BackgroundLogging         *bool `xml:"backgroundLogging,omitempty" json:"backgroundLogging,omitempty" yaml:"backgroundLogging,omitempty"`
	CommitLogging             *bool `xml:"commitLogging,omitempty" json:"commitLogging,omitempty" yaml:"commitLogging,omitempty"`
	ConfigurationLogging      *bool `xml:"configurationLogging,omitempty" json:"configurationLogging,omitempty" yaml:"configurationLogging,omitempty"`
	DatabaseLogging           *bool `xml:"databaseLogging,omitempty" json:"databaseLogging,omitempty" yaml:"databaseLogging,omitempty"`
	FrameworkLogging          *bool `xml:"frameworkLogging,omitempty" json:"frameworkLogging,omitempty" yaml:"frameworkLogging,omitempty"`
	InternalLogging           *bool `xml:"internalLogging,omitempty" json:"internalLogging,omitempty" yaml:"internalLogging,omitempty"`
	NotificationLogging       *bool `xml:"notificationLogging,omitempty" json:"notificationLogging,omitempty" yaml:"notificationLogging,omitempty"`
	PerformanceLogging        *bool `xml:"performanceLogging,omitempty" json:"performanceLogging,omitempty" yaml:"performanceLogging,omitempty"`
	PolicyManagerLogging      *bool `xml:"policyManagerLogging,omitempty" json:"policyManagerLogging,omitempty" yaml:"policyManagerLogging,omitempty"`
	RemoteConfigLogging       *bool `xml:"remoteConfigLogging,omitempty" json:"remoteConfigLogging,omitempty" yaml:"remoteConfigLogging,omitempty"`
	RequestPerformanceLogging *bool `xml:"requestPerformanceLogging,omitempty" json:"requestPerformanceLogging,omitempty" yaml:"requestPerformanceLogging,omitempty"`
	SkeletonizationLogging    *bool `xml:"skeletonizationLogging,omitempty" json:"skeletonizationLogging,omitempty" yaml:"skeletonizationLogging,omitempty"`
	TestAdvisorLogging        *bool `xml:"testAdvisorLogging,omitempty" json:"testAdvisorLogging,omitempty" yaml:"testAdvisorLogging,omitempty"`
	TriageLogging             *bool `xml:"triageLogging,omitempty" json:"triageLogging,omitempty" yaml:"triageLogging,omitempty"`
	TriageSynchLogging        *bool `xml:"triageSynchLogging,omitempty" json:"triageSynchLogging,omitempty" yaml:"triageSynchLogging,omitempty"`
	WebLogging                *bool `xml:"webLogging,omitempty" json:"webLogging,omitempty" yaml:"webLogging,omitempty"`
	WebServicesLogging        *bool `xml:"webServicesLogging,omitempty" json:"webServicesLogging,omitempty" yaml:"webServicesLogging,omitempty"`
}

// MergeTriageStores was auto-generated from WSDL.
type MergeTriageStores struct {
	SrcTriageStoreIds          []*TriageStoreIdDataObj `xml:"srcTriageStoreIds,omitempty" json:"srcTriageStoreIds,omitempty" yaml:"srcTriageStoreIds,omitempty"`
	TriageStoreId              *TriageStoreIdDataObj   `xml:"triageStoreId,omitempty" json:"triageStoreId,omitempty" yaml:"triageStoreId,omitempty"`
	DeleteSourceStores         *bool                   `xml:"deleteSourceStores,omitempty" json:"deleteSourceStores,omitempty" yaml:"deleteSourceStores,omitempty"`
	AssignStreamsToTargetStore *bool                   `xml:"assignStreamsToTargetStore,omitempty" json:"assignStreamsToTargetStore,omitempty" yaml:"assignStreamsToTargetStore,omitempty"`
}

// MergeTriageStoresResponse was auto-generated from WSDL.
type MergeTriageStoresResponse struct {
}

// Notify was auto-generated from WSDL.
type Notify struct {
	Usernames []*string `xml:"usernames,omitempty" json:"usernames,omitempty" yaml:"usernames,omitempty"`
	Subject   *string   `xml:"subject,omitempty" json:"subject,omitempty" yaml:"subject,omitempty"`
	Message   *string   `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// NotifyResponse was auto-generated from WSDL.
type NotifyResponse struct {
	Return []*string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// OutputFileDataObj was auto-generated from WSDL.
type OutputFileDataObj struct {
	Contents *[]byte `xml:"contents,omitempty" json:"contents,omitempty" yaml:"contents,omitempty"`
	Name     *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// PageSpecDataObj was auto-generated from WSDL.
type PageSpecDataObj struct {
	PageSize      *int    `xml:"pageSize,omitempty" json:"pageSize,omitempty" yaml:"pageSize,omitempty"`
	SortAscending *bool   `xml:"sortAscending,omitempty" json:"sortAscending,omitempty" yaml:"sortAscending,omitempty"`
	SortField     *string `xml:"sortField,omitempty" json:"sortField,omitempty" yaml:"sortField,omitempty"`
	StartIndex    *int    `xml:"startIndex,omitempty" json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
}

// PermissionDataObj was auto-generated from WSDL.
type PermissionDataObj struct {
	PermissionValue *string `xml:"permissionValue,omitempty" json:"permissionValue,omitempty" yaml:"permissionValue,omitempty"`
}

// ProjectDataObj was auto-generated from WSDL.
type ProjectDataObj struct {
	Streams         []*StreamDataObj         `xml:"streams,omitempty" json:"streams,omitempty" yaml:"streams,omitempty"`
	StreamLinks     []*StreamDataObj         `xml:"streamLinks,omitempty" json:"streamLinks,omitempty" yaml:"streamLinks,omitempty"`
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	DateCreated     *DateTime                `xml:"dateCreated,omitempty" json:"dateCreated,omitempty" yaml:"dateCreated,omitempty"`
	DateModified    *DateTime                `xml:"dateModified,omitempty" json:"dateModified,omitempty" yaml:"dateModified,omitempty"`
	Description     *string                  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Id              *ProjectIdDataObj        `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	ProjectKey      *int64                   `xml:"projectKey,omitempty" json:"projectKey,omitempty" yaml:"projectKey,omitempty"`
	UserCreated     *string                  `xml:"userCreated,omitempty" json:"userCreated,omitempty" yaml:"userCreated,omitempty"`
	UserModified    *string                  `xml:"userModified,omitempty" json:"userModified,omitempty" yaml:"userModified,omitempty"`
}

// ProjectFilterSpecDataObj was auto-generated from WSDL.
type ProjectFilterSpecDataObj struct {
	DescriptionPattern *string `xml:"descriptionPattern,omitempty" json:"descriptionPattern,omitempty" yaml:"descriptionPattern,omitempty"`
	IncludeChildren    *bool   `xml:"includeChildren,omitempty" json:"includeChildren,omitempty" yaml:"includeChildren,omitempty"`
	IncludeStreams     *bool   `xml:"includeStreams,omitempty" json:"includeStreams,omitempty" yaml:"includeStreams,omitempty"`
	NamePattern        *string `xml:"namePattern,omitempty" json:"namePattern,omitempty" yaml:"namePattern,omitempty"`
}

// ProjectIdDataObj was auto-generated from WSDL.
type ProjectIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// ProjectSpecDataObj was auto-generated from WSDL.
type ProjectSpecDataObj struct {
	Streams         []*StreamIdDataObj       `xml:"streams,omitempty" json:"streams,omitempty" yaml:"streams,omitempty"`
	StreamLinks     []*StreamIdDataObj       `xml:"streamLinks,omitempty" json:"streamLinks,omitempty" yaml:"streamLinks,omitempty"`
	Description     *string                  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Name            *string                  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
}

// RoleAssignmentDataObj was auto-generated from WSDL.
type RoleAssignmentDataObj struct {
	GroupId            *GroupIdDataObj `xml:"groupId,omitempty" json:"groupId,omitempty" yaml:"groupId,omitempty"`
	RoleAssignmentType *string         `xml:"roleAssignmentType,omitempty" json:"roleAssignmentType,omitempty" yaml:"roleAssignmentType,omitempty"`
	RoleId             *RoleIdDataObj  `xml:"roleId,omitempty" json:"roleId,omitempty" yaml:"roleId,omitempty"`
	Type               *string         `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Username           *string         `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
}

// RoleDataObj was auto-generated from WSDL.
type RoleDataObj struct {
	Deletable          *bool                `xml:"deletable,omitempty" json:"deletable,omitempty" yaml:"deletable,omitempty"`
	Description        *string              `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Editable           *bool                `xml:"editable,omitempty" json:"editable,omitempty" yaml:"editable,omitempty"`
	PermissionDataObjs []*PermissionDataObj `xml:"permissionDataObjs,omitempty" json:"permissionDataObjs,omitempty" yaml:"permissionDataObjs,omitempty"`
	RoleId             *RoleIdDataObj       `xml:"roleId,omitempty" json:"roleId,omitempty" yaml:"roleId,omitempty"`
}

// RoleIdDataObj was auto-generated from WSDL.
type RoleIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// RoleSpecDataObj was auto-generated from WSDL.
type RoleSpecDataObj struct {
	Deletable          *bool                `xml:"deletable,omitempty" json:"deletable,omitempty" yaml:"deletable,omitempty"`
	Description        *string              `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Editable           *bool                `xml:"editable,omitempty" json:"editable,omitempty" yaml:"editable,omitempty"`
	Name               *string              `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	PermissionDataObjs []*PermissionDataObj `xml:"permissionDataObjs,omitempty" json:"permissionDataObjs,omitempty" yaml:"permissionDataObjs,omitempty"`
}

// ServerDomainIdDataObj was auto-generated from WSDL.
type ServerDomainIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// SetAcceptingNewCommits was auto-generated from WSDL.
type SetAcceptingNewCommits struct {
	AcceptNewCommits *bool `xml:"acceptNewCommits,omitempty" json:"acceptNewCommits,omitempty" yaml:"acceptNewCommits,omitempty"`
}

// SetAcceptingNewCommitsResponse was auto-generated from WSDL.
type SetAcceptingNewCommitsResponse struct {
}

// SetArchitectureAnalysisConfiguration was auto-generated from
// WSDL.
type SetArchitectureAnalysisConfiguration struct {
	ArchitectureAnalysisConfiguration *string `xml:"architectureAnalysisConfiguration,omitempty" json:"architectureAnalysisConfiguration,omitempty" yaml:"architectureAnalysisConfiguration,omitempty"`
}

// SetArchitectureAnalysisConfigurationResponse was auto-generated
// from WSDL.
type SetArchitectureAnalysisConfigurationResponse struct {
}

// SetBackupConfiguration was auto-generated from WSDL.
type SetBackupConfiguration struct {
	BackupConfigurationDataObj *BackupConfigurationDataObj `xml:"backupConfigurationDataObj,omitempty" json:"backupConfigurationDataObj,omitempty" yaml:"backupConfigurationDataObj,omitempty"`
}

// SetBackupConfigurationResponse was auto-generated from WSDL.
type SetBackupConfigurationResponse struct {
}

// SetLoggingConfiguration was auto-generated from WSDL.
type SetLoggingConfiguration struct {
	LoggingConfigurationDataObj *LoggingConfigurationDataObj `xml:"loggingConfigurationDataObj,omitempty" json:"loggingConfigurationDataObj,omitempty" yaml:"loggingConfigurationDataObj,omitempty"`
}

// SetLoggingConfigurationResponse was auto-generated from WSDL.
type SetLoggingConfigurationResponse struct {
}

// SetMessageOfTheDay was auto-generated from WSDL.
type SetMessageOfTheDay struct {
	Message *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// SetMessageOfTheDayResponse was auto-generated from WSDL.
type SetMessageOfTheDayResponse struct {
}

// SetSkeletonizationConfiguration was auto-generated from WSDL.
type SetSkeletonizationConfiguration struct {
	SkeletonizationConfigurationDataObj *SkeletonizationConfigurationDataObj `xml:"skeletonizationConfigurationDataObj,omitempty" json:"skeletonizationConfigurationDataObj,omitempty" yaml:"skeletonizationConfigurationDataObj,omitempty"`
}

// SetSkeletonizationConfigurationResponse was auto-generated from
// WSDL.
type SetSkeletonizationConfigurationResponse struct {
}

// SetSnapshotPurgeDetails was auto-generated from WSDL.
type SetSnapshotPurgeDetails struct {
	PurgeDetailsSpec *SnapshotPurgeDetailsObj `xml:"purgeDetailsSpec,omitempty" json:"purgeDetailsSpec,omitempty" yaml:"purgeDetailsSpec,omitempty"`
}

// SetSnapshotPurgeDetailsResponse was auto-generated from WSDL.
type SetSnapshotPurgeDetailsResponse struct {
}

// SignInSettingsDataObj was auto-generated from WSDL.
type SignInSettingsDataObj struct {
	AllowPasswordRecovery      *bool   `xml:"allowPasswordRecovery,omitempty" json:"allowPasswordRecovery,omitempty" yaml:"allowPasswordRecovery,omitempty"`
	AuthenticationMethod       *string `xml:"authenticationMethod,omitempty" json:"authenticationMethod,omitempty" yaml:"authenticationMethod,omitempty"`
	DisableLocalPasswordAuth   *bool   `xml:"disableLocalPasswordAuth,omitempty" json:"disableLocalPasswordAuth,omitempty" yaml:"disableLocalPasswordAuth,omitempty"`
	EnableLdapAuth             *bool   `xml:"enableLdapAuth,omitempty" json:"enableLdapAuth,omitempty" yaml:"enableLdapAuth,omitempty"`
	EnableSessionTimeout       *bool   `xml:"enableSessionTimeout,omitempty" json:"enableSessionTimeout,omitempty" yaml:"enableSessionTimeout,omitempty"`
	LdapUserAutoCreate         *bool   `xml:"ldapUserAutoCreate,omitempty" json:"ldapUserAutoCreate,omitempty" yaml:"ldapUserAutoCreate,omitempty"`
	LimitFailedSignIns         *bool   `xml:"limitFailedSignIns,omitempty" json:"limitFailedSignIns,omitempty" yaml:"limitFailedSignIns,omitempty"`
	MaxFailedSignInAttempts    *int    `xml:"maxFailedSignInAttempts,omitempty" json:"maxFailedSignInAttempts,omitempty" yaml:"maxFailedSignInAttempts,omitempty"`
	MaxSessionIdleTime         *int    `xml:"maxSessionIdleTime,omitempty" json:"maxSessionIdleTime,omitempty" yaml:"maxSessionIdleTime,omitempty"`
	RequireLdapGroupMembership *bool   `xml:"requireLdapGroupMembership,omitempty" json:"requireLdapGroupMembership,omitempty" yaml:"requireLdapGroupMembership,omitempty"`
}

// SkeletonizationConfigurationDataObj was auto-generated from
// WSDL.
type SkeletonizationConfigurationDataObj struct {
	DaysBeforeSkeletonization *int    `xml:"daysBeforeSkeletonization,omitempty" json:"daysBeforeSkeletonization,omitempty" yaml:"daysBeforeSkeletonization,omitempty"`
	FridayEnabled             *bool   `xml:"fridayEnabled,omitempty" json:"fridayEnabled,omitempty" yaml:"fridayEnabled,omitempty"`
	MinSnapshotsToKeep        *int    `xml:"minSnapshotsToKeep,omitempty" json:"minSnapshotsToKeep,omitempty" yaml:"minSnapshotsToKeep,omitempty"`
	MondayEnabled             *bool   `xml:"mondayEnabled,omitempty" json:"mondayEnabled,omitempty" yaml:"mondayEnabled,omitempty"`
	SaturdayEnabled           *bool   `xml:"saturdayEnabled,omitempty" json:"saturdayEnabled,omitempty" yaml:"saturdayEnabled,omitempty"`
	SundayEnabled             *bool   `xml:"sundayEnabled,omitempty" json:"sundayEnabled,omitempty" yaml:"sundayEnabled,omitempty"`
	ThursdayEnabled           *bool   `xml:"thursdayEnabled,omitempty" json:"thursdayEnabled,omitempty" yaml:"thursdayEnabled,omitempty"`
	Time                      *string `xml:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
	TuesdayEnabled            *bool   `xml:"tuesdayEnabled,omitempty" json:"tuesdayEnabled,omitempty" yaml:"tuesdayEnabled,omitempty"`
	WednesdayEnabled          *bool   `xml:"wednesdayEnabled,omitempty" json:"wednesdayEnabled,omitempty" yaml:"wednesdayEnabled,omitempty"`
}

// SnapshotFilterSpecDataObj was auto-generated from WSDL.
type SnapshotFilterSpecDataObj struct {
	DescriptionPattern        *string   `xml:"descriptionPattern,omitempty" json:"descriptionPattern,omitempty" yaml:"descriptionPattern,omitempty"`
	EndDate                   *DateTime `xml:"endDate,omitempty" json:"endDate,omitempty" yaml:"endDate,omitempty"`
	HasSummaries              *bool     `xml:"hasSummaries,omitempty" json:"hasSummaries,omitempty" yaml:"hasSummaries,omitempty"`
	LastBeforeCodeVersionDate *DateTime `xml:"lastBeforeCodeVersionDate,omitempty" json:"lastBeforeCodeVersionDate,omitempty" yaml:"lastBeforeCodeVersionDate,omitempty"`
	StartDate                 *DateTime `xml:"startDate,omitempty" json:"startDate,omitempty" yaml:"startDate,omitempty"`
	TargetPattern             *string   `xml:"targetPattern,omitempty" json:"targetPattern,omitempty" yaml:"targetPattern,omitempty"`
	VersionPattern            *string   `xml:"versionPattern,omitempty" json:"versionPattern,omitempty" yaml:"versionPattern,omitempty"`
}

// SnapshotIdDataObj was auto-generated from WSDL.
type SnapshotIdDataObj struct {
	Id *int64 `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// SnapshotInfoDataObj was auto-generated from WSDL.
type SnapshotInfoDataObj struct {
	AnalysisCommandLine      *string            `xml:"analysisCommandLine,omitempty" json:"analysisCommandLine,omitempty" yaml:"analysisCommandLine,omitempty"`
	AnalysisConfiguration    *string            `xml:"analysisConfiguration,omitempty" json:"analysisConfiguration,omitempty" yaml:"analysisConfiguration,omitempty"`
	AnalysisHost             *string            `xml:"analysisHost,omitempty" json:"analysisHost,omitempty" yaml:"analysisHost,omitempty"`
	AnalysisIntermediateDir  *string            `xml:"analysisIntermediateDir,omitempty" json:"analysisIntermediateDir,omitempty" yaml:"analysisIntermediateDir,omitempty"`
	AnalysisInternalVersion  *string            `xml:"analysisInternalVersion,omitempty" json:"analysisInternalVersion,omitempty" yaml:"analysisInternalVersion,omitempty"`
	AnalysisTime             *int64             `xml:"analysisTime,omitempty" json:"analysisTime,omitempty" yaml:"analysisTime,omitempty"`
	AnalysisVersion          *string            `xml:"analysisVersion,omitempty" json:"analysisVersion,omitempty" yaml:"analysisVersion,omitempty"`
	BuildCommandLine         *string            `xml:"buildCommandLine,omitempty" json:"buildCommandLine,omitempty" yaml:"buildCommandLine,omitempty"`
	BuildConfiguration       *string            `xml:"buildConfiguration,omitempty" json:"buildConfiguration,omitempty" yaml:"buildConfiguration,omitempty"`
	BuildHost                *string            `xml:"buildHost,omitempty" json:"buildHost,omitempty" yaml:"buildHost,omitempty"`
	BuildIntermediateDir     *string            `xml:"buildIntermediateDir,omitempty" json:"buildIntermediateDir,omitempty" yaml:"buildIntermediateDir,omitempty"`
	BuildTime                *int64             `xml:"buildTime,omitempty" json:"buildTime,omitempty" yaml:"buildTime,omitempty"`
	CodeVersionDate          *DateTime          `xml:"codeVersionDate,omitempty" json:"codeVersionDate,omitempty" yaml:"codeVersionDate,omitempty"`
	CommitUser               *string            `xml:"commitUser,omitempty" json:"commitUser,omitempty" yaml:"commitUser,omitempty"`
	DateCreated              *DateTime          `xml:"dateCreated,omitempty" json:"dateCreated,omitempty" yaml:"dateCreated,omitempty"`
	Description              *string            `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	EnabledCheckers          []*string          `xml:"enabledCheckers,omitempty" json:"enabledCheckers,omitempty" yaml:"enabledCheckers,omitempty"`
	HasSummaries             *bool              `xml:"hasSummaries,omitempty" json:"hasSummaries,omitempty" yaml:"hasSummaries,omitempty"`
	ImpactHashVersion        *int               `xml:"impactHashVersion,omitempty" json:"impactHashVersion,omitempty" yaml:"impactHashVersion,omitempty"`
	PortableAnalysisSettings *string            `xml:"portableAnalysisSettings,omitempty" json:"portableAnalysisSettings,omitempty" yaml:"portableAnalysisSettings,omitempty"`
	PurgedOfDetails          *bool              `xml:"purgedOfDetails,omitempty" json:"purgedOfDetails,omitempty" yaml:"purgedOfDetails,omitempty"`
	SnapshotId               *SnapshotIdDataObj `xml:"snapshotId,omitempty" json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
	SourceVersion            *string            `xml:"sourceVersion,omitempty" json:"sourceVersion,omitempty" yaml:"sourceVersion,omitempty"`
	Target                   *string            `xml:"target,omitempty" json:"target,omitempty" yaml:"target,omitempty"`
}

// SnapshotPurgeDetailsObj was auto-generated from WSDL.
type SnapshotPurgeDetailsObj struct {
	DaysBeforeSkeletonization *int    `xml:"daysBeforeSkeletonization,omitempty" json:"daysBeforeSkeletonization,omitempty" yaml:"daysBeforeSkeletonization,omitempty"`
	FridayEnabled             *bool   `xml:"fridayEnabled,omitempty" json:"fridayEnabled,omitempty" yaml:"fridayEnabled,omitempty"`
	MinSnapshotsToKeep        *int    `xml:"minSnapshotsToKeep,omitempty" json:"minSnapshotsToKeep,omitempty" yaml:"minSnapshotsToKeep,omitempty"`
	MondayEnabled             *bool   `xml:"mondayEnabled,omitempty" json:"mondayEnabled,omitempty" yaml:"mondayEnabled,omitempty"`
	SaturdayEnabled           *bool   `xml:"saturdayEnabled,omitempty" json:"saturdayEnabled,omitempty" yaml:"saturdayEnabled,omitempty"`
	SundayEnabled             *bool   `xml:"sundayEnabled,omitempty" json:"sundayEnabled,omitempty" yaml:"sundayEnabled,omitempty"`
	ThursdayEnabled           *bool   `xml:"thursdayEnabled,omitempty" json:"thursdayEnabled,omitempty" yaml:"thursdayEnabled,omitempty"`
	TimeOfDay                 *string `xml:"timeOfDay,omitempty" json:"timeOfDay,omitempty" yaml:"timeOfDay,omitempty"`
	TuesdayEnabled            *bool   `xml:"tuesdayEnabled,omitempty" json:"tuesdayEnabled,omitempty" yaml:"tuesdayEnabled,omitempty"`
	WednesdayEnabled          *bool   `xml:"wednesdayEnabled,omitempty" json:"wednesdayEnabled,omitempty" yaml:"wednesdayEnabled,omitempty"`
}

// StreamDataObj was auto-generated from WSDL.
type StreamDataObj struct {
	AutoDeleteOnExpiry      *bool                    `xml:"autoDeleteOnExpiry,omitempty" json:"autoDeleteOnExpiry,omitempty" yaml:"autoDeleteOnExpiry,omitempty"`
	ComponentMapId          *ComponentMapIdDataObj   `xml:"componentMapId,omitempty" json:"componentMapId,omitempty" yaml:"componentMapId,omitempty"`
	Description             *string                  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Id                      *StreamIdDataObj         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Language                *string                  `xml:"language,omitempty" json:"language,omitempty" yaml:"language,omitempty"`
	Outdated                *bool                    `xml:"outdated,omitempty" json:"outdated,omitempty" yaml:"outdated,omitempty"`
	PrimaryProjectId        *ProjectIdDataObj        `xml:"primaryProjectId,omitempty" json:"primaryProjectId,omitempty" yaml:"primaryProjectId,omitempty"`
	TriageStoreId           *TriageStoreIdDataObj    `xml:"triageStoreId,omitempty" json:"triageStoreId,omitempty" yaml:"triageStoreId,omitempty"`
	RoleAssignments         []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	AnalysisVersionOverride *string                  `xml:"analysisVersionOverride,omitempty" json:"analysisVersionOverride,omitempty" yaml:"analysisVersionOverride,omitempty"`
	SummaryExpirationDays   *int                     `xml:"summaryExpirationDays,omitempty" json:"summaryExpirationDays,omitempty" yaml:"summaryExpirationDays,omitempty"`
	PluginVersionOverride   *string                  `xml:"pluginVersionOverride,omitempty" json:"pluginVersionOverride,omitempty" yaml:"pluginVersionOverride,omitempty"`
	VersionMismatchMessage  *string                  `xml:"versionMismatchMessage,omitempty" json:"versionMismatchMessage,omitempty" yaml:"versionMismatchMessage,omitempty"`
	EnableDesktopAnalysis   *bool                    `xml:"enableDesktopAnalysis,omitempty" json:"enableDesktopAnalysis,omitempty" yaml:"enableDesktopAnalysis,omitempty"`
	OwnerAssignmentOption   *string                  `xml:"ownerAssignmentOption,omitempty" json:"ownerAssignmentOption,omitempty" yaml:"ownerAssignmentOption,omitempty"`
}

// StreamFilterSpecDataObj was auto-generated from WSDL.
type StreamFilterSpecDataObj struct {
	LanguageList       []*string `xml:"languageList,omitempty" json:"languageList,omitempty" yaml:"languageList,omitempty"`
	DescriptionPattern *string   `xml:"descriptionPattern,omitempty" json:"descriptionPattern,omitempty" yaml:"descriptionPattern,omitempty"`
	NamePattern        *string   `xml:"namePattern,omitempty" json:"namePattern,omitempty" yaml:"namePattern,omitempty"`
}

// StreamIdDataObj was auto-generated from WSDL.
type StreamIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// StreamSpecDataObj was auto-generated from WSDL.
type StreamSpecDataObj struct {
	AllowCommitWithoutPassword *bool                    `xml:"allowCommitWithoutPassword,omitempty" json:"allowCommitWithoutPassword,omitempty" yaml:"allowCommitWithoutPassword,omitempty"`
	AnalysisVersionOverride    *string                  `xml:"analysisVersionOverride,omitempty" json:"analysisVersionOverride,omitempty" yaml:"analysisVersionOverride,omitempty"`
	AutoDeleteOnExpiry         *bool                    `xml:"autoDeleteOnExpiry,omitempty" json:"autoDeleteOnExpiry,omitempty" yaml:"autoDeleteOnExpiry,omitempty"`
	ComponentMapId             *ComponentMapIdDataObj   `xml:"componentMapId,omitempty" json:"componentMapId,omitempty" yaml:"componentMapId,omitempty"`
	Description                *string                  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	EnableDesktopAnalysis      *bool                    `xml:"enableDesktopAnalysis,omitempty" json:"enableDesktopAnalysis,omitempty" yaml:"enableDesktopAnalysis,omitempty"`
	Language                   *string                  `xml:"language,omitempty" json:"language,omitempty" yaml:"language,omitempty"`
	Name                       *string                  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Outdated                   *bool                    `xml:"outdated,omitempty" json:"outdated,omitempty" yaml:"outdated,omitempty"`
	OwnerAssignmentOption      *string                  `xml:"ownerAssignmentOption,omitempty" json:"ownerAssignmentOption,omitempty" yaml:"ownerAssignmentOption,omitempty"`
	PluginVersionOverride      *string                  `xml:"pluginVersionOverride,omitempty" json:"pluginVersionOverride,omitempty" yaml:"pluginVersionOverride,omitempty"`
	RoleAssignments            []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	SummaryExpirationDays      *int                     `xml:"summaryExpirationDays,omitempty" json:"summaryExpirationDays,omitempty" yaml:"summaryExpirationDays,omitempty"`
	TriageStoreId              *TriageStoreIdDataObj    `xml:"triageStoreId,omitempty" json:"triageStoreId,omitempty" yaml:"triageStoreId,omitempty"`
	VersionMismatchMessage     *string                  `xml:"versionMismatchMessage,omitempty" json:"versionMismatchMessage,omitempty" yaml:"versionMismatchMessage,omitempty"`
}

// TriageStoreDataObj was auto-generated from WSDL.
type TriageStoreDataObj struct {
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	Description     *string                  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Id              *TriageStoreIdDataObj    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	StreamIds       []*StreamIdDataObj       `xml:"streamIds,omitempty" json:"streamIds,omitempty" yaml:"streamIds,omitempty"`
}

// TriageStoreFilterSpecDataObj was auto-generated from WSDL.
type TriageStoreFilterSpecDataObj struct {
	DescriptionPattern *string `xml:"descriptionPattern,omitempty" json:"descriptionPattern,omitempty" yaml:"descriptionPattern,omitempty"`
	NamePattern        *string `xml:"namePattern,omitempty" json:"namePattern,omitempty" yaml:"namePattern,omitempty"`
}

// TriageStoreIdDataObj was auto-generated from WSDL.
type TriageStoreIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// TriageStoreSpecDataObj was auto-generated from WSDL.
type TriageStoreSpecDataObj struct {
	Description     *string                  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Name            *string                  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
}

// UpdateAttribute was auto-generated from WSDL.
type UpdateAttribute struct {
	AttributeDefinitionId   *AttributeDefinitionIdDataObj   `xml:"attributeDefinitionId,omitempty" json:"attributeDefinitionId,omitempty" yaml:"attributeDefinitionId,omitempty"`
	AttributeDefinitionSpec *AttributeDefinitionSpecDataObj `xml:"attributeDefinitionSpec,omitempty" json:"attributeDefinitionSpec,omitempty" yaml:"attributeDefinitionSpec,omitempty"`
}

// UpdateAttributeResponse was auto-generated from WSDL.
type UpdateAttributeResponse struct {
}

// UpdateComponentMap was auto-generated from WSDL.
type UpdateComponentMap struct {
	ComponentMapId   *ComponentMapIdDataObj   `xml:"componentMapId,omitempty" json:"componentMapId,omitempty" yaml:"componentMapId,omitempty"`
	ComponentMapSpec *ComponentMapSpecDataObj `xml:"componentMapSpec,omitempty" json:"componentMapSpec,omitempty" yaml:"componentMapSpec,omitempty"`
}

// UpdateComponentMapResponse was auto-generated from WSDL.
type UpdateComponentMapResponse struct {
}

// UpdateGroup was auto-generated from WSDL.
type UpdateGroup struct {
	GroupId   *GroupIdDataObj   `xml:"groupId,omitempty" json:"groupId,omitempty" yaml:"groupId,omitempty"`
	GroupSpec *GroupSpecDataObj `xml:"groupSpec,omitempty" json:"groupSpec,omitempty" yaml:"groupSpec,omitempty"`
}

// UpdateGroupResponse was auto-generated from WSDL.
type UpdateGroupResponse struct {
}

// UpdateLdapConfiguration was auto-generated from WSDL.
type UpdateLdapConfiguration struct {
	ServerDomainIdDataObj *ServerDomainIdDataObj        `xml:"serverDomainIdDataObj,omitempty" json:"serverDomainIdDataObj,omitempty" yaml:"serverDomainIdDataObj,omitempty"`
	LdapConfigurationSpec *LdapConfigurationSpecDataObj `xml:"ldapConfigurationSpec,omitempty" json:"ldapConfigurationSpec,omitempty" yaml:"ldapConfigurationSpec,omitempty"`
}

// UpdateLdapConfigurationResponse was auto-generated from WSDL.
type UpdateLdapConfigurationResponse struct {
}

// UpdateProject was auto-generated from WSDL.
type UpdateProject struct {
	ProjectId   *ProjectIdDataObj   `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	ProjectSpec *ProjectSpecDataObj `xml:"projectSpec,omitempty" json:"projectSpec,omitempty" yaml:"projectSpec,omitempty"`
}

// UpdateProjectResponse was auto-generated from WSDL.
type UpdateProjectResponse struct {
}

// UpdateRole was auto-generated from WSDL.
type UpdateRole struct {
	RoleId   *RoleIdDataObj   `xml:"roleId,omitempty" json:"roleId,omitempty" yaml:"roleId,omitempty"`
	RoleSpec *RoleSpecDataObj `xml:"roleSpec,omitempty" json:"roleSpec,omitempty" yaml:"roleSpec,omitempty"`
}

// UpdateRoleResponse was auto-generated from WSDL.
type UpdateRoleResponse struct {
}

// UpdateSignInConfiguration was auto-generated from WSDL.
type UpdateSignInConfiguration struct {
	SignInSettingsDataObj *SignInSettingsDataObj `xml:"signInSettingsDataObj,omitempty" json:"signInSettingsDataObj,omitempty" yaml:"signInSettingsDataObj,omitempty"`
}

// UpdateSignInConfigurationResponse was auto-generated from WSDL.
type UpdateSignInConfigurationResponse struct {
}

// UpdateSnapshotInfo was auto-generated from WSDL.
type UpdateSnapshotInfo struct {
	SnapshotId   *SnapshotIdDataObj   `xml:"snapshotId,omitempty" json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
	SnapshotData *SnapshotInfoDataObj `xml:"snapshotData,omitempty" json:"snapshotData,omitempty" yaml:"snapshotData,omitempty"`
}

// UpdateSnapshotInfoResponse was auto-generated from WSDL.
type UpdateSnapshotInfoResponse struct {
}

// UpdateStream was auto-generated from WSDL.
type UpdateStream struct {
	StreamId   *StreamIdDataObj   `xml:"streamId,omitempty" json:"streamId,omitempty" yaml:"streamId,omitempty"`
	StreamSpec *StreamSpecDataObj `xml:"streamSpec,omitempty" json:"streamSpec,omitempty" yaml:"streamSpec,omitempty"`
}

// UpdateStreamResponse was auto-generated from WSDL.
type UpdateStreamResponse struct {
}

// UpdateTriageStore was auto-generated from WSDL.
type UpdateTriageStore struct {
	TriageStoreId   *TriageStoreIdDataObj   `xml:"triageStoreId,omitempty" json:"triageStoreId,omitempty" yaml:"triageStoreId,omitempty"`
	TriageStoreSpec *TriageStoreSpecDataObj `xml:"triageStoreSpec,omitempty" json:"triageStoreSpec,omitempty" yaml:"triageStoreSpec,omitempty"`
}

// UpdateTriageStoreResponse was auto-generated from WSDL.
type UpdateTriageStoreResponse struct {
}

// UpdateUser was auto-generated from WSDL.
type UpdateUser struct {
	Username *string          `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	UserSpec *UserSpecDataObj `xml:"userSpec,omitempty" json:"userSpec,omitempty" yaml:"userSpec,omitempty"`
}

// UpdateUserResponse was auto-generated from WSDL.
type UpdateUserResponse struct {
}

// UserDataObj was auto-generated from WSDL.
type UserDataObj struct {
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	DateCreated     *DateTime                `xml:"dateCreated,omitempty" json:"dateCreated,omitempty" yaml:"dateCreated,omitempty"`
	DateDeleted     *DateTime                `xml:"dateDeleted,omitempty" json:"dateDeleted,omitempty" yaml:"dateDeleted,omitempty"`
	DateModified    *DateTime                `xml:"dateModified,omitempty" json:"dateModified,omitempty" yaml:"dateModified,omitempty"`
	Disabled        *bool                    `xml:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty"`
	Domain          *ServerDomainIdDataObj   `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	Email           *string                  `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	FamilyName      *string                  `xml:"familyName,omitempty" json:"familyName,omitempty" yaml:"familyName,omitempty"`
	GivenName       *string                  `xml:"givenName,omitempty" json:"givenName,omitempty" yaml:"givenName,omitempty"`
	Groups          []*string                `xml:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Local           *bool                    `xml:"local,omitempty" json:"local,omitempty" yaml:"local,omitempty"`
	Locale          *string                  `xml:"locale,omitempty" json:"locale,omitempty" yaml:"locale,omitempty"`
	Locked          *bool                    `xml:"locked,omitempty" json:"locked,omitempty" yaml:"locked,omitempty"`
	PasswordChanged *DateTime                `xml:"passwordChanged,omitempty" json:"passwordChanged,omitempty" yaml:"passwordChanged,omitempty"`
	SuperUser       *bool                    `xml:"superUser,omitempty" json:"superUser,omitempty" yaml:"superUser,omitempty"`
	UserCreated     *string                  `xml:"userCreated,omitempty" json:"userCreated,omitempty" yaml:"userCreated,omitempty"`
	UserDeleted     *string                  `xml:"userDeleted,omitempty" json:"userDeleted,omitempty" yaml:"userDeleted,omitempty"`
	UserModified    *string                  `xml:"userModified,omitempty" json:"userModified,omitempty" yaml:"userModified,omitempty"`
	Username        *string                  `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
}

// UserFilterSpecDataObj was auto-generated from WSDL.
type UserFilterSpecDataObj struct {
	Assignable       *bool             `xml:"assignable,omitempty" json:"assignable,omitempty" yaml:"assignable,omitempty"`
	Disabled         *bool             `xml:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty"`
	GroupsList       []*string         `xml:"groupsList,omitempty" json:"groupsList,omitempty" yaml:"groupsList,omitempty"`
	IncludeDetails   *bool             `xml:"includeDetails,omitempty" json:"includeDetails,omitempty" yaml:"includeDetails,omitempty"`
	Ldap             *bool             `xml:"ldap,omitempty" json:"ldap,omitempty" yaml:"ldap,omitempty"`
	Locked           *bool             `xml:"locked,omitempty" json:"locked,omitempty" yaml:"locked,omitempty"`
	NamePattern      *string           `xml:"namePattern,omitempty" json:"namePattern,omitempty" yaml:"namePattern,omitempty"`
	ProjectIdDataObj *ProjectIdDataObj `xml:"projectIdDataObj,omitempty" json:"projectIdDataObj,omitempty" yaml:"projectIdDataObj,omitempty"`
	StartId          *int64            `xml:"startId,omitempty" json:"startId,omitempty" yaml:"startId,omitempty"`
}

// UserSpecDataObj was auto-generated from WSDL.
type UserSpecDataObj struct {
	ClearRoles      *bool                    `xml:"clearRoles,omitempty" json:"clearRoles,omitempty" yaml:"clearRoles,omitempty"`
	Disabled        *bool                    `xml:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty"`
	Domain          *ServerDomainIdDataObj   `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	Email           *string                  `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	FamilyName      *string                  `xml:"familyName,omitempty" json:"familyName,omitempty" yaml:"familyName,omitempty"`
	GivenName       *string                  `xml:"givenName,omitempty" json:"givenName,omitempty" yaml:"givenName,omitempty"`
	GroupNames      []*GroupIdDataObj        `xml:"groupNames,omitempty" json:"groupNames,omitempty" yaml:"groupNames,omitempty"`
	Local           *bool                    `xml:"local,omitempty" json:"local,omitempty" yaml:"local,omitempty"`
	Locale          *string                  `xml:"locale,omitempty" json:"locale,omitempty" yaml:"locale,omitempty"`
	Locked          *bool                    `xml:"locked,omitempty" json:"locked,omitempty" yaml:"locked,omitempty"`
	Password        *string                  `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	RoleAssignments []*RoleAssignmentDataObj `xml:"roleAssignments,omitempty" json:"roleAssignments,omitempty" yaml:"roleAssignments,omitempty"`
	Username        *string                  `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
}

// UsersPageDataObj was auto-generated from WSDL.
type UsersPageDataObj struct {
	Users                []*UserDataObj `xml:"users,omitempty" json:"users,omitempty" yaml:"users,omitempty"`
	TotalNumberOfRecords *int           `xml:"totalNumberOfRecords,omitempty" json:"totalNumberOfRecords,omitempty" yaml:"totalNumberOfRecords,omitempty"`
}

// VersionDataObj was auto-generated from WSDL.
type VersionDataObj struct {
	ExternalVersion *string `xml:"externalVersion,omitempty" json:"externalVersion,omitempty" yaml:"externalVersion,omitempty"`
	InternalVersion *string `xml:"internalVersion,omitempty" json:"internalVersion,omitempty" yaml:"internalVersion,omitempty"`
}

// Operation wrapper for CopyStream.
// OperationCopyStream was auto-generated from WSDL.
type OperationCopyStream struct {
	CopyStream *CopyStream `xml:"tns:copyStream,omitempty" json:"tns:copyStream,omitempty" yaml:"tns:copyStream,omitempty"`
}

// Operation wrapper for CopyStream.
// OperationCopyStreamResponse was auto-generated from WSDL.
type OperationCopyStreamResponse struct {
	CopyStreamResponse *CopyStreamResponse `xml:"copyStreamResponse,omitempty" json:"copyStreamResponse,omitempty" yaml:"copyStreamResponse,omitempty"`
}

// Operation wrapper for CreateAttribute.
// OperationCreateAttribute was auto-generated from WSDL.
type OperationCreateAttribute struct {
	CreateAttribute *CreateAttribute `xml:"tns:createAttribute,omitempty" json:"tns:createAttribute,omitempty" yaml:"tns:createAttribute,omitempty"`
}

// Operation wrapper for CreateAttribute.
// OperationCreateAttributeResponse was auto-generated from WSDL.
type OperationCreateAttributeResponse struct {
	CreateAttributeResponse *CreateAttributeResponse `xml:"createAttributeResponse,omitempty" json:"createAttributeResponse,omitempty" yaml:"createAttributeResponse,omitempty"`
}

// Operation wrapper for CreateComponentMap.
// OperationCreateComponentMap was auto-generated from WSDL.
type OperationCreateComponentMap struct {
	CreateComponentMap *CreateComponentMap `xml:"tns:createComponentMap,omitempty" json:"tns:createComponentMap,omitempty" yaml:"tns:createComponentMap,omitempty"`
}

// Operation wrapper for CreateComponentMap.
// OperationCreateComponentMapResponse was auto-generated from
// WSDL.
type OperationCreateComponentMapResponse struct {
	CreateComponentMapResponse *CreateComponentMapResponse `xml:"createComponentMapResponse,omitempty" json:"createComponentMapResponse,omitempty" yaml:"createComponentMapResponse,omitempty"`
}

// Operation wrapper for CreateGroup.
// OperationCreateGroup was auto-generated from WSDL.
type OperationCreateGroup struct {
	CreateGroup *CreateGroup `xml:"tns:createGroup,omitempty" json:"tns:createGroup,omitempty" yaml:"tns:createGroup,omitempty"`
}

// Operation wrapper for CreateGroup.
// OperationCreateGroupResponse was auto-generated from WSDL.
type OperationCreateGroupResponse struct {
	CreateGroupResponse *CreateGroupResponse `xml:"createGroupResponse,omitempty" json:"createGroupResponse,omitempty" yaml:"createGroupResponse,omitempty"`
}

// Operation wrapper for CreateLdapConfiguration.
// OperationCreateLdapConfiguration was auto-generated from WSDL.
type OperationCreateLdapConfiguration struct {
	CreateLdapConfiguration *CreateLdapConfiguration `xml:"tns:createLdapConfiguration,omitempty" json:"tns:createLdapConfiguration,omitempty" yaml:"tns:createLdapConfiguration,omitempty"`
}

// Operation wrapper for CreateLdapConfiguration.
// OperationCreateLdapConfigurationResponse was auto-generated
// from WSDL.
type OperationCreateLdapConfigurationResponse struct {
	CreateLdapConfigurationResponse *CreateLdapConfigurationResponse `xml:"createLdapConfigurationResponse,omitempty" json:"createLdapConfigurationResponse,omitempty" yaml:"createLdapConfigurationResponse,omitempty"`
}

// Operation wrapper for CreateProject.
// OperationCreateProject was auto-generated from WSDL.
type OperationCreateProject struct {
	CreateProject *CreateProject `xml:"tns:createProject,omitempty" json:"tns:createProject,omitempty" yaml:"tns:createProject,omitempty"`
}

// Operation wrapper for CreateProject.
// OperationCreateProjectResponse was auto-generated from WSDL.
type OperationCreateProjectResponse struct {
	CreateProjectResponse *CreateProjectResponse `xml:"createProjectResponse,omitempty" json:"createProjectResponse,omitempty" yaml:"createProjectResponse,omitempty"`
}

// Operation wrapper for CreateRole.
// OperationCreateRole was auto-generated from WSDL.
type OperationCreateRole struct {
	CreateRole *CreateRole `xml:"tns:createRole,omitempty" json:"tns:createRole,omitempty" yaml:"tns:createRole,omitempty"`
}

// Operation wrapper for CreateRole.
// OperationCreateRoleResponse was auto-generated from WSDL.
type OperationCreateRoleResponse struct {
	CreateRoleResponse *CreateRoleResponse `xml:"createRoleResponse,omitempty" json:"createRoleResponse,omitempty" yaml:"createRoleResponse,omitempty"`
}

// Operation wrapper for CreateStream.
// OperationCreateStream was auto-generated from WSDL.
type OperationCreateStream struct {
	CreateStream *CreateStream `xml:"tns:createStream,omitempty" json:"tns:createStream,omitempty" yaml:"tns:createStream,omitempty"`
}

// Operation wrapper for CreateStream.
// OperationCreateStreamResponse was auto-generated from WSDL.
type OperationCreateStreamResponse struct {
	CreateStreamResponse *CreateStreamResponse `xml:"createStreamResponse,omitempty" json:"createStreamResponse,omitempty" yaml:"createStreamResponse,omitempty"`
}

// Operation wrapper for CreateStreamInProject.
// OperationCreateStreamInProject was auto-generated from WSDL.
type OperationCreateStreamInProject struct {
	CreateStreamInProject *CreateStreamInProject `xml:"tns:createStreamInProject,omitempty" json:"tns:createStreamInProject,omitempty" yaml:"tns:createStreamInProject,omitempty"`
}

// Operation wrapper for CreateStreamInProject.
// OperationCreateStreamInProjectResponse was auto-generated from
// WSDL.
type OperationCreateStreamInProjectResponse struct {
	CreateStreamInProjectResponse *CreateStreamInProjectResponse `xml:"createStreamInProjectResponse,omitempty" json:"createStreamInProjectResponse,omitempty" yaml:"createStreamInProjectResponse,omitempty"`
}

// Operation wrapper for CreateTriageStore.
// OperationCreateTriageStore was auto-generated from WSDL.
type OperationCreateTriageStore struct {
	CreateTriageStore *CreateTriageStore `xml:"tns:createTriageStore,omitempty" json:"tns:createTriageStore,omitempty" yaml:"tns:createTriageStore,omitempty"`
}

// Operation wrapper for CreateTriageStore.
// OperationCreateTriageStoreResponse was auto-generated from WSDL.
type OperationCreateTriageStoreResponse struct {
	CreateTriageStoreResponse *CreateTriageStoreResponse `xml:"createTriageStoreResponse,omitempty" json:"createTriageStoreResponse,omitempty" yaml:"createTriageStoreResponse,omitempty"`
}

// Operation wrapper for CreateUser.
// OperationCreateUser was auto-generated from WSDL.
type OperationCreateUser struct {
	CreateUser *CreateUser `xml:"tns:createUser,omitempty" json:"tns:createUser,omitempty" yaml:"tns:createUser,omitempty"`
}

// Operation wrapper for CreateUser.
// OperationCreateUserResponse was auto-generated from WSDL.
type OperationCreateUserResponse struct {
	CreateUserResponse *CreateUserResponse `xml:"createUserResponse,omitempty" json:"createUserResponse,omitempty" yaml:"createUserResponse,omitempty"`
}

// Operation wrapper for DeleteAttribute.
// OperationDeleteAttribute was auto-generated from WSDL.
type OperationDeleteAttribute struct {
	DeleteAttribute *DeleteAttribute `xml:"tns:deleteAttribute,omitempty" json:"tns:deleteAttribute,omitempty" yaml:"tns:deleteAttribute,omitempty"`
}

// Operation wrapper for DeleteAttribute.
// OperationDeleteAttributeResponse was auto-generated from WSDL.
type OperationDeleteAttributeResponse struct {
	DeleteAttributeResponse *DeleteAttributeResponse `xml:"deleteAttributeResponse,omitempty" json:"deleteAttributeResponse,omitempty" yaml:"deleteAttributeResponse,omitempty"`
}

// Operation wrapper for DeleteComponentMap.
// OperationDeleteComponentMap was auto-generated from WSDL.
type OperationDeleteComponentMap struct {
	DeleteComponentMap *DeleteComponentMap `xml:"tns:deleteComponentMap,omitempty" json:"tns:deleteComponentMap,omitempty" yaml:"tns:deleteComponentMap,omitempty"`
}

// Operation wrapper for DeleteComponentMap.
// OperationDeleteComponentMapResponse was auto-generated from
// WSDL.
type OperationDeleteComponentMapResponse struct {
	DeleteComponentMapResponse *DeleteComponentMapResponse `xml:"deleteComponentMapResponse,omitempty" json:"deleteComponentMapResponse,omitempty" yaml:"deleteComponentMapResponse,omitempty"`
}

// Operation wrapper for DeleteGroup.
// OperationDeleteGroup was auto-generated from WSDL.
type OperationDeleteGroup struct {
	DeleteGroup *DeleteGroup `xml:"tns:deleteGroup,omitempty" json:"tns:deleteGroup,omitempty" yaml:"tns:deleteGroup,omitempty"`
}

// Operation wrapper for DeleteGroup.
// OperationDeleteGroupResponse was auto-generated from WSDL.
type OperationDeleteGroupResponse struct {
	DeleteGroupResponse *DeleteGroupResponse `xml:"deleteGroupResponse,omitempty" json:"deleteGroupResponse,omitempty" yaml:"deleteGroupResponse,omitempty"`
}

// Operation wrapper for DeleteLdapConfiguration.
// OperationDeleteLdapConfiguration was auto-generated from WSDL.
type OperationDeleteLdapConfiguration struct {
	DeleteLdapConfiguration *DeleteLdapConfiguration `xml:"tns:deleteLdapConfiguration,omitempty" json:"tns:deleteLdapConfiguration,omitempty" yaml:"tns:deleteLdapConfiguration,omitempty"`
}

// Operation wrapper for DeleteLdapConfiguration.
// OperationDeleteLdapConfigurationResponse was auto-generated
// from WSDL.
type OperationDeleteLdapConfigurationResponse struct {
	DeleteLdapConfigurationResponse *DeleteLdapConfigurationResponse `xml:"deleteLdapConfigurationResponse,omitempty" json:"deleteLdapConfigurationResponse,omitempty" yaml:"deleteLdapConfigurationResponse,omitempty"`
}

// Operation wrapper for DeleteProject.
// OperationDeleteProject was auto-generated from WSDL.
type OperationDeleteProject struct {
	DeleteProject *DeleteProject `xml:"tns:deleteProject,omitempty" json:"tns:deleteProject,omitempty" yaml:"tns:deleteProject,omitempty"`
}

// Operation wrapper for DeleteProject.
// OperationDeleteProjectResponse was auto-generated from WSDL.
type OperationDeleteProjectResponse struct {
	DeleteProjectResponse *DeleteProjectResponse `xml:"deleteProjectResponse,omitempty" json:"deleteProjectResponse,omitempty" yaml:"deleteProjectResponse,omitempty"`
}

// Operation wrapper for DeleteRole.
// OperationDeleteRole was auto-generated from WSDL.
type OperationDeleteRole struct {
	DeleteRole *DeleteRole `xml:"tns:deleteRole,omitempty" json:"tns:deleteRole,omitempty" yaml:"tns:deleteRole,omitempty"`
}

// Operation wrapper for DeleteRole.
// OperationDeleteRoleResponse was auto-generated from WSDL.
type OperationDeleteRoleResponse struct {
	DeleteRoleResponse *DeleteRoleResponse `xml:"deleteRoleResponse,omitempty" json:"deleteRoleResponse,omitempty" yaml:"deleteRoleResponse,omitempty"`
}

// Operation wrapper for DeleteSnapshot.
// OperationDeleteSnapshot was auto-generated from WSDL.
type OperationDeleteSnapshot struct {
	DeleteSnapshot *DeleteSnapshot `xml:"tns:deleteSnapshot,omitempty" json:"tns:deleteSnapshot,omitempty" yaml:"tns:deleteSnapshot,omitempty"`
}

// Operation wrapper for DeleteSnapshot.
// OperationDeleteSnapshotResponse was auto-generated from WSDL.
type OperationDeleteSnapshotResponse struct {
	DeleteSnapshotResponse *DeleteSnapshotResponse `xml:"deleteSnapshotResponse,omitempty" json:"deleteSnapshotResponse,omitempty" yaml:"deleteSnapshotResponse,omitempty"`
}

// Operation wrapper for DeleteStream.
// OperationDeleteStream was auto-generated from WSDL.
type OperationDeleteStream struct {
	DeleteStream *DeleteStream `xml:"tns:deleteStream,omitempty" json:"tns:deleteStream,omitempty" yaml:"tns:deleteStream,omitempty"`
}

// Operation wrapper for DeleteStream.
// OperationDeleteStreamResponse was auto-generated from WSDL.
type OperationDeleteStreamResponse struct {
	DeleteStreamResponse *DeleteStreamResponse `xml:"deleteStreamResponse,omitempty" json:"deleteStreamResponse,omitempty" yaml:"deleteStreamResponse,omitempty"`
}

// Operation wrapper for DeleteTriageStore.
// OperationDeleteTriageStore was auto-generated from WSDL.
type OperationDeleteTriageStore struct {
	DeleteTriageStore *DeleteTriageStore `xml:"tns:deleteTriageStore,omitempty" json:"tns:deleteTriageStore,omitempty" yaml:"tns:deleteTriageStore,omitempty"`
}

// Operation wrapper for DeleteTriageStore.
// OperationDeleteTriageStoreResponse was auto-generated from WSDL.
type OperationDeleteTriageStoreResponse struct {
	DeleteTriageStoreResponse *DeleteTriageStoreResponse `xml:"deleteTriageStoreResponse,omitempty" json:"deleteTriageStoreResponse,omitempty" yaml:"deleteTriageStoreResponse,omitempty"`
}

// Operation wrapper for DeleteUser.
// OperationDeleteUser was auto-generated from WSDL.
type OperationDeleteUser struct {
	DeleteUser *DeleteUser `xml:"tns:deleteUser,omitempty" json:"tns:deleteUser,omitempty" yaml:"tns:deleteUser,omitempty"`
}

// Operation wrapper for DeleteUser.
// OperationDeleteUserResponse was auto-generated from WSDL.
type OperationDeleteUserResponse struct {
	DeleteUserResponse *DeleteUserResponse `xml:"deleteUserResponse,omitempty" json:"deleteUserResponse,omitempty" yaml:"deleteUserResponse,omitempty"`
}

// Operation wrapper for ExecuteNotification.
// OperationExecuteNotification was auto-generated from WSDL.
type OperationExecuteNotification struct {
	ExecuteNotification *ExecuteNotification `xml:"tns:executeNotification,omitempty" json:"tns:executeNotification,omitempty" yaml:"tns:executeNotification,omitempty"`
}

// Operation wrapper for ExecuteNotification.
// OperationExecuteNotificationResponse was auto-generated from
// WSDL.
type OperationExecuteNotificationResponse struct {
	ExecuteNotificationResponse *ExecuteNotificationResponse `xml:"executeNotificationResponse,omitempty" json:"executeNotificationResponse,omitempty" yaml:"executeNotificationResponse,omitempty"`
}

// Operation wrapper for GetAllLdapConfigurations.
// OperationGetAllLdapConfigurations was auto-generated from WSDL.
type OperationGetAllLdapConfigurations struct {
	GetAllLdapConfigurations *GetAllLdapConfigurations `xml:"tns:getAllLdapConfigurations,omitempty" json:"tns:getAllLdapConfigurations,omitempty" yaml:"tns:getAllLdapConfigurations,omitempty"`
}

// Operation wrapper for GetAllLdapConfigurations.
// OperationGetAllLdapConfigurationsResponse was auto-generated
// from WSDL.
type OperationGetAllLdapConfigurationsResponse struct {
	GetAllLdapConfigurationsResponse *GetAllLdapConfigurationsResponse `xml:"getAllLdapConfigurationsResponse,omitempty" json:"getAllLdapConfigurationsResponse,omitempty" yaml:"getAllLdapConfigurationsResponse,omitempty"`
}

// Operation wrapper for GetAllPermissions.
// OperationGetAllPermissions was auto-generated from WSDL.
type OperationGetAllPermissions struct {
	GetAllPermissions *GetAllPermissions `xml:"tns:getAllPermissions,omitempty" json:"tns:getAllPermissions,omitempty" yaml:"tns:getAllPermissions,omitempty"`
}

// Operation wrapper for GetAllPermissions.
// OperationGetAllPermissionsResponse was auto-generated from WSDL.
type OperationGetAllPermissionsResponse struct {
	GetAllPermissionsResponse *GetAllPermissionsResponse `xml:"getAllPermissionsResponse,omitempty" json:"getAllPermissionsResponse,omitempty" yaml:"getAllPermissionsResponse,omitempty"`
}

// Operation wrapper for GetAllRoles.
// OperationGetAllRoles was auto-generated from WSDL.
type OperationGetAllRoles struct {
	GetAllRoles *GetAllRoles `xml:"tns:getAllRoles,omitempty" json:"tns:getAllRoles,omitempty" yaml:"tns:getAllRoles,omitempty"`
}

// Operation wrapper for GetAllRoles.
// OperationGetAllRolesResponse was auto-generated from WSDL.
type OperationGetAllRolesResponse struct {
	GetAllRolesResponse *GetAllRolesResponse `xml:"getAllRolesResponse,omitempty" json:"getAllRolesResponse,omitempty" yaml:"getAllRolesResponse,omitempty"`
}

// Operation wrapper for GetArchitectureAnalysisConfiguration.
// OperationGetArchitectureAnalysisConfiguration was auto-generated
// from WSDL.
type OperationGetArchitectureAnalysisConfiguration struct {
	GetArchitectureAnalysisConfiguration *GetArchitectureAnalysisConfiguration `xml:"tns:getArchitectureAnalysisConfiguration,omitempty" json:"tns:getArchitectureAnalysisConfiguration,omitempty" yaml:"tns:getArchitectureAnalysisConfiguration,omitempty"`
}

// Operation wrapper for GetArchitectureAnalysisConfiguration.
// OperationGetArchitectureAnalysisConfigurationResponse was auto-generated
// from WSDL.
type OperationGetArchitectureAnalysisConfigurationResponse struct {
	GetArchitectureAnalysisConfigurationResponse *GetArchitectureAnalysisConfigurationResponse `xml:"getArchitectureAnalysisConfigurationResponse,omitempty" json:"getArchitectureAnalysisConfigurationResponse,omitempty" yaml:"getArchitectureAnalysisConfigurationResponse,omitempty"`
}

// Operation wrapper for GetAttribute.
// OperationGetAttribute was auto-generated from WSDL.
type OperationGetAttribute struct {
	GetAttribute *GetAttribute `xml:"tns:getAttribute,omitempty" json:"tns:getAttribute,omitempty" yaml:"tns:getAttribute,omitempty"`
}

// Operation wrapper for GetAttribute.
// OperationGetAttributeResponse was auto-generated from WSDL.
type OperationGetAttributeResponse struct {
	GetAttributeResponse *GetAttributeResponse `xml:"getAttributeResponse,omitempty" json:"getAttributeResponse,omitempty" yaml:"getAttributeResponse,omitempty"`
}

// Operation wrapper for GetAttributes.
// OperationGetAttributes was auto-generated from WSDL.
type OperationGetAttributes struct {
	GetAttributes *GetAttributes `xml:"tns:getAttributes,omitempty" json:"tns:getAttributes,omitempty" yaml:"tns:getAttributes,omitempty"`
}

// Operation wrapper for GetAttributes.
// OperationGetAttributesResponse was auto-generated from WSDL.
type OperationGetAttributesResponse struct {
	GetAttributesResponse *GetAttributesResponse `xml:"getAttributesResponse,omitempty" json:"getAttributesResponse,omitempty" yaml:"getAttributesResponse,omitempty"`
}

// Operation wrapper for GetBackupConfiguration.
// OperationGetBackupConfiguration was auto-generated from WSDL.
type OperationGetBackupConfiguration struct {
	GetBackupConfiguration *GetBackupConfiguration `xml:"tns:getBackupConfiguration,omitempty" json:"tns:getBackupConfiguration,omitempty" yaml:"tns:getBackupConfiguration,omitempty"`
}

// Operation wrapper for GetBackupConfiguration.
// OperationGetBackupConfigurationResponse was auto-generated from
// WSDL.
type OperationGetBackupConfigurationResponse struct {
	GetBackupConfigurationResponse *GetBackupConfigurationResponse `xml:"getBackupConfigurationResponse,omitempty" json:"getBackupConfigurationResponse,omitempty" yaml:"getBackupConfigurationResponse,omitempty"`
}

// Operation wrapper for GetCategoryNames.
// OperationGetCategoryNames was auto-generated from WSDL.
type OperationGetCategoryNames struct {
	GetCategoryNames *GetCategoryNames `xml:"tns:getCategoryNames,omitempty" json:"tns:getCategoryNames,omitempty" yaml:"tns:getCategoryNames,omitempty"`
}

// Operation wrapper for GetCategoryNames.
// OperationGetCategoryNamesResponse was auto-generated from WSDL.
type OperationGetCategoryNamesResponse struct {
	GetCategoryNamesResponse *GetCategoryNamesResponse `xml:"getCategoryNamesResponse,omitempty" json:"getCategoryNamesResponse,omitempty" yaml:"getCategoryNamesResponse,omitempty"`
}

// Operation wrapper for GetCheckerNames.
// OperationGetCheckerNames was auto-generated from WSDL.
type OperationGetCheckerNames struct {
	GetCheckerNames *GetCheckerNames `xml:"tns:getCheckerNames,omitempty" json:"tns:getCheckerNames,omitempty" yaml:"tns:getCheckerNames,omitempty"`
}

// Operation wrapper for GetCheckerNames.
// OperationGetCheckerNamesResponse was auto-generated from WSDL.
type OperationGetCheckerNamesResponse struct {
	GetCheckerNamesResponse *GetCheckerNamesResponse `xml:"getCheckerNamesResponse,omitempty" json:"getCheckerNamesResponse,omitempty" yaml:"getCheckerNamesResponse,omitempty"`
}

// Operation wrapper for GetCommitState.
// OperationGetCommitState was auto-generated from WSDL.
type OperationGetCommitState struct {
	GetCommitState *GetCommitState `xml:"tns:getCommitState,omitempty" json:"tns:getCommitState,omitempty" yaml:"tns:getCommitState,omitempty"`
}

// Operation wrapper for GetCommitState.
// OperationGetCommitStateResponse was auto-generated from WSDL.
type OperationGetCommitStateResponse struct {
	GetCommitStateResponse *GetCommitStateResponse `xml:"getCommitStateResponse,omitempty" json:"getCommitStateResponse,omitempty" yaml:"getCommitStateResponse,omitempty"`
}

// Operation wrapper for GetComponent.
// OperationGetComponent was auto-generated from WSDL.
type OperationGetComponent struct {
	GetComponent *GetComponent `xml:"tns:getComponent,omitempty" json:"tns:getComponent,omitempty" yaml:"tns:getComponent,omitempty"`
}

// Operation wrapper for GetComponent.
// OperationGetComponentResponse was auto-generated from WSDL.
type OperationGetComponentResponse struct {
	GetComponentResponse *GetComponentResponse `xml:"getComponentResponse,omitempty" json:"getComponentResponse,omitempty" yaml:"getComponentResponse,omitempty"`
}

// Operation wrapper for GetComponentMaps.
// OperationGetComponentMaps was auto-generated from WSDL.
type OperationGetComponentMaps struct {
	GetComponentMaps *GetComponentMaps `xml:"tns:getComponentMaps,omitempty" json:"tns:getComponentMaps,omitempty" yaml:"tns:getComponentMaps,omitempty"`
}

// Operation wrapper for GetComponentMaps.
// OperationGetComponentMapsResponse was auto-generated from WSDL.
type OperationGetComponentMapsResponse struct {
	GetComponentMapsResponse *GetComponentMapsResponse `xml:"getComponentMapsResponse,omitempty" json:"getComponentMapsResponse,omitempty" yaml:"getComponentMapsResponse,omitempty"`
}

// Operation wrapper for GetDefectStatuses.
// OperationGetDefectStatuses was auto-generated from WSDL.
type OperationGetDefectStatuses struct {
	GetDefectStatuses *GetDefectStatuses `xml:"tns:getDefectStatuses,omitempty" json:"tns:getDefectStatuses,omitempty" yaml:"tns:getDefectStatuses,omitempty"`
}

// Operation wrapper for GetDefectStatuses.
// OperationGetDefectStatusesResponse was auto-generated from WSDL.
type OperationGetDefectStatusesResponse struct {
	GetDefectStatusesResponse *GetDefectStatusesResponse `xml:"getDefectStatusesResponse,omitempty" json:"getDefectStatusesResponse,omitempty" yaml:"getDefectStatusesResponse,omitempty"`
}

// Operation wrapper for GetDeleteSnapshotJobInfo.
// OperationGetDeleteSnapshotJobInfo was auto-generated from WSDL.
type OperationGetDeleteSnapshotJobInfo struct {
	GetDeleteSnapshotJobInfo *GetDeleteSnapshotJobInfo `xml:"tns:getDeleteSnapshotJobInfo,omitempty" json:"tns:getDeleteSnapshotJobInfo,omitempty" yaml:"tns:getDeleteSnapshotJobInfo,omitempty"`
}

// Operation wrapper for GetDeleteSnapshotJobInfo.
// OperationGetDeleteSnapshotJobInfoResponse was auto-generated
// from WSDL.
type OperationGetDeleteSnapshotJobInfoResponse struct {
	GetDeleteSnapshotJobInfoResponse *GetDeleteSnapshotJobInfoResponse `xml:"getDeleteSnapshotJobInfoResponse,omitempty" json:"getDeleteSnapshotJobInfoResponse,omitempty" yaml:"getDeleteSnapshotJobInfoResponse,omitempty"`
}

// Operation wrapper for GetDeveloperStreamsProjects.
// OperationGetDeveloperStreamsProjects was auto-generated from
// WSDL.
type OperationGetDeveloperStreamsProjects struct {
	GetDeveloperStreamsProjects *GetDeveloperStreamsProjects `xml:"tns:getDeveloperStreamsProjects,omitempty" json:"tns:getDeveloperStreamsProjects,omitempty" yaml:"tns:getDeveloperStreamsProjects,omitempty"`
}

// Operation wrapper for GetDeveloperStreamsProjects.
// OperationGetDeveloperStreamsProjectsResponse was auto-generated
// from WSDL.
type OperationGetDeveloperStreamsProjectsResponse struct {
	GetDeveloperStreamsProjectsResponse *GetDeveloperStreamsProjectsResponse `xml:"getDeveloperStreamsProjectsResponse,omitempty" json:"getDeveloperStreamsProjectsResponse,omitempty" yaml:"getDeveloperStreamsProjectsResponse,omitempty"`
}

// Operation wrapper for GetGroup.
// OperationGetGroup was auto-generated from WSDL.
type OperationGetGroup struct {
	GetGroup *GetGroup `xml:"tns:getGroup,omitempty" json:"tns:getGroup,omitempty" yaml:"tns:getGroup,omitempty"`
}

// Operation wrapper for GetGroup.
// OperationGetGroupResponse was auto-generated from WSDL.
type OperationGetGroupResponse struct {
	GetGroupResponse *GetGroupResponse `xml:"getGroupResponse,omitempty" json:"getGroupResponse,omitempty" yaml:"getGroupResponse,omitempty"`
}

// Operation wrapper for GetGroups.
// OperationGetGroups was auto-generated from WSDL.
type OperationGetGroups struct {
	GetGroups *GetGroups `xml:"tns:getGroups,omitempty" json:"tns:getGroups,omitempty" yaml:"tns:getGroups,omitempty"`
}

// Operation wrapper for GetGroups.
// OperationGetGroupsResponse was auto-generated from WSDL.
type OperationGetGroupsResponse struct {
	GetGroupsResponse *GetGroupsResponse `xml:"getGroupsResponse,omitempty" json:"getGroupsResponse,omitempty" yaml:"getGroupsResponse,omitempty"`
}

// Operation wrapper for GetLastUpdateTimes.
// OperationGetLastUpdateTimes was auto-generated from WSDL.
type OperationGetLastUpdateTimes struct {
	GetLastUpdateTimes *GetLastUpdateTimes `xml:"tns:getLastUpdateTimes,omitempty" json:"tns:getLastUpdateTimes,omitempty" yaml:"tns:getLastUpdateTimes,omitempty"`
}

// Operation wrapper for GetLastUpdateTimes.
// OperationGetLastUpdateTimesResponse was auto-generated from
// WSDL.
type OperationGetLastUpdateTimesResponse struct {
	GetLastUpdateTimesResponse *GetLastUpdateTimesResponse `xml:"getLastUpdateTimesResponse,omitempty" json:"getLastUpdateTimesResponse,omitempty" yaml:"getLastUpdateTimesResponse,omitempty"`
}

// Operation wrapper for GetLdapServerDomains.
// OperationGetLdapServerDomains was auto-generated from WSDL.
type OperationGetLdapServerDomains struct {
	GetLdapServerDomains *GetLdapServerDomains `xml:"tns:getLdapServerDomains,omitempty" json:"tns:getLdapServerDomains,omitempty" yaml:"tns:getLdapServerDomains,omitempty"`
}

// Operation wrapper for GetLdapServerDomains.
// OperationGetLdapServerDomainsResponse was auto-generated from
// WSDL.
type OperationGetLdapServerDomainsResponse struct {
	GetLdapServerDomainsResponse *GetLdapServerDomainsResponse `xml:"getLdapServerDomainsResponse,omitempty" json:"getLdapServerDomainsResponse,omitempty" yaml:"getLdapServerDomainsResponse,omitempty"`
}

// Operation wrapper for GetLicenseConfiguration.
// OperationGetLicenseConfiguration was auto-generated from WSDL.
type OperationGetLicenseConfiguration struct {
	GetLicenseConfiguration *GetLicenseConfiguration `xml:"tns:getLicenseConfiguration,omitempty" json:"tns:getLicenseConfiguration,omitempty" yaml:"tns:getLicenseConfiguration,omitempty"`
}

// Operation wrapper for GetLicenseConfiguration.
// OperationGetLicenseConfigurationResponse was auto-generated
// from WSDL.
type OperationGetLicenseConfigurationResponse struct {
	GetLicenseConfigurationResponse *GetLicenseConfigurationResponse `xml:"getLicenseConfigurationResponse,omitempty" json:"getLicenseConfigurationResponse,omitempty" yaml:"getLicenseConfigurationResponse,omitempty"`
}

// Operation wrapper for GetLicenseState.
// OperationGetLicenseState was auto-generated from WSDL.
type OperationGetLicenseState struct {
	GetLicenseState *GetLicenseState `xml:"tns:getLicenseState,omitempty" json:"tns:getLicenseState,omitempty" yaml:"tns:getLicenseState,omitempty"`
}

// Operation wrapper for GetLicenseState.
// OperationGetLicenseStateResponse was auto-generated from WSDL.
type OperationGetLicenseStateResponse struct {
	GetLicenseStateResponse *GetLicenseStateResponse `xml:"getLicenseStateResponse,omitempty" json:"getLicenseStateResponse,omitempty" yaml:"getLicenseStateResponse,omitempty"`
}

// Operation wrapper for GetLoggingConfiguration.
// OperationGetLoggingConfiguration was auto-generated from WSDL.
type OperationGetLoggingConfiguration struct {
	GetLoggingConfiguration *GetLoggingConfiguration `xml:"tns:getLoggingConfiguration,omitempty" json:"tns:getLoggingConfiguration,omitempty" yaml:"tns:getLoggingConfiguration,omitempty"`
}

// Operation wrapper for GetLoggingConfiguration.
// OperationGetLoggingConfigurationResponse was auto-generated
// from WSDL.
type OperationGetLoggingConfigurationResponse struct {
	GetLoggingConfigurationResponse *GetLoggingConfigurationResponse `xml:"getLoggingConfigurationResponse,omitempty" json:"getLoggingConfigurationResponse,omitempty" yaml:"getLoggingConfigurationResponse,omitempty"`
}

// Operation wrapper for GetMessageOfTheDay.
// OperationGetMessageOfTheDay was auto-generated from WSDL.
type OperationGetMessageOfTheDay struct {
	GetMessageOfTheDay *GetMessageOfTheDay `xml:"tns:getMessageOfTheDay,omitempty" json:"tns:getMessageOfTheDay,omitempty" yaml:"tns:getMessageOfTheDay,omitempty"`
}

// Operation wrapper for GetMessageOfTheDay.
// OperationGetMessageOfTheDayResponse was auto-generated from
// WSDL.
type OperationGetMessageOfTheDayResponse struct {
	GetMessageOfTheDayResponse *GetMessageOfTheDayResponse `xml:"getMessageOfTheDayResponse,omitempty" json:"getMessageOfTheDayResponse,omitempty" yaml:"getMessageOfTheDayResponse,omitempty"`
}

// Operation wrapper for GetOutputFileForSnapshot.
// OperationGetOutputFileForSnapshot was auto-generated from WSDL.
type OperationGetOutputFileForSnapshot struct {
	GetOutputFileForSnapshot *GetOutputFileForSnapshot `xml:"tns:getOutputFileForSnapshot,omitempty" json:"tns:getOutputFileForSnapshot,omitempty" yaml:"tns:getOutputFileForSnapshot,omitempty"`
}

// Operation wrapper for GetOutputFileForSnapshot.
// OperationGetOutputFileForSnapshotResponse was auto-generated
// from WSDL.
type OperationGetOutputFileForSnapshotResponse struct {
	GetOutputFileForSnapshotResponse *GetOutputFileForSnapshotResponse `xml:"getOutputFileForSnapshotResponse,omitempty" json:"getOutputFileForSnapshotResponse,omitempty" yaml:"getOutputFileForSnapshotResponse,omitempty"`
}

// Operation wrapper for GetProjects.
// OperationGetProjects was auto-generated from WSDL.
type OperationGetProjects struct {
	GetProjects *GetProjects `xml:"tns:getProjects,omitempty" json:"tns:getProjects,omitempty" yaml:"tns:getProjects,omitempty"`
}

// Operation wrapper for GetProjects.
// OperationGetProjectsResponse was auto-generated from WSDL.
type OperationGetProjectsResponse struct {
	GetProjectsResponse *GetProjectsResponse `xml:"getProjectsResponse,omitempty" json:"getProjectsResponse,omitempty" yaml:"getProjectsResponse,omitempty"`
}

// Operation wrapper for GetRole.
// OperationGetRole was auto-generated from WSDL.
type OperationGetRole struct {
	GetRole *GetRole `xml:"tns:getRole,omitempty" json:"tns:getRole,omitempty" yaml:"tns:getRole,omitempty"`
}

// Operation wrapper for GetRole.
// OperationGetRoleResponse was auto-generated from WSDL.
type OperationGetRoleResponse struct {
	GetRoleResponse *GetRoleResponse `xml:"getRoleResponse,omitempty" json:"getRoleResponse,omitempty" yaml:"getRoleResponse,omitempty"`
}

// Operation wrapper for GetServerTime.
// OperationGetServerTime was auto-generated from WSDL.
type OperationGetServerTime struct {
	GetServerTime *GetServerTime `xml:"tns:getServerTime,omitempty" json:"tns:getServerTime,omitempty" yaml:"tns:getServerTime,omitempty"`
}

// Operation wrapper for GetServerTime.
// OperationGetServerTimeResponse was auto-generated from WSDL.
type OperationGetServerTimeResponse struct {
	GetServerTimeResponse *GetServerTimeResponse `xml:"getServerTimeResponse,omitempty" json:"getServerTimeResponse,omitempty" yaml:"getServerTimeResponse,omitempty"`
}

// Operation wrapper for GetSignInConfiguration.
// OperationGetSignInConfiguration was auto-generated from WSDL.
type OperationGetSignInConfiguration struct {
	GetSignInConfiguration *GetSignInConfiguration `xml:"tns:getSignInConfiguration,omitempty" json:"tns:getSignInConfiguration,omitempty" yaml:"tns:getSignInConfiguration,omitempty"`
}

// Operation wrapper for GetSignInConfiguration.
// OperationGetSignInConfigurationResponse was auto-generated from
// WSDL.
type OperationGetSignInConfigurationResponse struct {
	GetSignInConfigurationResponse *GetSignInConfigurationResponse `xml:"getSignInConfigurationResponse,omitempty" json:"getSignInConfigurationResponse,omitempty" yaml:"getSignInConfigurationResponse,omitempty"`
}

// Operation wrapper for GetSkeletonizationConfiguration.
// OperationGetSkeletonizationConfiguration was auto-generated
// from WSDL.
type OperationGetSkeletonizationConfiguration struct {
	GetSkeletonizationConfiguration *GetSkeletonizationConfiguration `xml:"tns:getSkeletonizationConfiguration,omitempty" json:"tns:getSkeletonizationConfiguration,omitempty" yaml:"tns:getSkeletonizationConfiguration,omitempty"`
}

// Operation wrapper for GetSkeletonizationConfiguration.
// OperationGetSkeletonizationConfigurationResponse was auto-generated
// from WSDL.
type OperationGetSkeletonizationConfigurationResponse struct {
	GetSkeletonizationConfigurationResponse *GetSkeletonizationConfigurationResponse `xml:"getSkeletonizationConfigurationResponse,omitempty" json:"getSkeletonizationConfigurationResponse,omitempty" yaml:"getSkeletonizationConfigurationResponse,omitempty"`
}

// Operation wrapper for GetSnapshotInformation.
// OperationGetSnapshotInformation was auto-generated from WSDL.
type OperationGetSnapshotInformation struct {
	GetSnapshotInformation *GetSnapshotInformation `xml:"tns:getSnapshotInformation,omitempty" json:"tns:getSnapshotInformation,omitempty" yaml:"tns:getSnapshotInformation,omitempty"`
}

// Operation wrapper for GetSnapshotInformation.
// OperationGetSnapshotInformationResponse was auto-generated from
// WSDL.
type OperationGetSnapshotInformationResponse struct {
	GetSnapshotInformationResponse *GetSnapshotInformationResponse `xml:"getSnapshotInformationResponse,omitempty" json:"getSnapshotInformationResponse,omitempty" yaml:"getSnapshotInformationResponse,omitempty"`
}

// Operation wrapper for GetSnapshotPurgeDetails.
// OperationGetSnapshotPurgeDetails was auto-generated from WSDL.
type OperationGetSnapshotPurgeDetails struct {
	GetSnapshotPurgeDetails *GetSnapshotPurgeDetails `xml:"tns:getSnapshotPurgeDetails,omitempty" json:"tns:getSnapshotPurgeDetails,omitempty" yaml:"tns:getSnapshotPurgeDetails,omitempty"`
}

// Operation wrapper for GetSnapshotPurgeDetails.
// OperationGetSnapshotPurgeDetailsResponse was auto-generated
// from WSDL.
type OperationGetSnapshotPurgeDetailsResponse struct {
	GetSnapshotPurgeDetailsResponse *GetSnapshotPurgeDetailsResponse `xml:"getSnapshotPurgeDetailsResponse,omitempty" json:"getSnapshotPurgeDetailsResponse,omitempty" yaml:"getSnapshotPurgeDetailsResponse,omitempty"`
}

// Operation wrapper for GetSnapshotsForStream.
// OperationGetSnapshotsForStream was auto-generated from WSDL.
type OperationGetSnapshotsForStream struct {
	GetSnapshotsForStream *GetSnapshotsForStream `xml:"tns:getSnapshotsForStream,omitempty" json:"tns:getSnapshotsForStream,omitempty" yaml:"tns:getSnapshotsForStream,omitempty"`
}

// Operation wrapper for GetSnapshotsForStream.
// OperationGetSnapshotsForStreamResponse was auto-generated from
// WSDL.
type OperationGetSnapshotsForStreamResponse struct {
	GetSnapshotsForStreamResponse *GetSnapshotsForStreamResponse `xml:"getSnapshotsForStreamResponse,omitempty" json:"getSnapshotsForStreamResponse,omitempty" yaml:"getSnapshotsForStreamResponse,omitempty"`
}

// Operation wrapper for GetStreams.
// OperationGetStreams was auto-generated from WSDL.
type OperationGetStreams struct {
	GetStreams *GetStreams `xml:"tns:getStreams,omitempty" json:"tns:getStreams,omitempty" yaml:"tns:getStreams,omitempty"`
}

// Operation wrapper for GetStreams.
// OperationGetStreamsResponse was auto-generated from WSDL.
type OperationGetStreamsResponse struct {
	GetStreamsResponse *GetStreamsResponse `xml:"getStreamsResponse,omitempty" json:"getStreamsResponse,omitempty" yaml:"getStreamsResponse,omitempty"`
}

// Operation wrapper for GetSystemConfig.
// OperationGetSystemConfig was auto-generated from WSDL.
type OperationGetSystemConfig struct {
	GetSystemConfig *GetSystemConfig `xml:"tns:getSystemConfig,omitempty" json:"tns:getSystemConfig,omitempty" yaml:"tns:getSystemConfig,omitempty"`
}

// Operation wrapper for GetSystemConfig.
// OperationGetSystemConfigResponse was auto-generated from WSDL.
type OperationGetSystemConfigResponse struct {
	GetSystemConfigResponse *GetSystemConfigResponse `xml:"getSystemConfigResponse,omitempty" json:"getSystemConfigResponse,omitempty" yaml:"getSystemConfigResponse,omitempty"`
}

// Operation wrapper for GetTriageStores.
// OperationGetTriageStores was auto-generated from WSDL.
type OperationGetTriageStores struct {
	GetTriageStores *GetTriageStores `xml:"tns:getTriageStores,omitempty" json:"tns:getTriageStores,omitempty" yaml:"tns:getTriageStores,omitempty"`
}

// Operation wrapper for GetTriageStores.
// OperationGetTriageStoresResponse was auto-generated from WSDL.
type OperationGetTriageStoresResponse struct {
	GetTriageStoresResponse *GetTriageStoresResponse `xml:"getTriageStoresResponse,omitempty" json:"getTriageStoresResponse,omitempty" yaml:"getTriageStoresResponse,omitempty"`
}

// Operation wrapper for GetTypeNames.
// OperationGetTypeNames was auto-generated from WSDL.
type OperationGetTypeNames struct {
	GetTypeNames *GetTypeNames `xml:"tns:getTypeNames,omitempty" json:"tns:getTypeNames,omitempty" yaml:"tns:getTypeNames,omitempty"`
}

// Operation wrapper for GetTypeNames.
// OperationGetTypeNamesResponse was auto-generated from WSDL.
type OperationGetTypeNamesResponse struct {
	GetTypeNamesResponse *GetTypeNamesResponse `xml:"getTypeNamesResponse,omitempty" json:"getTypeNamesResponse,omitempty" yaml:"getTypeNamesResponse,omitempty"`
}

// Operation wrapper for GetUser.
// OperationGetUser was auto-generated from WSDL.
type OperationGetUser struct {
	GetUser *GetUser `xml:"tns:getUser,omitempty" json:"tns:getUser,omitempty" yaml:"tns:getUser,omitempty"`
}

// Operation wrapper for GetUser.
// OperationGetUserResponse was auto-generated from WSDL.
type OperationGetUserResponse struct {
	GetUserResponse *GetUserResponse `xml:"getUserResponse,omitempty" json:"getUserResponse,omitempty" yaml:"getUserResponse,omitempty"`
}

// Operation wrapper for GetUsers.
// OperationGetUsers was auto-generated from WSDL.
type OperationGetUsers struct {
	GetUsers *GetUsers `xml:"tns:getUsers,omitempty" json:"tns:getUsers,omitempty" yaml:"tns:getUsers,omitempty"`
}

// Operation wrapper for GetUsers.
// OperationGetUsersResponse was auto-generated from WSDL.
type OperationGetUsersResponse struct {
	GetUsersResponse *GetUsersResponse `xml:"getUsersResponse,omitempty" json:"getUsersResponse,omitempty" yaml:"getUsersResponse,omitempty"`
}

// Operation wrapper for GetVersion.
// OperationGetVersion was auto-generated from WSDL.
type OperationGetVersion struct {
	GetVersion *GetVersion `xml:"tns:getVersion,omitempty" json:"tns:getVersion,omitempty" yaml:"tns:getVersion,omitempty"`
}

// Operation wrapper for GetVersion.
// OperationGetVersionResponse was auto-generated from WSDL.
type OperationGetVersionResponse struct {
	GetVersionResponse *GetVersionResponse `xml:"getVersionResponse,omitempty" json:"getVersionResponse,omitempty" yaml:"getVersionResponse,omitempty"`
}

// Operation wrapper for ImportLicense.
// OperationImportLicense was auto-generated from WSDL.
type OperationImportLicense struct {
	ImportLicense *ImportLicense `xml:"tns:importLicense,omitempty" json:"tns:importLicense,omitempty" yaml:"tns:importLicense,omitempty"`
}

// Operation wrapper for ImportLicense.
// OperationImportLicenseResponse was auto-generated from WSDL.
type OperationImportLicenseResponse struct {
	ImportLicenseResponse *ImportLicenseResponse `xml:"importLicenseResponse,omitempty" json:"importLicenseResponse,omitempty" yaml:"importLicenseResponse,omitempty"`
}

// Operation wrapper for MergeTriageStores.
// OperationMergeTriageStores was auto-generated from WSDL.
type OperationMergeTriageStores struct {
	MergeTriageStores *MergeTriageStores `xml:"tns:mergeTriageStores,omitempty" json:"tns:mergeTriageStores,omitempty" yaml:"tns:mergeTriageStores,omitempty"`
}

// Operation wrapper for MergeTriageStores.
// OperationMergeTriageStoresResponse was auto-generated from WSDL.
type OperationMergeTriageStoresResponse struct {
	MergeTriageStoresResponse *MergeTriageStoresResponse `xml:"mergeTriageStoresResponse,omitempty" json:"mergeTriageStoresResponse,omitempty" yaml:"mergeTriageStoresResponse,omitempty"`
}

// Operation wrapper for Notify.
// OperationNotify was auto-generated from WSDL.
type OperationNotify struct {
	Notify *Notify `xml:"tns:notify,omitempty" json:"tns:notify,omitempty" yaml:"tns:notify,omitempty"`
}

// Operation wrapper for Notify.
// OperationNotifyResponse was auto-generated from WSDL.
type OperationNotifyResponse struct {
	NotifyResponse *NotifyResponse `xml:"notifyResponse,omitempty" json:"notifyResponse,omitempty" yaml:"notifyResponse,omitempty"`
}

// Operation wrapper for SetAcceptingNewCommits.
// OperationSetAcceptingNewCommits was auto-generated from WSDL.
type OperationSetAcceptingNewCommits struct {
	SetAcceptingNewCommits *SetAcceptingNewCommits `xml:"tns:setAcceptingNewCommits,omitempty" json:"tns:setAcceptingNewCommits,omitempty" yaml:"tns:setAcceptingNewCommits,omitempty"`
}

// Operation wrapper for SetAcceptingNewCommits.
// OperationSetAcceptingNewCommitsResponse was auto-generated from
// WSDL.
type OperationSetAcceptingNewCommitsResponse struct {
	SetAcceptingNewCommitsResponse *SetAcceptingNewCommitsResponse `xml:"setAcceptingNewCommitsResponse,omitempty" json:"setAcceptingNewCommitsResponse,omitempty" yaml:"setAcceptingNewCommitsResponse,omitempty"`
}

// Operation wrapper for SetArchitectureAnalysisConfiguration.
// OperationSetArchitectureAnalysisConfiguration was auto-generated
// from WSDL.
type OperationSetArchitectureAnalysisConfiguration struct {
	SetArchitectureAnalysisConfiguration *SetArchitectureAnalysisConfiguration `xml:"tns:setArchitectureAnalysisConfiguration,omitempty" json:"tns:setArchitectureAnalysisConfiguration,omitempty" yaml:"tns:setArchitectureAnalysisConfiguration,omitempty"`
}

// Operation wrapper for SetArchitectureAnalysisConfiguration.
// OperationSetArchitectureAnalysisConfigurationResponse was auto-generated
// from WSDL.
type OperationSetArchitectureAnalysisConfigurationResponse struct {
	SetArchitectureAnalysisConfigurationResponse *SetArchitectureAnalysisConfigurationResponse `xml:"setArchitectureAnalysisConfigurationResponse,omitempty" json:"setArchitectureAnalysisConfigurationResponse,omitempty" yaml:"setArchitectureAnalysisConfigurationResponse,omitempty"`
}

// Operation wrapper for SetBackupConfiguration.
// OperationSetBackupConfiguration was auto-generated from WSDL.
type OperationSetBackupConfiguration struct {
	SetBackupConfiguration *SetBackupConfiguration `xml:"tns:setBackupConfiguration,omitempty" json:"tns:setBackupConfiguration,omitempty" yaml:"tns:setBackupConfiguration,omitempty"`
}

// Operation wrapper for SetBackupConfiguration.
// OperationSetBackupConfigurationResponse was auto-generated from
// WSDL.
type OperationSetBackupConfigurationResponse struct {
	SetBackupConfigurationResponse *SetBackupConfigurationResponse `xml:"setBackupConfigurationResponse,omitempty" json:"setBackupConfigurationResponse,omitempty" yaml:"setBackupConfigurationResponse,omitempty"`
}

// Operation wrapper for SetLoggingConfiguration.
// OperationSetLoggingConfiguration was auto-generated from WSDL.
type OperationSetLoggingConfiguration struct {
	SetLoggingConfiguration *SetLoggingConfiguration `xml:"tns:setLoggingConfiguration,omitempty" json:"tns:setLoggingConfiguration,omitempty" yaml:"tns:setLoggingConfiguration,omitempty"`
}

// Operation wrapper for SetLoggingConfiguration.
// OperationSetLoggingConfigurationResponse was auto-generated
// from WSDL.
type OperationSetLoggingConfigurationResponse struct {
	SetLoggingConfigurationResponse *SetLoggingConfigurationResponse `xml:"setLoggingConfigurationResponse,omitempty" json:"setLoggingConfigurationResponse,omitempty" yaml:"setLoggingConfigurationResponse,omitempty"`
}

// Operation wrapper for SetMessageOfTheDay.
// OperationSetMessageOfTheDay was auto-generated from WSDL.
type OperationSetMessageOfTheDay struct {
	SetMessageOfTheDay *SetMessageOfTheDay `xml:"tns:setMessageOfTheDay,omitempty" json:"tns:setMessageOfTheDay,omitempty" yaml:"tns:setMessageOfTheDay,omitempty"`
}

// Operation wrapper for SetMessageOfTheDay.
// OperationSetMessageOfTheDayResponse was auto-generated from
// WSDL.
type OperationSetMessageOfTheDayResponse struct {
	SetMessageOfTheDayResponse *SetMessageOfTheDayResponse `xml:"setMessageOfTheDayResponse,omitempty" json:"setMessageOfTheDayResponse,omitempty" yaml:"setMessageOfTheDayResponse,omitempty"`
}

// Operation wrapper for SetSkeletonizationConfiguration.
// OperationSetSkeletonizationConfiguration was auto-generated
// from WSDL.
type OperationSetSkeletonizationConfiguration struct {
	SetSkeletonizationConfiguration *SetSkeletonizationConfiguration `xml:"tns:setSkeletonizationConfiguration,omitempty" json:"tns:setSkeletonizationConfiguration,omitempty" yaml:"tns:setSkeletonizationConfiguration,omitempty"`
}

// Operation wrapper for SetSkeletonizationConfiguration.
// OperationSetSkeletonizationConfigurationResponse was auto-generated
// from WSDL.
type OperationSetSkeletonizationConfigurationResponse struct {
	SetSkeletonizationConfigurationResponse *SetSkeletonizationConfigurationResponse `xml:"setSkeletonizationConfigurationResponse,omitempty" json:"setSkeletonizationConfigurationResponse,omitempty" yaml:"setSkeletonizationConfigurationResponse,omitempty"`
}

// Operation wrapper for SetSnapshotPurgeDetails.
// OperationSetSnapshotPurgeDetails was auto-generated from WSDL.
type OperationSetSnapshotPurgeDetails struct {
	SetSnapshotPurgeDetails *SetSnapshotPurgeDetails `xml:"tns:setSnapshotPurgeDetails,omitempty" json:"tns:setSnapshotPurgeDetails,omitempty" yaml:"tns:setSnapshotPurgeDetails,omitempty"`
}

// Operation wrapper for SetSnapshotPurgeDetails.
// OperationSetSnapshotPurgeDetailsResponse was auto-generated
// from WSDL.
type OperationSetSnapshotPurgeDetailsResponse struct {
	SetSnapshotPurgeDetailsResponse *SetSnapshotPurgeDetailsResponse `xml:"setSnapshotPurgeDetailsResponse,omitempty" json:"setSnapshotPurgeDetailsResponse,omitempty" yaml:"setSnapshotPurgeDetailsResponse,omitempty"`
}

// Operation wrapper for UpdateAttribute.
// OperationUpdateAttribute was auto-generated from WSDL.
type OperationUpdateAttribute struct {
	UpdateAttribute *UpdateAttribute `xml:"tns:updateAttribute,omitempty" json:"tns:updateAttribute,omitempty" yaml:"tns:updateAttribute,omitempty"`
}

// Operation wrapper for UpdateAttribute.
// OperationUpdateAttributeResponse was auto-generated from WSDL.
type OperationUpdateAttributeResponse struct {
	UpdateAttributeResponse *UpdateAttributeResponse `xml:"updateAttributeResponse,omitempty" json:"updateAttributeResponse,omitempty" yaml:"updateAttributeResponse,omitempty"`
}

// Operation wrapper for UpdateComponentMap.
// OperationUpdateComponentMap was auto-generated from WSDL.
type OperationUpdateComponentMap struct {
	UpdateComponentMap *UpdateComponentMap `xml:"tns:updateComponentMap,omitempty" json:"tns:updateComponentMap,omitempty" yaml:"tns:updateComponentMap,omitempty"`
}

// Operation wrapper for UpdateComponentMap.
// OperationUpdateComponentMapResponse was auto-generated from
// WSDL.
type OperationUpdateComponentMapResponse struct {
	UpdateComponentMapResponse *UpdateComponentMapResponse `xml:"updateComponentMapResponse,omitempty" json:"updateComponentMapResponse,omitempty" yaml:"updateComponentMapResponse,omitempty"`
}

// Operation wrapper for UpdateGroup.
// OperationUpdateGroup was auto-generated from WSDL.
type OperationUpdateGroup struct {
	UpdateGroup *UpdateGroup `xml:"tns:updateGroup,omitempty" json:"tns:updateGroup,omitempty" yaml:"tns:updateGroup,omitempty"`
}

// Operation wrapper for UpdateGroup.
// OperationUpdateGroupResponse was auto-generated from WSDL.
type OperationUpdateGroupResponse struct {
	UpdateGroupResponse *UpdateGroupResponse `xml:"updateGroupResponse,omitempty" json:"updateGroupResponse,omitempty" yaml:"updateGroupResponse,omitempty"`
}

// Operation wrapper for UpdateLdapConfiguration.
// OperationUpdateLdapConfiguration was auto-generated from WSDL.
type OperationUpdateLdapConfiguration struct {
	UpdateLdapConfiguration *UpdateLdapConfiguration `xml:"tns:updateLdapConfiguration,omitempty" json:"tns:updateLdapConfiguration,omitempty" yaml:"tns:updateLdapConfiguration,omitempty"`
}

// Operation wrapper for UpdateLdapConfiguration.
// OperationUpdateLdapConfigurationResponse was auto-generated
// from WSDL.
type OperationUpdateLdapConfigurationResponse struct {
	UpdateLdapConfigurationResponse *UpdateLdapConfigurationResponse `xml:"updateLdapConfigurationResponse,omitempty" json:"updateLdapConfigurationResponse,omitempty" yaml:"updateLdapConfigurationResponse,omitempty"`
}

// Operation wrapper for UpdateProject.
// OperationUpdateProject was auto-generated from WSDL.
type OperationUpdateProject struct {
	UpdateProject *UpdateProject `xml:"tns:updateProject,omitempty" json:"tns:updateProject,omitempty" yaml:"tns:updateProject,omitempty"`
}

// Operation wrapper for UpdateProject.
// OperationUpdateProjectResponse was auto-generated from WSDL.
type OperationUpdateProjectResponse struct {
	UpdateProjectResponse *UpdateProjectResponse `xml:"updateProjectResponse,omitempty" json:"updateProjectResponse,omitempty" yaml:"updateProjectResponse,omitempty"`
}

// Operation wrapper for UpdateRole.
// OperationUpdateRole was auto-generated from WSDL.
type OperationUpdateRole struct {
	UpdateRole *UpdateRole `xml:"tns:updateRole,omitempty" json:"tns:updateRole,omitempty" yaml:"tns:updateRole,omitempty"`
}

// Operation wrapper for UpdateRole.
// OperationUpdateRoleResponse was auto-generated from WSDL.
type OperationUpdateRoleResponse struct {
	UpdateRoleResponse *UpdateRoleResponse `xml:"updateRoleResponse,omitempty" json:"updateRoleResponse,omitempty" yaml:"updateRoleResponse,omitempty"`
}

// Operation wrapper for UpdateSignInConfiguration.
// OperationUpdateSignInConfiguration was auto-generated from WSDL.
type OperationUpdateSignInConfiguration struct {
	UpdateSignInConfiguration *UpdateSignInConfiguration `xml:"tns:updateSignInConfiguration,omitempty" json:"tns:updateSignInConfiguration,omitempty" yaml:"tns:updateSignInConfiguration,omitempty"`
}

// Operation wrapper for UpdateSignInConfiguration.
// OperationUpdateSignInConfigurationResponse was auto-generated
// from WSDL.
type OperationUpdateSignInConfigurationResponse struct {
	UpdateSignInConfigurationResponse *UpdateSignInConfigurationResponse `xml:"updateSignInConfigurationResponse,omitempty" json:"updateSignInConfigurationResponse,omitempty" yaml:"updateSignInConfigurationResponse,omitempty"`
}

// Operation wrapper for UpdateSnapshotInfo.
// OperationUpdateSnapshotInfo was auto-generated from WSDL.
type OperationUpdateSnapshotInfo struct {
	UpdateSnapshotInfo *UpdateSnapshotInfo `xml:"tns:updateSnapshotInfo,omitempty" json:"tns:updateSnapshotInfo,omitempty" yaml:"tns:updateSnapshotInfo,omitempty"`
}

// Operation wrapper for UpdateSnapshotInfo.
// OperationUpdateSnapshotInfoResponse was auto-generated from
// WSDL.
type OperationUpdateSnapshotInfoResponse struct {
	UpdateSnapshotInfoResponse *UpdateSnapshotInfoResponse `xml:"updateSnapshotInfoResponse,omitempty" json:"updateSnapshotInfoResponse,omitempty" yaml:"updateSnapshotInfoResponse,omitempty"`
}

// Operation wrapper for UpdateStream.
// OperationUpdateStream was auto-generated from WSDL.
type OperationUpdateStream struct {
	UpdateStream *UpdateStream `xml:"tns:updateStream,omitempty" json:"tns:updateStream,omitempty" yaml:"tns:updateStream,omitempty"`
}

// Operation wrapper for UpdateStream.
// OperationUpdateStreamResponse was auto-generated from WSDL.
type OperationUpdateStreamResponse struct {
	UpdateStreamResponse *UpdateStreamResponse `xml:"updateStreamResponse,omitempty" json:"updateStreamResponse,omitempty" yaml:"updateStreamResponse,omitempty"`
}

// Operation wrapper for UpdateTriageStore.
// OperationUpdateTriageStore was auto-generated from WSDL.
type OperationUpdateTriageStore struct {
	UpdateTriageStore *UpdateTriageStore `xml:"tns:updateTriageStore,omitempty" json:"tns:updateTriageStore,omitempty" yaml:"tns:updateTriageStore,omitempty"`
}

// Operation wrapper for UpdateTriageStore.
// OperationUpdateTriageStoreResponse was auto-generated from WSDL.
type OperationUpdateTriageStoreResponse struct {
	UpdateTriageStoreResponse *UpdateTriageStoreResponse `xml:"updateTriageStoreResponse,omitempty" json:"updateTriageStoreResponse,omitempty" yaml:"updateTriageStoreResponse,omitempty"`
}

// Operation wrapper for UpdateUser.
// OperationUpdateUser was auto-generated from WSDL.
type OperationUpdateUser struct {
	UpdateUser *UpdateUser `xml:"tns:updateUser,omitempty" json:"tns:updateUser,omitempty" yaml:"tns:updateUser,omitempty"`
}

// Operation wrapper for UpdateUser.
// OperationUpdateUserResponse was auto-generated from WSDL.
type OperationUpdateUserResponse struct {
	UpdateUserResponse *UpdateUserResponse `xml:"updateUserResponse,omitempty" json:"updateUserResponse,omitempty" yaml:"updateUserResponse,omitempty"`
}

// configurationService implements the ConfigurationService interface.
type configurationService struct {
	cli *soap.Client
}

// CopyStream was auto-generated from WSDL.
func (p *configurationService) CopyStream(CopyStream *CopyStream) (*CopyStreamResponse, error) {
	α := struct {
		OperationCopyStream `xml:"tns:copyStream"`
	}{
		OperationCopyStream{
			CopyStream,
		},
	}

	γ := struct {
		OperationCopyStreamResponse `xml:"copyStreamResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CopyStream", α, &γ); err != nil {
		return nil, err
	}
	return γ.CopyStreamResponse, nil
}

// CreateAttribute was auto-generated from WSDL.
func (p *configurationService) CreateAttribute(CreateAttribute *CreateAttribute) (*CreateAttributeResponse, error) {
	α := struct {
		OperationCreateAttribute `xml:"tns:createAttribute"`
	}{
		OperationCreateAttribute{
			CreateAttribute,
		},
	}

	γ := struct {
		OperationCreateAttributeResponse `xml:"createAttributeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateAttribute", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateAttributeResponse, nil
}

// CreateComponentMap was auto-generated from WSDL.
func (p *configurationService) CreateComponentMap(CreateComponentMap *CreateComponentMap) (*CreateComponentMapResponse, error) {
	α := struct {
		OperationCreateComponentMap `xml:"tns:createComponentMap"`
	}{
		OperationCreateComponentMap{
			CreateComponentMap,
		},
	}

	γ := struct {
		OperationCreateComponentMapResponse `xml:"createComponentMapResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateComponentMap", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateComponentMapResponse, nil
}

// CreateGroup was auto-generated from WSDL.
func (p *configurationService) CreateGroup(CreateGroup *CreateGroup) (*CreateGroupResponse, error) {
	α := struct {
		OperationCreateGroup `xml:"tns:createGroup"`
	}{
		OperationCreateGroup{
			CreateGroup,
		},
	}

	γ := struct {
		OperationCreateGroupResponse `xml:"createGroupResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateGroup", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateGroupResponse, nil
}

// CreateLdapConfiguration was auto-generated from WSDL.
func (p *configurationService) CreateLdapConfiguration(CreateLdapConfiguration *CreateLdapConfiguration) (*CreateLdapConfigurationResponse, error) {
	α := struct {
		OperationCreateLdapConfiguration `xml:"tns:createLdapConfiguration"`
	}{
		OperationCreateLdapConfiguration{
			CreateLdapConfiguration,
		},
	}

	γ := struct {
		OperationCreateLdapConfigurationResponse `xml:"createLdapConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateLdapConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateLdapConfigurationResponse, nil
}

// CreateProject was auto-generated from WSDL.
func (p *configurationService) CreateProject(CreateProject *CreateProject) (*CreateProjectResponse, error) {
	α := struct {
		OperationCreateProject `xml:"tns:createProject"`
	}{
		OperationCreateProject{
			CreateProject,
		},
	}

	γ := struct {
		OperationCreateProjectResponse `xml:"createProjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateProject", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateProjectResponse, nil
}

// CreateRole was auto-generated from WSDL.
func (p *configurationService) CreateRole(CreateRole *CreateRole) (*CreateRoleResponse, error) {
	α := struct {
		OperationCreateRole `xml:"tns:createRole"`
	}{
		OperationCreateRole{
			CreateRole,
		},
	}

	γ := struct {
		OperationCreateRoleResponse `xml:"createRoleResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateRole", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateRoleResponse, nil
}

// CreateStream was auto-generated from WSDL.
func (p *configurationService) CreateStream(CreateStream *CreateStream) (*CreateStreamResponse, error) {
	α := struct {
		OperationCreateStream `xml:"tns:createStream"`
	}{
		OperationCreateStream{
			CreateStream,
		},
	}

	γ := struct {
		OperationCreateStreamResponse `xml:"createStreamResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateStream", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateStreamResponse, nil
}

// CreateStreamInProject was auto-generated from WSDL.
func (p *configurationService) CreateStreamInProject(CreateStreamInProject *CreateStreamInProject) (*CreateStreamInProjectResponse, error) {
	α := struct {
		OperationCreateStreamInProject `xml:"tns:createStreamInProject"`
	}{
		OperationCreateStreamInProject{
			CreateStreamInProject,
		},
	}

	γ := struct {
		OperationCreateStreamInProjectResponse `xml:"createStreamInProjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateStreamInProject", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateStreamInProjectResponse, nil
}

// CreateTriageStore was auto-generated from WSDL.
func (p *configurationService) CreateTriageStore(CreateTriageStore *CreateTriageStore) (*CreateTriageStoreResponse, error) {
	α := struct {
		OperationCreateTriageStore `xml:"tns:createTriageStore"`
	}{
		OperationCreateTriageStore{
			CreateTriageStore,
		},
	}

	γ := struct {
		OperationCreateTriageStoreResponse `xml:"createTriageStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateTriageStore", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateTriageStoreResponse, nil
}

// CreateUser was auto-generated from WSDL.
func (p *configurationService) CreateUser(CreateUser *CreateUser) (*CreateUserResponse, error) {
	α := struct {
		OperationCreateUser `xml:"tns:createUser"`
	}{
		OperationCreateUser{
			CreateUser,
		},
	}

	γ := struct {
		OperationCreateUserResponse `xml:"createUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateUserResponse, nil
}

// DeleteAttribute was auto-generated from WSDL.
func (p *configurationService) DeleteAttribute(DeleteAttribute *DeleteAttribute) (*DeleteAttributeResponse, error) {
	α := struct {
		OperationDeleteAttribute `xml:"tns:deleteAttribute"`
	}{
		OperationDeleteAttribute{
			DeleteAttribute,
		},
	}

	γ := struct {
		OperationDeleteAttributeResponse `xml:"deleteAttributeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteAttribute", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteAttributeResponse, nil
}

// DeleteComponentMap was auto-generated from WSDL.
func (p *configurationService) DeleteComponentMap(DeleteComponentMap *DeleteComponentMap) (*DeleteComponentMapResponse, error) {
	α := struct {
		OperationDeleteComponentMap `xml:"tns:deleteComponentMap"`
	}{
		OperationDeleteComponentMap{
			DeleteComponentMap,
		},
	}

	γ := struct {
		OperationDeleteComponentMapResponse `xml:"deleteComponentMapResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteComponentMap", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteComponentMapResponse, nil
}

// DeleteGroup was auto-generated from WSDL.
func (p *configurationService) DeleteGroup(DeleteGroup *DeleteGroup) (*DeleteGroupResponse, error) {
	α := struct {
		OperationDeleteGroup `xml:"tns:deleteGroup"`
	}{
		OperationDeleteGroup{
			DeleteGroup,
		},
	}

	γ := struct {
		OperationDeleteGroupResponse `xml:"deleteGroupResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteGroup", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteGroupResponse, nil
}

// DeleteLdapConfiguration was auto-generated from WSDL.
func (p *configurationService) DeleteLdapConfiguration(DeleteLdapConfiguration *DeleteLdapConfiguration) (*DeleteLdapConfigurationResponse, error) {
	α := struct {
		OperationDeleteLdapConfiguration `xml:"tns:deleteLdapConfiguration"`
	}{
		OperationDeleteLdapConfiguration{
			DeleteLdapConfiguration,
		},
	}

	γ := struct {
		OperationDeleteLdapConfigurationResponse `xml:"deleteLdapConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteLdapConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteLdapConfigurationResponse, nil
}

// DeleteProject was auto-generated from WSDL.
func (p *configurationService) DeleteProject(DeleteProject *DeleteProject) (*DeleteProjectResponse, error) {
	α := struct {
		OperationDeleteProject `xml:"tns:deleteProject"`
	}{
		OperationDeleteProject{
			DeleteProject,
		},
	}

	γ := struct {
		OperationDeleteProjectResponse `xml:"deleteProjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteProject", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteProjectResponse, nil
}

// DeleteRole was auto-generated from WSDL.
func (p *configurationService) DeleteRole(DeleteRole *DeleteRole) (*DeleteRoleResponse, error) {
	α := struct {
		OperationDeleteRole `xml:"tns:deleteRole"`
	}{
		OperationDeleteRole{
			DeleteRole,
		},
	}

	γ := struct {
		OperationDeleteRoleResponse `xml:"deleteRoleResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteRole", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteRoleResponse, nil
}

// DeleteSnapshot was auto-generated from WSDL.
func (p *configurationService) DeleteSnapshot(DeleteSnapshot *DeleteSnapshot) (*DeleteSnapshotResponse, error) {
	α := struct {
		OperationDeleteSnapshot `xml:"tns:deleteSnapshot"`
	}{
		OperationDeleteSnapshot{
			DeleteSnapshot,
		},
	}

	γ := struct {
		OperationDeleteSnapshotResponse `xml:"deleteSnapshotResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteSnapshot", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteSnapshotResponse, nil
}

// DeleteStream was auto-generated from WSDL.
func (p *configurationService) DeleteStream(DeleteStream *DeleteStream) (*DeleteStreamResponse, error) {
	α := struct {
		OperationDeleteStream `xml:"tns:deleteStream"`
	}{
		OperationDeleteStream{
			DeleteStream,
		},
	}

	γ := struct {
		OperationDeleteStreamResponse `xml:"deleteStreamResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteStream", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteStreamResponse, nil
}

// DeleteTriageStore was auto-generated from WSDL.
func (p *configurationService) DeleteTriageStore(DeleteTriageStore *DeleteTriageStore) (*DeleteTriageStoreResponse, error) {
	α := struct {
		OperationDeleteTriageStore `xml:"tns:deleteTriageStore"`
	}{
		OperationDeleteTriageStore{
			DeleteTriageStore,
		},
	}

	γ := struct {
		OperationDeleteTriageStoreResponse `xml:"deleteTriageStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteTriageStore", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteTriageStoreResponse, nil
}

// DeleteUser was auto-generated from WSDL.
func (p *configurationService) DeleteUser(DeleteUser *DeleteUser) (*DeleteUserResponse, error) {
	α := struct {
		OperationDeleteUser `xml:"tns:deleteUser"`
	}{
		OperationDeleteUser{
			DeleteUser,
		},
	}

	γ := struct {
		OperationDeleteUserResponse `xml:"deleteUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("DeleteUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteUserResponse, nil
}

// ExecuteNotification was auto-generated from WSDL.
func (p *configurationService) ExecuteNotification(ExecuteNotification *ExecuteNotification) (*ExecuteNotificationResponse, error) {
	α := struct {
		OperationExecuteNotification `xml:"tns:executeNotification"`
	}{
		OperationExecuteNotification{
			ExecuteNotification,
		},
	}

	γ := struct {
		OperationExecuteNotificationResponse `xml:"executeNotificationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ExecuteNotification", α, &γ); err != nil {
		return nil, err
	}
	return γ.ExecuteNotificationResponse, nil
}

// GetAllLdapConfigurations was auto-generated from WSDL.
func (p *configurationService) GetAllLdapConfigurations(GetAllLdapConfigurations *GetAllLdapConfigurations) (*GetAllLdapConfigurationsResponse, error) {
	α := struct {
		OperationGetAllLdapConfigurations `xml:"tns:getAllLdapConfigurations"`
	}{
		OperationGetAllLdapConfigurations{
			GetAllLdapConfigurations,
		},
	}

	γ := struct {
		OperationGetAllLdapConfigurationsResponse `xml:"getAllLdapConfigurationsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetAllLdapConfigurations", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAllLdapConfigurationsResponse, nil
}

// GetAllPermissions was auto-generated from WSDL.
func (p *configurationService) GetAllPermissions(GetAllPermissions *GetAllPermissions) (*GetAllPermissionsResponse, error) {
	α := struct {
		OperationGetAllPermissions `xml:"tns:getAllPermissions"`
	}{
		OperationGetAllPermissions{
			GetAllPermissions,
		},
	}

	γ := struct {
		OperationGetAllPermissionsResponse `xml:"getAllPermissionsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetAllPermissions", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAllPermissionsResponse, nil
}

// GetAllRoles was auto-generated from WSDL.
func (p *configurationService) GetAllRoles(GetAllRoles *GetAllRoles) (*GetAllRolesResponse, error) {
	α := struct {
		OperationGetAllRoles `xml:"tns:getAllRoles"`
	}{
		OperationGetAllRoles{
			GetAllRoles,
		},
	}

	γ := struct {
		OperationGetAllRolesResponse `xml:"getAllRolesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetAllRoles", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAllRolesResponse, nil
}

// GetArchitectureAnalysisConfiguration was auto-generated from
// WSDL.
func (p *configurationService) GetArchitectureAnalysisConfiguration(GetArchitectureAnalysisConfiguration *GetArchitectureAnalysisConfiguration) (*GetArchitectureAnalysisConfigurationResponse, error) {
	α := struct {
		OperationGetArchitectureAnalysisConfiguration `xml:"tns:getArchitectureAnalysisConfiguration"`
	}{
		OperationGetArchitectureAnalysisConfiguration{
			GetArchitectureAnalysisConfiguration,
		},
	}

	γ := struct {
		OperationGetArchitectureAnalysisConfigurationResponse `xml:"getArchitectureAnalysisConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetArchitectureAnalysisConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetArchitectureAnalysisConfigurationResponse, nil
}

// GetAttribute was auto-generated from WSDL.
func (p *configurationService) GetAttribute(GetAttribute *GetAttribute) (*GetAttributeResponse, error) {
	α := struct {
		OperationGetAttribute `xml:"tns:getAttribute"`
	}{
		OperationGetAttribute{
			GetAttribute,
		},
	}

	γ := struct {
		OperationGetAttributeResponse `xml:"getAttributeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetAttribute", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAttributeResponse, nil
}

// GetAttributes was auto-generated from WSDL.
func (p *configurationService) GetAttributes(GetAttributes *GetAttributes) (*GetAttributesResponse, error) {
	α := struct {
		OperationGetAttributes `xml:"tns:getAttributes"`
	}{
		OperationGetAttributes{
			GetAttributes,
		},
	}

	γ := struct {
		OperationGetAttributesResponse `xml:"getAttributesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetAttributes", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAttributesResponse, nil
}

// GetBackupConfiguration was auto-generated from WSDL.
func (p *configurationService) GetBackupConfiguration(GetBackupConfiguration *GetBackupConfiguration) (*GetBackupConfigurationResponse, error) {
	α := struct {
		OperationGetBackupConfiguration `xml:"tns:getBackupConfiguration"`
	}{
		OperationGetBackupConfiguration{
			GetBackupConfiguration,
		},
	}

	γ := struct {
		OperationGetBackupConfigurationResponse `xml:"getBackupConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetBackupConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetBackupConfigurationResponse, nil
}

// GetCategoryNames was auto-generated from WSDL.
func (p *configurationService) GetCategoryNames(GetCategoryNames *GetCategoryNames) (*GetCategoryNamesResponse, error) {
	α := struct {
		OperationGetCategoryNames `xml:"tns:getCategoryNames"`
	}{
		OperationGetCategoryNames{
			GetCategoryNames,
		},
	}

	γ := struct {
		OperationGetCategoryNamesResponse `xml:"getCategoryNamesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCategoryNames", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCategoryNamesResponse, nil
}

// GetCheckerNames was auto-generated from WSDL.
func (p *configurationService) GetCheckerNames(GetCheckerNames *GetCheckerNames) (*GetCheckerNamesResponse, error) {
	α := struct {
		OperationGetCheckerNames `xml:"tns:getCheckerNames"`
	}{
		OperationGetCheckerNames{
			GetCheckerNames,
		},
	}

	γ := struct {
		OperationGetCheckerNamesResponse `xml:"getCheckerNamesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCheckerNames", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCheckerNamesResponse, nil
}

// GetCommitState was auto-generated from WSDL.
func (p *configurationService) GetCommitState(GetCommitState *GetCommitState) (*GetCommitStateResponse, error) {
	α := struct {
		OperationGetCommitState `xml:"tns:getCommitState"`
	}{
		OperationGetCommitState{
			GetCommitState,
		},
	}

	γ := struct {
		OperationGetCommitStateResponse `xml:"getCommitStateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCommitState", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCommitStateResponse, nil
}

// GetComponent was auto-generated from WSDL.
func (p *configurationService) GetComponent(GetComponent *GetComponent) (*GetComponentResponse, error) {
	α := struct {
		OperationGetComponent `xml:"tns:getComponent"`
	}{
		OperationGetComponent{
			GetComponent,
		},
	}

	γ := struct {
		OperationGetComponentResponse `xml:"getComponentResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetComponent", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetComponentResponse, nil
}

// GetComponentMaps was auto-generated from WSDL.
func (p *configurationService) GetComponentMaps(GetComponentMaps *GetComponentMaps) (*GetComponentMapsResponse, error) {
	α := struct {
		OperationGetComponentMaps `xml:"tns:getComponentMaps"`
	}{
		OperationGetComponentMaps{
			GetComponentMaps,
		},
	}

	γ := struct {
		OperationGetComponentMapsResponse `xml:"getComponentMapsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetComponentMaps", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetComponentMapsResponse, nil
}

// GetDefectStatuses was auto-generated from WSDL.
func (p *configurationService) GetDefectStatuses(GetDefectStatuses *GetDefectStatuses) (*GetDefectStatusesResponse, error) {
	α := struct {
		OperationGetDefectStatuses `xml:"tns:getDefectStatuses"`
	}{
		OperationGetDefectStatuses{
			GetDefectStatuses,
		},
	}

	γ := struct {
		OperationGetDefectStatusesResponse `xml:"getDefectStatusesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetDefectStatuses", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDefectStatusesResponse, nil
}

// GetDeleteSnapshotJobInfo was auto-generated from WSDL.
func (p *configurationService) GetDeleteSnapshotJobInfo(GetDeleteSnapshotJobInfo *GetDeleteSnapshotJobInfo) (*GetDeleteSnapshotJobInfoResponse, error) {
	α := struct {
		OperationGetDeleteSnapshotJobInfo `xml:"tns:getDeleteSnapshotJobInfo"`
	}{
		OperationGetDeleteSnapshotJobInfo{
			GetDeleteSnapshotJobInfo,
		},
	}

	γ := struct {
		OperationGetDeleteSnapshotJobInfoResponse `xml:"getDeleteSnapshotJobInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetDeleteSnapshotJobInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDeleteSnapshotJobInfoResponse, nil
}

// GetDeveloperStreamsProjects was auto-generated from WSDL.
func (p *configurationService) GetDeveloperStreamsProjects(GetDeveloperStreamsProjects *GetDeveloperStreamsProjects) (*GetDeveloperStreamsProjectsResponse, error) {
	α := struct {
		OperationGetDeveloperStreamsProjects `xml:"tns:getDeveloperStreamsProjects"`
	}{
		OperationGetDeveloperStreamsProjects{
			GetDeveloperStreamsProjects,
		},
	}

	γ := struct {
		OperationGetDeveloperStreamsProjectsResponse `xml:"getDeveloperStreamsProjectsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetDeveloperStreamsProjects", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDeveloperStreamsProjectsResponse, nil
}

// GetGroup was auto-generated from WSDL.
func (p *configurationService) GetGroup(GetGroup *GetGroup) (*GetGroupResponse, error) {
	α := struct {
		OperationGetGroup `xml:"tns:getGroup"`
	}{
		OperationGetGroup{
			GetGroup,
		},
	}

	γ := struct {
		OperationGetGroupResponse `xml:"getGroupResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetGroup", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetGroupResponse, nil
}

// GetGroups was auto-generated from WSDL.
func (p *configurationService) GetGroups(GetGroups *GetGroups) (*GetGroupsResponse, error) {
	α := struct {
		OperationGetGroups `xml:"tns:getGroups"`
	}{
		OperationGetGroups{
			GetGroups,
		},
	}

	γ := struct {
		OperationGetGroupsResponse `xml:"getGroupsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetGroups", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetGroupsResponse, nil
}

// GetLastUpdateTimes was auto-generated from WSDL.
func (p *configurationService) GetLastUpdateTimes(GetLastUpdateTimes *GetLastUpdateTimes) (*GetLastUpdateTimesResponse, error) {
	α := struct {
		OperationGetLastUpdateTimes `xml:"tns:getLastUpdateTimes"`
	}{
		OperationGetLastUpdateTimes{
			GetLastUpdateTimes,
		},
	}

	γ := struct {
		OperationGetLastUpdateTimesResponse `xml:"getLastUpdateTimesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetLastUpdateTimes", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLastUpdateTimesResponse, nil
}

// GetLdapServerDomains was auto-generated from WSDL.
func (p *configurationService) GetLdapServerDomains(GetLdapServerDomains *GetLdapServerDomains) (*GetLdapServerDomainsResponse, error) {
	α := struct {
		OperationGetLdapServerDomains `xml:"tns:getLdapServerDomains"`
	}{
		OperationGetLdapServerDomains{
			GetLdapServerDomains,
		},
	}

	γ := struct {
		OperationGetLdapServerDomainsResponse `xml:"getLdapServerDomainsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetLdapServerDomains", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLdapServerDomainsResponse, nil
}

// GetLicenseConfiguration was auto-generated from WSDL.
func (p *configurationService) GetLicenseConfiguration(GetLicenseConfiguration *GetLicenseConfiguration) (*GetLicenseConfigurationResponse, error) {
	α := struct {
		OperationGetLicenseConfiguration `xml:"tns:getLicenseConfiguration"`
	}{
		OperationGetLicenseConfiguration{
			GetLicenseConfiguration,
		},
	}

	γ := struct {
		OperationGetLicenseConfigurationResponse `xml:"getLicenseConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetLicenseConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLicenseConfigurationResponse, nil
}

// GetLicenseState was auto-generated from WSDL.
func (p *configurationService) GetLicenseState(GetLicenseState *GetLicenseState) (*GetLicenseStateResponse, error) {
	α := struct {
		OperationGetLicenseState `xml:"tns:getLicenseState"`
	}{
		OperationGetLicenseState{
			GetLicenseState,
		},
	}

	γ := struct {
		OperationGetLicenseStateResponse `xml:"getLicenseStateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetLicenseState", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLicenseStateResponse, nil
}

// GetLoggingConfiguration was auto-generated from WSDL.
func (p *configurationService) GetLoggingConfiguration(GetLoggingConfiguration *GetLoggingConfiguration) (*GetLoggingConfigurationResponse, error) {
	α := struct {
		OperationGetLoggingConfiguration `xml:"tns:getLoggingConfiguration"`
	}{
		OperationGetLoggingConfiguration{
			GetLoggingConfiguration,
		},
	}

	γ := struct {
		OperationGetLoggingConfigurationResponse `xml:"getLoggingConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetLoggingConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLoggingConfigurationResponse, nil
}

// GetMessageOfTheDay was auto-generated from WSDL.
func (p *configurationService) GetMessageOfTheDay(GetMessageOfTheDay *GetMessageOfTheDay) (*GetMessageOfTheDayResponse, error) {
	α := struct {
		OperationGetMessageOfTheDay `xml:"tns:getMessageOfTheDay"`
	}{
		OperationGetMessageOfTheDay{
			GetMessageOfTheDay,
		},
	}

	γ := struct {
		OperationGetMessageOfTheDayResponse `xml:"getMessageOfTheDayResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMessageOfTheDay", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMessageOfTheDayResponse, nil
}

// GetOutputFileForSnapshot was auto-generated from WSDL.
func (p *configurationService) GetOutputFileForSnapshot(GetOutputFileForSnapshot *GetOutputFileForSnapshot) (*GetOutputFileForSnapshotResponse, error) {
	α := struct {
		OperationGetOutputFileForSnapshot `xml:"tns:getOutputFileForSnapshot"`
	}{
		OperationGetOutputFileForSnapshot{
			GetOutputFileForSnapshot,
		},
	}

	γ := struct {
		OperationGetOutputFileForSnapshotResponse `xml:"getOutputFileForSnapshotResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetOutputFileForSnapshot", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetOutputFileForSnapshotResponse, nil
}

// GetProjects was auto-generated from WSDL.
func (p *configurationService) GetProjects(GetProjects *GetProjects) (*GetProjectsResponse, error) {
	α := struct {
		OperationGetProjects `xml:"tns:getProjects"`
	}{
		OperationGetProjects{
			GetProjects,
		},
	}

	γ := struct {
		OperationGetProjectsResponse `xml:"getProjectsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetProjects", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetProjectsResponse, nil
}

// GetRole was auto-generated from WSDL.
func (p *configurationService) GetRole(GetRole *GetRole) (*GetRoleResponse, error) {
	α := struct {
		OperationGetRole `xml:"tns:getRole"`
	}{
		OperationGetRole{
			GetRole,
		},
	}

	γ := struct {
		OperationGetRoleResponse `xml:"getRoleResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetRole", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetRoleResponse, nil
}

// GetServerTime was auto-generated from WSDL.
func (p *configurationService) GetServerTime(GetServerTime *GetServerTime) (*GetServerTimeResponse, error) {
	α := struct {
		OperationGetServerTime `xml:"tns:getServerTime"`
	}{
		OperationGetServerTime{
			GetServerTime,
		},
	}

	γ := struct {
		OperationGetServerTimeResponse `xml:"getServerTimeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetServerTime", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetServerTimeResponse, nil
}

// GetSignInConfiguration was auto-generated from WSDL.
func (p *configurationService) GetSignInConfiguration(GetSignInConfiguration *GetSignInConfiguration) (*GetSignInConfigurationResponse, error) {
	α := struct {
		OperationGetSignInConfiguration `xml:"tns:getSignInConfiguration"`
	}{
		OperationGetSignInConfiguration{
			GetSignInConfiguration,
		},
	}

	γ := struct {
		OperationGetSignInConfigurationResponse `xml:"getSignInConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSignInConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetSignInConfigurationResponse, nil
}

// GetSkeletonizationConfiguration was auto-generated from WSDL.
func (p *configurationService) GetSkeletonizationConfiguration(GetSkeletonizationConfiguration *GetSkeletonizationConfiguration) (*GetSkeletonizationConfigurationResponse, error) {
	α := struct {
		OperationGetSkeletonizationConfiguration `xml:"tns:getSkeletonizationConfiguration"`
	}{
		OperationGetSkeletonizationConfiguration{
			GetSkeletonizationConfiguration,
		},
	}

	γ := struct {
		OperationGetSkeletonizationConfigurationResponse `xml:"getSkeletonizationConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSkeletonizationConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetSkeletonizationConfigurationResponse, nil
}

// GetSnapshotInformation was auto-generated from WSDL.
func (p *configurationService) GetSnapshotInformation(GetSnapshotInformation *GetSnapshotInformation) (*GetSnapshotInformationResponse, error) {
	α := struct {
		OperationGetSnapshotInformation `xml:"tns:getSnapshotInformation"`
	}{
		OperationGetSnapshotInformation{
			GetSnapshotInformation,
		},
	}

	γ := struct {
		OperationGetSnapshotInformationResponse `xml:"getSnapshotInformationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSnapshotInformation", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetSnapshotInformationResponse, nil
}

// GetSnapshotPurgeDetails was auto-generated from WSDL.
func (p *configurationService) GetSnapshotPurgeDetails(GetSnapshotPurgeDetails *GetSnapshotPurgeDetails) (*GetSnapshotPurgeDetailsResponse, error) {
	α := struct {
		OperationGetSnapshotPurgeDetails `xml:"tns:getSnapshotPurgeDetails"`
	}{
		OperationGetSnapshotPurgeDetails{
			GetSnapshotPurgeDetails,
		},
	}

	γ := struct {
		OperationGetSnapshotPurgeDetailsResponse `xml:"getSnapshotPurgeDetailsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSnapshotPurgeDetails", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetSnapshotPurgeDetailsResponse, nil
}

// GetSnapshotsForStream was auto-generated from WSDL.
func (p *configurationService) GetSnapshotsForStream(GetSnapshotsForStream *GetSnapshotsForStream) (*GetSnapshotsForStreamResponse, error) {
	α := struct {
		OperationGetSnapshotsForStream `xml:"tns:getSnapshotsForStream"`
	}{
		OperationGetSnapshotsForStream{
			GetSnapshotsForStream,
		},
	}

	γ := struct {
		OperationGetSnapshotsForStreamResponse `xml:"getSnapshotsForStreamResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSnapshotsForStream", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetSnapshotsForStreamResponse, nil
}

// GetStreams was auto-generated from WSDL.
func (p *configurationService) GetStreams(GetStreams *GetStreams) (*GetStreamsResponse, error) {
	α := struct {
		OperationGetStreams `xml:"tns:getStreams"`
	}{
		OperationGetStreams{
			GetStreams,
		},
	}

	γ := struct {
		OperationGetStreamsResponse `xml:"getStreamsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStreams", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStreamsResponse, nil
}

// GetSystemConfig was auto-generated from WSDL.
func (p *configurationService) GetSystemConfig(GetSystemConfig *GetSystemConfig) (*GetSystemConfigResponse, error) {
	α := struct {
		OperationGetSystemConfig `xml:"tns:getSystemConfig"`
	}{
		OperationGetSystemConfig{
			GetSystemConfig,
		},
	}

	γ := struct {
		OperationGetSystemConfigResponse `xml:"getSystemConfigResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSystemConfig", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetSystemConfigResponse, nil
}

// GetTriageStores was auto-generated from WSDL.
func (p *configurationService) GetTriageStores(GetTriageStores *GetTriageStores) (*GetTriageStoresResponse, error) {
	α := struct {
		OperationGetTriageStores `xml:"tns:getTriageStores"`
	}{
		OperationGetTriageStores{
			GetTriageStores,
		},
	}

	γ := struct {
		OperationGetTriageStoresResponse `xml:"getTriageStoresResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetTriageStores", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTriageStoresResponse, nil
}

// GetTypeNames was auto-generated from WSDL.
func (p *configurationService) GetTypeNames(GetTypeNames *GetTypeNames) (*GetTypeNamesResponse, error) {
	α := struct {
		OperationGetTypeNames `xml:"tns:getTypeNames"`
	}{
		OperationGetTypeNames{
			GetTypeNames,
		},
	}

	γ := struct {
		OperationGetTypeNamesResponse `xml:"getTypeNamesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetTypeNames", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTypeNamesResponse, nil
}

// GetUser was auto-generated from WSDL.
func (p *configurationService) GetUser(GetUser *GetUser) (*GetUserResponse, error) {
	α := struct {
		OperationGetUser `xml:"tns:getUser"`
	}{
		OperationGetUser{
			GetUser,
		},
	}

	γ := struct {
		OperationGetUserResponse `xml:"getUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetUserResponse, nil
}

// GetUsers was auto-generated from WSDL.
func (p *configurationService) GetUsers(GetUsers *GetUsers) (*GetUsersResponse, error) {
	α := struct {
		OperationGetUsers `xml:"tns:getUsers"`
	}{
		OperationGetUsers{
			GetUsers,
		},
	}

	γ := struct {
		OperationGetUsersResponse `xml:"getUsersResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetUsers", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetUsersResponse, nil
}

// GetVersion was auto-generated from WSDL.
func (p *configurationService) GetVersion(GetVersion *GetVersion) (*GetVersionResponse, error) {
	α := struct {
		OperationGetVersion `xml:"tns:getVersion"`
	}{
		OperationGetVersion{
			GetVersion,
		},
	}

	γ := struct {
		OperationGetVersionResponse `xml:"getVersionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetVersion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetVersionResponse, nil
}

// ImportLicense was auto-generated from WSDL.
func (p *configurationService) ImportLicense(ImportLicense *ImportLicense) (*ImportLicenseResponse, error) {
	α := struct {
		OperationImportLicense `xml:"tns:importLicense"`
	}{
		OperationImportLicense{
			ImportLicense,
		},
	}

	γ := struct {
		OperationImportLicenseResponse `xml:"importLicenseResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ImportLicense", α, &γ); err != nil {
		return nil, err
	}
	return γ.ImportLicenseResponse, nil
}

// MergeTriageStores was auto-generated from WSDL.
func (p *configurationService) MergeTriageStores(MergeTriageStores *MergeTriageStores) (*MergeTriageStoresResponse, error) {
	α := struct {
		OperationMergeTriageStores `xml:"tns:mergeTriageStores"`
	}{
		OperationMergeTriageStores{
			MergeTriageStores,
		},
	}

	γ := struct {
		OperationMergeTriageStoresResponse `xml:"mergeTriageStoresResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("MergeTriageStores", α, &γ); err != nil {
		return nil, err
	}
	return γ.MergeTriageStoresResponse, nil
}

// Notify was auto-generated from WSDL.
func (p *configurationService) Notify(Notify *Notify) (*NotifyResponse, error) {
	α := struct {
		OperationNotify `xml:"tns:notify"`
	}{
		OperationNotify{
			Notify,
		},
	}

	γ := struct {
		OperationNotifyResponse `xml:"notifyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("Notify", α, &γ); err != nil {
		return nil, err
	}
	return γ.NotifyResponse, nil
}

// SetAcceptingNewCommits was auto-generated from WSDL.
func (p *configurationService) SetAcceptingNewCommits(SetAcceptingNewCommits *SetAcceptingNewCommits) (*SetAcceptingNewCommitsResponse, error) {
	α := struct {
		OperationSetAcceptingNewCommits `xml:"tns:setAcceptingNewCommits"`
	}{
		OperationSetAcceptingNewCommits{
			SetAcceptingNewCommits,
		},
	}

	γ := struct {
		OperationSetAcceptingNewCommitsResponse `xml:"setAcceptingNewCommitsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SetAcceptingNewCommits", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetAcceptingNewCommitsResponse, nil
}

// SetArchitectureAnalysisConfiguration was auto-generated from
// WSDL.
func (p *configurationService) SetArchitectureAnalysisConfiguration(SetArchitectureAnalysisConfiguration *SetArchitectureAnalysisConfiguration) (*SetArchitectureAnalysisConfigurationResponse, error) {
	α := struct {
		OperationSetArchitectureAnalysisConfiguration `xml:"tns:setArchitectureAnalysisConfiguration"`
	}{
		OperationSetArchitectureAnalysisConfiguration{
			SetArchitectureAnalysisConfiguration,
		},
	}

	γ := struct {
		OperationSetArchitectureAnalysisConfigurationResponse `xml:"setArchitectureAnalysisConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SetArchitectureAnalysisConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetArchitectureAnalysisConfigurationResponse, nil
}

// SetBackupConfiguration was auto-generated from WSDL.
func (p *configurationService) SetBackupConfiguration(SetBackupConfiguration *SetBackupConfiguration) (*SetBackupConfigurationResponse, error) {
	α := struct {
		OperationSetBackupConfiguration `xml:"tns:setBackupConfiguration"`
	}{
		OperationSetBackupConfiguration{
			SetBackupConfiguration,
		},
	}

	γ := struct {
		OperationSetBackupConfigurationResponse `xml:"setBackupConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SetBackupConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetBackupConfigurationResponse, nil
}

// SetLoggingConfiguration was auto-generated from WSDL.
func (p *configurationService) SetLoggingConfiguration(SetLoggingConfiguration *SetLoggingConfiguration) (*SetLoggingConfigurationResponse, error) {
	α := struct {
		OperationSetLoggingConfiguration `xml:"tns:setLoggingConfiguration"`
	}{
		OperationSetLoggingConfiguration{
			SetLoggingConfiguration,
		},
	}

	γ := struct {
		OperationSetLoggingConfigurationResponse `xml:"setLoggingConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SetLoggingConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetLoggingConfigurationResponse, nil
}

// SetMessageOfTheDay was auto-generated from WSDL.
func (p *configurationService) SetMessageOfTheDay(SetMessageOfTheDay *SetMessageOfTheDay) (*SetMessageOfTheDayResponse, error) {
	α := struct {
		OperationSetMessageOfTheDay `xml:"tns:setMessageOfTheDay"`
	}{
		OperationSetMessageOfTheDay{
			SetMessageOfTheDay,
		},
	}

	γ := struct {
		OperationSetMessageOfTheDayResponse `xml:"setMessageOfTheDayResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SetMessageOfTheDay", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetMessageOfTheDayResponse, nil
}

// SetSkeletonizationConfiguration was auto-generated from WSDL.
func (p *configurationService) SetSkeletonizationConfiguration(SetSkeletonizationConfiguration *SetSkeletonizationConfiguration) (*SetSkeletonizationConfigurationResponse, error) {
	α := struct {
		OperationSetSkeletonizationConfiguration `xml:"tns:setSkeletonizationConfiguration"`
	}{
		OperationSetSkeletonizationConfiguration{
			SetSkeletonizationConfiguration,
		},
	}

	γ := struct {
		OperationSetSkeletonizationConfigurationResponse `xml:"setSkeletonizationConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SetSkeletonizationConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetSkeletonizationConfigurationResponse, nil
}

// SetSnapshotPurgeDetails was auto-generated from WSDL.
func (p *configurationService) SetSnapshotPurgeDetails(SetSnapshotPurgeDetails *SetSnapshotPurgeDetails) (*SetSnapshotPurgeDetailsResponse, error) {
	α := struct {
		OperationSetSnapshotPurgeDetails `xml:"tns:setSnapshotPurgeDetails"`
	}{
		OperationSetSnapshotPurgeDetails{
			SetSnapshotPurgeDetails,
		},
	}

	γ := struct {
		OperationSetSnapshotPurgeDetailsResponse `xml:"setSnapshotPurgeDetailsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SetSnapshotPurgeDetails", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetSnapshotPurgeDetailsResponse, nil
}

// UpdateAttribute was auto-generated from WSDL.
func (p *configurationService) UpdateAttribute(UpdateAttribute *UpdateAttribute) (*UpdateAttributeResponse, error) {
	α := struct {
		OperationUpdateAttribute `xml:"tns:updateAttribute"`
	}{
		OperationUpdateAttribute{
			UpdateAttribute,
		},
	}

	γ := struct {
		OperationUpdateAttributeResponse `xml:"updateAttributeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateAttribute", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateAttributeResponse, nil
}

// UpdateComponentMap was auto-generated from WSDL.
func (p *configurationService) UpdateComponentMap(UpdateComponentMap *UpdateComponentMap) (*UpdateComponentMapResponse, error) {
	α := struct {
		OperationUpdateComponentMap `xml:"tns:updateComponentMap"`
	}{
		OperationUpdateComponentMap{
			UpdateComponentMap,
		},
	}

	γ := struct {
		OperationUpdateComponentMapResponse `xml:"updateComponentMapResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateComponentMap", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateComponentMapResponse, nil
}

// UpdateGroup was auto-generated from WSDL.
func (p *configurationService) UpdateGroup(UpdateGroup *UpdateGroup) (*UpdateGroupResponse, error) {
	α := struct {
		OperationUpdateGroup `xml:"tns:updateGroup"`
	}{
		OperationUpdateGroup{
			UpdateGroup,
		},
	}

	γ := struct {
		OperationUpdateGroupResponse `xml:"updateGroupResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateGroup", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateGroupResponse, nil
}

// UpdateLdapConfiguration was auto-generated from WSDL.
func (p *configurationService) UpdateLdapConfiguration(UpdateLdapConfiguration *UpdateLdapConfiguration) (*UpdateLdapConfigurationResponse, error) {
	α := struct {
		OperationUpdateLdapConfiguration `xml:"tns:updateLdapConfiguration"`
	}{
		OperationUpdateLdapConfiguration{
			UpdateLdapConfiguration,
		},
	}

	γ := struct {
		OperationUpdateLdapConfigurationResponse `xml:"updateLdapConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateLdapConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateLdapConfigurationResponse, nil
}

// UpdateProject was auto-generated from WSDL.
func (p *configurationService) UpdateProject(UpdateProject *UpdateProject) (*UpdateProjectResponse, error) {
	α := struct {
		OperationUpdateProject `xml:"tns:updateProject"`
	}{
		OperationUpdateProject{
			UpdateProject,
		},
	}

	γ := struct {
		OperationUpdateProjectResponse `xml:"updateProjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateProject", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateProjectResponse, nil
}

// UpdateRole was auto-generated from WSDL.
func (p *configurationService) UpdateRole(UpdateRole *UpdateRole) (*UpdateRoleResponse, error) {
	α := struct {
		OperationUpdateRole `xml:"tns:updateRole"`
	}{
		OperationUpdateRole{
			UpdateRole,
		},
	}

	γ := struct {
		OperationUpdateRoleResponse `xml:"updateRoleResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateRole", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateRoleResponse, nil
}

// UpdateSignInConfiguration was auto-generated from WSDL.
func (p *configurationService) UpdateSignInConfiguration(UpdateSignInConfiguration *UpdateSignInConfiguration) (*UpdateSignInConfigurationResponse, error) {
	α := struct {
		OperationUpdateSignInConfiguration `xml:"tns:updateSignInConfiguration"`
	}{
		OperationUpdateSignInConfiguration{
			UpdateSignInConfiguration,
		},
	}

	γ := struct {
		OperationUpdateSignInConfigurationResponse `xml:"updateSignInConfigurationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateSignInConfiguration", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateSignInConfigurationResponse, nil
}

// UpdateSnapshotInfo was auto-generated from WSDL.
func (p *configurationService) UpdateSnapshotInfo(UpdateSnapshotInfo *UpdateSnapshotInfo) (*UpdateSnapshotInfoResponse, error) {
	α := struct {
		OperationUpdateSnapshotInfo `xml:"tns:updateSnapshotInfo"`
	}{
		OperationUpdateSnapshotInfo{
			UpdateSnapshotInfo,
		},
	}

	γ := struct {
		OperationUpdateSnapshotInfoResponse `xml:"updateSnapshotInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateSnapshotInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateSnapshotInfoResponse, nil
}

// UpdateStream was auto-generated from WSDL.
func (p *configurationService) UpdateStream(UpdateStream *UpdateStream) (*UpdateStreamResponse, error) {
	α := struct {
		OperationUpdateStream `xml:"tns:updateStream"`
	}{
		OperationUpdateStream{
			UpdateStream,
		},
	}

	γ := struct {
		OperationUpdateStreamResponse `xml:"updateStreamResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateStream", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateStreamResponse, nil
}

// UpdateTriageStore was auto-generated from WSDL.
func (p *configurationService) UpdateTriageStore(UpdateTriageStore *UpdateTriageStore) (*UpdateTriageStoreResponse, error) {
	α := struct {
		OperationUpdateTriageStore `xml:"tns:updateTriageStore"`
	}{
		OperationUpdateTriageStore{
			UpdateTriageStore,
		},
	}

	γ := struct {
		OperationUpdateTriageStoreResponse `xml:"updateTriageStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateTriageStore", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateTriageStoreResponse, nil
}

// UpdateUser was auto-generated from WSDL.
func (p *configurationService) UpdateUser(UpdateUser *UpdateUser) (*UpdateUserResponse, error) {
	α := struct {
		OperationUpdateUser `xml:"tns:updateUser"`
	}{
		OperationUpdateUser{
			UpdateUser,
		},
	}

	γ := struct {
		OperationUpdateUserResponse `xml:"updateUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateUserResponse, nil
}
