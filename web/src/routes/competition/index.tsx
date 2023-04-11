import { Route } from "@tanstack/router";
import { rootRoute } from "../root";
import {baseLayout} from "../../layouts/baseLayout";
import {authenticatedLayout} from "../../layouts/authenticatedLayout";
import {scoretrak} from "../../lib/queries";


export const competitionIndexRoute = new Route({
  getParentRoute: () => authenticatedLayout,
  path: "competitions",
  component: () => {
    const {mutate, mutateAsync} = scoretrak.mutations.useCreateCompetition()

    return (
      <>
        <h1>Competitions</h1>
      </>
    )
  }
})
