import { Route } from "@tanstack/router";
import { rootRoute } from "./root";
import { indexRoute } from ".";
import {baseLayout} from "../layouts/baseLayout";


export const notFoundRoute = new Route({
  getParentRoute: () => baseLayout,
  path: "*",
  component: () => {
    return (
      <>
        <p>Not Found</p>
      </>
    )
  }
})
