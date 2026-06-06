import { LitElement, html, css } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';
import { api } from '../api';
import type { Profile } from '../types';

interface ProfileOption {
  name: string;
  description: string;
}

@customElement('profile-selector')
export class ProfileSelector extends LitElement {
  static override styles = css`
    :host {
      display: inline-block;
    }

    .selector {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
    }

    label {
      font-size: 0.875rem;
      color: var(--color-text-secondary);
    }

    select {
      padding: var(--spacing-xs) var(--spacing-sm);
      font-family: var(--font-sans);
      font-size: 0.875rem;
      background: var(--color-bg);
      color: var(--color-text);
      border: 1px solid var(--color-border);
      border-radius: var(--radius-md);
      cursor: pointer;
      outline: none;
      transition: border-color 0.15s;
    }

    select:hover {
      border-color: var(--color-primary);
    }

    select:focus {
      border-color: var(--color-primary);
      box-shadow: 0 0 0 2px rgba(79, 70, 229, 0.1);
    }
  `;

  @property({ type: String })
  value: string = 'default';

  @state()
  private profiles: ProfileOption[] = [
    { name: 'default', description: 'Common REST best practices' },
    { name: 'azure', description: 'Microsoft Azure guidelines' },
    { name: 'google', description: 'Google API Design Guide' },
    { name: 'zalando', description: 'Zalando RESTful Guidelines' },
    { name: 'vacuum', description: 'Vacuum built-in rules' },
  ];

  override connectedCallback() {
    super.connectedCallback();
    this.fetchProfiles();
  }

  private async fetchProfiles() {
    try {
      const profiles = await api.getProfiles();
      this.profiles = profiles.map((p: Profile) => ({
        name: p.name,
        description: p.description,
      }));
    } catch {
      // Keep default profiles on error
      console.warn('Failed to fetch profiles, using defaults');
    }
  }

  override render() {
    return html`
      <div class="selector">
        <label for="profile">Profile:</label>
        <select id="profile" .value=${this.value} @change=${this.handleChange}>
          ${this.profiles.map(
            (profile) => html`
              <option value=${profile.name} ?selected=${profile.name === this.value}>
                ${profile.name} - ${profile.description}
              </option>
            `
          )}
        </select>
      </div>
    `;
  }

  private handleChange(e: Event) {
    const select = e.target as HTMLSelectElement;
    this.dispatchEvent(
      new CustomEvent('profile-change', {
        detail: select.value,
        bubbles: true,
        composed: true,
      })
    );
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'profile-selector': ProfileSelector;
  }
}
