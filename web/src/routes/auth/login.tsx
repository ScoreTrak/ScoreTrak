import { Route } from "@tanstack/router"
import { authRoute } from "."
import { useNavigate, useSearch } from "@tanstack/react-router"
import { useCallback, useEffect, useState } from "react"
import { LoginFlow, UpdateLoginFlowBody } from "@ory/client"
import { ory } from "../../lib/auth/ory"
import { UserAuthCard } from "@ory/elements"

import {z} from "zod"

const loginSearchSchema = z.object({
  flow: z.string().uuid().nullable().optional(),
  aal1: z.string().nullable().optional(),
  aal2: z.string().nullable().optional(),
})

type LoginSearch = z.infer<typeof loginSearchSchema>

export const authLoginRoute = new Route({
  getParentRoute: () => authRoute,
  path: "login",
  validateSearch: loginSearchSchema,
  component: () => {
    const [flow, setFlow] = useState<LoginFlow | null>(null)
    const { flow: flowSearchParam, aal1, aal2 } = useSearch({ from: authLoginRoute.id })

    const navigate = useNavigate({from: authLoginRoute.id })

    const getFlow = useCallback((flowId: string) => {
      return ory.getLoginFlow({ id: flowId })
      .then(({ data: flowData }) => setFlow(flowData))
      .catch((err) => {console.error(err)})
    }, [])
    

    const createFlow = () => {
      ory.createBrowserLoginFlow({ refresh: true, aal: aal2 ? "aal2" : "aal1"})
      .then(({ data: flowData }) => {
        setFlow(flowData)
      })
      .catch((err) => {console.error(err)})
    }

    const submitFlow = (body: UpdateLoginFlowBody) => {
      if (!flow) return navigate({to: '/'})

      ory.updateLoginFlow({ flow: flow.id, updateLoginFlowBody: body})
      .then(() => {
        navigate({to: "/"})
      })
      .catch((err) => {console.error(err)})
    }

    useEffect(() => {
      if (flowSearchParam) {
        getFlow(flowSearchParam).catch(createFlow)
        return
      }

      createFlow()
    }, [])


    return flow ? (
      <>
        <UserAuthCard
          title="Login"
          flowType="login"
          flow={flow}
          additionalProps={{
            signupURL: "/auth/register"
          }}
          onSubmit={({ body }) => submitFlow(body as UpdateLoginFlowBody)}
        />
      </>
    ) : (
      <div>Loading...</div>
    )
  }
})
