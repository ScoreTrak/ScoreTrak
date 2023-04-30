import {NavLink, Outlet} from "react-router-dom";

export function DashboardPage() {
  return (
    <>
      <nav>
        <NavLink to={"/dashboard/competitions"}>Competitions</NavLink>
        <NavLink to={"/dashboard/services"}>Services</NavLink>
        <NavLink to={"/dashboard/teams"}>Teams</NavLink>
        <NavLink to={"/dashboard/hosts"}>Hosts</NavLink>
        <NavLink to={"/dashboard/hostservices"}>Host Services</NavLink>
        <NavLink to={"/dashboard/properties"}>Properties</NavLink>
      </nav>
      <Outlet/>
    </>
  )
}
