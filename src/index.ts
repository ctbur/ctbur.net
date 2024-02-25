import * as contentful from "contentful";
import { remark } from "remark";
import remarkHtml from "remark-html";
import "@shoelace-style/shoelace/dist/themes/light.css";
import "@shoelace-style/shoelace/dist/components/card/card.js";
import "@shoelace-style/shoelace/dist/components/breadcrumb/breadcrumb.js";
import { Router } from "@vaadin/router";

import "./views/home";
import "./views/tils";

const client = contentful.createClient({
  space: "6d4vhsbxh0yj",
  environment: "master",
  accessToken: "6S26NCjEdWLehauOCnlaijvcbjfd3gashEfh4Bnwjpc",
});

const fetchContent = async (entryId: string) => {
  try {
    const entry = await client.getEntry(entryId);
    const renderedContent = await remark()
      .use(remarkHtml)
      .process(entry.fields.content as string);
    console.log(renderedContent.toString());
  } catch (error) {
    console.error("Error processing content:", error);
  }
};

fetchContent("7CHb54awH0gjxM69qIrUG3");

const router = new Router(document.getElementById("app"));
router.setRoutes([
  { path: "/", component: "x-view-home" },
  { path: "/tils", component: "x-view-tils" },
]);
