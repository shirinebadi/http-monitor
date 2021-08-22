package config

const defaultConfig = `
server:
  address: :21345
jwt:
  secret: jdnfksdmfks
  expiration: 60
nats:
  host: nats://localhost:4222
  topic: url
  name: httpMonitor
common:
  period: 3
  `
