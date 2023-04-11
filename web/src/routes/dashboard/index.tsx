import {Route} from "@tanstack/router";
import {baseLayout} from "../../layouts/baseLayout";
import {scoretrak} from "../../lib/queries";


export const dashboardRoute = new Route({
  getParentRoute: () => baseLayout,
  path: "dashboard",
  component: () => {
    const { data, isLoading, isError } = scoretrak.queries.useListCompetition()

    return (
      <>
        { data &&
          data?.map((competition) => {
            return (
              <>
                <p>{competition.id}</p>
              </>
            )
          })
        }
      </>
    )
  }
})
