resource "null_resource" "sh_example_1" {
  provisioner "local-exec" {
    command = <<EOF
      curl -XPOST 'https://realtimelog.herokuapp.com:443/p0xqfmznqa' \
      -H 'Content-Type: application/json' -d '{"code":"200", "status": "OK"}'
    EOF
  }
}
