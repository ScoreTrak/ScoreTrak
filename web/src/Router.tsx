import { Router } from "@tanstack/router";
import { rootRoute } from "./routes/root";
import { indexRoute } from "./routes";
import { authRoute } from "./routes/auth";
import { authLoginRoute } from "./routes/auth/login";
import { notFoundRoute } from "./routes/notfound";
import {competitionIndexRoute, competitionRoute} from "./routes/competition";
import { authRegisterRoute } from "./routes/auth/register";
import {baseLayout} from "./routes/layouts/baseLayout";
import {authenticatedLayout} from "./routes/layouts/authenticatedLayout";
import {dashboardIndexRoute, dashboardRoute} from "./routes/dashboard";
import {unauthenticatedLayout} from "./routes/layouts/unauthenticatedLayout";
import {competitionDashboardRoute} from "./routes/dashboard/competitionDashboard";


const routeTree = rootRoute.addChildren([
  baseLayout.addChildren([
    indexRoute,
    notFoundRoute,
    unauthenticatedLayout.addChildren([
      authRoute.addChildren([authLoginRoute, authRegisterRoute]),
    ]),
    authenticatedLayout.addChildren([
      dashboardRoute.addChildren([dashboardIndexRoute, competitionDashboardRoute]),
      competitionRoute.addChildren([competitionIndexRoute]),
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
