import * as contentful from "contentful";
import { remark } from "remark";
import remarkHtml from "remark-html";
import { LitElement, html } from "lit";
import { customElement, property } from "lit/decorators.js";
import { unsafeHTML } from "lit/directives/unsafe-html.js";
import { Task } from "@lit/task";

const client = contentful.createClient({
  space: "6d4vhsbxh0yj",
  environment: "master",
  accessToken: "6S26NCjEdWLehauOCnlaijvcbjfd3gashEfh4Bnwjpc",
});

const fetchTil = async (entryId: string): Promise<Til> => {
  const entry = await client.getEntry(entryId);
  const renderedContent = await remark()
    .use(remarkHtml)
    .process(entry.fields.content as string);

  return {
    title: entry.fields.title as string,
    content: renderedContent.toString(),
  };
};

interface Til {
  title: string;
  content: string;
}

@customElement("x-til-card")
export class TilCard extends LitElement {
  @property() tilId?: string;

  private _fetchTilTask = new Task(this, {
    task: async ([tilId]) => {
      if (tilId === undefined) {
        return Promise.reject("No TIL ID provided");
      }
      return await fetchTil(tilId);
    },
    args: () => [this.tilId],
  });

  render() {
    return this._fetchTilTask.render({
      pending: () => html`<sl-card><p>Loading...</p></sl-card>`,
      complete: (til) => html`
        <sl-card class="card-header">
          <div slot="header">${til.title}</div>

          <p>${unsafeHTML(til.content)}</p>

          <div slot="footer">
            <sl-tag variant="neutral">tag TODO</sl-tag>
          </div>
        </sl-card>
      `,
      error: (e) => html`<sl-card><p>Error: ${e}</p></sl-card>`,
    });
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "x-til-card": TilCard;
  }
}
