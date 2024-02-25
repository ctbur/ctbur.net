import { LitElement, html } from "lit";
import { customElement } from "lit/decorators.js";

import "../components/nav-bar";
import "../components/til-card";

@customElement("x-view-tils")
export class Home extends LitElement {
  render() {
    return html`
      <x-nav-bar current-view="tils"></x-nav-bar>
      <main>
        <x-til-card></x-til-card>
      </main>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "x-view-tils": Home;
  }
}
