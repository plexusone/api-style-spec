import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('spec-editor')
export class SpecEditor extends LitElement {
  static override styles = css`
    :host {
      display: block;
    }

    .editor-container {
      display: flex;
      flex-direction: column;
      gap: var(--spacing-md);
    }

    .editor-wrapper {
      position: relative;
      border: 1px solid var(--color-border);
      border-radius: var(--radius-md);
      overflow: hidden;
    }

    textarea {
      width: 100%;
      min-height: 500px;
      padding: var(--spacing-md);
      font-family: var(--font-mono);
      font-size: 0.875rem;
      line-height: 1.6;
      background: var(--color-bg-secondary);
      color: var(--color-text);
      border: none;
      outline: none;
      resize: vertical;
    }

    textarea::placeholder {
      color: var(--color-text-muted);
    }

    .actions {
      display: flex;
      gap: var(--spacing-sm);
      justify-content: flex-end;
    }

    button {
      display: inline-flex;
      align-items: center;
      gap: var(--spacing-xs);
      padding: var(--spacing-sm) var(--spacing-md);
      font-family: var(--font-sans);
      font-size: 0.875rem;
      font-weight: 500;
      border-radius: var(--radius-md);
      cursor: pointer;
      transition: all 0.15s;
    }

    .btn-primary {
      background: var(--color-primary);
      color: white;
      border: none;
    }

    .btn-primary:hover {
      background: var(--color-primary-dark);
    }

    .btn-secondary {
      background: transparent;
      color: var(--color-text-secondary);
      border: 1px solid var(--color-border);
    }

    .btn-secondary:hover {
      background: var(--color-bg-tertiary);
      color: var(--color-text);
    }

    .example-specs {
      display: flex;
      gap: var(--spacing-sm);
      flex-wrap: wrap;
    }

    .example-btn {
      padding: var(--spacing-xs) var(--spacing-sm);
      font-size: 0.75rem;
      background: var(--color-bg-tertiary);
      border: 1px solid var(--color-border);
      border-radius: var(--radius-sm);
      color: var(--color-text-secondary);
      cursor: pointer;
      transition: all 0.15s;
    }

    .example-btn:hover {
      background: var(--color-primary);
      color: white;
      border-color: var(--color-primary);
    }
  `;

  @property({ type: String })
  value: string = '';

  private readonly sampleSpec = `openapi: "3.1.0"
info:
  title: Sample API
  description: A sample API for testing
  version: "1.0.0"
servers:
  - url: https://api.example.com/v1
paths:
  /users:
    get:
      summary: List users
      operationId: listUsers
      responses:
        "200":
          description: Successful response
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/User"
components:
  schemas:
    User:
      type: object
      properties:
        id:
          type: integer
        name:
          type: string
        email:
          type: string
          format: email
`;

  override render() {
    return html`
      <div class="editor-container">
        <div class="example-specs">
          <span style="color: var(--color-text-secondary); font-size: 0.75rem;">Load example:</span>
          <button class="example-btn" @click=${this.loadSample}>Basic API</button>
          <button class="example-btn" @click=${this.loadPetstore}>Petstore</button>
        </div>

        <div class="editor-wrapper">
          <textarea
            .value=${this.value}
            @input=${this.handleInput}
            placeholder="Paste your OpenAPI specification here (YAML or JSON)..."
          ></textarea>
        </div>

        <div class="actions">
          <button class="btn-secondary" @click=${this.handleClear}>
            Clear
          </button>
          <button class="btn-primary" @click=${this.handleLint}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
              <polyline points="22 4 12 14.01 9 11.01"></polyline>
            </svg>
            Lint Specification
          </button>
        </div>
      </div>
    `;
  }

  private handleInput(e: Event) {
    const textarea = e.target as HTMLTextAreaElement;
    this.dispatchEvent(
      new CustomEvent('spec-change', {
        detail: textarea.value,
        bubbles: true,
        composed: true,
      })
    );
  }

  private handleLint() {
    this.dispatchEvent(
      new CustomEvent('lint-request', {
        bubbles: true,
        composed: true,
      })
    );
  }

  private handleClear() {
    this.dispatchEvent(
      new CustomEvent('spec-change', {
        detail: '',
        bubbles: true,
        composed: true,
      })
    );
  }

  private loadSample() {
    this.dispatchEvent(
      new CustomEvent('spec-change', {
        detail: this.sampleSpec,
        bubbles: true,
        composed: true,
      })
    );
  }

  private async loadPetstore() {
    // In production, this would fetch the actual petstore spec
    const petstoreSpec = `openapi: "3.1.0"
info:
  title: Swagger Petstore
  description: A sample Pet Store Server
  version: "1.0.0"
servers:
  - url: https://petstore.swagger.io/api/v3
paths:
  /pets:
    get:
      tags:
        - pets
      summary: List all pets
      operationId: listPets
      parameters:
        - name: limit
          in: query
          description: Maximum number of pets
          schema:
            type: integer
            maximum: 100
      responses:
        "200":
          description: A list of pets
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pets"
    post:
      tags:
        - pets
      summary: Create a pet
      operationId: createPet
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Pet"
      responses:
        "201":
          description: Pet created
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        tag:
          type: string
    Pets:
      type: array
      items:
        $ref: "#/components/schemas/Pet"
`;

    this.dispatchEvent(
      new CustomEvent('spec-change', {
        detail: petstoreSpec,
        bubbles: true,
        composed: true,
      })
    );
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'spec-editor': SpecEditor;
  }
}
