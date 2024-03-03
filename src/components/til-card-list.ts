import * as contentful from "contentful";
import { remark } from "remark";
import remarkHtml from "remark-html";
import { LitElement, css, html } from "lit";
import { customElement } from "lit/decorators.js";
import { unsafeHTML } from "lit/directives/unsafe-html.js";
import { Task } from "@lit/task";

interface Til {
  title: string;
  content: string;
}

const client = contentful.createClient({
  space: "6d4vhsbxh0yj",
  environment: "master",
  accessToken: "6S26NCjEdWLehauOCnlaijvcbjfd3gashEfh4Bnwjpc",
});

const fetchTils = async (): Promise<Array<Til>> => {
  const entries = await client.getEntries();
  console.log(entries);
  const tils = entries.items.map(async (entry) => {
    const renderedContent = await remark()
      .use(remarkHtml)
      .process(entry.fields.content as string);

    return {
      title: entry.fields.title as string,
      content: renderedContent.toString(),
    };
  });
  return Promise.all(tils);
};

@customElement("x-til-card-list")
export class TilCardList extends LitElement {
  private _fetchTilsTask = new Task(this, { task: fetchTils, args: () => [] });

  private renderTil = (til: Til) => {
    return html`
      <div class="til-card">
        <sl-card class="card-header til-card-card">
          <div slot="header">${til.title}</div>

          <p>${unsafeHTML(til.content)}</p>

          <div slot="footer">
            <sl-tag variant="neutral">tag TODO</sl-tag>
          </div>
        </sl-card>
      </div>
    `;
  };

  render() {
    return this._fetchTilsTask.render({
      pending: () => html`<p>Loading...</p>`,
      complete: (tils) =>
        html`<div class="til-card-list">${tils.map(this.renderTil)}</div>`,
      error: (e) => html`<p>Error loading TILs: ${e}</p>`,
    });
  }

  static styles = css`
    .til-card-list {
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }
    .til-card {
      width: 100%;
    }
    .til-card-card {
      width: 100%;
    }
  `;
}

declare global {
  interface HTMLElementTagNameMap {
    "x-til-card-list": TilCardList;
  }
}
