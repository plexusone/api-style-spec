import { LitElement, html, css, nothing } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import type { LintResult, Violation } from '../types';

@customElement('lint-results')
export class LintResults extends LitElement {
  static override styles = css`
    :host {
      display: block;
    }

    .results-container {
      background: var(--color-bg-secondary);
      border: 1px solid var(--color-border);
      border-radius: var(--radius-lg);
      overflow: hidden;
    }

    .results-header {
      padding: var(--spacing-md);
      border-bottom: 1px solid var(--color-border);
      display: flex;
      align-items: center;
      justify-content: space-between;
    }

    .results-title {
      font-weight: 600;
      font-size: 1rem;
    }

    .status-badge {
      padding: var(--spacing-xs) var(--spacing-sm);
      border-radius: var(--radius-sm);
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
    }

    .status-pass {
      background: rgba(16, 185, 129, 0.1);
      color: var(--color-success);
    }

    .status-fail {
      background: rgba(239, 68, 68, 0.1);
      color: var(--color-error);
    }

    .results-body {
      padding: var(--spacing-md);
    }

    .empty-state {
      text-align: center;
      padding: var(--spacing-xl);
      color: var(--color-text-secondary);
    }

    .empty-icon {
      width: 48px;
      height: 48px;
      margin: 0 auto var(--spacing-md);
      color: var(--color-text-muted);
    }

    .loading {
      display: flex;
      align-items: center;
      justify-content: center;
      padding: var(--spacing-xl);
      gap: var(--spacing-sm);
      color: var(--color-text-secondary);
    }

    .spinner {
      width: 20px;
      height: 20px;
      border: 2px solid var(--color-border);
      border-top-color: var(--color-primary);
      border-radius: 50%;
      animation: spin 0.8s linear infinite;
    }

    @keyframes spin {
      to {
        transform: rotate(360deg);
      }
    }

    .error-message {
      padding: var(--spacing-md);
      background: rgba(239, 68, 68, 0.1);
      border-radius: var(--radius-md);
      color: var(--color-error);
      display: flex;
      align-items: start;
      gap: var(--spacing-sm);
    }

    .summary {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      gap: var(--spacing-sm);
      margin-bottom: var(--spacing-md);
    }

    .summary-item {
      text-align: center;
      padding: var(--spacing-sm);
      background: var(--color-bg-tertiary);
      border-radius: var(--radius-md);
    }

    .summary-value {
      font-size: 1.5rem;
      font-weight: 700;
    }

    .summary-label {
      font-size: 0.75rem;
      color: var(--color-text-secondary);
    }

    .summary-errors .summary-value {
      color: var(--color-error);
    }

    .summary-warnings .summary-value {
      color: var(--color-warning);
    }

    .summary-infos .summary-value {
      color: var(--color-info);
    }

    .violations-list {
      display: flex;
      flex-direction: column;
      gap: var(--spacing-sm);
    }

    .violation {
      padding: var(--spacing-sm) var(--spacing-md);
      background: var(--color-bg);
      border-radius: var(--radius-md);
      border-left: 3px solid;
    }

    .violation-error {
      border-left-color: var(--color-error);
    }

    .violation-warn {
      border-left-color: var(--color-warning);
    }

    .violation-info {
      border-left-color: var(--color-info);
    }

    .violation-hint {
      border-left-color: var(--color-text-muted);
    }

    .violation-header {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      margin-bottom: var(--spacing-xs);
    }

    .violation-rule {
      font-family: var(--font-mono);
      font-size: 0.75rem;
      font-weight: 600;
      padding: 2px 6px;
      background: var(--color-bg-tertiary);
      border-radius: var(--radius-sm);
    }

    .violation-message {
      font-size: 0.875rem;
    }

    .violation-path {
      font-family: var(--font-mono);
      font-size: 0.75rem;
      color: var(--color-text-secondary);
      margin-top: var(--spacing-xs);
    }

    .success-message {
      text-align: center;
      padding: var(--spacing-lg);
      color: var(--color-success);
    }

    .success-icon {
      width: 48px;
      height: 48px;
      margin: 0 auto var(--spacing-sm);
    }
  `;

  @property({ type: Object })
  result: LintResult | null = null;

  @property({ type: Boolean })
  loading: boolean = false;

  @property({ type: String })
  error: string | null = null;

  override render() {
    return html`
      <div class="results-container">
        <div class="results-header">
          <h3 class="results-title">Lint Results</h3>
          ${this.result
            ? html`
                <span class="status-badge ${this.result.status === 'pass' ? 'status-pass' : 'status-fail'}">
                  ${this.result.status}
                </span>
              `
            : nothing}
        </div>
        <div class="results-body">
          ${this.renderContent()}
        </div>
      </div>
    `;
  }

  private renderContent() {
    if (this.loading) {
      return html`
        <div class="loading">
          <div class="spinner"></div>
          <span>Linting specification...</span>
        </div>
      `;
    }

    if (this.error) {
      return html`
        <div class="error-message">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="10"></circle>
            <line x1="12" y1="8" x2="12" y2="12"></line>
            <line x1="12" y1="16" x2="12.01" y2="16"></line>
          </svg>
          <span>${this.error}</span>
        </div>
      `;
    }

    if (!this.result) {
      return html`
        <div class="empty-state">
          <svg class="empty-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <path d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          <p>Enter an OpenAPI specification and click "Lint" to see results</p>
        </div>
      `;
    }

    return html`
      ${this.renderSummary()}
      ${this.result.violations.length > 0 ? this.renderViolations() : this.renderSuccess()}
    `;
  }

  private renderSummary() {
    if (!this.result) return nothing;

    const { summary } = this.result;

    return html`
      <div class="summary">
        <div class="summary-item summary-errors">
          <div class="summary-value">${summary.errors}</div>
          <div class="summary-label">Errors</div>
        </div>
        <div class="summary-item summary-warnings">
          <div class="summary-value">${summary.warnings}</div>
          <div class="summary-label">Warnings</div>
        </div>
        <div class="summary-item summary-infos">
          <div class="summary-value">${summary.infos}</div>
          <div class="summary-label">Info</div>
        </div>
        <div class="summary-item">
          <div class="summary-value">${summary.hints}</div>
          <div class="summary-label">Hints</div>
        </div>
      </div>
    `;
  }

  private renderViolations() {
    if (!this.result) return nothing;

    return html`
      <div class="violations-list">
        ${this.result.violations.map((violation) => this.renderViolation(violation))}
      </div>
    `;
  }

  private renderViolation(violation: Violation) {
    return html`
      <div class="violation violation-${violation.severity}">
        <div class="violation-header">
          <span class="violation-rule">${violation.ruleId}</span>
        </div>
        <div class="violation-message">${violation.message}</div>
        <div class="violation-path">${violation.path}</div>
      </div>
    `;
  }

  private renderSuccess() {
    return html`
      <div class="success-message">
        <svg class="success-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
          <polyline points="22 4 12 14.01 9 11.01"></polyline>
        </svg>
        <p>No violations found!</p>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'lint-results': LintResults;
  }
}
