import {Outlet} from "react-router-dom";
import {CompetitionNav} from "../../components/CompetitionNav";

export function CompetitionPortalPage() {
  return (
    <>
      <CompetitionNav />
      <div className={"container mx-auto"}>
        <div></div>
        <Outlet />
        <div></div>
      </div>
    </>
  )
}
