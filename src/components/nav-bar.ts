import { LitElement, css, html } from "lit";
import { customElement, property } from "lit/decorators.js";

export type View = "home" | "tils";

@customElement("x-nav-bar")
export class NavBar extends LitElement {
  @property() currentView: View = "home";

  render() {
    return html`
      <nav>
        <ul>
          <li>
            <a class=${this.currentView === "home" ? "active" : ""} href="/"
              >Home</a
            >
          </li>
          <li>
            <a class=${this.currentView === "tils" ? "active" : ""} href="/tils"
              >TILs</a
            >
          </li>
        </ul>
      </nav>
    `;
  }

  static styles = css`
    ul {
      list-style-type: none;
      margin: 0;
      padding: 0;
      overflow: hidden;
      background-color: #333;
    }
    li {
      float: left;
      border-right: 1px solid #bbb;
    }
    li a {
      display: block;
      color: white;
      text-align: center;
      padding: 14px 16px;
      text-decoration: none;
    }
    li a:hover:not(.active) {
      background-color: #111;
    }
    .active {
      background-color: blue;
    }
  `;
}

declare global {
  interface HTMLElementTagNameMap {
    "x-nav-bar": NavBar;
  }
}
