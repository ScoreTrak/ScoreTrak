import { Route } from "@tanstack/router"
import { authRoute } from "."
import { ory } from "../../lib/auth/ory"
import { RegistrationFlow, UiNode, UiNodeInputAttributes } from "@ory/client"
import {
  filterNodesByGroups,
  isUiNodeInputAttributes,
} from "@ory/integrations/ui"
import { useState, useEffect } from "react"

interface RegisterSearch {
  flow: string
}

export const authRegisterRoute = new Route({
  getParentRoute: () => authRoute,
  path: "register",
  component: () => {
    const [flow, setFlow] = useState<RegistrationFlow>()

    useEffect(() => {
      ory.createBrowserRegistrationFlow()
        .then(({ data: flowData}) => {
          setFlow(flowData)
        })
        .catch((err) => {
          console.error(err);
        })
    }, [])
    
    const mapUINode = (node: UiNode, key: number) => {
      if (isUiNodeInputAttributes(node.attributes)) {
        const attrs = node.attributes as UiNodeInputAttributes
        const nodeType = attrs.type
  
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
      <form action={flow.ui.action} method={flow.ui.method}>
        {filterNodesByGroups({
          nodes: flow.ui.nodes,
          // we will also map default fields here such as csrf_token
          // this only maps the `password` method
          // other methods can also be mapped such as `oidc` or `webauthn`
          groups: ["default", "password"],
        }).map((node, idx) => mapUINode(node, idx))}
      </form>
    ) : (
      <div>Loading...</div>
    )

    return (
      <>
        <p>register</p>
      </>
    )
  }
})
