import * as contentful from "contentful";
import { remark } from "remark";
import remarkHtml from "remark-html";

console.log(contentful);
const client = contentful.createClient({
  space: "6d4vhsbxh0yj",
  environment: "master",
  accessToken: "6S26NCjEdWLehauOCnlaijvcbjfd3gashEfh4Bnwjpc",
});

const entry = await client.getEntry("7CHb54awH0gjxM69qIrUG3");

const renderedContent = await remark()
  .use(remarkHtml)
  .process(entry.fields.content as string);
console.log(renderedContent.toString());
