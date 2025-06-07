# Test Coverage Report

Generated on: Sat Jun  7 22:43:18 UTC 2025

## Coverage Summary

```
github.com/georgetaylor/spotctl/cmd/cloudspaces/cloudspaces.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:14:		getCloudSpacesTableConfig	100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:33:		outputCloudSpaces		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:79:		outputCreatedCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:106:		outputCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:13:		NewCreateCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:44:		runCreate			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:13:		NewDeleteCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:36:		runDelete			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:22:		loadPatchOperations		100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:37:		displayPatchOperations		92.9%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:67:		promptForConfirmation		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:75:		NewEditCommand			100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:126:		runEdit				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:42:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:39:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/common.go:13:			AddOutputFlag			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:18:			GetOutputFormat			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:24:			ConfirmAction			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:32:			CheckError			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:143:			init				0.0%
github.com/georgetaylor/spotctl/cmd/config.go:154:			contains			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:164:			min				0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:11:		getOrganizationsTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:25:		outputOrganizations		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:30:		runOrganizationsList		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/organizations.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/pricing.go:36:			init				0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:11:		getRegionsTableConfig		0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:26:		outputRegions			0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:57:		outputRegion			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:13:			NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:35:			runGet				0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:13:			NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:31:			runList				0.0%
github.com/georgetaylor/spotctl/cmd/regions/regions.go:6:		NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/root.go:31:				Execute				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:35:				init				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:59:				initConfig			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11:		getServerClassesTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33:		outputServerClasses		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64:		outputServerClass		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:35:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:30:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/serverclasses.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/version.go:27:			init				0.0%
github.com/georgetaylor/spotctl/main.go:9:				main				0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:46:			NewTokenManager			100.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:55:			GetValidAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:70:			refreshAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:124:			IsValid				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:32:		String				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:38:		NewAPIVersion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:43:		IsValid				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:53:		GetAllAPIVersions		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:74:		Error				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:82:		NewClient			100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:106:		prepareRequest			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:150:		doRequest			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:188:		MakeRequest			71.4%
github.com/georgetaylor/spotctl/pkg/client/client.go:210:		Get				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:214:		Post				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:218:		Put				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:222:		Delete				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:226:		Patch				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:230:		PatchWithContentType		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:235:		GetAuth				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:239:		PostAuth			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:243:		PutAuth				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:247:		DeleteAuth			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:252:		ListRegions			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:272:		GetRegion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:296:		ListServerClasses		90.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:316:		GetServerClass			83.3%
github.com/georgetaylor/spotctl/pkg/client/client.go:340:		ListOrganizations		90.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:361:		ListCloudSpaces			84.6%
github.com/georgetaylor/spotctl/pkg/client/client.go:386:		CreateCloudSpace		86.7%
github.com/georgetaylor/spotctl/pkg/client/client.go:414:		DeleteCloudSpace		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:456:		GetCloudSpace			86.7%
github.com/georgetaylor/spotctl/pkg/client/client.go:484:		EditCloudSpace			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:512:		HandleAPIError			16.7%
github.com/georgetaylor/spotctl/pkg/config/config.go:36:		ValidateConfig			100.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:62:		GetConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:84:		SaveConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:123:		InitConfig			0.0%
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
total:									(statements)			28.3%
```

## Coverage Status

| Package | Coverage | Status |
|---------|----------|--------|
| github.com/georgetaylor/spotctl/cmd/cloudspaces/cloudspaces.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:14: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:33: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:79: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:106: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:44: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:36: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:22: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:37: | 92.9% | 🟡 |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:67: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:75: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:126: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:42: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:39: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:18: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:24: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/common.go:32: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:143: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:154: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:164: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:25: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/list.go:30: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/organizations.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/pricing.go:36: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:26: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:57: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/get.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/get.go:35: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/list.go:31: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/regions.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/root.go:31: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/root.go:35: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/root.go:59: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:35: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:13: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:30: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/serverclasses.go:6: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/version.go:27: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/main.go:9: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:46: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:55: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:70: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:124: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:32: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:38: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:43: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:53: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:74: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:82: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:106: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:150: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:188: | 71.4% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:210: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:214: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:218: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:222: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:226: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:230: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:235: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:239: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:243: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:247: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:252: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:272: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:296: | 90.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:316: | 83.3% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:340: | 90.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:361: | 84.6% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:386: | 86.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:414: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:456: | 86.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:484: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:512: | 16.7% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:36: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:62: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:84: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:123: | 0.0% | ❌ |
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
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:14: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:33: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:79: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:106: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:143: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:154: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/config.go:164: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/organizations/config.go:25: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:26: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/regions/config.go:57: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:46: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:55: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:70: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/auth.go:124: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:32: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:38: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:43: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:53: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:74: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:82: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:106: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:150: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:188: | 71.4% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:210: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:214: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:218: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:222: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:226: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:230: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:235: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:239: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:243: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:247: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:252: | 80.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:272: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:296: | 90.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:316: | 83.3% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:340: | 90.0% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:361: | 84.6% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:386: | 86.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:414: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:456: | 86.7% | 🟡 |
| github.com/georgetaylor/spotctl/pkg/client/client.go:484: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/client/client.go:512: | 16.7% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:36: | 100.0% | ✅ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:62: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:84: | 0.0% | ❌ |
| github.com/georgetaylor/spotctl/pkg/config/config.go:123: | 0.0% | ❌ |

