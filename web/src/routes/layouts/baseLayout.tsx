import { Route } from "@tanstack/router";
import { rootRoute } from "../root";
import {Outlet} from "@tanstack/react-router";
import Nav from "../../components/Nav";


export const baseLayout = new Route({
  getParentRoute: () => rootRoute,
  id: "base",
  component: () => {
    return (
      <>
        <Nav />
        <Outlet />
      </>
    )
  }
})
