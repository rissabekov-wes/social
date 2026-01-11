# API Visualization Implementation Plan

## Overview

This document outlines a comprehensive plan for building an automated API visualization and documentation tool. The tool will automatically detect API routes in web applications, extract their metadata, and generate interactive visual documentation similar to Swagger/OpenAPI interfaces.

**Vision**: Create a framework-agnostic tool that developers can drop into any project to instantly visualize and test their API endpoints without manual configuration.

---

## Goals

### Primary Objectives
1. **Automatic Detection**: Scan codebases to identify API routes across multiple frameworks
2. **Zero Configuration**: Work out-of-the-box with sensible defaults
3. **Interactive UI**: Provide a beautiful, user-friendly interface for API exploration
4. **Framework Support**: Support Express, Fastify, NestJS, Next.js API routes, and more
5. **Type Safety**: Leverage TypeScript types for enhanced documentation

### Secondary Objectives
- Request/Response validation
- API testing capabilities
- Export to OpenAPI/Swagger format
- Authentication flow support
- Customizable themes and branding

---

## Architecture

```
┌─────────────────────────────────────────────────────┐
│                   CLI Interface                     │
│              (api-viz scan, serve)                  │
└─────────────────┬───────────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────────┐
│              Route Scanner Engine                   │
│  ┌──────────────────────────────────────────────┐  │
│  │  Framework Detectors (Express, Fastify, etc) │  │
│  └──────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────┐  │
│  │     AST Parser (TypeScript/JavaScript)       │  │
│  └──────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────┐  │
│  │      Metadata Extractor (JSDoc, TS Types)    │  │
│  └──────────────────────────────────────────────┘  │
└─────────────────┬───────────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────────┐
│            Route Registry (JSON/DB)                 │
└─────────────────┬───────────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────────┐
│              Visualization Server                   │
│  ┌──────────────────────────────────────────────┐  │
│  │        REST API (Serve route data)           │  │
│  └──────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────┐  │
│  │    React UI (Interactive Documentation)      │  │
│  └──────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────┘
```

---

## Implementation Phases

### Phase 1: Core Scanner (Weeks 1-3)

#### Deliverables
- CLI tool initialization (`npx create-api-viz`)
- Framework detection system
- Basic route extraction for Express.js
- JSON output format specification

#### Tasks
1. Set up monorepo structure (Turborepo or Nx)
2. Implement framework detection heuristics
3. Build Express.js route parser using AST analysis
4. Create standardized route metadata schema
5. Write unit tests for scanner

---

### Phase 2: Multi-Framework Support (Weeks 4-6)

#### Deliverables
- Support for Fastify, NestJS, Next.js API routes
- TypeScript type extraction
- JSDoc comment parsing
- Request/Response schema inference

#### Tasks
1. Implement Fastify route parser
2. Implement NestJS decorator parser
3. Implement Next.js API route scanner
4. Build TypeScript type analyzer
5. Extract and parse JSDoc annotations
6. Enhance metadata schema with types and validation

---

### Phase 3: Visualization UI (Weeks 7-10)

#### Deliverables
- React-based interactive UI
- Route browser with search/filter
- Request builder interface
- Response viewer
- Dark/Light theme support

#### Tasks
1. Design UI/UX mockups
2. Build React component library
3. Implement route listing and navigation
4. Create request form with validation
5. Build response display with syntax highlighting
6. Add authentication configuration UI
7. Implement theme system

---

### Phase 4: Advanced Features (Weeks 11-14)

#### Deliverables
- OpenAPI/Swagger export
- API testing suite
- Request history and collections
- Environment variable management
- Plugin system for extensibility

#### Tasks
1. Build OpenAPI 3.0 export functionality
2. Implement API testing framework
3. Create request collection system
4. Add environment configuration
5. Design and implement plugin architecture
6. Write comprehensive documentation
7. Create example projects for each framework

---

## Technical Details

### Framework Detection

