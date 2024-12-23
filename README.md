This functionality has been abosorbed into [JSC Provider](https://github.com/danjamf/terraform-provider-jsctfprovider)

This repo is no longer maintained 

# terraform-provider-jamfprotect

Example TF for Jamf Protect

## Building

`go build -o terraform-provider-jamfprotect`

## Configuring

Add `~/.terraformrc` for dev overrides. Example config

```provider_installation {

  dev_overrides {
        "jamfprotect" = "/Users/userpath/golang/jamfprotect"
}
```
