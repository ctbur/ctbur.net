import "@shoelace-style/shoelace";
import "@shoelace-style/shoelace/dist/themes/light.css";
import { Router } from "@vaadin/router";

import "./views/home";
import "./views/tils";

const router = new Router(document.getElementById("app"));
router.setRoutes([
  { path: "/", component: "x-view-home" },
  { path: "/tils", component: "x-view-tils" },
]).catch((err) => console.error(err));
