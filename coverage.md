# Test Coverage Report

Generated on: Thu Jun 26 21:13:21 UTC 2025

## Coverage Summary

```
github.com/georgetaylor/spotctl/cmd/cloudspaces/cloudspaces.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:16:		getCloudSpacesTableConfig	100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:35:		outputCloudSpaces		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:81:		outputCreatedCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:108:		outputCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:135:		getNamespace			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:15:		NewCreateCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:65:		runCreate			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:154:		loadCloudSpaceSpecFromFile	0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:13:		NewDeleteCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:44:		runDelete			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:13:		NewEditCommand			100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:71:		runEdit				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:50:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:47:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/common.go:13:			AddOutputFlag			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:18:			GetOutputFormat			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:24:			ConfirmAction			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:32:			CheckError			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:164:			init				0.0%
github.com/georgetaylor/spotctl/cmd/config.go:175:			contains			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:185:			min				0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:11:		getOrganizationsTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:25:		outputOrganizations		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:30:		runOrganizationsList		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/organizations.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:11:		getRegionsTableConfig		0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:26:		outputRegions			0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:57:		outputRegion			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:13:			NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:35:			runGet				0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:13:			NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:31:			runList				0.0%
github.com/georgetaylor/spotctl/cmd/regions/regions.go:6:		NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/root.go:32:				Execute				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:36:				init				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:63:				initConfig			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11:		getServerClassesTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33:		outputServerClasses		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64:		outputServerClass		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:35:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:30:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/serverclasses.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:16:		getSpotNodePoolTableConfig	100.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:37:		outputSpotNodePool		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:64:		outputSpotNodePools		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:110:		outputCreatedSpotNodePool	0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:137:		getNamespace			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:15:		NewCreateCommand		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:52:		runCreate			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:156:		loadSpecFromFile		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete-all.go:13:	NewDeleteAllCommand		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete-all.go:47:	runDeleteAll			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete.go:13:		NewDeleteCommand		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete.go:44:		runDelete			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/edit.go:13:		NewEditCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/edit.go:69:		runEdit				0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/get.go:42:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/list.go:39:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/spotnodepool.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/version.go:27:			init				0.0%
github.com/georgetaylor/spotctl/main.go:9:				main				0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:46:			NewTokenManager			100.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:55:			GetValidAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:70:			refreshAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:124:			IsValid				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:28:		String				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:34:		NewAPIVersion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:39:		IsValid				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:49:		GetAllAPIVersions		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:70:		Error				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:78:		NewClient			100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:102:		prepareRequest			77.3%
github.com/georgetaylor/spotctl/pkg/client/client.go:142:		doRequest			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:182:		MakeRequest			71.4%
github.com/georgetaylor/spotctl/pkg/client/client.go:204:		Get				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:212:		Post				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:220:		Put				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:228:		Delete				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:236:		Patch				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:245:		ListRegions			75.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:254:		GetRegion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:267:		ListServerClasses		75.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:276:		GetServerClass			66.7%
github.com/georgetaylor/spotctl/pkg/client/client.go:289:		ListOrganizations		75.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:299:		ListCloudSpaces			85.7%
github.com/georgetaylor/spotctl/pkg/client/client.go:313:		CreateCloudSpace		88.9%
github.com/georgetaylor/spotctl/pkg/client/client.go:330:		DeleteCloudSpace		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:352:		GetCloudSpace			88.9%
github.com/georgetaylor/spotctl/pkg/client/client.go:369:		EditCloudSpace			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:389:		ListSpotNodePools		85.7%
github.com/georgetaylor/spotctl/pkg/client/client.go:403:		CreateSpotNodePool		88.9%
github.com/georgetaylor/spotctl/pkg/client/client.go:420:		EditSpotNodePool		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:440:		GetSpotNodePool			88.9%
github.com/georgetaylor/spotctl/pkg/client/client.go:457:		DeleteSpotNodePool		88.9%
github.com/georgetaylor/spotctl/pkg/client/client.go:479:		DeleteAllSpotNodePools		85.7%
github.com/georgetaylor/spotctl/pkg/client/client.go:497:		HandleAPIError			17.6%
github.com/georgetaylor/spotctl/pkg/client/generic.go:48:		genericList			84.6%
github.com/georgetaylor/spotctl/pkg/client/generic.go:74:		genericGet			84.6%
github.com/georgetaylor/spotctl/pkg/client/generic.go:100:		genericCreate			76.9%
github.com/georgetaylor/spotctl/pkg/client/generic.go:126:		genericEdit			0.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:152:		genericDelete			47.8%
github.com/georgetaylor/spotctl/pkg/client/generic.go:208:		validateNamespace		100.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:215:		validateName			100.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:222:		validateCreateInput		80.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:236:		validatePatchOperations		0.0%
github.com/georgetaylor/spotctl/pkg/client/patch.go:17:			LoadPatchOperations		0.0%
github.com/georgetaylor/spotctl/pkg/client/patch.go:32:			DisplayPatchOperations		0.0%
github.com/georgetaylor/spotctl/pkg/client/patch.go:62:			PromptForConfirmation		0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:38:		ValidateConfig			100.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:57:		GetConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:80:		SaveConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:120:		InitConfig			0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:30:		Error				0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:37:		Unwrap				0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:42:		NewAPIError			0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:52:		NewConfigError			0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:62:		NewValidationError		0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:72:		NewInternalError		0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:54:		NewFormatter			100.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:62:		NewFormatterWithPager		0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:71:		Output				100.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:78:		OutputToWriter			55.6%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:99:		outputJSON			0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:104:		outputJSONToWriter		100.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:111:		outputYAML			0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:116:		outputYAMLToWriter		0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:124:		outputTable			0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:129:		outputTableToWriter		86.2%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:185:		extractItems			57.9%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:225:		getFieldValue			57.7%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:281:		findField			82.4%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:16:			max				0.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:34:			NewPager			100.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:41:			Write				35.3%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:80:			usePager			0.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:132:			getPagerCommand			100.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:151:			isTerminal			100.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:156:			getTerminalHeight		44.4%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:179:			getTerminalHeightStty		50.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:198:			WriteToWriter			0.0%
total:									(statements)			21.9%
```

