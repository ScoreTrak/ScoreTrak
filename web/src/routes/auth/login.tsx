import { Route } from "@tanstack/router"
import { authRoute } from "."
import { useNavigate, useSearch } from "@tanstack/react-router"
import { useCallback, useEffect, useState } from "react"
import {LoginFlow, UiNode, UiNodeInputAttributes, UpdateLoginFlowBody} from "@ory/client"
import {ory, oryError} from "../../lib/auth/ory"

import {z} from "zod"
import {filterNodesByGroups, isUiNodeInputAttributes } from "@ory/integrations/ui";
import {AxiosError} from "axios";
import {useAuth} from "../../contexts/AuthContext";

const loginSearchSchema = z.object({
  flow: z.string().uuid().nullable().optional(),
  aal1: z.string().nullable().optional(),
  aal2: z.string().nullable().optional(),
  return_to: z.string().default("/"),
})

type LoginSearch = z.infer<typeof loginSearchSchema>

export const authLoginRoute = new Route({
  getParentRoute: () => authRoute,
  path: "login",
  validateSearch: loginSearchSchema,
  component: () => {
    const [flow, setFlow] = useState<LoginFlow>()
    const { flow: flowSearchParam, aal1, aal2, return_to } = useSearch({ from: authLoginRoute.id })
    const {refreshSession} = useAuth()


    const navigate = useNavigate({from: authLoginRoute.id })

    const getFlow = useCallback((flowId: string) => {
      return ory.getLoginFlow({ id: flowId })
        .then(({ data: flowData }) => setFlow(flowData))
    }, [])

    const sdkErrorHandler = oryError(getFlow, setFlow, navigate, "/auth/login",  true)

    const createFlow = () => {
      ory.createBrowserLoginFlow({
        refresh: true,
        aal: aal2 ? "aal2" : "aal1",
        // returnTo: return_to
      })
        .then(({ data: flowData }) => {
          setFlow(flowData)
        })
        .catch(sdkErrorHandler)
    }

    const submitFlow = (body: UpdateLoginFlowBody) => {
      // something unexpected went wrong and the flow was not set
      if (!flow) return navigate( { to: "/auth/login" })

      ory
        .updateLoginFlow({
          flow: flow.id,
          updateLoginFlowBody: body,
        })
        .then(() => {
          refreshSession()
          navigate({ to: "/" })
        })
        .catch(sdkErrorHandler)
    }

    useEffect(() => {
      const flowId = flowSearchParam
      if (flowId) {
        getFlow(flowId).catch(createFlow)
        return
      }
      createFlow()
    }, [])

    const submit = (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault()

      const form = event.currentTarget
      const formData = new FormData(form)

      // map the entire form data to JSON for the request body
      let body = Object.fromEntries(formData) as unknown as UpdateLoginFlowBody

      // We need the method specified from the name and value of the submit button.
      // when multiple submit buttons are present, the clicked one's value is used.
      if ("submitter" in event.nativeEvent) {
        const method = (
          event.nativeEvent as unknown as { submitter: HTMLInputElement }
        ).submitter
        body = {
          ...body,
          ...{ [method.name]: method.value },
        }
      }

      submitFlow(body)
    }

    const mapUINode = (node: UiNode, key: number) => {
      // other node types are also supported
      // if (isUiNodeTextAttributes(node.attributes)) {
      // if (isUiNodeImageAttributes(node.attributes)) {
      // if (isUiNodeAnchorAttributes(node.attributes)) {
      if (isUiNodeInputAttributes(node.attributes)) {
        const attrs = node.attributes as UiNodeInputAttributes
        const nodeType = attrs.type
        const messages = node.messages

        switch (nodeType) {
          case "button":
          case "submit":
            return (
              <button
                type={attrs.type as "submit" | "reset" | "button" | undefined}
                name={attrs.name}
                value={attrs.value}
                key={key}
              >
                Login
              </button>
            )
          default:
            return (
              <input
                name={attrs.name}
                type={attrs.type}
                autoComplete={
                  attrs.autocomplete || attrs.name === "identifier"
                    ? "username"
                    : ""
                }
                defaultValue={attrs.value}
                required={attrs.required}
                disabled={attrs.disabled}
                key={key}
              />
            )
        }
      }
    }

    return flow ? (
      <>
        <form action={flow.ui.action} method={flow.ui.method} onSubmit={submit}>
          {filterNodesByGroups({
            nodes: flow.ui.nodes,
            // we will also map default fields here such as csrf_token
            // this only maps the `password` method
            // other methods can also be mapped such as `oidc` or `webauthn`
            groups: ["default", "password"],
          }).map((node, idx) => mapUINode(node, idx))}
        </form>
        {flow.ui.messages?.map((message, idx, arr) => {
          return (
            <>
              <div key={idx}>
                <p>{message.type}</p>
                <p>{message.text}</p>
              </div>
            </>
          )
        })}
      </>
    ) : (
      <div>Loading...</div>
    )
  }
})
