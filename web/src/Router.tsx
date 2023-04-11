import { Router } from "@tanstack/router";
import { rootRoute } from "./routes/root";
import { indexRoute } from "./routes";
import { authRoute } from "./routes/auth";
import { authLoginRoute } from "./routes/auth/login";
import { notFoundRoute } from "./routes/notfound";
import { competitionIndexRoute } from "./routes/competition";
import { authRegisterRoute } from "./routes/auth/register";
import {baseLayout} from "./layouts/baseLayout";
import {authenticatedLayout} from "./layouts/authenticatedLayout";
import {dashboardRoute} from "./routes/dashboard";


const routeTree = rootRoute.addChildren([
  baseLayout.addChildren([
    indexRoute,
    notFoundRoute,
    dashboardRoute,
    authRoute.addChildren([authLoginRoute, authRegisterRoute]),
    authenticatedLayout.addChildren([
      competitionIndexRoute.addChildren([]),
    ]),
  ]),
])

export const router = new Router({
  routeTree,
  defaultPreload: 'intent'
})


declare module '@tanstack/router' {
  interface Register {
    router: typeof router
  }
}
