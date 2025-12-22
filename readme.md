<p align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&color=10b981&height=250&section=header&text=CIPHER-OSINT%20v7.0&fontSize=80&animation=fadeIn&fontAlignY=35&desc=ADVANCED%20OSINT%20AUTOMATION&descAlignY=55&descSize=20" />
</p>

<p align="center">
  <a href="https://github.com/cipher-attack">
    <svg width="220" height="220" viewBox="0 0 100 100" fill="none" xmlns="http://www.w3.org/2000/svg" style="filter: drop-shadow(0px 0px 15px rgba(16, 185, 129, 0.6));">
      <circle cx="50" cy="50" r="48" stroke="#10b981" stroke-width="0.5" stroke-dasharray="4 2">
        <animateTransform attributeName="transform" type="rotate" from="0 50 50" to="360 50 50" dur="15s" repeatCount="indefinite" />
      </circle>
      <path d="M 75 30 L 35 30 L 15 50 L 35 70 L 75 70" stroke="#10b981" stroke-width="4" stroke-linecap="round" stroke-linejoin="round">
        <animate attributeName="stroke-opacity" values="1;0.4;1" dur="2s" repeatCount="indefinite" />
      </path>
      <circle cx="45" cy="50" r="8" stroke="#ffffff" stroke-width="3">
        <animate attributeName="r" values="7;9;7" dur="1.5s" repeatCount="indefinite" />
      </circle>
      <path d="M 60 50 L 85 50" stroke="#10b981" stroke-width="5" stroke-linecap="round"/>
      <path d="M 70 50 L 70 65" stroke="#10b981" stroke-width="3" stroke-linecap="round"/>
      <path d="M 80 50 L 80 60" stroke="#10b981" stroke-width="3" stroke-linecap="round"/>
      <circle cx="50" cy="50" r="3" fill="#10b981">
        <animate attributeName="r" values="2;5;2" dur="1.2s" repeatCount="indefinite" />
        <animate attributeName="opacity" values="1;0.2;1" dur="1.2s" repeatCount="indefinite" />
      </circle>
    </svg>
  </a>
</p>

<h1 align="center">CIPHER-OSINT</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Engine-Golang-00ADD8?style=flat-square&logo=go" />
  <img src="https://img.shields.io/badge/Framework-Fiber-black?style=flat-square&logo=go" />
  <img src="https://img.shields.io/badge/Intelligence-AI%20Augmented-10b981?style=flat-square&logo=google-gemini" />
  <img src="https://img.shields.io/badge/Security-Tor%20Gateway-7D4698?style=flat-square&logo=tor-browser" />
</p>

<p align="center">
  <b>High-Performance OSINT Discovery & Intelligence Framework</b><br>
  <i>Scalable reconnaissance engine leveraging concurrent Golang architecture and neural synthesis for actionable data extraction.</i>
</p>

---

### Project Overview
**CIPHER-OSINT v7.0** is a sophisticated Open Source Intelligence (OSINT) framework designed for rapid data acquisition and analysis. Built with a focus on high concurrency using **Golang**, it streamlines the process of scanning surface web and dark web targets. By integrating advanced LLMs, it transforms raw technical indicators into structured intelligence reports with minimal manual intervention.

---

### Core Specifications

| Component | Technology | Functional Capability |
| :--- | :--- | :--- |
| **Neural Nexus** | `Gemini / GPT-4` | Automated synthesis of raw data into high-fidelity intelligence. |
| **Privacy Layer** | `Tor SOCKS5` | Stealth routing for .onion service reconnaissance and leak analysis. |
| **Control Plane** | `Go-Fiber` | High-performance dashboard with real-time markdown visualization. |
| **Data Integrity** | `AES-256` | Secure, encrypted archival of intelligence gathering sessions. |
| **Efficiency** | `Preprocessing` | Advanced metadata stripping to optimize AI processing and token cost. |

---

### System Architecture

```mermaid
graph LR
    A[Target Input] --> B{Core Orchestrator}
    B -->|Concurrent Scrape| C[Surface Web Nodes]
    B -->|SOCKS5 Proxy| D[Tor Service Nodes]
    C --> E[Data Normalization]
    D --> E
    E --> F{AI Analytical Engine}
    F -->|Reasoning| G[Interface Dashboard]
    G -->|Persistence| H[Intelligence Vault]
    style B fill:#333,stroke:#10b981,stroke-width:2px,color:#fff
    style F fill:#10b981,stroke:#000,stroke-width:2px,color:#000
```

---

### Deployment Protocol

Environment setup for Linux and Termux nodes.

```bash
# 1. Acquire Framework
git clone [https://github.com/cipher-attack/cipher-osint.git](https://github.com/cipher-attack/cipher-osint.git)

# 2. Resolve Module Dependencies
cd cipher-osint && go mod tidy

# 3. Environment Configuration
# Define your secure API credentials in the .env file
echo "GEMINI_API_KEY=your_secure_key" > .env

# 4. Initialize Core
go run .
```

> **Engineering Note:** Ensure your Tor relay is active on `127.0.0.1:9050` before initiating dark-web modules.

---

### 👤 The Architect

<div align="left">
  <img src="https://github.com/cipher-attack.png" width="150" align="left" style="border-radius: 15px; border: 3px solid #10b981; margin-right: 20px;" />
  <h3>Biruk Getachew</h3>
  <p><i>Principal Software Engineer & Cybersecurity Researcher</i></p>
  <p>Specializing in high-performance system architecture, neural interaction models, and sovereign intelligence systems. Engineering tools that bridge the gap between raw data and strategic insight.</p>
  <p>
    <a href="https://www.youtube.com/@cipher-attack"><img src="https://img.shields.io/badge/YouTube-FF0000?style=flat-square&logo=youtube&logoColor=white" /></a>
    <a href="https://github.com/cipher-attack"><img src="https://img.shields.io/badge/GitHub-181717?style=flat-square&logo=github&logoColor=white" /></a>
    <a href="https://t.me/cipher_attacks"><img src="https://img.shields.io/badge/Telegram-26A6E1?style=flat-square&logo=telegram&logoColor=white" /></a>
  </p>
</div>

<br clear="left"/>

---

### Ethical Disclosure
This framework is provided for authorized cybersecurity research and institutional intelligence purposes only. Unauthorized use of this tool for malicious activities is strictly prohibited. The developer assumes no liability for misuse.

<p align="center">
  <br>
  <img src="https://capsule-render.vercel.app/api?type=waving&color=10b981&height=100&section=footer" />
</p>
