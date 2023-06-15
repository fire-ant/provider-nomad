job "example" {
  region = "global"
  datacenters = ["us-east-1"]
  type = "service"
  update {
    stagger      = "3s"
    max_parallel = 3
  }
  group "alpha" {
    # Specify the number of these tasks we want.
    count = 5
    task "sleep" {
      driver = "exec"

      config {
        command = "/bin/sleep"
        args    = ["300"]
      }
    }
  }
  group "beta" {
    # Specify the number of these tasks we want.
    count = 4
    task "sleep" {
      driver = "exec"

      config {
        command = "/bin/sleep"
        args    = ["300"]
      }
    }
  }
  group "delta" {
    # Specify the number of these tasks we want.
    count = 3
    task "sleep" {
      driver = "exec"

      config {
        command = "/bin/sleep"
        args    = ["300"]
      }
    }
  }
}