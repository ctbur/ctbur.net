import { LitElement, html } from "lit";
import { customElement } from "lit/decorators.js";

@customElement("x-til-card")
export class TilCard extends LitElement {
  render() {
    return html`
      <sl-card class="card-header">
        <div slot="header">Header Title</div>

        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat.
        </p>

        <div slot="footer">
          <sl-tag variant="neutral">tag</sl-tag>
        </div>
      </sl-card>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "x-til-card": TilCard;
  }
}