## Untested Files

```
github.com/georgetaylor/spotctl/cmd/cloudspaces/cloudspaces.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:14:		getCloudSpacesTableConfig	100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:33:		outputCloudSpaces		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:79:		outputCreatedCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/config.go:106:		outputCloudSpace		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:13:		NewCreateCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/create.go:44:		runCreate			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:13:		NewDeleteCommand		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/delete.go:36:		runDelete			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:22:		loadPatchOperations		100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:67:		promptForConfirmation		0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:75:		NewEditCommand			100.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/edit.go:126:		runEdit				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/get.go:42:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/cloudspaces/list.go:39:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/common.go:13:			AddOutputFlag			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:18:			GetOutputFormat			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:24:			ConfirmAction			0.0%
github.com/georgetaylor/spotctl/cmd/common.go:32:			CheckError			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:143:			init				0.0%
github.com/georgetaylor/spotctl/cmd/config.go:154:			contains			0.0%
github.com/georgetaylor/spotctl/cmd/config.go:164:			min				0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:11:		getOrganizationsTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/organizations/config.go:25:		outputOrganizations		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/organizations/list.go:30:		runOrganizationsList		0.0%
github.com/georgetaylor/spotctl/cmd/organizations/organizations.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/pricing.go:36:			init				0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:11:		getRegionsTableConfig		0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:26:		outputRegions			0.0%
github.com/georgetaylor/spotctl/cmd/regions/config.go:57:		outputRegion			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:13:			NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/get.go:35:			runGet				0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:13:			NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/regions/list.go:31:			runList				0.0%
github.com/georgetaylor/spotctl/cmd/regions/regions.go:6:		NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/root.go:31:				Execute				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:35:				init				0.0%
github.com/georgetaylor/spotctl/cmd/root.go:59:				initConfig			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:11:		getServerClassesTableConfig	0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:33:		outputServerClasses		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/config.go:64:		outputServerClass		0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:13:		NewGetCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/get.go:35:		runGet				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:13:		NewListCommand			0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/list.go:30:		runList				0.0%
github.com/georgetaylor/spotctl/cmd/serverclasses/serverclasses.go:6:	NewCommand			0.0%
github.com/georgetaylor/spotctl/cmd/version.go:27:			init				0.0%
github.com/georgetaylor/spotctl/main.go:9:				main				0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:46:			NewTokenManager			100.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:55:			GetValidAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:70:			refreshAccessToken		0.0%
github.com/georgetaylor/spotctl/pkg/client/auth.go:124:			IsValid				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:32:		String				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:38:		NewAPIVersion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:43:		IsValid				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:53:		GetAllAPIVersions		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:74:		Error				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:82:		NewClient			100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:106:		prepareRequest			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:150:		doRequest			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:210:		Get				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:214:		Post				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:218:		Put				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:222:		Delete				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:226:		Patch				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:230:		PatchWithContentType		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:235:		GetAuth				100.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:239:		PostAuth			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:243:		PutAuth				0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:247:		DeleteAuth			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:252:		ListRegions			80.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:272:		GetRegion			0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:296:		ListServerClasses		90.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:340:		ListOrganizations		90.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:414:		DeleteCloudSpace		0.0%
github.com/georgetaylor/spotctl/pkg/client/client.go:484:		EditCloudSpace			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:36:		ValidateConfig			100.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:62:		GetConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:84:		SaveConfig			0.0%
github.com/georgetaylor/spotctl/pkg/config/config.go:123:		InitConfig			0.0%
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
