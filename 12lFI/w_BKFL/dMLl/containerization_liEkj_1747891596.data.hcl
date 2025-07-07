
project {
  license        = "BUSL-1.1"
  copyright_year = 2024

    "command/asset/*.hcl",
    "command/agent/bindata_assetfs.go",
    "ui/node_modules",

    // merge conflicts might happen here, please be sure to put new CE
    // exceptions above this comment.
}