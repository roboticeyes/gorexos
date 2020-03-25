package rexos

import (
	"fmt"
	"testing"
)

func TestGenerateRefTree(t *testing.T) {

	// tree, err := reconstructReferenceTreefromJSON(jsonInput)
	// if err != nil {
	// 	t.Errorf("Cannot reconstruct reference tree: %v", err)
	// }
	// fmt.Println(tree.Head)
}

func TestSimpleRefTree(t *testing.T) {

	var refs []Reference
	refs = append(refs, Reference{Urn: "6", ParentReferenceUrn: "2"})
	refs = append(refs, Reference{Urn: "5", ParentReferenceUrn: "2"})
	refs = append(refs, Reference{Urn: "4", ParentReferenceUrn: "3"})
	refs = append(refs, Reference{Urn: "3", ParentReferenceUrn: "1"})
	refs = append(refs, Reference{Urn: "2", ParentReferenceUrn: "1"})
	refs = append(refs, Reference{Urn: "1", RootReference: true})

	tree, err := reconstructReferenceTreefromJSON(refs)
	if err != nil {
		t.Errorf("Cannot reconstruct reference tree: %v", err)
	}
	fmt.Println(tree.Head)
}

const jsonInput = `
{
  "_embedded" : {
    "rexReferences" : [ {
      "urn" : "robotic-eyes:rex-reference:10050",
      "category" : "file",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10050/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:10047",
      "category" : "track",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10047/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:10045",
      "category" : "file",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9930",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10045/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:10009",
      "category" : "file",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9930",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10009/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9948",
      "category" : "file",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9948/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9946",
      "category" : "file",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9946/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9944",
      "category" : "rex",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9944/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9942",
      "category" : "data",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9927",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9942/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9938",
      "category" : "route",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9938/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9935",
      "category" : "file",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9935/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9930",
      "category" : "activity",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9929",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9930/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9929",
      "category" : "inspection",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9927",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9929/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9928",
      "category" : null,
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9927",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9928/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:9927",
      "category" : null,
      "parentReferenceUrn" : null,
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/9927/rexReferenceConfigs{?projection}",
          "templated" : true
        }
      }
    }, {
      "urn" : "robotic-eyes:rex-reference:10105",
      "category" : "file",
      "parentReferenceUrn" : "robotic-eyes:rex-reference:9942",
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
        "projectFile" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/projectFile{?projection}",
          "templated" : true
        },
        "project" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/project{?projection}",
          "templated" : true
        },
        "projectFiles" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/projectFiles{?projection}",
          "templated" : true
        },
        "parentReference" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/parentReference{?projection}",
          "templated" : true
        },
        "childReferences" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/childReferences{?projection}",
          "templated" : true
        },
        "rexReferenceConfigs" : {
          "href" : "https://api-dev-01.rexos.cloud/rex-gateway/api/v2/rexReferences/10105/rexReferenceConfigs{?projection}",
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
