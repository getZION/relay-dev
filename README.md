# Zion Relay v2

The new Zion relay server.

## Features
- Save messages and media via [Identity Hubs](https://identity.foundation/identity-hub/spec/) and IPFS
- Lightning user account system

## Who runs relays?

Zion runs a series of relays. The Zion v2 app connects to them by default.

Anyone can run their own relay. Users can connect their app to these other relays.

## Media hosting

Relays can set their own rules for:

- What sorts of media are accepted
- How much media uploads cost (in sats)
- Who may upload media
  - Zion-hosted relays run may restrict media uploads to paying customers

## Tech Stack
- [Go](https://go.dev/) 1.17
- [IPFS](https://docs.ipfs.io/concepts/what-is-ipfs/)
- [ION](https://identity.foundation/ion/)

## Specifications
We will gradually implement the following specifications:
- [Identity Hub](https://identity.foundation/identity-hub/spec) (DIF)
- [Decentralized Identifiers](https://www.w3.org/TR/did-core/) (W3C)
- [Verifiable Credentials](https://www.w3.org/TR/vc-data-model/) (W3C)
