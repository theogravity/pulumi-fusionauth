## Publishing new versions

- Update `provider/go.mod` with the appropriate `github.com/gpsinsight/terraform-provider-fusionauth` version
- Run `cd provider && go mod tidy && cd -`
- Run `make tfgen` in the root of this directory
- Check the warnings for any new mappings that need to be added to `provider/resources.go` and run `make tfgen` again after
    corrections are made.
  * See `Adding Mappings, Building the Provider and SDKs` in `README-BUILD.md` for more info.
  * You'll have to run `gofmt -s provider/resources.go > provider/resources.go.new` and delete `resources.go` and rename `resources.go.new` to `resources.go` to get the formatting correct. 
- Run `cd sdk && go mod tidy && cd -`
- Run `make build_sdks`
- Commit all the generated files and push.
- Make sure the CI passes green.
- Create a release by making a git tag in the format of `vX.Y.Z` and push that tag to origin via `git push origin <tag>`
- Look at the action tab in the github repo to monitor the build for that version
