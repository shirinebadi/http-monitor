package config

const defaultConfig = `
server:
  address: :61432
jwt:
  secret: jdnfksdmfks
  expiration: 60
nats:
  host: nats://localhost:4222
  topic: url
  queue: httpMonitor
common:
  period: 3
  `
