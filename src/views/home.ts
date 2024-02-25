import { LitElement, html } from "lit";
import { customElement } from "lit/decorators.js";

import "../components/nav-bar";

@customElement("x-view-home")
export class Home extends LitElement {
  render() {
    return html`
      <x-nav-bar currentview="home"></x-nav-bar>
      <main>
        <p>this is the home page</p>
      </main>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "x-view-home": Home;
  }
}
