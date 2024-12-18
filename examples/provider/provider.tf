terraform {
  required_providers {
    jamfprotect = {
      source  = "jamfprotect"
      version = "1.0.0"
    }

  }
}

provider "jamfprotect" {

  clientid       = "clientidhere"
  clientpassword = "clientpasswordhere"
  domainname     = "yourdomain.protect.jamfcloud.com"

}
