resource "jamfprotect_preventlist" "made001" {
  name        = "testtfnew"
  description = "test"
  type        = "CDHASH"
  list        = ["test", "test100"]
}
