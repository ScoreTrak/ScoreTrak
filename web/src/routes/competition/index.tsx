import { Route } from "@tanstack/router";
import {authenticatedLayout} from "../layouts/authenticatedLayout";
import {CreateCompetitionForm} from "../../components/forms/competitions";


export const competitionRoute = new Route({
  getParentRoute: () => authenticatedLayout,
  path: "competitions"
})

export const competitionIndexRoute = new Route({
  getParentRoute: () => competitionRoute,
  path: "/",
  component: () => {
    return (
      <>
        <CreateCompetitionForm />
      </>
    )
  }
})
