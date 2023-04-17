import { Route } from "@tanstack/router";
import { rootRoute } from "../root";
import { baseLayout } from "./baseLayout";
import {Navigate, Outlet, useNavigate} from "@tanstack/react-router";
import {useAuth} from "../../contexts/AuthContext";
import {indexRoute} from "..";
import Nav from "../../components/Nav";


export const unauthenticatedLayout = new Route({
  getParentRoute: () => baseLayout,
  id: "unauthenticated",
  component: () => {

    // Check if authenticated
    // Send a toast notification about needing to be authenticated to access this route
    // send user to the previous route
    const {session} = useAuth()

    if (session) {
      return <Navigate to={"/"} />
    }

    return (
      <>
        <Nav />
        <Outlet />
      </>
    )

  }
})
