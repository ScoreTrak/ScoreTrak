import { useCallback, useEffect, useState } from "react"
import {LoginFlow, UiNode, UiNodeInputAttributes, UpdateLoginFlowBody} from "@ory/client"
import {ory, oryError} from "@lib/auth/ory"

import {filterNodesByGroups, isUiNodeInputAttributes } from "@ory/integrations/ui";
import {useAuth} from "../../contexts/AuthContext";
import {useNavigate, useParams, useSearchParams} from "react-router-dom";

type LoginVariables = {
  flowId: string
  updateLoginFlowBody: UpdateLoginFlowBody
}

export function AuthLoginPage() {
  const [flow, setFlow] = useState<LoginFlow>()
  const {refreshSession} = useAuth()
  // const { mutate: login } = useLogin<LoginVariables>()

  const navigate = useNavigate()

  const createFlow = () => {
    ory.createBrowserLoginFlow({
      refresh: true,
      aal: "aal1",
      // returnTo: return_to
    })
      .then(({ data: flowData }) => {
        setFlow(flowData)
      })
      // .catch(sdkErrorHandler)
  }

  const submitFlow = (body: UpdateLoginFlowBody) => {
    // something unexpected went wrong and the flow was not set
    if (!flow) return navigate("/auth/login")

    // login({flowId: flow.id, updateLoginFlowBody: body})
    ory
      .updateLoginFlow({
        flow: flow.id,
        updateLoginFlowBody: body,
      })
      .then(() => {
        refreshSession()
        navigate("/")
      })
    //   .catch(sdkErrorHandler)

  }

  useEffect(() => {
    createFlow()
  }, [])

  const submit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    const form = event.currentTarget
    const formData = new FormData(form)
    console.log(formData)

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
            <>
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
              {node.messages.map((message, idx) => {
                return (
                  <p key={idx}>{message.text}</p>
                )
              })}
          </>
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