## Coverage Status

| Package | Coverage | Status |
|---------|----------|--------|
| github.com/georgetaylor/spotctl/cmd/cloudspaces/cloudspaces.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:16: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:35: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:81: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:108: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:135: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:15: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:65: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:154: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:44: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:13: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:71: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:50: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:47: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:18: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:24: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:32: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:164: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:175: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:185: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:25: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/list.go:30: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/organizations.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:26: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:57: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/get.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/get.go:35: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/list.go:31: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/regions.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/root.go:32: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/root.go:36: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/root.go:63: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:35: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:30: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/serverclasses.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:16: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:37: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:64: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:110: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:137: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:15: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:52: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:156: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/delete-all.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/delete-all.go:47: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/delete.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/delete.go:44: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/edit.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/edit.go:69: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/get.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/get.go:42: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/list.go:39: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/spotnodepool.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/version.go:27: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/main.go:9: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:46: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:55: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:70: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:124: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:28: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:34: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:39: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:49: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:70: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:78: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:102: | 77.3% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:142: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:182: | 71.4% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:204: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:212: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:220: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:228: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:236: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:245: | 75.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:254: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:267: | 75.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:276: | 66.7% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:289: | 75.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:299: | 85.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:313: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:330: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:352: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:369: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:389: | 85.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:403: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:420: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:440: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:457: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:479: | 85.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:497: | 17.6% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:48: | 84.6% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:74: | 84.6% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:100: | 76.9% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:126: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:152: | 47.8% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:208: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:215: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:222: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:236: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/patch.go:17: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/patch.go:32: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/patch.go:62: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:38: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:57: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:80: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:120: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/errors/errors.go:30: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/errors/errors.go:37: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/errors/errors.go:42: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/errors/errors.go:52: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/errors/errors.go:62: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/errors/errors.go:72: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:54: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:62: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:71: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:78: | 55.6% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:99: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:104: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:111: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:116: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:124: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:129: | 86.2% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:185: | 57.9% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:225: | 57.7% | ❌ |
| github.com/georgetaylor/spotctl/pkg/output/formatter.go:281: | 82.4% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:16: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:34: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:41: | 35.3% | ❌ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:80: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:132: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:151: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:156: | 44.4% | ❌ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:179: | 50.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/pager/pager.go:198: | 0.0% | ❌ |

## Critical Paths Coverage

| File | Coverage | Status |
|------|----------|--------|
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:16: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:35: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:81: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:108: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:135: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:164: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:175: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:185: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:25: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:26: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:57: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:16: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:37: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:64: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:110: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:137: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:46: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:55: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:70: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:124: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:28: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:34: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:39: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:49: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:70: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:78: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:102: | 77.3% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:142: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:182: | 71.4% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:204: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:212: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:220: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:228: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:236: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:245: | 75.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:254: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:267: | 75.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:276: | 66.7% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:289: | 75.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:299: | 85.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:313: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:330: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:352: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:369: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:389: | 85.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:403: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:420: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:440: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:457: | 88.9% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:479: | 85.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:497: | 17.6% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:48: | 84.6% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:74: | 84.6% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:100: | 76.9% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:126: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:152: | 47.8% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:208: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:215: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:222: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/generic.go:236: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/patch.go:17: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/patch.go:32: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/patch.go:62: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:38: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:57: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:80: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:120: | 0.0% | ❌ |

