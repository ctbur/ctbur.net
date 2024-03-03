import { LitElement, css, html } from "lit";
import { customElement } from "lit/decorators.js";

import "../components/nav-bar";
import "../components/til-card-list";

@customElement("x-view-tils")
export class Tils extends LitElement {
  render() {
    return html`
      <x-nav-bar currentview="tils"></x-nav-bar>
      <main>
        <div class="til-card-list">
          <x-til-card-list></x-til-card-list>
        </div>
      </main>
    `;
  }

  static styles = css`
    .til-card-list {
      margin: 1rem auto;
      width: 50rem;
      max-width: 100%;
    }
  `;
}

declare global {
  interface HTMLElementTagNameMap {
    "x-view-tils": Tils;
  }
}
