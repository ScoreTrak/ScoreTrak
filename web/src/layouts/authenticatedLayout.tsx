import { Route } from "@tanstack/router";
import { rootRoute } from "../routes/root";
import { baseLayout } from "./baseLayout";
import {Navigate, Outlet, useNavigate} from "@tanstack/react-router";
import {useAuth} from "../contexts/AuthContext";
import {indexRoute} from "../routes";


export const authenticatedLayout = new Route({
  getParentRoute: () => baseLayout,
  id: "authenticated",
  component: () => {

    // Check if authenticated
    // Send a toast notification about needing to be authenticated to acccess this route
    // send user to the previous route
    const {session} = useAuth()

    if (!session) {
      return <Navigate to={"/"} />
    }

    return (
      <>
        <Outlet />
      </>
    )

  }
})
