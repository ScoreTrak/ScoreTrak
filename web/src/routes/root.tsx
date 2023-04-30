import {Link, NavLink, Outlet} from "react-router-dom";
import {queryClient} from "../lib/query-client";
import {QueryClientProvider} from "@tanstack/react-query";
import AuthProvider from "../contexts/AuthContext";
import Nav from "../components/Nav";

export function RootPage() {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <Nav />
          <Outlet />
        </AuthProvider>
      </QueryClientProvider>
    </>
  )
}