```typescript
// packages/scanner/src/detectors/framework-detector.ts

import { readFile } from 'fs/promises';
import { join } from 'path';

export enum Framework {
  EXPRESS = 'express',
  FASTIFY = 'fastify',
  NESTJS = 'nestjs',
  NEXTJS = 'nextjs',
  KOA = 'koa',
  UNKNOWN = 'unknown'
}

export interface DetectionResult {
  framework: Framework;
  version?: string;
  confidence: number;
  entryPoints: string[];
}

export class FrameworkDetector {
  async detect(projectRoot: string): Promise<DetectionResult> {
    const packageJson = await this.readPackageJson(projectRoot);
    
    // Check dependencies for framework signatures
    const deps = {
      ...packageJson.dependencies,
      ...packageJson.devDependencies
    };
    
    if (deps['@nestjs/core']) {
      return {
        framework: Framework.NESTJS,
        version: deps['@nestjs/core'],
        confidence: 0.95,
        entryPoints: await this.findNestJSModules(projectRoot)
      };
    }
    
    if (deps['next']) {
      return {
        framework: Framework.NEXTJS,
        version: deps['next'],
        confidence: 0.95,
        entryPoints: [join(projectRoot, 'pages/api'), join(projectRoot, 'app/api')]
      };
    }
    
    if (deps['fastify']) {
      return {
        framework: Framework.FASTIFY,
        version: deps['fastify'],
        confidence: 0.90,
        entryPoints: await this.findFastifyRoutes(projectRoot)
      };
    }
    
    if (deps['express']) {
      return {
        framework: Framework.EXPRESS,
        version: deps['express'],
        confidence: 0.85,
        entryPoints: await this.findExpressRoutes(projectRoot)
      };
    }
    
    return {
      framework: Framework.UNKNOWN,
      confidence: 0,
      entryPoints: []
    };
  }
  
  private async readPackageJson(projectRoot: string): Promise<any> {
    const content = await readFile(join(projectRoot, 'package.json'), 'utf-8');
    return JSON.parse(content);
  }
  
  private async findNestJSModules(projectRoot: string): Promise<string[]> {
    // Implementation: Scan for @Module decorators
    return [];
  }
  
  private async findFastifyRoutes(projectRoot: string): Promise<string[]> {
    // Implementation: Scan for fastify.register patterns
    return [];
  }
  
  private async findExpressRoutes(projectRoot: string): Promise<string[]> {
    // Implementation: Scan for app.use, Router() patterns
    return [];
  }
}
```

### Route Metadata Schema

```typescript
// packages/core/src/types/route.ts

export interface RouteMetadata {
  id: string;
  path: string;
  method: HttpMethod;
  handler: HandlerInfo;
  metadata: {
    summary?: string;
    description?: string;
    tags?: string[];
    deprecated?: boolean;
  };
  parameters: Parameter[];
  requestBody?: RequestBody;
  responses: Response[];
  security?: SecurityRequirement[];
  sourceLocation: SourceLocation;
}

export enum HttpMethod {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  PATCH = 'PATCH',
  DELETE = 'DELETE',
  OPTIONS = 'OPTIONS',
  HEAD = 'HEAD'
}

export interface HandlerInfo {
  name: string;
  file: string;
  line: number;
  column: number;
}

export interface Parameter {
  name: string;
  in: 'path' | 'query' | 'header' | 'cookie';
  required: boolean;
  schema: JSONSchema;
  description?: string;
}

export interface RequestBody {
  required: boolean;
  contentType: string;
  schema: JSONSchema;
}

export interface Response {
  statusCode: number;
  description: string;
  contentType?: string;
  schema?: JSONSchema;
}

export interface JSONSchema {
  type: string;
  properties?: Record<string, JSONSchema>;
  items?: JSONSchema;
  required?: string[];
  enum?: any[];
  format?: string;
  description?: string;
}

export interface SourceLocation {
  file: string;
  line: number;
  column: number;
}

export interface SecurityRequirement {
  type: 'apiKey' | 'http' | 'oauth2' | 'openIdConnect';
  scheme?: string;
  in?: 'header' | 'query' | 'cookie';
  name?: string;
}
```

### Express Route Parser Example

