import "@shoelace-style/shoelace/dist/themes/light.css";
import "@shoelace-style/shoelace/dist/components/card/card.js";
import "@shoelace-style/shoelace/dist/components/breadcrumb/breadcrumb.js";
import { Router } from "@vaadin/router";

import "./views/home";
import "./views/tils";

const router = new Router(document.getElementById("app"));
router.setRoutes([
  { path: "/", component: "x-view-home" },
  { path: "/tils", component: "x-view-tils" },
]);
