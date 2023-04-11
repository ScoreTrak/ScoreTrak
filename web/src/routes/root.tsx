import { Outlet } from "@tanstack/react-router";
import { RootRoute } from "@tanstack/router";
import Nav from "../components/Nav";


export const rootRoute = new RootRoute({
  component: () => {
    return (
      <>
        <Outlet />
      </>
    )
  }
})
