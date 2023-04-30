import {createContext, PropsWithChildren, useContext} from "react";
import {useLocalStorage} from "react-use";
import {useParams} from "react-router-dom";

type TeamContextType = {
  teamId: string | undefined
}

const TeamContext = createContext<TeamContextType>({teamId: undefined})

export const useTeam = () => {
  return useContext(TeamContext)
}

export function TeamProvider({ children }: PropsWithChildren) {
  const {teamId} = useParams()

  return (
    <>
      <TeamContext.Provider value={{teamId}}>
        {children}
      </TeamContext.Provider>
    </>
  )
}
