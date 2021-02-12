package main

import "fmt"

var sample = fmt.Sprintf(`
This is the **desc**.

%s
apiVersion: apps/v1
%s
%s
apiVersion: apps/v1
%s
`, "```", "```", "```yaml", "```")