import {ory, oryError} from "@lib/auth/ory"
import {
  RegistrationFlow, UiNode, UiNodeInputAttributes, UpdateLoginFlowBody,
  UpdateRegistrationFlowBody
} from "@ory/client"
import {useState, useEffect} from "react"

import {filterNodesByGroups, isUiNodeInputAttributes} from "@ory/integrations/ui";
import {useAuth} from "../../contexts/AuthContext";
import {useNavigate} from "react-router-dom";

type RegisterVariables = {
  flowId: string
  updateRegistrationFlowBody: UpdateRegistrationFlowBody
}

export function AuthRegisterPage() {
  const [flow, setFlow] = useState<RegistrationFlow>()
  const {refreshSession} = useAuth()

  const navigate = useNavigate()

  const createFlow = () => {
    ory.createBrowserRegistrationFlow()
      .then(({data: flowData}) => {
        setFlow(flowData)
      })
      // .catch(sdkErrorHandler)
  }

  const submitFlow = (body: UpdateRegistrationFlowBody) => {
    // something unexpected went wrong and the flow was not set
    if (!flow) { // @ts-ignore
      return navigate("/auth/register")
    }

    // register({flowId: flow.id, updateRegistrationFlowBody: body})
    ory.updateRegistrationFlow({
      flow: flow.id,
      updateRegistrationFlowBody: body,
    })
      .then(() => {
        refreshSession()
        // @ts-ignore
        navigate("/")
      })
    //   // .catch(sdkErrorHandler)
  }


  useEffect(() => {
    createFlow()
  }, [])

  const submit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    const form = event.currentTarget
    const formData = new FormData(form)

    // map the entire form data to JSON for the request body
    let body = Object.fromEntries(
      formData,
    ) as unknown as UpdateRegistrationFlowBody

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
              {node.meta.label?.text}
            </button>
          )
        default:
          let autoComplete = attrs.name == "traits.username" ? "username" : attrs.autocomplete
          return (
             <div key={key}>
            <input
              name={attrs.name}
              type={attrs.type}
              autoComplete={
                autoComplete
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
              </div>
          )
      }
    }
  }

  return flow ? (
    <>
      <form action={flow.ui.action} method={flow.ui.method} onSubmit={submit}>
        {filterNodesByGroups({
          nodes: flow.ui.nodes,
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
