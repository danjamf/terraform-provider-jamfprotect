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
