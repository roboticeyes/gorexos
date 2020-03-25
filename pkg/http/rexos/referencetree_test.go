package rexos

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGenerateRefTree(t *testing.T) {

	var halTree referenceTreeHal
	err := json.Unmarshal([]byte(references), &halTree)
	if err != nil {
		t.Fatal(err)
	}

	var files projectFilesHal
	err = json.Unmarshal([]byte(projectFiles), &files)
	if err != nil {
		t.Fatal(err)
	}

	refTree := ReferenceTree{
		ProjectUrn:   "roboticeyes:project:9926",
		ProjectType:  "inspection.target",
		ProjectFiles: files.Embedded.ProjectFiles,
	}
	refTree.References = halTree.Embedded.RexReferences
	refTree.Tree, err = reconstructReferenceTreefromJSON(halTree.Embedded.RexReferences)
	if err != nil {
		t.Errorf("Cannot reconstruct reference tree: %v", err)
	}
	if refTree.Tree == nil {
		t.Errorf("Tree is nil")
	}

	refTree.Beautify()
	err = refTree.WriteToDot(os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimpleRefTree(t *testing.T) {

	var refs []Reference
	refs = append(refs, Reference{Urn: "6", ParentReferenceUrn: "2"})
	refs = append(refs, Reference{Urn: "5", ParentReferenceUrn: "2"})
	refs = append(refs, Reference{Urn: "4", ParentReferenceUrn: "3"})
	refs = append(refs, Reference{Urn: "3", ParentReferenceUrn: "1"})
	refs = append(refs, Reference{Urn: "2", ParentReferenceUrn: "1"})
	refs = append(refs, Reference{Urn: "1", RootReference: true})

	refTree, err := reconstructReferenceTreefromJSON(refs)
	if err != nil {
		t.Errorf("Cannot reconstruct reference tree: %v", err)
	}
	if refTree == nil {
		t.Errorf("Referencetree is nil")
	}
}

const projectFiles = `
{
  "_embedded" : {
    "projectFiles" : [ {
      "lastModified" : "2020-03-23T14:56:19.716+0000",
      "contentType" : null,
      "urn" : "robotic-eyes:project-file:10104",
      "fileSize" : null,
      "contentHash" : null,
      "name" : "rex_e5afa329-e65b-46ba-8795-ef3067ad2ac5",
      "type" : "rex",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10104/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T18:26:31.814+0000",
      "contentType" : "application/octet-stream",
      "urn" : "robotic-eyes:project-file:10049",
      "fileSize" : 282,
      "contentHash" : "72e240260a80163962ff6df4ef453c43",
      "name" : "rex_d84e58ab-8b12-485b-ad78-fd8df43404de",
      "type" : "rex",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049"
        },
        "file.download" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049/file?contentHash=72e240260a80163962ff6df4ef453c43"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10049/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T18:25:47.630+0000",
      "contentType" : "application/octet-stream",
      "urn" : "robotic-eyes:project-file:10046",
      "fileSize" : 282,
      "contentHash" : "72e240260a80163962ff6df4ef453c43",
      "name" : "track_f606bf34-46c5-467e-b92a-2870437a8191",
      "type" : "track",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046"
        },
        "file.download" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046/file?contentHash=72e240260a80163962ff6df4ef453c43"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10046/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T18:21:59.385+0000",
      "contentType" : null,
      "urn" : "robotic-eyes:project-file:10044",
      "fileSize" : null,
      "contentHash" : null,
      "name" : "b0b1b7fb-57bb-41ec-a9f3-e2ccd362883a",
      "type" : "image",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10044/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T18:21:12.667+0000",
      "contentType" : null,
      "urn" : "robotic-eyes:project-file:10043",
      "fileSize" : null,
      "contentHash" : null,
      "name" : "b039a242-f6ef-46e2-a86b-888b61fbddca",
      "type" : "image",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10043/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T18:16:31.781+0000",
      "contentType" : null,
      "urn" : "robotic-eyes:project-file:10042",
      "fileSize" : null,
      "contentHash" : null,
      "name" : "79e18c4e-f821-4f5a-a3ff-5a424cf862bd",
      "type" : "image",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10042/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T16:49:39.654+0000",
      "contentType" : "image/jpeg",
      "urn" : "robotic-eyes:project-file:10008",
      "fileSize" : 22964,
      "contentHash" : "c2f9411a387a3041475ba49cea9520dd",
      "name" : "6ea13b27-3ffa-4613-99d3-8f2aa06f14a5",
      "type" : "image",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008"
        },
        "file.download" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008/file?contentHash=c2f9411a387a3041475ba49cea9520dd"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/10008/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T14:22:41.665+0000",
      "contentType" : "application/octet-stream",
      "urn" : "robotic-eyes:project-file:9947",
      "fileSize" : 282,
      "contentHash" : "72e240260a80163962ff6df4ef453c43",
      "name" : "rex_3050cf5d-2e71-431a-8aff-64c903190f61",
      "type" : "rex",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947"
        },
        "file.download" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947/file?contentHash=72e240260a80163962ff6df4ef453c43"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9947/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T14:11:47.842+0000",
      "contentType" : null,
      "urn" : "robotic-eyes:project-file:9945",
      "fileSize" : null,
      "contentHash" : null,
      "name" : "rex_4171c96e-d48a-435e-827c-009329c15939",
      "type" : "rex",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9945/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T12:50:05.801+0000",
      "contentType" : null,
      "urn" : "robotic-eyes:project-file:9943",
      "fileSize" : null,
      "contentHash" : null,
      "name" : "rex_254f339c-c801-48ec-83ae-7b93a3bdc24b",
      "type" : "rex",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9943/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T10:40:24.803+0000",
      "contentType" : "application/octet-stream",
      "urn" : "robotic-eyes:project-file:9937",
      "fileSize" : 282,
      "contentHash" : "72e240260a80163962ff6df4ef453c43",
      "name" : "route_65b540da-3f81-4080-9efe-fa187c411587",
      "type" : "route",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937"
        },
        "file.download" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937/file?contentHash=72e240260a80163962ff6df4ef453c43"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9937/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "lastModified" : "2020-03-20T10:18:52.133+0000",
      "contentType" : "application/octet-stream",
      "urn" : "robotic-eyes:project-file:9934",
      "fileSize" : 282,
      "contentHash" : "72e240260a80163962ff6df4ef453c43",
      "name" : "route_115c7bd4-d80f-4742-a397-541b87e608fc",
      "type" : "route",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934"
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934{?projection}",
          "templated" : true
        },
        "constructionSiteDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionSiteDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934"
        },
        "constructionLayoutVersionDocument" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersionDocuments/search/findByProjectFile?projectFile=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934"
        },
        "file.download" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934/file?contentHash=72e240260a80163962ff6df4ef453c43"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934/project{?projection}",
          "templated" : true
        },
        "rexReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934/rexReferences{?projection}",
          "templated" : true
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934/rexReference{?projection}",
          "templated" : true
        },
        "projectFileConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projectFiles/9934/projectFileConfigs{?projection}",
          "templated" : true
        }
      }
    } ]
  },
  "_links" : {
    "self" : {
      "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projects/9926/projectFiles"
    }
  }
}
`

const references = `
{
  "_embedded" : {
    "rexReferences" : [ {
      "category" : "file",
      "urn" : "robotic-eyes:rex-reference:10050",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
      "projectFileUrn" : "robotic-eyes:project-file:10049",
      "rootReference" : false,
      "name" : "",
      "key" : "827b0da2-75e3-4cce-9419-68b19439dd6b",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "track",
      "urn" : "robotic-eyes:rex-reference:10047",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
      "projectFileUrn" : "robotic-eyes:project-file:10046",
      "rootReference" : false,
      "name" : "",
      "key" : "80f02e74-5e09-4ef6-908c-b86d72a39dcc",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "file",
      "urn" : "robotic-eyes:rex-reference:10045",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9930",
      "projectFileUrn" : "robotic-eyes:project-file:10044",
      "rootReference" : false,
      "name" : "",
      "key" : "c7b2db1e-db73-4da0-a529-e436c53a675b",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "file",
      "urn" : "robotic-eyes:rex-reference:10009",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9930",
      "projectFileUrn" : "robotic-eyes:project-file:10008",
      "rootReference" : false,
      "name" : "",
      "key" : "43c92efe-5416-4bf0-a140-8aec346f8d0c",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "file",
      "urn" : "robotic-eyes:rex-reference:9948",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
      "projectFileUrn" : "robotic-eyes:project-file:9947",
      "rootReference" : false,
      "name" : "",
      "key" : "409fa9e1-5139-4452-b516-2bf0187efc69",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "file",
      "urn" : "robotic-eyes:rex-reference:9946",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
      "projectFileUrn" : "robotic-eyes:project-file:9945",
      "rootReference" : false,
      "name" : "",
      "key" : "3e989ba5-818d-468a-9b90-1a87a7c1c451",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "rex",
      "urn" : "robotic-eyes:rex-reference:9944",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
      "projectFileUrn" : "robotic-eyes:project-file:9943",
      "rootReference" : false,
      "name" : "",
      "key" : "b70658b4-18bd-4f38-9360-b759baee0331",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "data",
      "urn" : "robotic-eyes:rex-reference:9942",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9927",
      "projectFileUrn" : null,
      "rootReference" : false,
      "name" : "dataGroup",
      "key" : "0ff6316d-950d-403c-b832-5971aa30f3e6",
      "type" : "group",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "route",
      "urn" : "robotic-eyes:rex-reference:9938",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
      "projectFileUrn" : "robotic-eyes:project-file:9937",
      "rootReference" : false,
      "name" : "",
      "key" : "b4f8b481-79ee-470b-8030-e69948a0b030",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "file",
      "urn" : "robotic-eyes:rex-reference:9935",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
      "projectFileUrn" : "robotic-eyes:project-file:9934",
      "rootReference" : false,
      "name" : "",
      "key" : "4429287f-c841-43b9-a078-c037723920cd",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "activity",
      "urn" : "robotic-eyes:rex-reference:9930",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
      "projectFileUrn" : null,
      "rootReference" : false,
      "name" : "Check 1",
      "key" : "6e5888ee-438f-46c0-b04a-c430076aef77",
      "type" : "group",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930{?projection}",
          "templated" : true
        },
        "dataResource" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/activities/search/findByUrn?urn=rexos:activity:2294"
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "inspection",
      "urn" : "robotic-eyes:rex-reference:9929",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9927",
      "projectFileUrn" : null,
      "rootReference" : false,
      "name" : "Inspection 1",
      "key" : "26442e1d-b18d-4980-ac7a-a6e3950edb14",
      "type" : "group",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929{?projection}",
          "templated" : true
        },
        "dataResource" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/inspections/search/findByUrn?urn=rexos:inspection:2333"
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : null,
      "urn" : "robotic-eyes:rex-reference:9928",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9927",
      "projectFileUrn" : null,
      "rootReference" : false,
      "name" : "",
      "key" : "fe98f4bf-57d7-47e7-9891-59a2e6374c19",
      "type" : "portal",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : null,
      "urn" : "robotic-eyes:rex-reference:9927",
      "parentReferenceUrn" : null,
      "projectFileUrn" : null,
      "rootReference" : true,
      "name" : "",
      "key" : "30e338f9-974a-442a-981b-5dc98243a2b7",
      "type" : "root",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/projectFiles{?projection}",
          "templated" : true
        }
      }
    }, {
      "category" : "file",
      "urn" : "robotic-eyes:rex-reference:10105",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
      "projectFileUrn" : "robotic-eyes:project-file:10104",
      "rootReference" : false,
      "name" : "",
      "key" : "af0013d0-53f0-48dc-a915-a4a57cef5c18",
      "type" : "file",
      "_links" : {
        "self" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105"
        },
        "rexReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105{?projection}",
          "templated" : true
        },
        "constructionLayout" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayouts/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105"
        },
        "constructionLayoutVersion" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutVersions/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105"
        },
        "constructionLayoutItem" : {
          "href" : "https://api-dev-01.rexos.cloud/api/v2/constructionLayoutItems/search/findByRexReference?rexReference=https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105"
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/project{?projection}",
          "templated" : true
        },
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/projectFile{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/rexReferenceConfigs{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/childReferences{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/parentReference{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/projectFiles{?projection}",
          "templated" : true
        }
      }
    } ]
  },
  "_links" : {
    "self" : {
      "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/projects/9926/rexReferences?projection=linkedList"
    }
  }
}
`
