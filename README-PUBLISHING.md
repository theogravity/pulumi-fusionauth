## Publishing new versions

- Update `provider/go.mod` with the appropriate `github.com/gpsinsight/terraform-provider-fusionauth` version
- Run `cd provider && go mod tidy && cd -`
- Run `make tfgen` in the root of this directory
- Run `cd sdk && go mod tidy && cd -`
- Run `make build_sdks`
- Check the warnings for any new mappings that need to be added to `provider/resources.go` and run `make tfgen` again after
  corrections are made.
  * See `Adding Mappings, Building the Provider and SDKs` in `README-BUILD.md` for more info.
- Commit all the generated files and push.
- Make sure the CI passes green.
- Create a release by making a git tag in the format of `vX.Y.Z` and push that tag to origin via `git push origin <tag>`
- Look at the action tab in the github repo to monitor the build for that version
