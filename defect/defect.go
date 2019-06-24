package defect

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://ws.coverity.com/v9"

// NewDefectService creates an initializes a DefectService.
func NewDefectService(cli *soap.Client) DefectService {
	return &defectService{cli}
}

// DefectService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type DefectService interface {
	// CreateMergedDefect was auto-generated from WSDL.
	CreateMergedDefect(CreateMergedDefect *CreateMergedDefect) (*CreateMergedDefectResponse, error)

	// GetComponentMetricsForProject was auto-generated from WSDL.
	GetComponentMetricsForProject(GetComponentMetricsForProject *GetComponentMetricsForProject) (*GetComponentMetricsForProjectResponse, error)

	// GetFileContents was auto-generated from WSDL.
	GetFileContents(GetFileContents *GetFileContents) (*GetFileContentsResponse, error)

	// GetMergedDefectDetectionHistory was auto-generated from WSDL.
	GetMergedDefectDetectionHistory(GetMergedDefectDetectionHistory *GetMergedDefectDetectionHistory) (*GetMergedDefectDetectionHistoryResponse, error)

	// GetMergedDefectHistory was auto-generated from WSDL.
	GetMergedDefectHistory(GetMergedDefectHistory *GetMergedDefectHistory) (*GetMergedDefectHistoryResponse, error)

	// GetMergedDefectsForProjectScope was auto-generated from WSDL.
	GetMergedDefectsForProjectScope(GetMergedDefectsForProjectScope *GetMergedDefectsForProjectScope) (*GetMergedDefectsForProjectScopeResponse, error)

	// GetMergedDefectsForSnapshotScope was auto-generated from WSDL.
	GetMergedDefectsForSnapshotScope(GetMergedDefectsForSnapshotScope *GetMergedDefectsForSnapshotScope) (*GetMergedDefectsForSnapshotScopeResponse, error)

	// GetMergedDefectsForStreams was auto-generated from WSDL.
	GetMergedDefectsForStreams(GetMergedDefectsForStreams *GetMergedDefectsForStreams) (*GetMergedDefectsForStreamsResponse, error)

	// GetStreamDefects was auto-generated from WSDL.
	GetStreamDefects(GetStreamDefects *GetStreamDefects) (*GetStreamDefectsResponse, error)

	// GetTrendRecordsForProject was auto-generated from WSDL.
	GetTrendRecordsForProject(GetTrendRecordsForProject *GetTrendRecordsForProject) (*GetTrendRecordsForProjectResponse, error)

	// GetTriageHistory was auto-generated from WSDL.
	GetTriageHistory(GetTriageHistory *GetTriageHistory) (*GetTriageHistoryResponse, error)

	// UpdateDefectInstanceProperties was auto-generated from WSDL.
	UpdateDefectInstanceProperties(UpdateDefectInstanceProperties *UpdateDefectInstanceProperties) (*UpdateDefectInstancePropertiesResponse, error)

	// UpdateStreamDefects was auto-generated from WSDL.
	UpdateStreamDefects(UpdateStreamDefects *UpdateStreamDefects) (*UpdateStreamDefectsResponse, error)

	// UpdateTriageForCIDsInTriageStore was auto-generated from WSDL.
	UpdateTriageForCIDsInTriageStore(UpdateTriageForCIDsInTriageStore *UpdateTriageForCIDsInTriageStore) (*UpdateTriageForCIDsInTriageStoreResponse, error)
}

// DateTime in WSDL format.
type DateTime string

