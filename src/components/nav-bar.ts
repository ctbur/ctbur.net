import { LitElement, html } from "lit";
import { customElement, property } from "lit/decorators.js";

export type View = "home" | "tils";

@customElement("x-nav-bar")
export class NavBar extends LitElement {
  @property() currentView: View = "home";

  render() {
    return html`
      <nav>
        <a href="/">
          <sl-tab ?active=${this.currentView === "home"}>Home</sl-tab>
        </a>
        <a href="/tils">
          <sl-tab ?active=${this.currentView === "tils"}>TILs</sl-tab>
        </a>
      </nav>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "x-nav-bar": NavBar;
  }
}
