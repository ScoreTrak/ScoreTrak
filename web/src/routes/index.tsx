import { Route } from "@tanstack/router"
import { rootRoute } from "./root"
import {baseLayout} from "../layouts/baseLayout";

export const indexRoute = new Route({
  getParentRoute: () => baseLayout,
  path: '/',
  component: () => {
    return (
      <>
        <h1>Index Route</h1>
      </>
    )
  }
})