// CovRemoteServiceException was auto-generated from WSDL.
type CovRemoteServiceException struct {
	ErrorCode *int    `xml:"errorCode,omitempty" json:"errorCode,omitempty" yaml:"errorCode,omitempty"`
	Message   *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// AttributeDefinitionIdDataObj was auto-generated from WSDL.
type AttributeDefinitionIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// AttributeDefinitionValueFilterMapDataObj was auto-generated
// from WSDL.
type AttributeDefinitionValueFilterMapDataObj struct {
	AttributeDefinitionId *AttributeDefinitionIdDataObj `xml:"attributeDefinitionId,omitempty" json:"attributeDefinitionId,omitempty" yaml:"attributeDefinitionId,omitempty"`
	AttributeValueIds     []*AttributeValueIdDataObj    `xml:"attributeValueIds,omitempty" json:"attributeValueIds,omitempty" yaml:"attributeValueIds,omitempty"`
}

// AttributeValueIdDataObj was auto-generated from WSDL.
type AttributeValueIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// ComponentIdDataObj was auto-generated from WSDL.
type ComponentIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// ComponentMetricsDataObj was auto-generated from WSDL.
type ComponentMetricsDataObj struct {
	BlankLineCount   *int                `xml:"blankLineCount,omitempty" json:"blankLineCount,omitempty" yaml:"blankLineCount,omitempty"`
	CodeLineCount    *int                `xml:"codeLineCount,omitempty" json:"codeLineCount,omitempty" yaml:"codeLineCount,omitempty"`
	CommentLineCount *int                `xml:"commentLineCount,omitempty" json:"commentLineCount,omitempty" yaml:"commentLineCount,omitempty"`
	ComponentId      *ComponentIdDataObj `xml:"componentId,omitempty" json:"componentId,omitempty" yaml:"componentId,omitempty"`
	DismissedCount   *int                `xml:"dismissedCount,omitempty" json:"dismissedCount,omitempty" yaml:"dismissedCount,omitempty"`
	FixedCount       *int                `xml:"fixedCount,omitempty" json:"fixedCount,omitempty" yaml:"fixedCount,omitempty"`
	MetricsDate      *DateTime           `xml:"metricsDate,omitempty" json:"metricsDate,omitempty" yaml:"metricsDate,omitempty"`
	NewCount         *int                `xml:"newCount,omitempty" json:"newCount,omitempty" yaml:"newCount,omitempty"`
	OutstandingCount *int                `xml:"outstandingCount,omitempty" json:"outstandingCount,omitempty" yaml:"outstandingCount,omitempty"`
	TotalCount       *int                `xml:"totalCount,omitempty" json:"totalCount,omitempty" yaml:"totalCount,omitempty"`
	TriagedCount     *int                `xml:"triagedCount,omitempty" json:"triagedCount,omitempty" yaml:"triagedCount,omitempty"`
}

// CreateMergedDefect was auto-generated from WSDL.
type CreateMergedDefect struct {
	MergeKey               *string   `xml:"mergeKey,omitempty" json:"mergeKey,omitempty" yaml:"mergeKey,omitempty"`
	DateOriginated         *DateTime `xml:"dateOriginated,omitempty" json:"dateOriginated,omitempty" yaml:"dateOriginated,omitempty"`
	ExternalPreventVersion *string   `xml:"externalPreventVersion,omitempty" json:"externalPreventVersion,omitempty" yaml:"externalPreventVersion,omitempty"`
	InternalPreventVersion *string   `xml:"internalPreventVersion,omitempty" json:"internalPreventVersion,omitempty" yaml:"internalPreventVersion,omitempty"`
	CheckerName            *string   `xml:"checkerName,omitempty" json:"checkerName,omitempty" yaml:"checkerName,omitempty"`
	DomainName             *string   `xml:"domainName,omitempty" json:"domainName,omitempty" yaml:"domainName,omitempty"`
}

// CreateMergedDefectResponse was auto-generated from WSDL.
type CreateMergedDefectResponse struct {
}

// DefectChangeDataObj was auto-generated from WSDL.
type DefectChangeDataObj struct {
	AffectedStreams  []*StreamIdDataObj    `xml:"affectedStreams,omitempty" json:"affectedStreams,omitempty" yaml:"affectedStreams,omitempty"`
	AttributeChanges []*FieldChangeDataObj `xml:"attributeChanges,omitempty" json:"attributeChanges,omitempty" yaml:"attributeChanges,omitempty"`
	Comments         *string               `xml:"comments,omitempty" json:"comments,omitempty" yaml:"comments,omitempty"`
	DateModified     *DateTime             `xml:"dateModified,omitempty" json:"dateModified,omitempty" yaml:"dateModified,omitempty"`
	UserModified     *string               `xml:"userModified,omitempty" json:"userModified,omitempty" yaml:"userModified,omitempty"`
}

// DefectDetectionHistoryDataObj was auto-generated from WSDL.
type DefectDetectionHistoryDataObj struct {
	DefectDetection   *string            `xml:"defectDetection,omitempty" json:"defectDetection,omitempty" yaml:"defectDetection,omitempty"`
	Detection         *DateTime          `xml:"detection,omitempty" json:"detection,omitempty" yaml:"detection,omitempty"`
	InCurrentSnapshot *bool              `xml:"inCurrentSnapshot,omitempty" json:"inCurrentSnapshot,omitempty" yaml:"inCurrentSnapshot,omitempty"`
	SnapshotId        *int64             `xml:"snapshotId,omitempty" json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
	Streams           []*StreamIdDataObj `xml:"streams,omitempty" json:"streams,omitempty" yaml:"streams,omitempty"`
	UserName          *string            `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
}

// DefectInstanceDataObj was auto-generated from WSDL.
type DefectInstanceDataObj struct {
	Events           []*EventDataObj          `xml:"events,omitempty" json:"events,omitempty" yaml:"events,omitempty"`
	Properties       []*PropertyDataObj       `xml:"properties,omitempty" json:"properties,omitempty" yaml:"properties,omitempty"`
	Category         *LocalizedValueDataObj   `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	CheckerName      *string                  `xml:"checkerName,omitempty" json:"checkerName,omitempty" yaml:"checkerName,omitempty"`
	Component        *string                  `xml:"component,omitempty" json:"component,omitempty" yaml:"component,omitempty"`
	Cwe              *int                     `xml:"cwe,omitempty" json:"cwe,omitempty" yaml:"cwe,omitempty"`
	Domain           *string                  `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	EventSetCaptions []*string                `xml:"eventSetCaptions,omitempty" json:"eventSetCaptions,omitempty" yaml:"eventSetCaptions,omitempty"`
	Extra            *string                  `xml:"extra,omitempty" json:"extra,omitempty" yaml:"extra,omitempty"`
	Function         *FunctionInfoDataObj     `xml:"function,omitempty" json:"function,omitempty" yaml:"function,omitempty"`
	Id               *DefectInstanceIdDataObj `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Impact           *LocalizedValueDataObj   `xml:"impact,omitempty" json:"impact,omitempty" yaml:"impact,omitempty"`
	IssueKinds       []*LocalizedValueDataObj `xml:"issueKinds,omitempty" json:"issueKinds,omitempty" yaml:"issueKinds,omitempty"`
	LocalEffect      *string                  `xml:"localEffect,omitempty" json:"localEffect,omitempty" yaml:"localEffect,omitempty"`
	LongDescription  *string                  `xml:"longDescription,omitempty" json:"longDescription,omitempty" yaml:"longDescription,omitempty"`
	Subcategory      *string                  `xml:"subcategory,omitempty" json:"subcategory,omitempty" yaml:"subcategory,omitempty"`
	Type             *LocalizedValueDataObj   `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// DefectInstanceIdDataObj was auto-generated from WSDL.
type DefectInstanceIdDataObj struct {
	Id *int64 `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// DefectStateAttributeValueDataObj was auto-generated from WSDL.
type DefectStateAttributeValueDataObj struct {
	AttributeDefinitionId *AttributeDefinitionIdDataObj `xml:"attributeDefinitionId,omitempty" json:"attributeDefinitionId,omitempty" yaml:"attributeDefinitionId,omitempty"`
	AttributeValueId      *AttributeValueIdDataObj      `xml:"attributeValueId,omitempty" json:"attributeValueId,omitempty" yaml:"attributeValueId,omitempty"`
}

// DefectStateDataObj was auto-generated from WSDL.
type DefectStateDataObj struct {
	DateCreated                *DateTime                           `xml:"dateCreated,omitempty" json:"dateCreated,omitempty" yaml:"dateCreated,omitempty"`
	DefectStateAttributeValues []*DefectStateAttributeValueDataObj `xml:"defectStateAttributeValues,omitempty" json:"defectStateAttributeValues,omitempty" yaml:"defectStateAttributeValues,omitempty"`
	UserCreated                *string                             `xml:"userCreated,omitempty" json:"userCreated,omitempty" yaml:"userCreated,omitempty"`
}

// DefectStateSpecDataObj was auto-generated from WSDL.
type DefectStateSpecDataObj struct {
	DefectStateAttributeValues []*DefectStateAttributeValueDataObj `xml:"defectStateAttributeValues,omitempty" json:"defectStateAttributeValues,omitempty" yaml:"defectStateAttributeValues,omitempty"`
}

// EventDataObj was auto-generated from WSDL.
type EventDataObj struct {
	Events            []*EventDataObj `xml:"events,omitempty" json:"events,omitempty" yaml:"events,omitempty"`
	EventDescription  *string         `xml:"eventDescription,omitempty" json:"eventDescription,omitempty" yaml:"eventDescription,omitempty"`
	EventKind         *string         `xml:"eventKind,omitempty" json:"eventKind,omitempty" yaml:"eventKind,omitempty"`
	EventNumber       *int            `xml:"eventNumber,omitempty" json:"eventNumber,omitempty" yaml:"eventNumber,omitempty"`
	EventSet          *int            `xml:"eventSet,omitempty" json:"eventSet,omitempty" yaml:"eventSet,omitempty"`
	EventTag          *string         `xml:"eventTag,omitempty" json:"eventTag,omitempty" yaml:"eventTag,omitempty"`
	FileId            *FileIdDataObj  `xml:"fileId,omitempty" json:"fileId,omitempty" yaml:"fileId,omitempty"`
	Id                *int64          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	LineNumber        *int            `xml:"lineNumber,omitempty" json:"lineNumber,omitempty" yaml:"lineNumber,omitempty"`
	Main              *bool           `xml:"main,omitempty" json:"main,omitempty" yaml:"main,omitempty"`
	MoreInformationId *string         `xml:"moreInformationId,omitempty" json:"moreInformationId,omitempty" yaml:"moreInformationId,omitempty"`
	PathCondition     *string         `xml:"pathCondition,omitempty" json:"pathCondition,omitempty" yaml:"pathCondition,omitempty"`
	Polarity          *bool           `xml:"polarity,omitempty" json:"polarity,omitempty" yaml:"polarity,omitempty"`
}

// FieldChangeDataObj was auto-generated from WSDL.
type FieldChangeDataObj struct {
	FieldName *string `xml:"fieldName,omitempty" json:"fieldName,omitempty" yaml:"fieldName,omitempty"`
	NewValue  *string `xml:"newValue,omitempty" json:"newValue,omitempty" yaml:"newValue,omitempty"`
	OldValue  *string `xml:"oldValue,omitempty" json:"oldValue,omitempty" yaml:"oldValue,omitempty"`
}

// FileContentsDataObj was auto-generated from WSDL.
type FileContentsDataObj struct {
	Contents *[]byte        `xml:"contents,omitempty" json:"contents,omitempty" yaml:"contents,omitempty"`
	FileId   *FileIdDataObj `xml:"fileId,omitempty" json:"fileId,omitempty" yaml:"fileId,omitempty"`
}

// FileIdDataObj was auto-generated from WSDL.
type FileIdDataObj struct {
	ContentsMD5  *string `xml:"contentsMD5,omitempty" json:"contentsMD5,omitempty" yaml:"contentsMD5,omitempty"`
	FilePathname *string `xml:"filePathname,omitempty" json:"filePathname,omitempty" yaml:"filePathname,omitempty"`
}

// FunctionInfoDataObj was auto-generated from WSDL.
type FunctionInfoDataObj struct {
	FileId              *FileIdDataObj `xml:"fileId,omitempty" json:"fileId,omitempty" yaml:"fileId,omitempty"`
	FunctionDisplayName *string        `xml:"functionDisplayName,omitempty" json:"functionDisplayName,omitempty" yaml:"functionDisplayName,omitempty"`
	FunctionMangledName *string        `xml:"functionMangledName,omitempty" json:"functionMangledName,omitempty" yaml:"functionMangledName,omitempty"`
	FunctionMergeName   *string        `xml:"functionMergeName,omitempty" json:"functionMergeName,omitempty" yaml:"functionMergeName,omitempty"`
}

// GetComponentMetricsForProject was auto-generated from WSDL.
type GetComponentMetricsForProject struct {
	ProjectId    *ProjectIdDataObj     `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	ComponentIds []*ComponentIdDataObj `xml:"componentIds,omitempty" json:"componentIds,omitempty" yaml:"componentIds,omitempty"`
}

// GetComponentMetricsForProjectResponse was auto-generated from
// WSDL.
type GetComponentMetricsForProjectResponse struct {
	Return []*ComponentMetricsDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetFileContents was auto-generated from WSDL.
type GetFileContents struct {
	StreamId *StreamIdDataObj `xml:"streamId,omitempty" json:"streamId,omitempty" yaml:"streamId,omitempty"`
	FileId   *FileIdDataObj   `xml:"fileId,omitempty" json:"fileId,omitempty" yaml:"fileId,omitempty"`
}

// GetFileContentsResponse was auto-generated from WSDL.
type GetFileContentsResponse struct {
	Return *FileContentsDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMergedDefectDetectionHistory was auto-generated from WSDL.
type GetMergedDefectDetectionHistory struct {
	MergedDefectIdDataObj *MergedDefectIdDataObj `xml:"mergedDefectIdDataObj,omitempty" json:"mergedDefectIdDataObj,omitempty" yaml:"mergedDefectIdDataObj,omitempty"`
	StreamIds             []*StreamIdDataObj     `xml:"streamIds,omitempty" json:"streamIds,omitempty" yaml:"streamIds,omitempty"`
}

// GetMergedDefectDetectionHistoryResponse was auto-generated from
// WSDL.
type GetMergedDefectDetectionHistoryResponse struct {
	Return []*DefectDetectionHistoryDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMergedDefectHistory was auto-generated from WSDL.
type GetMergedDefectHistory struct {
	MergedDefectIdDataObj *MergedDefectIdDataObj `xml:"mergedDefectIdDataObj,omitempty" json:"mergedDefectIdDataObj,omitempty" yaml:"mergedDefectIdDataObj,omitempty"`
	StreamIds             []*StreamIdDataObj     `xml:"streamIds,omitempty" json:"streamIds,omitempty" yaml:"streamIds,omitempty"`
}

// GetMergedDefectHistoryResponse was auto-generated from WSDL.
type GetMergedDefectHistoryResponse struct {
	Return []*DefectChangeDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMergedDefectsForProjectScope was auto-generated from WSDL.
type GetMergedDefectsForProjectScope struct {
	ProjectId  *ProjectIdDataObj                    `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	FilterSpec *ProjectScopeDefectFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
	PageSpec   *PageSpecDataObj                     `xml:"pageSpec,omitempty" json:"pageSpec,omitempty" yaml:"pageSpec,omitempty"`
}

// GetMergedDefectsForProjectScopeResponse was auto-generated from
// WSDL.
type GetMergedDefectsForProjectScopeResponse struct {
	Return *MergedDefectsPageDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMergedDefectsForSnapshotScope was auto-generated from WSDL.
type GetMergedDefectsForSnapshotScope struct {
	ProjectId     *ProjectIdDataObj                     `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	FilterSpec    *SnapshotScopeDefectFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
	PageSpec      *PageSpecDataObj                      `xml:"pageSpec,omitempty" json:"pageSpec,omitempty" yaml:"pageSpec,omitempty"`
	SnapshotScope *SnapshotScopeSpecDataObj             `xml:"snapshotScope,omitempty" json:"snapshotScope,omitempty" yaml:"snapshotScope,omitempty"`
}

// GetMergedDefectsForSnapshotScopeResponse was auto-generated
// from WSDL.
type GetMergedDefectsForSnapshotScopeResponse struct {
	Return *MergedDefectsPageDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMergedDefectsForStreams was auto-generated from WSDL.
type GetMergedDefectsForStreams struct {
	StreamIds     []*StreamIdDataObj             `xml:"streamIds,omitempty" json:"streamIds,omitempty" yaml:"streamIds,omitempty"`
	FilterSpec    *MergedDefectFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
	PageSpec      *PageSpecDataObj               `xml:"pageSpec,omitempty" json:"pageSpec,omitempty" yaml:"pageSpec,omitempty"`
	SnapshotScope *SnapshotScopeSpecDataObj      `xml:"snapshotScope,omitempty" json:"snapshotScope,omitempty" yaml:"snapshotScope,omitempty"`
}

// GetMergedDefectsForStreamsResponse was auto-generated from WSDL.
type GetMergedDefectsForStreamsResponse struct {
	Return *MergedDefectsPageDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetStreamDefects was auto-generated from WSDL.
type GetStreamDefects struct {
	MergedDefectIdDataObjs []*MergedDefectIdDataObj       `xml:"mergedDefectIdDataObjs,omitempty" json:"mergedDefectIdDataObjs,omitempty" yaml:"mergedDefectIdDataObjs,omitempty"`
	FilterSpec             *StreamDefectFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetStreamDefectsResponse was auto-generated from WSDL.
type GetStreamDefectsResponse struct {
	Return []*StreamDefectDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetTrendRecordsForProject was auto-generated from WSDL.
type GetTrendRecordsForProject struct {
	ProjectId  *ProjectIdDataObj                    `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	FilterSpec *ProjectTrendRecordFilterSpecDataObj `xml:"filterSpec,omitempty" json:"filterSpec,omitempty" yaml:"filterSpec,omitempty"`
}

// GetTrendRecordsForProjectResponse was auto-generated from WSDL.
type GetTrendRecordsForProjectResponse struct {
	Return []*ProjectMetricsDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetTriageHistory was auto-generated from WSDL.
type GetTriageHistory struct {
	MergedDefectIdDataObj *MergedDefectIdDataObj  `xml:"mergedDefectIdDataObj,omitempty" json:"mergedDefectIdDataObj,omitempty" yaml:"mergedDefectIdDataObj,omitempty"`
	TriageStoreIds        []*TriageStoreIdDataObj `xml:"triageStoreIds,omitempty" json:"triageStoreIds,omitempty" yaml:"triageStoreIds,omitempty"`
}

// GetTriageHistoryResponse was auto-generated from WSDL.
type GetTriageHistoryResponse struct {
	Return []*TriageHistoryDataObj `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LocalizedValueDataObj was auto-generated from WSDL.
type LocalizedValueDataObj struct {
	DisplayName *string `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Name        *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// MergedDefectDataObj was auto-generated from WSDL.
type MergedDefectDataObj struct {
	CheckerName                *string                             `xml:"checkerName,omitempty" json:"checkerName,omitempty" yaml:"checkerName,omitempty"`
	Cid                        *int64                              `xml:"cid,omitempty" json:"cid,omitempty" yaml:"cid,omitempty"`
	ComponentName              *string                             `xml:"componentName,omitempty" json:"componentName,omitempty" yaml:"componentName,omitempty"`
	Cwe                        *int                                `xml:"cwe,omitempty" json:"cwe,omitempty" yaml:"cwe,omitempty"`
	DefectStateAttributeValues []*DefectStateAttributeValueDataObj `xml:"defectStateAttributeValues,omitempty" json:"defectStateAttributeValues,omitempty" yaml:"defectStateAttributeValues,omitempty"`
	DisplayCategory            *string                             `xml:"displayCategory,omitempty" json:"displayCategory,omitempty" yaml:"displayCategory,omitempty"`
	DisplayImpact              *string                             `xml:"displayImpact,omitempty" json:"displayImpact,omitempty" yaml:"displayImpact,omitempty"`
	DisplayIssueKind           *string                             `xml:"displayIssueKind,omitempty" json:"displayIssueKind,omitempty" yaml:"displayIssueKind,omitempty"`
	DisplayType                *string                             `xml:"displayType,omitempty" json:"displayType,omitempty" yaml:"displayType,omitempty"`
	Domain                     *string                             `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	FilePathname               *string                             `xml:"filePathname,omitempty" json:"filePathname,omitempty" yaml:"filePathname,omitempty"`
	FirstDetected              *DateTime                           `xml:"firstDetected,omitempty" json:"firstDetected,omitempty" yaml:"firstDetected,omitempty"`
	FirstDetectedBy            *string                             `xml:"firstDetectedBy,omitempty" json:"firstDetectedBy,omitempty" yaml:"firstDetectedBy,omitempty"`
	FirstDetectedDescription   *string                             `xml:"firstDetectedDescription,omitempty" json:"firstDetectedDescription,omitempty" yaml:"firstDetectedDescription,omitempty"`
	FirstDetectedSnapshotId    *int64                              `xml:"firstDetectedSnapshotId,omitempty" json:"firstDetectedSnapshotId,omitempty" yaml:"firstDetectedSnapshotId,omitempty"`
	FirstDetectedStream        *string                             `xml:"firstDetectedStream,omitempty" json:"firstDetectedStream,omitempty" yaml:"firstDetectedStream,omitempty"`
	FirstDetectedTarget        *string                             `xml:"firstDetectedTarget,omitempty" json:"firstDetectedTarget,omitempty" yaml:"firstDetectedTarget,omitempty"`
	FirstDetectedVersion       *string                             `xml:"firstDetectedVersion,omitempty" json:"firstDetectedVersion,omitempty" yaml:"firstDetectedVersion,omitempty"`
	FunctionDisplayName        *string                             `xml:"functionDisplayName,omitempty" json:"functionDisplayName,omitempty" yaml:"functionDisplayName,omitempty"`
	FunctionMergeName          *string                             `xml:"functionMergeName,omitempty" json:"functionMergeName,omitempty" yaml:"functionMergeName,omitempty"`
	FunctionName               *string                             `xml:"functionName,omitempty" json:"functionName,omitempty" yaml:"functionName,omitempty"`
	IssueKind                  *string                             `xml:"issueKind,omitempty" json:"issueKind,omitempty" yaml:"issueKind,omitempty"`
	LastDetected               *DateTime                           `xml:"lastDetected,omitempty" json:"lastDetected,omitempty" yaml:"lastDetected,omitempty"`
	LastDetectedDescription    *string                             `xml:"lastDetectedDescription,omitempty" json:"lastDetectedDescription,omitempty" yaml:"lastDetectedDescription,omitempty"`
	LastDetectedSnapshotId     *int64                              `xml:"lastDetectedSnapshotId,omitempty" json:"lastDetectedSnapshotId,omitempty" yaml:"lastDetectedSnapshotId,omitempty"`
	LastDetectedStream         *string                             `xml:"lastDetectedStream,omitempty" json:"lastDetectedStream,omitempty" yaml:"lastDetectedStream,omitempty"`
	LastDetectedTarget         *string                             `xml:"lastDetectedTarget,omitempty" json:"lastDetectedTarget,omitempty" yaml:"lastDetectedTarget,omitempty"`
	LastDetectedVersion        *string                             `xml:"lastDetectedVersion,omitempty" json:"lastDetectedVersion,omitempty" yaml:"lastDetectedVersion,omitempty"`
	LastFixed                  *DateTime                           `xml:"lastFixed,omitempty" json:"lastFixed,omitempty" yaml:"lastFixed,omitempty"`
	LastTriaged                *DateTime                           `xml:"lastTriaged,omitempty" json:"lastTriaged,omitempty" yaml:"lastTriaged,omitempty"`
	MergeKey                   *string                             `xml:"mergeKey,omitempty" json:"mergeKey,omitempty" yaml:"mergeKey,omitempty"`
	OccurrenceCount            *int                                `xml:"occurrenceCount,omitempty" json:"occurrenceCount,omitempty" yaml:"occurrenceCount,omitempty"`
}

// MergedDefectFilterSpecDataObj was auto-generated from WSDL.
type MergedDefectFilterSpecDataObj struct {
	CidList                           []*int64                                    `xml:"cidList,omitempty" json:"cidList,omitempty" yaml:"cidList,omitempty"`
	FilenamePatternList               []*string                                   `xml:"filenamePatternList,omitempty" json:"filenamePatternList,omitempty" yaml:"filenamePatternList,omitempty"`
	ComponentIdList                   []*ComponentIdDataObj                       `xml:"componentIdList,omitempty" json:"componentIdList,omitempty" yaml:"componentIdList,omitempty"`
	StatusNameList                    []*string                                   `xml:"statusNameList,omitempty" json:"statusNameList,omitempty" yaml:"statusNameList,omitempty"`
	ClassificationNameList            []*string                                   `xml:"classificationNameList,omitempty" json:"classificationNameList,omitempty" yaml:"classificationNameList,omitempty"`
	ActionNameList                    []*string                                   `xml:"actionNameList,omitempty" json:"actionNameList,omitempty" yaml:"actionNameList,omitempty"`
	FixTargetNameList                 []*string                                   `xml:"fixTargetNameList,omitempty" json:"fixTargetNameList,omitempty" yaml:"fixTargetNameList,omitempty"`
	SeverityNameList                  []*string                                   `xml:"severityNameList,omitempty" json:"severityNameList,omitempty" yaml:"severityNameList,omitempty"`
	LegacyNameList                    []*string                                   `xml:"legacyNameList,omitempty" json:"legacyNameList,omitempty" yaml:"legacyNameList,omitempty"`
	OwnerNameList                     []*string                                   `xml:"ownerNameList,omitempty" json:"ownerNameList,omitempty" yaml:"ownerNameList,omitempty"`
	CheckerList                       []*string                                   `xml:"checkerList,omitempty" json:"checkerList,omitempty" yaml:"checkerList,omitempty"`
	CweList                           []*int                                      `xml:"cweList,omitempty" json:"cweList,omitempty" yaml:"cweList,omitempty"`
	CheckerCategoryList               []*string                                   `xml:"checkerCategoryList,omitempty" json:"checkerCategoryList,omitempty" yaml:"checkerCategoryList,omitempty"`
	CheckerTypeList                   []*string                                   `xml:"checkerTypeList,omitempty" json:"checkerTypeList,omitempty" yaml:"checkerTypeList,omitempty"`
	ImpactList                        []*string                                   `xml:"impactList,omitempty" json:"impactList,omitempty" yaml:"impactList,omitempty"`
	IssueKindList                     []*string                                   `xml:"issueKindList,omitempty" json:"issueKindList,omitempty" yaml:"issueKindList,omitempty"`
	AttributeDefinitionValueFilterMap []*AttributeDefinitionValueFilterMapDataObj `xml:"attributeDefinitionValueFilterMap,omitempty" json:"attributeDefinitionValueFilterMap,omitempty" yaml:"attributeDefinitionValueFilterMap,omitempty"`
	ComponentIdExclude                *bool                                       `xml:"componentIdExclude,omitempty" json:"componentIdExclude,omitempty" yaml:"componentIdExclude,omitempty"`
	DefectPropertyKey                 *string                                     `xml:"defectPropertyKey,omitempty" json:"defectPropertyKey,omitempty" yaml:"defectPropertyKey,omitempty"`
	DefectPropertyPattern             *string                                     `xml:"defectPropertyPattern,omitempty" json:"defectPropertyPattern,omitempty" yaml:"defectPropertyPattern,omitempty"`
	ExternalReferencePattern          *string                                     `xml:"externalReferencePattern,omitempty" json:"externalReferencePattern,omitempty" yaml:"externalReferencePattern,omitempty"`
	FirstDetectedEndDate              *DateTime                                   `xml:"firstDetectedEndDate,omitempty" json:"firstDetectedEndDate,omitempty" yaml:"firstDetectedEndDate,omitempty"`
	FirstDetectedStartDate            *DateTime                                   `xml:"firstDetectedStartDate,omitempty" json:"firstDetectedStartDate,omitempty" yaml:"firstDetectedStartDate,omitempty"`
	FunctionNamePattern               *string                                     `xml:"functionNamePattern,omitempty" json:"functionNamePattern,omitempty" yaml:"functionNamePattern,omitempty"`
	LastDetectedEndDate               *DateTime                                   `xml:"lastDetectedEndDate,omitempty" json:"lastDetectedEndDate,omitempty" yaml:"lastDetectedEndDate,omitempty"`
	LastDetectedStartDate             *DateTime                                   `xml:"lastDetectedStartDate,omitempty" json:"lastDetectedStartDate,omitempty" yaml:"lastDetectedStartDate,omitempty"`
	LastFixedEndDate                  *DateTime                                   `xml:"lastFixedEndDate,omitempty" json:"lastFixedEndDate,omitempty" yaml:"lastFixedEndDate,omitempty"`
	LastFixedStartDate                *DateTime                                   `xml:"lastFixedStartDate,omitempty" json:"lastFixedStartDate,omitempty" yaml:"lastFixedStartDate,omitempty"`
	LastTriagedEndDate                *DateTime                                   `xml:"lastTriagedEndDate,omitempty" json:"lastTriagedEndDate,omitempty" yaml:"lastTriagedEndDate,omitempty"`
	LastTriagedStartDate              *DateTime                                   `xml:"lastTriagedStartDate,omitempty" json:"lastTriagedStartDate,omitempty" yaml:"lastTriagedStartDate,omitempty"`
	MaxCid                            *int64                                      `xml:"maxCid,omitempty" json:"maxCid,omitempty" yaml:"maxCid,omitempty"`
	MaxOccurrenceCount                *int                                        `xml:"maxOccurrenceCount,omitempty" json:"maxOccurrenceCount,omitempty" yaml:"maxOccurrenceCount,omitempty"`
	MergedDefectIdDataObjs            []*MergedDefectIdDataObj                    `xml:"mergedDefectIdDataObjs,omitempty" json:"mergedDefectIdDataObjs,omitempty" yaml:"mergedDefectIdDataObjs,omitempty"`
	MinCid                            *int64                                      `xml:"minCid,omitempty" json:"minCid,omitempty" yaml:"minCid,omitempty"`
	MinOccurrenceCount                *int                                        `xml:"minOccurrenceCount,omitempty" json:"minOccurrenceCount,omitempty" yaml:"minOccurrenceCount,omitempty"`
	OwnerNamePattern                  *string                                     `xml:"ownerNamePattern,omitempty" json:"ownerNamePattern,omitempty" yaml:"ownerNamePattern,omitempty"`
	SnapshotComparisonField           *string                                     `xml:"snapshotComparisonField,omitempty" json:"snapshotComparisonField,omitempty" yaml:"snapshotComparisonField,omitempty"`
	StreamExcludeNameList             []*StreamIdDataObj                          `xml:"streamExcludeNameList,omitempty" json:"streamExcludeNameList,omitempty" yaml:"streamExcludeNameList,omitempty"`
	StreamExcludeQualifier            *string                                     `xml:"streamExcludeQualifier,omitempty" json:"streamExcludeQualifier,omitempty" yaml:"streamExcludeQualifier,omitempty"`
	StreamIncludeNameList             []*StreamIdDataObj                          `xml:"streamIncludeNameList,omitempty" json:"streamIncludeNameList,omitempty" yaml:"streamIncludeNameList,omitempty"`
	StreamIncludeQualifier            *string                                     `xml:"streamIncludeQualifier,omitempty" json:"streamIncludeQualifier,omitempty" yaml:"streamIncludeQualifier,omitempty"`
}

// MergedDefectIdDataObj was auto-generated from WSDL.
type MergedDefectIdDataObj struct {
	Cid      *int64  `xml:"cid,omitempty" json:"cid,omitempty" yaml:"cid,omitempty"`
	MergeKey *string `xml:"mergeKey,omitempty" json:"mergeKey,omitempty" yaml:"mergeKey,omitempty"`
}

// MergedDefectsPageDataObj was auto-generated from WSDL.
type MergedDefectsPageDataObj struct {
	MergedDefectIds      []*MergedDefectIdDataObj `xml:"mergedDefectIds,omitempty" json:"mergedDefectIds,omitempty" yaml:"mergedDefectIds,omitempty"`
	MergedDefects        []*MergedDefectDataObj   `xml:"mergedDefects,omitempty" json:"mergedDefects,omitempty" yaml:"mergedDefects,omitempty"`
	TotalNumberOfRecords *int                     `xml:"totalNumberOfRecords,omitempty" json:"totalNumberOfRecords,omitempty" yaml:"totalNumberOfRecords,omitempty"`
}

// PageSpecDataObj was auto-generated from WSDL.
type PageSpecDataObj struct {
	PageSize      *int    `xml:"pageSize,omitempty" json:"pageSize,omitempty" yaml:"pageSize,omitempty"`
	SortAscending *bool   `xml:"sortAscending,omitempty" json:"sortAscending,omitempty" yaml:"sortAscending,omitempty"`
	SortField     *string `xml:"sortField,omitempty" json:"sortField,omitempty" yaml:"sortField,omitempty"`
	StartIndex    *int    `xml:"startIndex,omitempty" json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
}

// ProjectIdDataObj was auto-generated from WSDL.
type ProjectIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// ProjectMetricsDataObj was auto-generated from WSDL.
type ProjectMetricsDataObj struct {
	BlankLineCount   *int              `xml:"blankLineCount,omitempty" json:"blankLineCount,omitempty" yaml:"blankLineCount,omitempty"`
	CodeLineCount    *int              `xml:"codeLineCount,omitempty" json:"codeLineCount,omitempty" yaml:"codeLineCount,omitempty"`
	CommentLineCount *int              `xml:"commentLineCount,omitempty" json:"commentLineCount,omitempty" yaml:"commentLineCount,omitempty"`
	DismissedCount   *int              `xml:"dismissedCount,omitempty" json:"dismissedCount,omitempty" yaml:"dismissedCount,omitempty"`
	FixedCount       *int              `xml:"fixedCount,omitempty" json:"fixedCount,omitempty" yaml:"fixedCount,omitempty"`
	InspectedCount   *int              `xml:"inspectedCount,omitempty" json:"inspectedCount,omitempty" yaml:"inspectedCount,omitempty"`
	MetricsDate      *DateTime         `xml:"metricsDate,omitempty" json:"metricsDate,omitempty" yaml:"metricsDate,omitempty"`
	NewCount         *int              `xml:"newCount,omitempty" json:"newCount,omitempty" yaml:"newCount,omitempty"`
	OutstandingCount *int              `xml:"outstandingCount,omitempty" json:"outstandingCount,omitempty" yaml:"outstandingCount,omitempty"`
	ProjectId        *ProjectIdDataObj `xml:"projectId,omitempty" json:"projectId,omitempty" yaml:"projectId,omitempty"`
	ResolvedCount    *int              `xml:"resolvedCount,omitempty" json:"resolvedCount,omitempty" yaml:"resolvedCount,omitempty"`
	TotalCount       *int              `xml:"totalCount,omitempty" json:"totalCount,omitempty" yaml:"totalCount,omitempty"`
	TriagedCount     *int              `xml:"triagedCount,omitempty" json:"triagedCount,omitempty" yaml:"triagedCount,omitempty"`
}

// ProjectScopeDefectFilterSpecDataObj was auto-generated from
// WSDL.
type ProjectScopeDefectFilterSpecDataObj struct {
	ActionNameList         []*string             `xml:"actionNameList,omitempty" json:"actionNameList,omitempty" yaml:"actionNameList,omitempty"`
	CheckerCategoryList    []*string             `xml:"checkerCategoryList,omitempty" json:"checkerCategoryList,omitempty" yaml:"checkerCategoryList,omitempty"`
	CheckerList            []*string             `xml:"checkerList,omitempty" json:"checkerList,omitempty" yaml:"checkerList,omitempty"`
	CheckerTypeList        []*string             `xml:"checkerTypeList,omitempty" json:"checkerTypeList,omitempty" yaml:"checkerTypeList,omitempty"`
	CidList                []*int64              `xml:"cidList,omitempty" json:"cidList,omitempty" yaml:"cidList,omitempty"`
	ClassificationNameList []*string             `xml:"classificationNameList,omitempty" json:"classificationNameList,omitempty" yaml:"classificationNameList,omitempty"`
	ComponentIdExclude     *bool                 `xml:"componentIdExclude,omitempty" json:"componentIdExclude,omitempty" yaml:"componentIdExclude,omitempty"`
	ComponentIdList        []*ComponentIdDataObj `xml:"componentIdList,omitempty" json:"componentIdList,omitempty" yaml:"componentIdList,omitempty"`
	CweList                []*int64              `xml:"cweList,omitempty" json:"cweList,omitempty" yaml:"cweList,omitempty"`
	FirstDetectedBy        []*string             `xml:"firstDetectedBy,omitempty" json:"firstDetectedBy,omitempty" yaml:"firstDetectedBy,omitempty"`
	FirstDetectedEndDate   *DateTime             `xml:"firstDetectedEndDate,omitempty" json:"firstDetectedEndDate,omitempty" yaml:"firstDetectedEndDate,omitempty"`
	FirstDetectedStartDate *DateTime             `xml:"firstDetectedStartDate,omitempty" json:"firstDetectedStartDate,omitempty" yaml:"firstDetectedStartDate,omitempty"`
	FixTargetNameList      []*string             `xml:"fixTargetNameList,omitempty" json:"fixTargetNameList,omitempty" yaml:"fixTargetNameList,omitempty"`
	ImpactNameList         []*string             `xml:"impactNameList,omitempty" json:"impactNameList,omitempty" yaml:"impactNameList,omitempty"`
	IssueKindList          []*string             `xml:"issueKindList,omitempty" json:"issueKindList,omitempty" yaml:"issueKindList,omitempty"`
	LastDetectedEndDate    *DateTime             `xml:"lastDetectedEndDate,omitempty" json:"lastDetectedEndDate,omitempty" yaml:"lastDetectedEndDate,omitempty"`
	LastDetectedStartDate  *DateTime             `xml:"lastDetectedStartDate,omitempty" json:"lastDetectedStartDate,omitempty" yaml:"lastDetectedStartDate,omitempty"`
	LegacyNameList         []*string             `xml:"legacyNameList,omitempty" json:"legacyNameList,omitempty" yaml:"legacyNameList,omitempty"`
	OwnerNameList          []*string             `xml:"ownerNameList,omitempty" json:"ownerNameList,omitempty" yaml:"ownerNameList,omitempty"`
	OwnerNamePattern       *string               `xml:"ownerNamePattern,omitempty" json:"ownerNamePattern,omitempty" yaml:"ownerNamePattern,omitempty"`
	SeverityNameList       []*string             `xml:"severityNameList,omitempty" json:"severityNameList,omitempty" yaml:"severityNameList,omitempty"`
}

// ProjectTrendRecordFilterSpecDataObj was auto-generated from
// WSDL.
type ProjectTrendRecordFilterSpecDataObj struct {
	EndDate   *DateTime `xml:"endDate,omitempty" json:"endDate,omitempty" yaml:"endDate,omitempty"`
	StartDate *DateTime `xml:"startDate,omitempty" json:"startDate,omitempty" yaml:"startDate,omitempty"`
}

// PropertyDataObj was auto-generated from WSDL.
type PropertyDataObj struct {
	Key   *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// PropertySpecDataObj was auto-generated from WSDL.
type PropertySpecDataObj struct {
	Key   *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// SnapshotScopeDefectFilterSpecDataObj was auto-generated from
// WSDL.
type SnapshotScopeDefectFilterSpecDataObj struct {
	ActionNameList                    []*string                                   `xml:"actionNameList,omitempty" json:"actionNameList,omitempty" yaml:"actionNameList,omitempty"`
	AttributeDefinitionValueFilterMap []*AttributeDefinitionValueFilterMapDataObj `xml:"attributeDefinitionValueFilterMap,omitempty" json:"attributeDefinitionValueFilterMap,omitempty" yaml:"attributeDefinitionValueFilterMap,omitempty"`
	CheckerCategoryList               []*string                                   `xml:"checkerCategoryList,omitempty" json:"checkerCategoryList,omitempty" yaml:"checkerCategoryList,omitempty"`
	CheckerList                       []*string                                   `xml:"checkerList,omitempty" json:"checkerList,omitempty" yaml:"checkerList,omitempty"`
	CheckerTypeList                   []*string                                   `xml:"checkerTypeList,omitempty" json:"checkerTypeList,omitempty" yaml:"checkerTypeList,omitempty"`
	CidList                           []*int64                                    `xml:"cidList,omitempty" json:"cidList,omitempty" yaml:"cidList,omitempty"`
	ClassificationNameList            []*string                                   `xml:"classificationNameList,omitempty" json:"classificationNameList,omitempty" yaml:"classificationNameList,omitempty"`
	ComponentIdExclude                *bool                                       `xml:"componentIdExclude,omitempty" json:"componentIdExclude,omitempty" yaml:"componentIdExclude,omitempty"`
	ComponentIdList                   []*ComponentIdDataObj                       `xml:"componentIdList,omitempty" json:"componentIdList,omitempty" yaml:"componentIdList,omitempty"`
	CweList                           []*int64                                    `xml:"cweList,omitempty" json:"cweList,omitempty" yaml:"cweList,omitempty"`
	ExternalReference                 *string                                     `xml:"externalReference,omitempty" json:"externalReference,omitempty" yaml:"externalReference,omitempty"`
	FileName                          *string                                     `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
	FirstDetectedEndDate              *DateTime                                   `xml:"firstDetectedEndDate,omitempty" json:"firstDetectedEndDate,omitempty" yaml:"firstDetectedEndDate,omitempty"`
	FirstDetectedStartDate            *DateTime                                   `xml:"firstDetectedStartDate,omitempty" json:"firstDetectedStartDate,omitempty" yaml:"firstDetectedStartDate,omitempty"`
	FixTargetNameList                 []*string                                   `xml:"fixTargetNameList,omitempty" json:"fixTargetNameList,omitempty" yaml:"fixTargetNameList,omitempty"`
	FunctionMergeName                 *string                                     `xml:"functionMergeName,omitempty" json:"functionMergeName,omitempty" yaml:"functionMergeName,omitempty"`
	FunctionName                      *string                                     `xml:"functionName,omitempty" json:"functionName,omitempty" yaml:"functionName,omitempty"`
	ImpactNameList                    []*string                                   `xml:"impactNameList,omitempty" json:"impactNameList,omitempty" yaml:"impactNameList,omitempty"`
	IssueComparison                   *string                                     `xml:"issueComparison,omitempty" json:"issueComparison,omitempty" yaml:"issueComparison,omitempty"`
	IssueKindList                     []*string                                   `xml:"issueKindList,omitempty" json:"issueKindList,omitempty" yaml:"issueKindList,omitempty"`
	LastDetectedEndDate               *DateTime                                   `xml:"lastDetectedEndDate,omitempty" json:"lastDetectedEndDate,omitempty" yaml:"lastDetectedEndDate,omitempty"`
	LastDetectedStartDate             *DateTime                                   `xml:"lastDetectedStartDate,omitempty" json:"lastDetectedStartDate,omitempty" yaml:"lastDetectedStartDate,omitempty"`
	LegacyNameList                    []*string                                   `xml:"legacyNameList,omitempty" json:"legacyNameList,omitempty" yaml:"legacyNameList,omitempty"`
	MaxOccurrenceCount                *int                                        `xml:"maxOccurrenceCount,omitempty" json:"maxOccurrenceCount,omitempty" yaml:"maxOccurrenceCount,omitempty"`
	MergeExtra                        *string                                     `xml:"mergeExtra,omitempty" json:"mergeExtra,omitempty" yaml:"mergeExtra,omitempty"`
	MergeKey                          *string                                     `xml:"mergeKey,omitempty" json:"mergeKey,omitempty" yaml:"mergeKey,omitempty"`
	MinOccurrenceCount                *int                                        `xml:"minOccurrenceCount,omitempty" json:"minOccurrenceCount,omitempty" yaml:"minOccurrenceCount,omitempty"`
	OwnerNameList                     []*string                                   `xml:"ownerNameList,omitempty" json:"ownerNameList,omitempty" yaml:"ownerNameList,omitempty"`
	OwnerNamePattern                  *string                                     `xml:"ownerNamePattern,omitempty" json:"ownerNamePattern,omitempty" yaml:"ownerNamePattern,omitempty"`
	SeverityNameList                  []*string                                   `xml:"severityNameList,omitempty" json:"severityNameList,omitempty" yaml:"severityNameList,omitempty"`
	StatusNameList                    []*string                                   `xml:"statusNameList,omitempty" json:"statusNameList,omitempty" yaml:"statusNameList,omitempty"`
	StreamExcludeNameList             []*StreamIdDataObj                          `xml:"streamExcludeNameList,omitempty" json:"streamExcludeNameList,omitempty" yaml:"streamExcludeNameList,omitempty"`
	StreamExcludeQualifier            *string                                     `xml:"streamExcludeQualifier,omitempty" json:"streamExcludeQualifier,omitempty" yaml:"streamExcludeQualifier,omitempty"`
	StreamIncludeNameList             []*StreamIdDataObj                          `xml:"streamIncludeNameList,omitempty" json:"streamIncludeNameList,omitempty" yaml:"streamIncludeNameList,omitempty"`
	StreamIncludeQualifier            *string                                     `xml:"streamIncludeQualifier,omitempty" json:"streamIncludeQualifier,omitempty" yaml:"streamIncludeQualifier,omitempty"`
}

// SnapshotScopeSpecDataObj was auto-generated from WSDL.
type SnapshotScopeSpecDataObj struct {
	CompareOutdatedStreams *bool   `xml:"compareOutdatedStreams,omitempty" json:"compareOutdatedStreams,omitempty" yaml:"compareOutdatedStreams,omitempty"`
	CompareSelector        *string `xml:"compareSelector,omitempty" json:"compareSelector,omitempty" yaml:"compareSelector,omitempty"`
	ShowOutdatedStreams    *bool   `xml:"showOutdatedStreams,omitempty" json:"showOutdatedStreams,omitempty" yaml:"showOutdatedStreams,omitempty"`
	ShowSelector           *string `xml:"showSelector,omitempty" json:"showSelector,omitempty" yaml:"showSelector,omitempty"`
}

// StreamDefectDataObj was auto-generated from WSDL.
type StreamDefectDataObj struct {
	CheckerName                *string                             `xml:"checkerName,omitempty" json:"checkerName,omitempty" yaml:"checkerName,omitempty"`
	Cid                        *int64                              `xml:"cid,omitempty" json:"cid,omitempty" yaml:"cid,omitempty"`
	DefectInstances            []*DefectInstanceDataObj            `xml:"defectInstances,omitempty" json:"defectInstances,omitempty" yaml:"defectInstances,omitempty"`
	DefectStateAttributeValues []*DefectStateAttributeValueDataObj `xml:"defectStateAttributeValues,omitempty" json:"defectStateAttributeValues,omitempty" yaml:"defectStateAttributeValues,omitempty"`
	Domain                     *string                             `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	History                    []*DefectStateDataObj               `xml:"history,omitempty" json:"history,omitempty" yaml:"history,omitempty"`
	Id                         *StreamDefectIdDataObj              `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	StreamId                   *StreamIdDataObj                    `xml:"streamId,omitempty" json:"streamId,omitempty" yaml:"streamId,omitempty"`
}

// StreamDefectFilterSpecDataObj was auto-generated from WSDL.
type StreamDefectFilterSpecDataObj struct {
	DefectStateEndDate     *DateTime          `xml:"defectStateEndDate,omitempty" json:"defectStateEndDate,omitempty" yaml:"defectStateEndDate,omitempty"`
	DefectStateStartDate   *DateTime          `xml:"defectStateStartDate,omitempty" json:"defectStateStartDate,omitempty" yaml:"defectStateStartDate,omitempty"`
	IncludeDefectInstances *bool              `xml:"includeDefectInstances,omitempty" json:"includeDefectInstances,omitempty" yaml:"includeDefectInstances,omitempty"`
	IncludeHistory         *bool              `xml:"includeHistory,omitempty" json:"includeHistory,omitempty" yaml:"includeHistory,omitempty"`
	MaxDefectInstances     *int               `xml:"maxDefectInstances,omitempty" json:"maxDefectInstances,omitempty" yaml:"maxDefectInstances,omitempty"`
	StreamIdList           []*StreamIdDataObj `xml:"streamIdList,omitempty" json:"streamIdList,omitempty" yaml:"streamIdList,omitempty"`
}

// StreamDefectIdDataObj was auto-generated from WSDL.
type StreamDefectIdDataObj struct {
	DefectTriageId     *int64 `xml:"defectTriageId,omitempty" json:"defectTriageId,omitempty" yaml:"defectTriageId,omitempty"`
	DefectTriageVerNum *int   `xml:"defectTriageVerNum,omitempty" json:"defectTriageVerNum,omitempty" yaml:"defectTriageVerNum,omitempty"`
	Id                 *int64 `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	VerNum             *int   `xml:"verNum,omitempty" json:"verNum,omitempty" yaml:"verNum,omitempty"`
}

// StreamIdDataObj was auto-generated from WSDL.
type StreamIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// TriageHistoryDataObj was auto-generated from WSDL.
type TriageHistoryDataObj struct {
	Attributes []*DefectStateAttributeValueDataObj `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Id         *int64                              `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// TriageStoreIdDataObj was auto-generated from WSDL.
type TriageStoreIdDataObj struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// UpdateDefectInstanceProperties was auto-generated from WSDL.
type UpdateDefectInstanceProperties struct {
	DefectInstanceId *DefectInstanceIdDataObj `xml:"defectInstanceId,omitempty" json:"defectInstanceId,omitempty" yaml:"defectInstanceId,omitempty"`
	Properties       []*PropertySpecDataObj   `xml:"properties,omitempty" json:"properties,omitempty" yaml:"properties,omitempty"`
}

// UpdateDefectInstancePropertiesResponse was auto-generated from
// WSDL.
type UpdateDefectInstancePropertiesResponse struct {
}

// UpdateStreamDefects was auto-generated from WSDL.
type UpdateStreamDefects struct {
	StreamDefectIds []*StreamDefectIdDataObj `xml:"streamDefectIds,omitempty" json:"streamDefectIds,omitempty" yaml:"streamDefectIds,omitempty"`
	DefectStateSpec *DefectStateSpecDataObj  `xml:"defectStateSpec,omitempty" json:"defectStateSpec,omitempty" yaml:"defectStateSpec,omitempty"`
}

// UpdateStreamDefectsResponse was auto-generated from WSDL.
type UpdateStreamDefectsResponse struct {
}

// UpdateTriageForCIDsInTriageStore was auto-generated from WSDL.
type UpdateTriageForCIDsInTriageStore struct {
	TriageStore            *TriageStoreIdDataObj    `xml:"triageStore,omitempty" json:"triageStore,omitempty" yaml:"triageStore,omitempty"`
	MergedDefectIdDataObjs []*MergedDefectIdDataObj `xml:"mergedDefectIdDataObjs,omitempty" json:"mergedDefectIdDataObjs,omitempty" yaml:"mergedDefectIdDataObjs,omitempty"`
	DefectState            *DefectStateSpecDataObj  `xml:"defectState,omitempty" json:"defectState,omitempty" yaml:"defectState,omitempty"`
}

// UpdateTriageForCIDsInTriageStoreResponse was auto-generated
// from WSDL.
type UpdateTriageForCIDsInTriageStoreResponse struct {
}

// Operation wrapper for CreateMergedDefect.
// OperationCreateMergedDefect was auto-generated from WSDL.
type OperationCreateMergedDefect struct {
	CreateMergedDefect *CreateMergedDefect `xml:"tns:createMergedDefect,omitempty" json:"tns:createMergedDefect,omitempty" yaml:"tns:createMergedDefect,omitempty"`
}

// Operation wrapper for CreateMergedDefect.
// OperationCreateMergedDefectResponse was auto-generated from
// WSDL.
type OperationCreateMergedDefectResponse struct {
	CreateMergedDefectResponse *CreateMergedDefectResponse `xml:"createMergedDefectResponse,omitempty" json:"createMergedDefectResponse,omitempty" yaml:"createMergedDefectResponse,omitempty"`
}

// Operation wrapper for GetComponentMetricsForProject.
// OperationGetComponentMetricsForProject was auto-generated from
// WSDL.
type OperationGetComponentMetricsForProject struct {
	GetComponentMetricsForProject *GetComponentMetricsForProject `xml:"tns:getComponentMetricsForProject,omitempty" json:"tns:getComponentMetricsForProject,omitempty" yaml:"tns:getComponentMetricsForProject,omitempty"`
}

// Operation wrapper for GetComponentMetricsForProject.
// OperationGetComponentMetricsForProjectResponse was auto-generated
// from WSDL.
type OperationGetComponentMetricsForProjectResponse struct {
	GetComponentMetricsForProjectResponse *GetComponentMetricsForProjectResponse `xml:"getComponentMetricsForProjectResponse,omitempty" json:"getComponentMetricsForProjectResponse,omitempty" yaml:"getComponentMetricsForProjectResponse,omitempty"`
}

// Operation wrapper for GetFileContents.
// OperationGetFileContents was auto-generated from WSDL.
type OperationGetFileContents struct {
	GetFileContents *GetFileContents `xml:"tns:getFileContents,omitempty" json:"tns:getFileContents,omitempty" yaml:"tns:getFileContents,omitempty"`
}

// Operation wrapper for GetFileContents.
// OperationGetFileContentsResponse was auto-generated from WSDL.
type OperationGetFileContentsResponse struct {
	GetFileContentsResponse *GetFileContentsResponse `xml:"getFileContentsResponse,omitempty" json:"getFileContentsResponse,omitempty" yaml:"getFileContentsResponse,omitempty"`
}

// Operation wrapper for GetMergedDefectDetectionHistory.
// OperationGetMergedDefectDetectionHistory was auto-generated
// from WSDL.
type OperationGetMergedDefectDetectionHistory struct {
	GetMergedDefectDetectionHistory *GetMergedDefectDetectionHistory `xml:"tns:getMergedDefectDetectionHistory,omitempty" json:"tns:getMergedDefectDetectionHistory,omitempty" yaml:"tns:getMergedDefectDetectionHistory,omitempty"`
}

// Operation wrapper for GetMergedDefectDetectionHistory.
// OperationGetMergedDefectDetectionHistoryResponse was auto-generated
// from WSDL.
type OperationGetMergedDefectDetectionHistoryResponse struct {
	GetMergedDefectDetectionHistoryResponse *GetMergedDefectDetectionHistoryResponse `xml:"getMergedDefectDetectionHistoryResponse,omitempty" json:"getMergedDefectDetectionHistoryResponse,omitempty" yaml:"getMergedDefectDetectionHistoryResponse,omitempty"`
}

// Operation wrapper for GetMergedDefectHistory.
// OperationGetMergedDefectHistory was auto-generated from WSDL.
type OperationGetMergedDefectHistory struct {
	GetMergedDefectHistory *GetMergedDefectHistory `xml:"tns:getMergedDefectHistory,omitempty" json:"tns:getMergedDefectHistory,omitempty" yaml:"tns:getMergedDefectHistory,omitempty"`
}

// Operation wrapper for GetMergedDefectHistory.
// OperationGetMergedDefectHistoryResponse was auto-generated from
// WSDL.
type OperationGetMergedDefectHistoryResponse struct {
	GetMergedDefectHistoryResponse *GetMergedDefectHistoryResponse `xml:"getMergedDefectHistoryResponse,omitempty" json:"getMergedDefectHistoryResponse,omitempty" yaml:"getMergedDefectHistoryResponse,omitempty"`
}

// Operation wrapper for GetMergedDefectsForProjectScope.
// OperationGetMergedDefectsForProjectScope was auto-generated
// from WSDL.
type OperationGetMergedDefectsForProjectScope struct {
	GetMergedDefectsForProjectScope *GetMergedDefectsForProjectScope `xml:"tns:getMergedDefectsForProjectScope,omitempty" json:"tns:getMergedDefectsForProjectScope,omitempty" yaml:"tns:getMergedDefectsForProjectScope,omitempty"`
}

// Operation wrapper for GetMergedDefectsForProjectScope.
// OperationGetMergedDefectsForProjectScopeResponse was auto-generated
// from WSDL.
type OperationGetMergedDefectsForProjectScopeResponse struct {
	GetMergedDefectsForProjectScopeResponse *GetMergedDefectsForProjectScopeResponse `xml:"getMergedDefectsForProjectScopeResponse,omitempty" json:"getMergedDefectsForProjectScopeResponse,omitempty" yaml:"getMergedDefectsForProjectScopeResponse,omitempty"`
}

// Operation wrapper for GetMergedDefectsForSnapshotScope.
// OperationGetMergedDefectsForSnapshotScope was auto-generated
// from WSDL.
type OperationGetMergedDefectsForSnapshotScope struct {
	GetMergedDefectsForSnapshotScope *GetMergedDefectsForSnapshotScope `xml:"tns:getMergedDefectsForSnapshotScope,omitempty" json:"tns:getMergedDefectsForSnapshotScope,omitempty" yaml:"tns:getMergedDefectsForSnapshotScope,omitempty"`
}

// Operation wrapper for GetMergedDefectsForSnapshotScope.
// OperationGetMergedDefectsForSnapshotScopeResponse was auto-generated
// from WSDL.
type OperationGetMergedDefectsForSnapshotScopeResponse struct {
	GetMergedDefectsForSnapshotScopeResponse *GetMergedDefectsForSnapshotScopeResponse `xml:"getMergedDefectsForSnapshotScopeResponse,omitempty" json:"getMergedDefectsForSnapshotScopeResponse,omitempty" yaml:"getMergedDefectsForSnapshotScopeResponse,omitempty"`
}

// Operation wrapper for GetMergedDefectsForStreams.
// OperationGetMergedDefectsForStreams was auto-generated from
// WSDL.
type OperationGetMergedDefectsForStreams struct {
	GetMergedDefectsForStreams *GetMergedDefectsForStreams `xml:"tns:getMergedDefectsForStreams,omitempty" json:"tns:getMergedDefectsForStreams,omitempty" yaml:"tns:getMergedDefectsForStreams,omitempty"`
}

// Operation wrapper for GetMergedDefectsForStreams.
// OperationGetMergedDefectsForStreamsResponse was auto-generated
// from WSDL.
type OperationGetMergedDefectsForStreamsResponse struct {
	GetMergedDefectsForStreamsResponse *GetMergedDefectsForStreamsResponse `xml:"getMergedDefectsForStreamsResponse,omitempty" json:"getMergedDefectsForStreamsResponse,omitempty" yaml:"getMergedDefectsForStreamsResponse,omitempty"`
}

// Operation wrapper for GetStreamDefects.
// OperationGetStreamDefects was auto-generated from WSDL.
type OperationGetStreamDefects struct {
	GetStreamDefects *GetStreamDefects `xml:"tns:getStreamDefects,omitempty" json:"tns:getStreamDefects,omitempty" yaml:"tns:getStreamDefects,omitempty"`
}

// Operation wrapper for GetStreamDefects.
// OperationGetStreamDefectsResponse was auto-generated from WSDL.
type OperationGetStreamDefectsResponse struct {
	GetStreamDefectsResponse *GetStreamDefectsResponse `xml:"getStreamDefectsResponse,omitempty" json:"getStreamDefectsResponse,omitempty" yaml:"getStreamDefectsResponse,omitempty"`
}

// Operation wrapper for GetTrendRecordsForProject.
// OperationGetTrendRecordsForProject was auto-generated from WSDL.
type OperationGetTrendRecordsForProject struct {
	GetTrendRecordsForProject *GetTrendRecordsForProject `xml:"tns:getTrendRecordsForProject,omitempty" json:"tns:getTrendRecordsForProject,omitempty" yaml:"tns:getTrendRecordsForProject,omitempty"`
}

// Operation wrapper for GetTrendRecordsForProject.
// OperationGetTrendRecordsForProjectResponse was auto-generated
// from WSDL.
type OperationGetTrendRecordsForProjectResponse struct {
	GetTrendRecordsForProjectResponse *GetTrendRecordsForProjectResponse `xml:"getTrendRecordsForProjectResponse,omitempty" json:"getTrendRecordsForProjectResponse,omitempty" yaml:"getTrendRecordsForProjectResponse,omitempty"`
}

// Operation wrapper for GetTriageHistory.
// OperationGetTriageHistory was auto-generated from WSDL.
type OperationGetTriageHistory struct {
	GetTriageHistory *GetTriageHistory `xml:"tns:getTriageHistory,omitempty" json:"tns:getTriageHistory,omitempty" yaml:"tns:getTriageHistory,omitempty"`
}

// Operation wrapper for GetTriageHistory.
// OperationGetTriageHistoryResponse was auto-generated from WSDL.
type OperationGetTriageHistoryResponse struct {
	GetTriageHistoryResponse *GetTriageHistoryResponse `xml:"getTriageHistoryResponse,omitempty" json:"getTriageHistoryResponse,omitempty" yaml:"getTriageHistoryResponse,omitempty"`
}

// Operation wrapper for UpdateDefectInstanceProperties.
// OperationUpdateDefectInstanceProperties was auto-generated from
// WSDL.
type OperationUpdateDefectInstanceProperties struct {
	UpdateDefectInstanceProperties *UpdateDefectInstanceProperties `xml:"tns:updateDefectInstanceProperties,omitempty" json:"tns:updateDefectInstanceProperties,omitempty" yaml:"tns:updateDefectInstanceProperties,omitempty"`
}

// Operation wrapper for UpdateDefectInstanceProperties.
// OperationUpdateDefectInstancePropertiesResponse was auto-generated
// from WSDL.
type OperationUpdateDefectInstancePropertiesResponse struct {
	UpdateDefectInstancePropertiesResponse *UpdateDefectInstancePropertiesResponse `xml:"updateDefectInstancePropertiesResponse,omitempty" json:"updateDefectInstancePropertiesResponse,omitempty" yaml:"updateDefectInstancePropertiesResponse,omitempty"`
}

// Operation wrapper for UpdateStreamDefects.
// OperationUpdateStreamDefects was auto-generated from WSDL.
type OperationUpdateStreamDefects struct {
	UpdateStreamDefects *UpdateStreamDefects `xml:"tns:updateStreamDefects,omitempty" json:"tns:updateStreamDefects,omitempty" yaml:"tns:updateStreamDefects,omitempty"`
}

// Operation wrapper for UpdateStreamDefects.
// OperationUpdateStreamDefectsResponse was auto-generated from
// WSDL.
type OperationUpdateStreamDefectsResponse struct {
	UpdateStreamDefectsResponse *UpdateStreamDefectsResponse `xml:"updateStreamDefectsResponse,omitempty" json:"updateStreamDefectsResponse,omitempty" yaml:"updateStreamDefectsResponse,omitempty"`
}

// Operation wrapper for UpdateTriageForCIDsInTriageStore.
// OperationUpdateTriageForCIDsInTriageStore was auto-generated
// from WSDL.
type OperationUpdateTriageForCIDsInTriageStore struct {
	UpdateTriageForCIDsInTriageStore *UpdateTriageForCIDsInTriageStore `xml:"tns:updateTriageForCIDsInTriageStore,omitempty" json:"tns:updateTriageForCIDsInTriageStore,omitempty" yaml:"tns:updateTriageForCIDsInTriageStore,omitempty"`
}

// Operation wrapper for UpdateTriageForCIDsInTriageStore.
// OperationUpdateTriageForCIDsInTriageStoreResponse was auto-generated
// from WSDL.
type OperationUpdateTriageForCIDsInTriageStoreResponse struct {
	UpdateTriageForCIDsInTriageStoreResponse *UpdateTriageForCIDsInTriageStoreResponse `xml:"updateTriageForCIDsInTriageStoreResponse,omitempty" json:"updateTriageForCIDsInTriageStoreResponse,omitempty" yaml:"updateTriageForCIDsInTriageStoreResponse,omitempty"`
}

// defectService implements the DefectService interface.
type defectService struct {
	cli *soap.Client
}

// CreateMergedDefect was auto-generated from WSDL.
func (p *defectService) CreateMergedDefect(CreateMergedDefect *CreateMergedDefect) (*CreateMergedDefectResponse, error) {
	α := struct {
		OperationCreateMergedDefect `xml:"tns:createMergedDefect"`
	}{
		OperationCreateMergedDefect{
			CreateMergedDefect,
		},
	}

	γ := struct {
		OperationCreateMergedDefectResponse `xml:"createMergedDefectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateMergedDefect", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateMergedDefectResponse, nil
}

// GetComponentMetricsForProject was auto-generated from WSDL.
func (p *defectService) GetComponentMetricsForProject(GetComponentMetricsForProject *GetComponentMetricsForProject) (*GetComponentMetricsForProjectResponse, error) {
	α := struct {
		OperationGetComponentMetricsForProject `xml:"tns:getComponentMetricsForProject"`
	}{
		OperationGetComponentMetricsForProject{
			GetComponentMetricsForProject,
		},
	}

	γ := struct {
		OperationGetComponentMetricsForProjectResponse `xml:"getComponentMetricsForProjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetComponentMetricsForProject", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetComponentMetricsForProjectResponse, nil
}

// GetFileContents was auto-generated from WSDL.
func (p *defectService) GetFileContents(GetFileContents *GetFileContents) (*GetFileContentsResponse, error) {
	α := struct {
		OperationGetFileContents `xml:"tns:getFileContents"`
	}{
		OperationGetFileContents{
			GetFileContents,
		},
	}

	γ := struct {
		OperationGetFileContentsResponse `xml:"getFileContentsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetFileContents", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetFileContentsResponse, nil
}

// GetMergedDefectDetectionHistory was auto-generated from WSDL.
func (p *defectService) GetMergedDefectDetectionHistory(GetMergedDefectDetectionHistory *GetMergedDefectDetectionHistory) (*GetMergedDefectDetectionHistoryResponse, error) {
	α := struct {
		OperationGetMergedDefectDetectionHistory `xml:"tns:getMergedDefectDetectionHistory"`
	}{
		OperationGetMergedDefectDetectionHistory{
			GetMergedDefectDetectionHistory,
		},
	}

	γ := struct {
		OperationGetMergedDefectDetectionHistoryResponse `xml:"getMergedDefectDetectionHistoryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMergedDefectDetectionHistory", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMergedDefectDetectionHistoryResponse, nil
}

// GetMergedDefectHistory was auto-generated from WSDL.
func (p *defectService) GetMergedDefectHistory(GetMergedDefectHistory *GetMergedDefectHistory) (*GetMergedDefectHistoryResponse, error) {
	α := struct {
		OperationGetMergedDefectHistory `xml:"tns:getMergedDefectHistory"`
	}{
		OperationGetMergedDefectHistory{
			GetMergedDefectHistory,
		},
	}

	γ := struct {
		OperationGetMergedDefectHistoryResponse `xml:"getMergedDefectHistoryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMergedDefectHistory", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMergedDefectHistoryResponse, nil
}

// GetMergedDefectsForProjectScope was auto-generated from WSDL.
func (p *defectService) GetMergedDefectsForProjectScope(GetMergedDefectsForProjectScope *GetMergedDefectsForProjectScope) (*GetMergedDefectsForProjectScopeResponse, error) {
	α := struct {
		OperationGetMergedDefectsForProjectScope `xml:"tns:getMergedDefectsForProjectScope"`
	}{
		OperationGetMergedDefectsForProjectScope{
			GetMergedDefectsForProjectScope,
		},
	}

	γ := struct {
		OperationGetMergedDefectsForProjectScopeResponse `xml:"getMergedDefectsForProjectScopeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMergedDefectsForProjectScope", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMergedDefectsForProjectScopeResponse, nil
}

// GetMergedDefectsForSnapshotScope was auto-generated from WSDL.
func (p *defectService) GetMergedDefectsForSnapshotScope(GetMergedDefectsForSnapshotScope *GetMergedDefectsForSnapshotScope) (*GetMergedDefectsForSnapshotScopeResponse, error) {
	α := struct {
		OperationGetMergedDefectsForSnapshotScope `xml:"tns:getMergedDefectsForSnapshotScope"`
	}{
		OperationGetMergedDefectsForSnapshotScope{
			GetMergedDefectsForSnapshotScope,
		},
	}

	γ := struct {
		OperationGetMergedDefectsForSnapshotScopeResponse `xml:"getMergedDefectsForSnapshotScopeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMergedDefectsForSnapshotScope", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMergedDefectsForSnapshotScopeResponse, nil
}

// GetMergedDefectsForStreams was auto-generated from WSDL.
func (p *defectService) GetMergedDefectsForStreams(GetMergedDefectsForStreams *GetMergedDefectsForStreams) (*GetMergedDefectsForStreamsResponse, error) {
	α := struct {
		OperationGetMergedDefectsForStreams `xml:"tns:getMergedDefectsForStreams"`
	}{
		OperationGetMergedDefectsForStreams{
			GetMergedDefectsForStreams,
		},
	}

	γ := struct {
		OperationGetMergedDefectsForStreamsResponse `xml:"getMergedDefectsForStreamsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMergedDefectsForStreams", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMergedDefectsForStreamsResponse, nil
}

// GetStreamDefects was auto-generated from WSDL.
func (p *defectService) GetStreamDefects(GetStreamDefects *GetStreamDefects) (*GetStreamDefectsResponse, error) {
	α := struct {
		OperationGetStreamDefects `xml:"tns:getStreamDefects"`
	}{
		OperationGetStreamDefects{
			GetStreamDefects,
		},
	}

	γ := struct {
		OperationGetStreamDefectsResponse `xml:"getStreamDefectsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStreamDefects", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStreamDefectsResponse, nil
}

// GetTrendRecordsForProject was auto-generated from WSDL.
func (p *defectService) GetTrendRecordsForProject(GetTrendRecordsForProject *GetTrendRecordsForProject) (*GetTrendRecordsForProjectResponse, error) {
	α := struct {
		OperationGetTrendRecordsForProject `xml:"tns:getTrendRecordsForProject"`
	}{
		OperationGetTrendRecordsForProject{
			GetTrendRecordsForProject,
		},
	}

	γ := struct {
		OperationGetTrendRecordsForProjectResponse `xml:"getTrendRecordsForProjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetTrendRecordsForProject", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTrendRecordsForProjectResponse, nil
}

// GetTriageHistory was auto-generated from WSDL.
func (p *defectService) GetTriageHistory(GetTriageHistory *GetTriageHistory) (*GetTriageHistoryResponse, error) {
	α := struct {
		OperationGetTriageHistory `xml:"tns:getTriageHistory"`
	}{
		OperationGetTriageHistory{
			GetTriageHistory,
		},
	}

	γ := struct {
		OperationGetTriageHistoryResponse `xml:"getTriageHistoryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetTriageHistory", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTriageHistoryResponse, nil
}

// UpdateDefectInstanceProperties was auto-generated from WSDL.
func (p *defectService) UpdateDefectInstanceProperties(UpdateDefectInstanceProperties *UpdateDefectInstanceProperties) (*UpdateDefectInstancePropertiesResponse, error) {
	α := struct {
		OperationUpdateDefectInstanceProperties `xml:"tns:updateDefectInstanceProperties"`
	}{
		OperationUpdateDefectInstanceProperties{
			UpdateDefectInstanceProperties,
		},
	}

	γ := struct {
		OperationUpdateDefectInstancePropertiesResponse `xml:"updateDefectInstancePropertiesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateDefectInstanceProperties", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateDefectInstancePropertiesResponse, nil
}

// UpdateStreamDefects was auto-generated from WSDL.
func (p *defectService) UpdateStreamDefects(UpdateStreamDefects *UpdateStreamDefects) (*UpdateStreamDefectsResponse, error) {
	α := struct {
		OperationUpdateStreamDefects `xml:"tns:updateStreamDefects"`
	}{
		OperationUpdateStreamDefects{
			UpdateStreamDefects,
		},
	}

	γ := struct {
		OperationUpdateStreamDefectsResponse `xml:"updateStreamDefectsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateStreamDefects", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateStreamDefectsResponse, nil
}

// UpdateTriageForCIDsInTriageStore was auto-generated from WSDL.
func (p *defectService) UpdateTriageForCIDsInTriageStore(UpdateTriageForCIDsInTriageStore *UpdateTriageForCIDsInTriageStore) (*UpdateTriageForCIDsInTriageStoreResponse, error) {
	α := struct {
		OperationUpdateTriageForCIDsInTriageStore `xml:"tns:updateTriageForCIDsInTriageStore"`
	}{
		OperationUpdateTriageForCIDsInTriageStore{
			UpdateTriageForCIDsInTriageStore,
		},
	}

	γ := struct {
		OperationUpdateTriageForCIDsInTriageStoreResponse `xml:"updateTriageForCIDsInTriageStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateTriageForCIDsInTriageStore", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateTriageForCIDsInTriageStoreResponse, nil
}
