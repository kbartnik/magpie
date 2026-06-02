---
title: "Digital Surveillance Resistance"
type: concept
status: active
created: 2026-06-01
updated: 2026-06-01
sources:
  - "[[archive/clippings/2026-06-01-wired-surveillance-guide]]"
related:
  - "[[Security Culture]]"
  - "[[Community Self-Defense]]"
  - "[[Threat Modeling]]"
tags:
  - security
  - privacy
  - digital-hygiene
---

# Digital Surveillance Resistance

Practical tools and practices for protecting personal data and communications from government surveillance. As of 2024–25, legal and policy protections have collapsed in the US; technology is the last line of defense for at-risk communities.

## The Realistic Goal: Compartmentalization, Not Anonymity

True anonymity is essentially unattainable for most people — data brokers hold years of accumulated history, and every slip in operational security undermines a burner phone or alias. The achievable goal is **compartmentalization**: separating digital identities so that no single breach or data request exposes everything.

A second phone with a separate digital identity (different accounts, different apps) is more effective in practice than a "burner" phone that requires near-perfect discipline to not accidentally link to your real identity.

## Threat Modeling First

Before choosing tools, identify your actual risk profile:
- Who might target you? (law enforcement, employer, far-right actors, stalker)
- What are they likely to demand? (records from a company, device search, location data)
- What data currently exists about you, and where?

Tools should match the threat, not the most paranoid possible scenario.

## Tool Hierarchy

### Communications
| Tool | Protection | Notes |
|------|-----------|-------|
| **Signal** | End-to-end encrypted, no metadata, no cloud backup, disappearing messages | Best overall; "With Signal, there are no records to seize" |
| iMessage / FaceTime | End-to-end encrypted | Backed up to iCloud by default — disable this |
| WhatsApp | E2E encrypted (with E2E backup enabled) | Meta owns it; enable encrypted backup |
| Telegram | Encrypted in transit only (default) | Server-accessible; avoid for sensitive comms |
| Facebook Messenger | Encrypted in transit only (default) | Server-accessible; Nebraska abortion case |

**Enable disappearing messages on Signal.** Set to the shortest duration that fits your use.

### Devices
- **Phones:** Full disk encryption is on by default on modern iOS and Android. Use a long alphanumeric passphrase, not a 6-digit PIN.
- **Biometrics:** Convenient, but you can be compelled to use your face. Know the gesture to temporarily disable FaceID/fingerprint (hold side button + volume on iPhone; long-press power on Android) before border crossings or protests.
- **Powered off = harder to crack.** Cryptographic keys leave memory when the device shuts down.
- **Laptops:** Enable FileVault (Mac) or BitLocker (Windows Pro) or VeraCrypt (Windows Home).

### Cloud Storage
- Cloud data is accessible to providers and subject to government requests.
- Use end-to-end encrypted iCloud backup (opt-in, not default).
- Assume anything in unencrypted cloud is accessible to law enforcement.
- Simplest rule: keep sensitive data off the cloud entirely.

### Browsing Anonymity
- **Tor Browser** — triple-layer encrypted proxy; strongest protection; slow.
- **Brave with private window** — stripped-down Tor routing; practical for daily use.
- **VPN** — one-hop proxy; choose carefully (many log activity). Refer to Freedom of the Press Foundation's VPN guide.

### Location Data
- Your phone pings cell towers continuously — this can't be fully disabled while on.
- Simplest protection: leave phone at home when you need location privacy.
- Alternatives: airplane mode, Faraday bag (blocks all electromagnetic signals).
- Review and revoke app location permissions regularly.
- Use cash. Credit/debit transactions are transparent to law enforcement.

### Financial Privacy
- Cash is the strongest protection.
- Cryptocurrency: bitcoin has no real anonymity (public blockchain, traceable). Monero/Zcash are harder to trace but still imperfect.
- Payment apps (Venmo, Cash App) are as transparent as banks to law enforcement.

## The Past Problem

Data brokers hold years of location, behavioral, and financial records on almost everyone. Changing behavior now protects your future self but doesn't erase the past. Begin anyway — the goal is to reduce exposure incrementally.

## Relationship to Security Culture

[[Security Culture]]'s principles (need-to-know, minimize trails, compartmentalize) map directly onto digital hygiene. This article is the practical, modern implementation of those 2004 principles — updated for smartphones, cloud services, and the current threat environment.

## AI Scraping Resistance and Behavioral Poisoning

A related but distinct domain: active poisoning rather than defensive hygiene. See [[AI Scraping Resistance]] for two layers:
- **Training-time**: subtitle poisoning, Nightshade, Glaze — corrupt what scrapers collect for model training
- **Inference-time**: deliberate behavioral noise (out-of-context engagement, false signals) — disrupt the profiling predictions a running surveillance system makes about you

The adversary differs — corporate AI scrapers and profiling systems vs. state surveillance — but the posture is shared: technical resistance when policy fails.

## Sources

- [[archive/clippings/2026-06-01-wired-surveillance-guide]] — Andy Greenberg, WIRED (2024)
