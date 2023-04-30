import { Identity, Session } from "@ory/client";
import { createContext, useContext, useEffect, useState } from "react";
import { ory } from "../lib/auth/ory";
import { PropsWithChildren } from "react"


type AuthContextType = {
  isLoading: boolean,
  session: Session | undefined,
  user: Identity | undefined,
  logoutUrl: string | undefined,
  refreshSession: () => void,
  isAdmin: boolean,
  competitions: string[],
  teams: string[],
}

const AuthContext = createContext<AuthContextType>({isLoading: true, user: undefined, session: undefined, logoutUrl: "", competitions: [], teams: [], refreshSession: () => {}, isAdmin: false})

export const useAuth = () => {
  return useContext(AuthContext)
}

export default function AuthProvider({ children }: PropsWithChildren) {
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const [session, setSession] = useState<Session | undefined>()
  const [identity, setIdentity] = useState<Identity | undefined>()
  const [competitions, setCompetitions] = useState<string[]>([])
  const [teams, setTeams] = useState<string[]>([])
  const [logoutUrl, setLogoutUrl] = useState<string | undefined>()
  const [isAdmin, setIsAdmin] = useState<boolean>(false)

  const setUserData = (sessionData: Session) => {
    setSession(sessionData)
    setIdentity(sessionData.identity)
    // @ts-ignore
    setCompetitions(sessionData.identity.metadata_public?.competitions ?? [])
    // @ts-ignore
    setTeams(sessionData.identity.metadata_public?.teams ?? [])
  }

  const clearUserData = () => {
    setSession(undefined)
    setIdentity(undefined)
    setCompetitions([])
    setTeams([])
  }

  const getSessionAndLogoutUrl = () => {
    setIsLoading(true)
    ory
      .toSession()
      .then(({ data: sessionData }) => {
        setUserData(sessionData)
        ory.createBrowserLogoutFlow()
          .then(({ data: logoutFlowData }) => {
            setLogoutUrl(logoutFlowData.logout_url)
          })
          .catch((err) => {
          })
      })
      .catch((err) => {
        clearUserData()
      })
      .finally(() => setIsLoading(false))
  }

  const refresh = () => {
    getSessionAndLogoutUrl()
  }

  useEffect(() => {
    refresh()
  }, [])

  return (
    <>
      <AuthContext.Provider value={{session, user: identity, logoutUrl, refreshSession: refresh, isLoading, isAdmin, competitions, teams}}>
        {children}
      </AuthContext.Provider>
    </>
  )
}