```typescript
// packages/scanner/src/parsers/express-parser.ts

import * as ts from 'typescript';
import { RouteMetadata, HttpMethod } from '@api-viz/core';

export class ExpressParser {
  parse(sourceFile: ts.SourceFile): RouteMetadata[] {
    const routes: RouteMetadata[] = [];
    
    const visit = (node: ts.Node) => {
      // Look for app.get(), app.post(), router.put(), etc.
      if (ts.isCallExpression(node)) {
        const route = this.extractRoute(node, sourceFile);
        if (route) {
          routes.push(route);
        }
      }
      
      ts.forEachChild(node, visit);
    };
    
    visit(sourceFile);
    return routes;
  }
  
  private extractRoute(
    node: ts.CallExpression, 
    sourceFile: ts.SourceFile
  ): RouteMetadata | null {
    const expression = node.expression;
    
    if (!ts.isPropertyAccessExpression(expression)) {
      return null;
    }
    
    const method = expression.name.text.toUpperCase();
    if (!this.isHttpMethod(method)) {
      return null;
    }
    
    const [pathArg, handlerArg] = node.arguments;
    
    if (!pathArg || !handlerArg) {
      return null;
    }
    
    const path = this.extractPath(pathArg);
    if (!path) {
      return null;
    }
    
    const handler = this.extractHandler(handlerArg, sourceFile);
    const parameters = this.extractParameters(path, handlerArg);
    const jsdoc = this.extractJSDoc(handlerArg);
    
    return {
      id: `${method}-${path}`,
      path,
      method: method as HttpMethod,
      handler,
      metadata: {
        summary: jsdoc?.summary,
        description: jsdoc?.description,
        tags: jsdoc?.tags
      },
      parameters,
      requestBody: this.extractRequestBody(handlerArg),
      responses: this.extractResponses(handlerArg),
      sourceLocation: this.getSourceLocation(node, sourceFile)
    };
  }
  
  private isHttpMethod(method: string): boolean {
    return ['GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'OPTIONS', 'HEAD'].includes(method);
  }
  
  private extractPath(node: ts.Node): string | null {
    if (ts.isStringLiteral(node)) {
      return node.text;
    }
    return null;
  }
  
  private extractHandler(node: ts.Node, sourceFile: ts.SourceFile): any {
    // Implementation: Extract handler function details
    return {};
  }
  
  private extractParameters(path: string, handler: ts.Node): any[] {
    // Implementation: Extract path params, query params from types
    return [];
  }
  
  private extractJSDoc(node: ts.Node): any {
    // Implementation: Parse JSDoc comments
    return null;
  }
  
  private extractRequestBody(handler: ts.Node): any {
    // Implementation: Infer request body schema from handler
    return undefined;
  }
  
  private extractResponses(handler: ts.Node): any[] {
    // Implementation: Infer response schemas
    return [];
  }
  
  private getSourceLocation(node: ts.Node, sourceFile: ts.SourceFile): any {
    const { line, character } = sourceFile.getLineAndCharacterOfPosition(node.getStart());
    return {
      file: sourceFile.fileName,
      line: line + 1,
      column: character + 1
    };
  }
}
```

---

## Effort Estimation

| Phase | Component | Estimated Time | Priority |
|-------|-----------|----------------|----------|
| 1 | CLI Setup & Framework Detection | 1 week | Critical |
| 1 | Express Parser | 1.5 weeks | Critical |
| 1 | Route Schema Design | 0.5 weeks | Critical |
| 2 | Fastify Parser | 1 week | High |
| 2 | NestJS Parser | 1 week | High |
| 2 | Next.js Parser | 1 week | High |
| 2 | TypeScript Type Extraction | 1 week | Medium |
| 3 | UI Component Library | 2 weeks | Critical |
| 3 | Route Browser | 1 week | Critical |
| 3 | Request Builder | 1.5 weeks | Critical |
| 3 | Theme System | 0.5 weeks | Low |
| 4 | OpenAPI Export | 1 week | Medium |
| 4 | Testing Suite | 1.5 weeks | Medium |
| 4 | Plugin System | 1 week | Low |
| 4 | Documentation | 1 week | High |
| **Total** | | **14-16 weeks** | |

---

## Success Metrics

### Technical Metrics
- **Framework Coverage**: Support 5+ popular frameworks
- **Accuracy**: 95%+ route detection accuracy
- **Performance**: Scan 1000 routes in < 10 seconds
- **Type Safety**: Full TypeScript support with generated types

### User Metrics
- **Time to First Value**: < 2 minutes from install to visualization
- **Adoption**: 1000+ GitHub stars in first 6 months
- **Community**: 50+ community contributions
- **Documentation**: 90%+ API coverage documented

### Business Metrics
- **Cost Savings**: Reduce API documentation time by 80%
- **Developer Satisfaction**: 4.5+ rating on npm
- **Integration**: Used in 100+ production projects

---

## MVP Quick Start Guide

### For End Users

```bash
# Install the CLI globally
npm install -g @api-viz/cli

# Navigate to your project
cd my-express-app

# Scan and visualize your APIs
api-viz scan
api-viz serve

# Open browser to http://localhost:3001
```

### For Contributors

```bash
# Clone the repository
git clone https://github.com/rissabekov-wes/social.git
cd social

# Install dependencies
pnpm install

# Build all packages
pnpm build

# Run tests
pnpm test

# Start development
pnpm dev
```

---

## Reference Implementation

This repository serves as a reference implementation and working example of the API visualization tool.

### Repository Structure

```
rissabekov-wes/social/
├── packages/
│   ├── cli/              # Command-line interface
│   ├── core/             # Core types and utilities
│   ├── scanner/          # Route scanning engine
│   ├── server/           # Visualization server
│   └── ui/               # React UI components
├── examples/
│   ├── express-app/      # Example Express application
│   ├── fastify-app/      # Example Fastify application
│   ├── nestjs-app/       # Example NestJS application
│   └── nextjs-app/       # Example Next.js application
├── docs/
│   ├── API_VISUALIZATION_PLAN.md  # This document
│   ├── ARCHITECTURE.md   # Detailed architecture
│   └── CONTRIBUTING.md   # Contribution guidelines
└── README.md
```

