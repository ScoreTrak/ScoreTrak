import { Route } from "@tanstack/router";
import { rootRoute } from "../root";
import { baseLayout } from "./baseLayout";
import {Navigate, Outlet, useNavigate} from "@tanstack/react-router";
import {useAuth} from "../../contexts/AuthContext";
import {indexRoute} from "..";
import Nav from "../../components/Nav";
import {MainLayout} from "../../layouts/mainLayout";


export const authenticatedLayout = new Route({
  getParentRoute: () => baseLayout,
  id: "authenticated",
  beforeLoad: () => {

  },
  component: () => {

    // Check if authenticated
    // Send a toast notification about needing to be authenticated to access this route
    // send user to the previous route
    const {isLoading, session} = useAuth()

    if (!session) {
      return <Navigate to={"/"} />
    }

    return (
      <>
        <MainLayout>
          <Outlet />
        </MainLayout>
      </>
    )

  }
})
