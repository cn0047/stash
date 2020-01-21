variable "public_key_path" {
  type = "string"
  default = "/Users/kovpakvolodymyr/web/kovpak/gh/ed/sh/ssh/examples/nopwd/id_rsa.pub"
}

resource "aws_key_pair" "mysshkey" {
  key_name = "mysshkey"
  public_key = "${file(var.public_key_path)}"
  # public_key = "ssh-rsa AAAAB3Nz... email@email.local"
}