### Key Files to Study

1. **packages/scanner/src/detectors/framework-detector.ts** - Learn how frameworks are detected
2. **packages/scanner/src/parsers/express-parser.ts** - See how Express routes are parsed
3. **packages/core/src/types/route.ts** - Understand the route metadata schema
4. **packages/ui/src/components/RouteExplorer.tsx** - Explore the UI implementation

---

## Contributing Guidelines

We welcome contributions from the community! Here's how you can help:

### Areas for Contribution

1. **Framework Parsers**: Add support for new frameworks (Hono, Elysia, etc.)
2. **UI Enhancements**: Improve the visualization interface
3. **Documentation**: Write guides and tutorials
4. **Bug Fixes**: Report and fix issues
5. **Testing**: Add test coverage

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Write tests for your changes
5. Ensure all tests pass (`pnpm test`)
6. Commit with conventional commits (`feat: add Hono parser`)
7. Push to your fork
8. Open a Pull Request

### Code Style

- Use TypeScript for all new code
- Follow the existing code style (Prettier + ESLint)
- Write meaningful commit messages
- Add JSDoc comments for public APIs
- Include tests for new features

### Testing Requirements

- Unit tests for all parsers and utilities
- Integration tests for end-to-end workflows
- Minimum 80% code coverage
- All tests must pass before merging

---

## Questions & Next Steps

### Open Questions

1. **Authentication**: How should we handle various auth schemes (JWT, OAuth, API Keys)?
2. **Versioning**: Should we support API versioning detection?
3. **Performance**: What's the best caching strategy for large codebases?
4. **Plugins**: What's the ideal plugin API for extensibility?
5. **Hosting**: Should we offer a hosted version or keep it self-hosted only?

### Immediate Next Steps

- [ ] Set up repository structure and CI/CD
- [ ] Implement framework detector
- [ ] Build Express parser MVP
- [ ] Create basic CLI commands
- [ ] Design route metadata schema
- [ ] Set up documentation site
- [ ] Create example Express application
- [ ] Write contribution guidelines

### Future Considerations

- **GraphQL Support**: Extend to visualize GraphQL schemas
- **WebSocket Routes**: Support WebSocket endpoint documentation
- **Performance Monitoring**: Integrate API performance metrics
- **Collaboration**: Add team collaboration features
- **CI/CD Integration**: Generate docs in CI pipeline
- **IDE Extensions**: Create VSCode extension for inline visualization

---

## Appendix

### Similar Tools Comparison

| Tool | Pros | Cons | Our Differentiation |
|------|------|------|---------------------|
| **Swagger/OpenAPI** | Industry standard, extensive ecosystem | Requires manual schema writing | Auto-detection, zero config |
| **Postman** | Great testing UI, collaboration features | Not code-first, separate tool | Integrated with codebase |
| **Insomnia** | Beautiful UI, open source | Manual API definition | Automatic route discovery |
| **ApiDoc** | Comment-based, multi-language | Requires annotations | Type inference from code |
| **TypeDoc** | Excellent TypeScript support | Not API-focused | Specialized for REST APIs |
| **Stoplight** | Design-first approach, validation | Paid, external tool | Free, embedded in project |

### Technology Stack

- **Language**: TypeScript
- **CLI**: Commander.js
- **Parser**: TypeScript Compiler API
- **UI**: React, TailwindCSS, Radix UI
- **Build**: Turborepo, Vite
- **Testing**: Vitest, Testing Library
- **Documentation**: VitePress

### Resources & References

- [TypeScript Compiler API Documentation](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API)
- [OpenAPI Specification 3.0](https://swagger.io/specification/)
- [Express.js Routing Guide](https://expressjs.com/en/guide/routing.html)
- [NestJS Controllers](https://docs.nestjs.com/controllers)
- [Next.js API Routes](https://nextjs.org/docs/api-routes/introduction)

---

## Conclusion

This API visualization tool aims to revolutionize how developers document and explore their APIs. By combining automatic detection, beautiful visualization, and zero-configuration setup, we can save developers countless hours while improving API documentation quality.

**Timeline**: 14-16 weeks to MVP
**Team Size**: 2-3 developers
**Investment**: Medium complexity, high impact

---

*Document Version*: 1.0  
*Last Updated*: 2026-01-11  
*Author*: rissabekov-wes  
*Status*: Planning Phase
