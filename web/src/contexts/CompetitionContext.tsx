import {createContext, PropsWithChildren, useContext, useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {CompetitionRead} from "../lib/scoretrak-queries";
import {scoretrak} from "../lib/queries";

type CompetitionContextType = {
  competitionId: string | undefined
  competition: CompetitionRead | undefined
}

const CompetitionContext = createContext<CompetitionContextType>({competitionId: undefined, competition: undefined})

export const useCompetition = () => {
  return useContext(CompetitionContext)
}

export function CompetitionProvider({ children }: PropsWithChildren) {
  const {competitionId} = useParams()
  const {data: competition} = scoretrak.queries.useReadCompetition(competitionId ?? "")

  return (
    <>
      <CompetitionContext.Provider value={{competitionId, competition}}>
        {children}
      </CompetitionContext.Provider>
    </>
  )
}
