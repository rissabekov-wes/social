# API Visualization - Implementation Plan

## Overview
This document outlines the plan to create an automated API visualization tool for Wesfarmers Go APIs, with initial focus on the `one_http` framework.

## Current State
- ✅ Manual HTML visualization created for `social` API
- ✅ Shows server config, middleware chain, routes, and router hierarchy
- ✅ Interactive tabs with real data from codebase
- 📄 Reference implementation: `docs/chi-visualization.html`

---

## 🎯 Goals

1. **Automate generation** of API visualizations from source code
2. **Standardize** across all Wesfarmers APIs
3. **Framework-aware** - detect one_http, chi, gin, echo, etc.
4. **Multi-IDE support** - VSCode, IntelliJ, Web, CLI
5. **Self-updating** - regenerate on code changes

---

## 🏗️ Proposed Architecture

### Central Repository Structure
```
Wesfarmers-Digital/developer-tools/
├── agents/
│   └── api-visualizer/
│       ├── README.md
│       ├── agent.yml                    # GitHub Copilot Agent definition
│       ├── src/
│       │   ├── analyzers/
│       │   │   ├── one_http.ts          # Wesfarmers one_http framework
│       │   │   ├── chi.ts               # go-chi/chi
│       │   │   ├── gin.ts               # gin-gonic/gin
│       │   │   └── echo.ts              # labstack/echo
│       │   ├── generators/
│       │   │   ├── html.ts              # HTML generation
│       │   │   └── markdown.ts          # Markdown generation
│       │   ├── parsers/
│       │   │   ├── ast.ts               # Go AST parser
│       │   │   └── routes.ts            # Route extractor
│       │   └── templates/
│       │       ├── chi-viz.html         # Based on social API viz
│       │       ├── gin-viz.html
│       │       └── components/
│       │           ├── middleware.html
│       │           ├── routes.html
│       │           └── hierarchy.html
│       ├── tests/
│       │   ├── fixtures/
│       │   │   └── social-api/          # This repo as test case
│       │   └── analyzers.test.ts
│       └── package.json
├── mcp-server/
│   ├── server.ts                        # MCP protocol implementation
│   ├── tools/
│   │   ├── analyze-routes.ts
│   │   ├── generate-viz.ts
│   │   └── detect-framework.ts
│   └── package.json
└── ide-extensions/
    ├── vscode/
    │   ├── extension.js
    │   └── package.json
    └── intellij/
        └── plugin.xml
```

---

## 📋 Implementation Phases

### **Phase 1: GitHub Copilot Agent** (2-3 weeks)

#### Week 1: Core Analyzer
- [ ] Create `Wesfarmers-Digital/developer-tools` repository
- [ ] Set up TypeScript project structure
- [ ] Build Go AST parser
- [ ] Implement `one_http` framework detector
- [ ] Extract routes from `RegisterRoute()` calls
- [ ] Extract middlewares from `generateDefaultMiddlewares()`
- [ ] Parse server configuration (port, TLS, service name)

#### Week 2: Template & Agent
- [ ] Convert `chi-visualization.html` to dynamic template
- [ ] Implement HTML generator with data injection
- [ ] Create `agent.yml` for GitHub Copilot
- [ ] Add command: `@api-visualizer analyze <file>`
- [ ] Test with `rissabekov-wes/social` repository

#### Week 3: Testing & Documentation
- [ ] Unit tests for all analyzers
- [ ] Integration test with real repos
- [ ] Write comprehensive README
- [ ] Register agent with GitHub
- [ ] Deploy to Wesfarmers organization

**Deliverables:**
- Working Copilot agent
- Supports one_http framework
- Generates interactive HTML
- Available via `@api-visualizer`

---

### **Phase 2: Multi-Framework Support** (2 weeks)

#### Framework Analyzers
- [ ] **chi/v5** - Pure chi without one_http wrapper
- [ ] **gin** - gin-gonic/gin framework
- [ ] **echo** - labstack/echo framework
- [ ] **net/http** - Standard library

#### Auto-Detection Logic
```typescript
function detectFramework(code: string): Framework {
  if (code.includes('one_http.NewServer')) return 'one_http';
  if (code.includes('chi.NewRouter')) return 'chi';
  if (code.includes('gin.Default')) return 'gin';
  if (code.includes('echo.New')) return 'echo';
  return 'http';
}
```

---

### **Phase 3: MCP Server** (2 weeks)

Expose functionality via Model Context Protocol for broader tool support.

---

### **Phase 4: IDE Extensions** (3-4 weeks)

#### VSCode Extension
- [ ] Command palette integration
- [ ] Right-click context menu
- [ ] Auto-generate on save
- [ ] WebView panel for inline viewing

---

## 📊 Effort Estimation

| Phase | Duration | Complexity | Priority |
|-------|----------|------------|----------|
| **Phase 1: Copilot Agent** | 2-3 weeks | Medium | High |
| **Phase 2: Multi-Framework** | 2 weeks | Medium | Medium |
| **Phase 3: MCP Server** | 2 weeks | Medium | Low |
| **Phase 4: VSCode Extension** | 2 weeks | Low | Medium |
| **Phase 4: IntelliJ Plugin** | 2 weeks | Medium | Low |

**Total: 10-13 weeks for complete implementation**

---

## 🎯 Success Metrics

### Adoption
- [ ] Used in >50% of Wesfarmers Go APIs
- [ ] Integrated into API starter kits
- [ ] Part of standard onboarding docs

### Quality
- [ ] 95%+ accuracy in route detection
- [ ] 100% accuracy in one_http framework detection
- [ ] <2 seconds generation time

---

## 🚀 Quick Start (MVP)

### Minimal Viable Product (1 week)

Create a simple CLI tool first:

```bash
# Install
npm install -g @wesfarmers/api-visualizer

# Use
api-visualizer analyze cmd/api-wd/main.go
# Output: docs/api-visualization.html
```

---

## 📚 Reference Implementation

This repository (`rissabekov-wes/social`) serves as the reference implementation:

- **Main file**: `cmd/api-wd/main.go`
- **Framework**: one_http (wraps chi/v5)
- **Visualization**: `docs/chi-visualization.html`
- **Config**: `internal/config/`
- **Routes**: `internal/api/`

### Current Visualization Features
- ✅ Server configuration display
- ✅ Middleware chain with execution order
- ✅ Route table with handlers
- ✅ Mux hierarchy visualization
- ✅ Interactive tabs
- ✅ Source code snippets
- ✅ curl command examples