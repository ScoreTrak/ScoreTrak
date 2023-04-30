import {NavLink} from "react-router-dom";
import {useAuth} from "../contexts/AuthContext";
import {useCompetition} from "../contexts/CompetitionContext";
import {EditCompetitionForm} from "./forms/competitions";


export function CompetitionNav() {
  const {user, teams} = useAuth()
  const {competitionId, competition} = useCompetition()
  return (
    <>
      <div>
        {
          competitionId && competition &&
          <>
            <p>Icon?</p>
            <h2 className={"text-lg font-bold"}>{competition.display_name}</h2>
            <h3 className={"text-md"}>c/{competition.name}</h3>
            {/*<EditCompetitionForm competitionId={competitionId} />*/}
            {/*<p>Do not show anything flashy. Have link to scoreboard and ranks. Possibly a blank screen showing the teams that are competing</p>*/}
            {/*<p>have an asie nav bar showing teams competing and its current rank</p>*/}
          </>
        }
        <nav aria-label={"scoreboard-competitor-nav"}>
          <NavLink to={`/c/${competitionId}/scoreboard`}>Scoreboard</NavLink>
          {/* Table health */}
          <NavLink to={`/c/${competitionId}/ranks`}>Rank</NavLink>
          {/* Bar chart */}
          <NavLink to={`/c/${competitionId}/health`}>Health</NavLink>
        {/*  Health of any team you are a part of*/}
        {/*  this includes recent checks, which services are failing. opportunity to change properties*/}
        {/*  <NavLink to={`/c/${competitionId}/shop`}>Shop</NavLink>*/}
        </nav>
      </div>
    </>
  )
}
