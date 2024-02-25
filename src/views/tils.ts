import { LitElement, html } from "lit";
import { customElement } from "lit/decorators.js";

import "../components/nav-bar";
import "../components/til-card";

@customElement("x-view-tils")
export class Home extends LitElement {
  render() {
    return html`
      <x-nav-bar currentview="tils"></x-nav-bar>
      <main>
        <x-til-card tilid="7CHb54awH0gjxM69qIrUG3"></x-til-card>
      </main>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "x-view-tils": Home;
  }
}
