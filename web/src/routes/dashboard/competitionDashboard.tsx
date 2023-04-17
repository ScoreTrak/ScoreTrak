import {Route} from "@tanstack/router";
import {dashboardRoute} from "./index";
import {useParams} from "@tanstack/react-router";
import {scoretrak} from "../../lib/queries";
import {CreateCompetitionTeamForm} from "../../components/forms/teams";


export const competitionDashboardRoute = new Route({
  getParentRoute: () => dashboardRoute,
  path: "$competitionId",
  component: () => {
    const { competitionId } = useParams({ from: competitionDashboardRoute.id })
    const {data, isLoading} = scoretrak.queries.useReadCompetition(Number(competitionId))
    const {data: teams} = scoretrak.queries.useListCompetitionTeams(Number(competitionId))

    return (
      <>
        <h1>Competition</h1>
        <p>{competitionId}</p>
        {teams && teams.map((team, idx) => {
          return (
            <p key={idx}>{team.name}</p>
          )
        })}
        <CreateCompetitionTeamForm competitionId={Number(competitionId)} />
      </>
    )
  }
})
