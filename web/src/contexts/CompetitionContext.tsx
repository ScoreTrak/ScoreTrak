import {Competition} from "../lib/scoretrak-queries";
import {createContext, PropsWithChildren, useContext} from "react";


type CompetitionContextType = {
  competitionId: string | undefined,
  competition: Competition | undefined,
  setCompetition: (id: string) => void
}

const CompetitionContext = createContext<CompetitionContextType>({competitionId: undefined, competition: undefined, setCompetition: () => {}})

export const useCompetition = () => {
  return useContext(CompetitionContext)
}

export default function CompetitionProvider({ children }: PropsWithChildren) {

}