## Untested Files

```
github.com/georgetaylor/spotctl/cmd/cloudspaces/cloudspaces.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:16:		getCloudSpacesTableConfig	100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:35:		outputCloudSpaces		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:81:		outputCreatedCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:108:		outputCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:135:		getNamespace			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:15:		NewCreateCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:65:		runCreate			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:154:		loadCloudSpaceSpecFromFile	0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:13:		NewDeleteCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:44:		runDelete			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:13:		NewEditCommand			100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:71:		runEdit				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:50:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:47:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/common.go:13:			AddOutputFlag			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:18:			GetOutputFormat			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:24:			ConfirmAction			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:32:			CheckError			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:164:			init				0.0%
github.com/georgetaylor/spotctl/cmd/config.go:175:			contains			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:185:			min				0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:11:		getOrganizationsTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:25:		outputOrganizations		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:30:		runOrganizationsList		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/organizations.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:11:		getRegionsTableConfig		0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:26:		outputRegions			0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:57:		outputRegion			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:13:			NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:35:			runGet				0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:13:			NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:31:			runList				0.0%
github.com/georgetaylor/spotctl/cmd/regions/regions.go:6:		NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/root.go:32:				Execute				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:36:				init				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:63:				initConfig			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11:		getServerClassesTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33:		outputServerClasses		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64:		outputServerClass		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:35:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:30:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/serverclasses.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:16:		getSpotNodePoolTableConfig	100.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:37:		outputSpotNodePool		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:64:		outputSpotNodePools		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:110:		outputCreatedSpotNodePool	0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/config.go:137:		getNamespace			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:15:		NewCreateCommand		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:52:		runCreate			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/create.go:156:		loadSpecFromFile		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete-all.go:13:	NewDeleteAllCommand		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete-all.go:47:	runDeleteAll			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete.go:13:		NewDeleteCommand		0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/delete.go:44:		runDelete			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/edit.go:13:		NewEditCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/edit.go:69:		runEdit				0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/get.go:42:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/list.go:39:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/spotnodepool/spotnodepool.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/version.go:27:			init				0.0%
github.com/georgetaylor/spotctl/main.go:9:				main				0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:46:			NewTokenManager			100.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:55:			GetValidAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:70:			refreshAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:124:			IsValid				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:28:		String				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:34:		NewAPIVersion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:39:		IsValid				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:49:		GetAllAPIVersions		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:70:		Error				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:78:		NewClient			100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:142:		doRequest			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:204:		Get				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:212:		Post				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:220:		Put				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:228:		Delete				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:236:		Patch				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:254:		GetRegion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:330:		DeleteCloudSpace		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:369:		EditCloudSpace			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:420:		EditSpotNodePool		0.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:126:		genericEdit			0.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:208:		validateNamespace		100.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:215:		validateName			100.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:222:		validateCreateInput		80.0%
github.com/georgetaylor/spotctl/pkg/client/generic.go:236:		validatePatchOperations		0.0%
github.com/georgetaylor/spotctl/pkg/client/patch.go:17:			LoadPatchOperations		0.0%
github.com/georgetaylor/spotctl/pkg/client/patch.go:32:			DisplayPatchOperations		0.0%
github.com/georgetaylor/spotctl/pkg/client/patch.go:62:			PromptForConfirmation		0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:38:		ValidateConfig			100.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:57:		GetConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:80:		SaveConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:120:		InitConfig			0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:30:		Error				0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:37:		Unwrap				0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:42:		NewAPIError			0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:52:		NewConfigError			0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:62:		NewValidationError		0.0%
github.com/georgetaylor/spotctl/pkg/errors/errors.go:72:		NewInternalError		0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:54:		NewFormatter			100.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:62:		NewFormatterWithPager		0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:71:		Output				100.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:99:		outputJSON			0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:104:		outputJSONToWriter		100.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:111:		outputYAML			0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:116:		outputYAMLToWriter		0.0%
github.com/georgetaylor/spotctl/pkg/output/formatter.go:124:		outputTable			0.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:16:			max				0.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:34:			NewPager			100.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:80:			usePager			0.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:132:			getPagerCommand			100.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:151:			isTerminal			100.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:179:			getTerminalHeightStty		50.0%
github.com/georgetaylor/spotctl/pkg/pager/pager.go:198:			WriteToWriter			0.0%
```

## Notes

- ✅ 100% coverage
- 🟡 Meets threshold (80%) but not perfect
- ❌ Below threshold (80%)
