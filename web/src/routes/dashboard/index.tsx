import {Route} from "@tanstack/router";
import {baseLayout} from "../layouts/baseLayout";
import {scoretrak} from "../../lib/queries";
import {authenticatedLayout} from "../layouts/authenticatedLayout";
import {queryClient} from "../../App";
import {queryKeys} from "../../lib/scoretrak-queries";
import {CreateCompetitionForm} from "../../components/forms/competitions";
import {Link} from "@tanstack/react-router";


export const dashboardRoute = new Route({
  getParentRoute: () => authenticatedLayout,
  path: "dashboard"
})


export const dashboardIndexRoute = new Route({
  getParentRoute: () => dashboardRoute,
  path: '/',
  component: () => {
    const { data, isLoading, isError } = scoretrak.queries.useListCompetition()

    return (
      <>
        <h1>Dashboard</h1>
        { data &&
          data?.map((competition) => {
            return (
              <Link key={competition.id} to={"/dashboard/$competitionId"} params={{competitionId: competition.id.toString()}}>{competition.display_name}</Link>
            )
          })
        }
        <CreateCompetitionForm />
      </>
    )
  }
})
