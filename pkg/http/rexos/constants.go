package rexos

var (
	// REXos core services
	apiProjectByNameAndOwner           = "/api/v2/projects/search/findByNameAndOwner?"
	apiProjectByOwner                  = "/api/v2/projects/search/findAllByOwner?"
	apiStatisticsByUserID              = "/api/v2/projects/statisticsByUser?userId="
	apiProjectFiles                    = "/api/v2/projectFiles/"
	apiProjectFileFindByProjectAndName = "/api/v2/projectFiles/search/findByProjectAndName?"
	apiProjectsByUrn                   = "/api/v2/projects/search/findByUrn?"
	apiProjects                        = "/api/v2/projects"
	apiReferences                      = "/api/v2/rexReferences"
	apiReferenceByKey                  = "/api/v2/rexReferences/search/findByKey?"

	// Construction service
	apiConstructionLayoutItems            = "/construction/v1/constructionLayoutItems"
	apiConstructionLayoutVersions         = "/construction/v1/constructionLayoutVersions"
	apiConstructionLayoutVersionDocuments = "/construction/v1/constructionLayoutVersionDocuments"
	apiConstructionLayouts                = "/construction/v1/constructionLayouts"
	apiConstructionLayoutsByUrn           = "/construction/v1/constructionLayouts/search/findByUrn?urn="
	apiConstructionLayoutVersionsByUrn    = "/construction/v1/constructionLayoutVersions/search/findByUrn?urn="
	apiConstructionProducts               = "/construction/v1/constructionProducts"
	apiConstructionProductsByKey          = "/construction/v1/constructionProducts/search/findByKey?key="
	apiConstructionProductsByUrn          = "/construction/v1/constructionProducts/search/findByUrn?urn="
	apiConstructionSites                  = "/construction/v1/constructionSites/"
	apiConstructionSitesByUrn             = "/construction/v1/constructionSites/search/findByUrn?urn="
)

const (
	rexSchemeV1                = "rexos.scheme.v1"
	rexConstructionProductType = "construction.product"
	rexConstructionSiteType    = "construction.site"
	rexFileType                = "rex"

	// This is the name of the ProjectFile in the Rex Project which is used to identify the proper
	// REX file for the Construction Product
	rexConstructionProductGeometryName = "geometry"
)
