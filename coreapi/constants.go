package coreapi

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
)

const (
	rexSchemeV1 = "rexos.scheme.v1"
	rexFileType = "rex"
)
