package main

workload_resources = [
  "ReplicaSet",
  "Deployment",
  "DaemonSet",
  "StatefulSet",
  "Job",
]

deny[msg] {
  input.kind == workload_resources[_]
  not input.spec.selector.matchLabels.service
  not input.spec.template.metadata.labels.service
  not input.spec.selector.matchLabels.role
  not input.spec.template.metadata.labels.role
  msg = sprintf("Please set [service] [role] pod template label and selector (spec.template.metadata.labels.app, spec.selector.matchLabels.app): [Resource=%s, Name=%s]", [input.kind, input.metadata.name])
}

expected_labels(labels) {
  labels.service
  labels.role
}

deny[msg] {
  not expected_labels(input.metadata.labels)
  msg = sprintf("Please set [service] [role] labels (metadata.labels): [Resource=%s, Name=%s]", [input.kind, input.metadata.name])
}

deny[msg] {
  input.kind == workload_resources[_]
  terminationGracePeriodSeconds := 30
  not input.spec.template.spec.terminationGracePeriodSeconds
  input.spec.template.spec.terminationGracePeriodSeconds > terminationGracePeriodSeconds

  msg = sprintf("Please set terminationGracePeriodSeconds lower than %d: [Resource=%s, Name=%s, terminationGracePeriodSeconds=%d]", [terminationGracePeriodSeconds, input.kind, input.metadata.name, input.spec.template.spec.terminationGracePeriodSeconds])
}

deny[msg] {
  input.kind == workload_resources[_]
  container := input.spec.template.spec.containers[_]
  not container.resources.requests.cpu

  msg = sprintf("Please set CPU Requests: [Resource=%s, Name=%s, Container=%s]", [input.kind, input.metadata.name, container.name])
}

deny[msg] {
  input.kind == workload_resources[_]
  container := input.spec.template.spec.containers[_]
  not container.resources.requests.memory

  msg = sprintf("Please set Memory Requests: [Resource=%s, Name=%s, Container=%s]", [input.kind, input.metadata.name, container.name])
}

deny[msg] {
  input.kind == workload_resources[_]
  ratio := 1.5
  container := input.spec.template.spec.containers[_]
  container.resources.requests.cpu
  container.resources.limits.cpu
  formated_requests := split(container.resources.requests.cpu, "m")
  formated_limits := split(container.resources.limits.cpu, "m")
  to_number(formated_limits[0]) / to_number(formated_requests[0]) > ratio

  msg = sprintf("CPU limits/requests ratio must be lower than %2.2f: [Resource=%s, Name=%s, Container=%s, limits=%s requests=%s]", [ratio, input.kind, input.metadata.name, container.name, container.resources.limits.cpu, container.resources.requests.cpu])
}

deny[msg] {
  input.kind == workload_resources[_]
  ratio := 1.5
  container := input.spec.template.spec.containers[_]
  container.resources.requests.memory
  container.resources.limits.memory
  formated_requests := split(container.resources.requests.memory, "m")
  formated_limits := split(container.resources.limits.memory, "m")
  to_number(formated_limits[0]) / to_number(formated_requests[0]) > ratio

  msg = sprintf("Memory limits/requests ratio must be lower than %2.2f: [Resource=%s, Name=%s, Container=%s, limits=%s requests=%s]", [ratio, input.kind, input.metadata.name, container.name, container.resources.limits.cpu, container.resources.requests.cpu])
}

warn[msg] {
  input.kind == workload_resources[_]
  container := input.spec.template.spec.containers[_]
  not container.readinessProbe

  msg = sprintf("Please set readinessProbe: [Resource=%s, Name=%s, Container=%s]", [input.kind, input.metadata.name, container.name])
}

warn[msg] {
  input.kind == workload_resources[_]
  container := input.spec.template.spec.containers[_]
  not contains(container.readinessProbe.httpGet.path, "/healthz")
  not contains(container.readinessProbe.exec.command, "grpc-health-probe")

  msg = sprintf("readinessProbe with httpGet.path=/healthz or exec.command=grpc-health-probe: [Resource=%s, Name=%s, Container=%s]", [input.kind, input.metadata.name, container.name])
}

